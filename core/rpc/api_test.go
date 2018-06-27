package rpc

import (
	"testing"
	"log"
	"math/big"

	"github.com/vigozhang/neb-go/utils/httprequest"
	"github.com/vigozhang/neb-go/utils"
	"github.com/vigozhang/neb-go/core/transaction"
	"github.com/vigozhang/neb-go/core/account"
)

var requesttest = httprequest.NewHttpRequest(httprequest.MainNet, httprequest.APIVersion1)
var nebtest = NewNeb(requesttest)
var api = nebtest.Api

const TestNetAccountAddress = "n1UHqTFvng8vXbcoWxYECwc4shXKnrcXwdz"
const TestNetContractAddress = "n1f5rgBtVKVEjxPBwrDcaV8H8QqxniFyhPk"

func TestApi_GetNebState(t *testing.T) {
	resp, err := api.GetNebState()
	if err != nil {
		t.Error("TestApi_GetNebState failed")
	} else {
		t.Log("TestApi_GetNebState Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_LatestIrreversibleBlock(t *testing.T) {
	resp, err := api.LatestIrreversibleBlock()
	if err != nil {
		t.Error("TestApi_LatestIrreversibleBlock failed")
	} else {
		t.Log("TestApi_LatestIrreversibleBlock Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_GetAccountState(t *testing.T) {
	req := GetAccountStateRequest{
		Address: TestNetAccountAddress,
	}
	resp, err := api.GetAccountState(req)
	if err != nil {
		t.Error("TestApi_GetAccountState failed")
	} else {
		t.Log("TestApi_GetAccountState Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_Call(t *testing.T) {
	contract := ContractRequest{
		Function: "getCreatedList",
		Args:     string(utils.EncodeToJsonBytes([]string{TestNetAccountAddress})),
	}

	req := TransactionRequest{
		From:     TestNetAccountAddress,
		To:       TestNetContractAddress,
		Value:    "0",
		Nonce:    3,
		GasPrice: "1000000",
		GasLimit: "2000000",
		Contract: &contract,
	}

	resp, err := api.Call(req)
	if err != nil {
		t.Error("TestApi_Call failed")
	} else {
		t.Log("TestApi_Call Resp:", string(utils.EncodeToJsonBytes(resp)))
	}

}

func TestApi_SendRawTransaction(t *testing.T) {
	acc := loadAccount()

	contract := transaction.Contract{
		Function: "getCreatedList",
		Args:     string(utils.EncodeToJsonBytes([]string{"n1UHqTFvng8vXbcoWxYECwc4shXKnrcXwdz"})),
	}

	txOpts := transaction.TransactionOptions{
		ChainID:  1001,
		From:     acc,
		To:       "n1f5rgBtVKVEjxPBwrDcaV8H8QqxniFyhPk",
		Value:    big.NewInt(0),
		Nonce:    6,
		GasPrice: big.NewInt(1000000),
		GasLimit: big.NewInt(2000000),
		Contract: &contract,
	}

	tx := transaction.NewTransaction(txOpts)
	tx.SignTransaction()
	raw, _ := tx.ToProtoString()

	req := SendRawTransactionRequest{
		Data: raw,
	}

	resp, err := api.SendRawTransaction(req)
	if err != nil {
		t.Error("TestApi_SendRawTransaction failed")
	} else {
		t.Log("TestApi_SendRawTransaction Resp:", string(utils.EncodeToJsonBytes(resp)))
	}

}

func TestApi_GetBlockByHash(t *testing.T) {
	testNetBlockHash := "8371d448bdc86218a38c9e89c17c8665a4bbb0bebb0dcbf094712447a3d8fcf1"
	req := GetBlockByHashRequest{
		Hash: testNetBlockHash,
	}
	resp, err := api.GetBlockByHash(req)
	if err != nil {
		t.Error("TestApi_GetBlockByHash failed")
	} else {
		t.Log("TestApi_GetBlockByHash Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_GetBlockByHeight(t *testing.T) {
	testNetBlockHeight := uint64(377161)
	req := GetBlockByHeightRequest{
		Height: testNetBlockHeight,
	}
	resp, err := api.GetBlockByHeight(req)
	if err != nil {
		t.Error("TestApi_GetBlockByHeight failed")
	} else {
		t.Log("TestApi_GetBlockByHeight Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_Subscribe(t *testing.T) {
	req := SubscribeRequest{
		Topics: []string{"chain.pendingTransaction", "chain.sendTransaction", "chain.linkBlock"},
	}

	err := api.Subscribe(req, subscribeCallback)
	if err != nil {
		t.Error("TestApi_Subscribe failed")
	}
}

func TestApi_GasPrice(t *testing.T) {
	resp, err := api.GasPrice()
	if err != nil {
		t.Error("TestApi_GasPrice failed")
	} else {
		t.Log("TestApi_GasPrice Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_EstimateGas(t *testing.T) {
	contract := ContractRequest{
		Function: "getCreatedList",
		Args:     string(utils.EncodeToJsonBytes([]string{TestNetAccountAddress})),
	}

	req := TransactionRequest{
		From:     TestNetAccountAddress,
		To:       TestNetContractAddress,
		Value:    "0",
		Nonce:    3,
		GasPrice: "1000000",
		GasLimit: "2000000",
		Contract: &contract,
	}

	resp, err := api.EstimateGas(req)
	if err != nil {
		t.Error("TestApi_EstimateGas failed")
	} else {
		t.Log("TestApi_EstimateGas Resp:", string(utils.EncodeToJsonBytes(resp)))
	}

}

func TestApi_GetEventsByHash(t *testing.T) {
	testNetTxHash := "0b4239206842b6ec2fd12a94f3370946f9246b2660d54c721202ba22c42ad146"
	req := HashRequest{
		Hash: testNetTxHash,
	}
	resp, err := api.GetEventsByHash(req)
	if err != nil {
		t.Error("TestApi_GetEventsByHash failed")
	} else {
		t.Log("TestApi_GetEventsByHash Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestApi_GetDynasty(t *testing.T) {
	testNetBlockHeight := uint64(377161)
	req := ByBlockHeightRequest{
		Height: testNetBlockHeight,
	}
	resp, err := api.GetDynasty(req)
	if err != nil {
		t.Error("TestApi_GetDynasty failed")
	} else {
		t.Log("TestApi_GetDynasty Resp:", string(utils.EncodeToJsonBytes(resp)))
	}

}

func loadAccount() *account.Account {
	acc := account.NewAccount()
	// keyjson
	keyjson := ""
	loadAcc, _ := acc.FromKey(keyjson, "passphrase", true)
	return loadAcc
}

func subscribeCallback(line *SubscribeResponse) {
	log.Println(line.Result.Topic)
	log.Println(line.Result.Data)
}
