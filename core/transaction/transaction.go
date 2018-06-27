package transaction

import (
	"math/big"
	"time"
	"errors"
	"encoding/hex"
	"encoding/json"
	"encoding/base64"

	"github.com/vigozhang/neb-go/core/account"
	"github.com/vigozhang/neb-go/utils"
	"github.com/vigozhang/neb-go/utils/hash"
	"github.com/nebulasio/go-nebulas/core/pb"
	"github.com/golang/protobuf/proto"
	"github.com/vigozhang/neb-go/utils/byteutils"
	"github.com/vigozhang/neb-go/utils/secp256k1"
	"github.com/vigozhang/neb-go/utils/base58"
)

const (
	TxPayloadBinaryType = "binary"
	TxPayloadDeployType = "deploy"
	TxPayloadCallType   = "call"

	SECP256K1 = 1
)

type TransactionOptions struct {
	ChainID  uint32
	From     *account.Account
	To       string
	Value    *big.Int
	Nonce    uint64
	GasPrice *big.Int
	GasLimit *big.Int
	Contract *Contract
}

type Transaction struct {
	Hash      []byte
	From      *account.Account
	To        *account.Account
	Value     *big.Int
	Nonce     uint64
	Timestamp int64
	Data      *TxPayload
	ChainID   uint32
	GasPrice  *big.Int
	GasLimit  *big.Int

	Alg  uint32
	Sign []byte
}

type Contract struct {
	Source     string
	SourceType string
	Args       string
	Function   string
	Binary     []byte
}

type TxPayload struct {
	Type    string
	Payload []byte
}

type TransactionBinaryPayload struct {
	Data []byte
}

type TransactionCallPayload struct {
	Function string
	Args     string
}

type TransactionDeployPayload struct {
	SourceType string
	Source     string
	Args       string
}

func NewTransaction(opts TransactionOptions) *Transaction {
	transaction := new(Transaction)
	transaction.ChainID = opts.ChainID
	transaction.From = opts.From
	transaction.To, _ = account.FromAddress(opts.To)
	transaction.Value = opts.Value
	transaction.Nonce = opts.Nonce
	transaction.Timestamp = time.Now().Unix()
	transaction.GasPrice = opts.GasPrice
	transaction.GasLimit = opts.GasLimit
	transaction.Data = parseContract(opts.Contract)

	if transaction.GasPrice.Cmp(big.NewInt(0)) == -1 {
		transaction.GasPrice = big.NewInt(1000000)
	}

	if transaction.GasLimit.Cmp(big.NewInt(0)) == -1 {
		transaction.GasLimit = big.NewInt(20000)
	}

	return transaction
}

func (tx *Transaction) HashTransaction() []byte {

	dataBytes := SerializeDataToProto(tx.Data)

	hashValue := hash.Sha3256(
		tx.From.GetAddress(),
		tx.To.GetAddress(),
		utils.To128BitSlice(tx.Value),
		byteutils.FromUint64(tx.Nonce),
		byteutils.FromInt64(tx.Timestamp),
		dataBytes,
		byteutils.FromUint32(tx.ChainID),
		utils.To128BitSlice(tx.GasPrice),
		utils.To128BitSlice(tx.GasLimit),
	)

	return hashValue
}

func (tx *Transaction) SignTransaction() error {
	if tx.From.GetPrivateKey() == nil {
		return errors.New("transaction from address's private key is invalid")
	}

	tx.Hash = tx.HashTransaction()
	tx.Alg = SECP256K1
	var err error
	tx.Sign, err = secp256k1.Sign(tx.Hash, tx.From.GetPrivateKey())
	if err != nil {
		return err
	}
	return nil
}

func (tx *Transaction) ToString() (string, error) {
	if tx.Sign == nil {
		return "", errors.New("you should sign transaction before this operation")
	}

	var payloadMap map[string]interface{}
	if tx.Data.Payload != nil {
		json.Unmarshal(tx.Data.Payload, &payloadMap)
	}

	mapData := make(map[string]interface{})
	mapData["payloadType"] = tx.Data.Type
	mapData["payload"] = payloadMap

	mapTx := make(map[string]interface{})
	mapTx["chainID"] = tx.ChainID
	mapTx["from"] = tx.From.GetAddressString()
	mapTx["to"] = tx.To.GetAddressString()
	mapTx["value"] = tx.Value.String()
	mapTx["nonce"] = tx.Nonce
	mapTx["timestamp"] = tx.Timestamp
	mapTx["data"] = mapData
	mapTx["gasPrice"] = tx.GasPrice.String()
	mapTx["gasLimit"] = tx.GasLimit.String()
	mapTx["hash"] = hex.EncodeToString(tx.Hash)
	mapTx["alg"] = tx.Alg
	mapTx["sign"] = hex.EncodeToString(tx.Sign)

	txBytes, err := json.Marshal(&mapTx)
	if err != nil {
		return "", err
	}
	return string(txBytes), nil
}

