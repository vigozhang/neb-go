package transaction

import (
	"testing"
	"encoding/hex"
	"github.com/vigozhang/neb-go/core/account"
	"math/big"
)

func TestTransaction_HashTransaction(t *testing.T) {
	tx := newTransaction()

	t.Logf("HashTransaction: %s", hex.EncodeToString(tx.HashTransaction()))
}

func TestTransaction_SignTransaction(t *testing.T) {
	tx := newTransaction()

	err := tx.SignTransaction()

	if err != nil {
		t.Error("TestTransaction_SignTransaction failed")
	} else {
		t.Logf("TestTransaction_SignTransaction success,sign:%s", hex.EncodeToString(tx.Sign))
	}
}

func TestTransaction_ToString(t *testing.T) {
	tx := newTransaction()

	tx.SignTransaction()

	txStr, err := tx.ToString()
	if err != nil {
		t.Error("TestTransaction_ToString failed")
	} else {
		t.Logf("TestTransaction_ToString success,result:%s", txStr)
	}

}

func TestTransaction_ToProtoString(t *testing.T) {
	tx := newTransaction()

	tx.SignTransaction()

	protoStr, err := tx.ToProtoString()
	if err != nil {
		t.Error("TestTransaction_ToProtoString failed")
	} else {
		t.Logf("TestTransaction_ToProtoString success,result:%s", protoStr)
	}
}

func TestTransaction_FromProto(t *testing.T) {
	protoStr := "CiAuTz4bhJj54X/doi6EXmIf3f1H1oOil7r2U/nOTGNX9hIaGVdDhNxJ4+OzYNWr2if95MASrtEj0U0nmgYaGhlXf89CeLWgHFjKu9/6tn4KNbelsMDAIIi2IhAAAAAAAAAAAAAAAAAAAAAKKAww3d7p2AU6KAoEY2FsbBIgeyJGdW5jdGlvbiI6InNhdmUiLCJBcmdzIjoiWzBdIn1AAUoQAAAAAAAAAAAAAAAAAA9CQFIQAAAAAAAAAAAAAAAAAB6EgFgBYkGkVEUhcFggQZVmN+2C5c6UPOgqF/pFWTk/g3HKqDBjpRq3Gz/Rtp7znhRTQCYZ2bp6FzX5OaVv90LTuIW3kwaIAA=="

	tx := newTransaction()

	tx, err := tx.FromProto(protoStr)

	if err != nil {
		t.Error("TestTransaction_FromProto failed")
	} else {
		newProStr, _ := tx.ToProtoString()
		if newProStr != protoStr {
			t.Error("TestTransaction_FromProto failed")
		} else {
			t.Logf("TestTransaction_FromProto success")
		}
	}

}

func newAccount() *account.Account {
	acc := account.NewAccount()

	accPriv, _ := big.NewInt(0).SetString("ac3773e06ae74c0fa566b0e421d4e391333f31aef90b383f0c0e83e4873609d6", 16)
	acc.SetPrivateKey(accPriv.Bytes())

	return acc
}

func newTransaction() *Transaction {
	contract := new(Contract)
	contract.Function = "save"
	contract.Args = `[0]`

	txopts := TransactionOptions{
		ChainID:  1,
		From:     newAccount(),
		To:       "n1SAeQRVn33bamxN4ehWUT7JGdxipwn8b17",
		Value:    10,
		Nonce:    12,
		GasPrice: 1000000,
		GasLimit: 2000000,
		Contract: contract,
	}

	return NewTransaction(txopts)
}