func (tx *Transaction) ToProto() ([]byte, error) {
	if tx.Sign == nil {
		return nil, errors.New("you should sign transaction before this operation")
	}

	data := corepb.Data{
		Type:    tx.Data.Type,
		Payload: tx.Data.Payload,
	}

	txData := corepb.Transaction{
		Hash:      tx.Hash,
		From:      tx.From.GetAddress(),
		To:        tx.To.GetAddress(),
		Value:     utils.To128BitSlice(tx.Value),
		Nonce:     tx.Nonce,
		Timestamp: tx.Timestamp,
		Data:      &data,
		ChainId:   tx.ChainID,
		GasPrice:  utils.To128BitSlice(tx.GasPrice),
		GasLimit:  utils.To128BitSlice(tx.GasLimit),
		Alg:       tx.Alg,
		Sign:      tx.Sign,
	}

	txBytes, err := proto.Marshal(&txData)
	if err != nil {
		return nil, err
	}
	return txBytes, nil
}

func (tx *Transaction) ToProtoString() (string, error) {
	txBytes, err := tx.ToProto()
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(txBytes), nil
}

func (tx *Transaction) FromProto(protoString string) (*Transaction, error) {
	txBytes, err := base64.StdEncoding.DecodeString(protoString)
	if err != nil {
		return nil, err
	}

	txProto := corepb.Transaction{}
	err = proto.Unmarshal(txBytes, &txProto)
	if err != nil {
		return nil, err
	}

	tx.Hash = txProto.Hash
	tx.From, err = account.FromAddress(base58.Encode(txProto.From))
	if err != nil {
		return nil, err
	}

	tx.To, err = account.FromAddress(base58.Encode(txProto.To))
	if err != nil {
		return nil, err
	}

	tx.Value = big.NewInt(1).SetBytes(txProto.Value)
	tx.Nonce = txProto.Nonce
	tx.Timestamp = txProto.Timestamp

	txPayload := TxPayload{
		Type:    txProto.Data.Type,
		Payload: txProto.Data.Payload,
	}
	tx.Data = &txPayload
	if len(tx.Data.Payload) == 0 {
		tx.Data.Payload = nil
	}

	tx.ChainID = txProto.ChainId
	tx.GasPrice = big.NewInt(1).SetBytes(txProto.GasPrice)
	tx.GasLimit = big.NewInt(1).SetBytes(txProto.GasLimit)
	tx.Alg = txProto.Alg
	tx.Sign = txProto.Sign

	return tx, nil
}

func SerializeDataToProto(txpayload *TxPayload) []byte {
	data := corepb.Data{
		Type:    txpayload.Type,
		Payload: txpayload.Payload,
	}

	protoBytes, _ := proto.Marshal(&data)

	return protoBytes
}

func parseContract(contract *Contract) *TxPayload {
	var payloadType string
	var payload []byte
	if len(contract.Source) > 0 {
		payloadType = TxPayloadDeployType
		transactionDeployPayload := TransactionDeployPayload{
			SourceType: contract.SourceType,
			Source:     contract.Source,
			Args:       contract.Args,
		}

		payload = utils.EncodeToJsonBytes(transactionDeployPayload)

	} else if len(contract.Function) > 0 {
		payloadType = TxPayloadCallType
		transactionCallPayload := TransactionCallPayload{
			Function: contract.Function,
			Args:     contract.Args,
		}

		payload = utils.EncodeToJsonBytes(transactionCallPayload)

	} else {
		payloadType = TxPayloadBinaryType
		if len(contract.Binary) > 0 {
			transactionBinaryPayload := TransactionBinaryPayload{contract.Binary}

			payload = utils.EncodeToJsonBytes(transactionBinaryPayload)
		} else {
			payload = nil
		}
	}

	txPayload := TxPayload{payloadType, payload}

	return &txPayload

}
