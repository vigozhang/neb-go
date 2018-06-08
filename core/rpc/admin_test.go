package rpc

import (
	"testing"
	"github.com/vigozhang/neb-go/utils/httprequest"
	"github.com/vigozhang/neb-go/utils"
)

var requestlocal = httprequest.NewHttpRequest(httprequest.LocalNet, httprequest.APIVersion1)
var neblocal = NewNeb(requestlocal)
var admin = neblocal.Admin

const TestLocalAddress = "n1KkgcnqQAapJPSrevvAbC1Ze9SuiAVtYwD"

func TestAdmin_NodeInfo(t *testing.T) {
	resp, err := admin.NodeInfo()
	if err != nil {
		t.Error("TestAdmin_NodeInfo failed")
	} else {
		t.Log("TestAdmin_NodeInfo Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_Accounts(t *testing.T) {
	resp, err := admin.Accounts()
	if err != nil {
		t.Error("TestAdmin_Accounts failed")
	} else {
		t.Log("TestAdmin_Accounts Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_NewAccount(t *testing.T) {
	// n1KkgcnqQAapJPSrevvAbC1Ze9SuiAVtYwD
	req := NewAccountRequest{
		Passphrase: "123456",
	}
	resp, err := admin.NewAccount(req)
	if err != nil {
		t.Error("TestAdmin_NewAccount failed")
	} else {
		t.Log("TestAdmin_NewAccount Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_UnlockAccount(t *testing.T) {
	req := UnlockAccountRequest{
		Address:    TestLocalAddress,
		Passphrase: "123456",
		// The unit is ns
		Duration: 100000000000,
	}
	resp, err := admin.UnlockAccount(req)
	if err != nil {
		t.Error("TestAdmin_UnlockAccount failed")
	} else {
		t.Log("TestAdmin_UnlockAccount Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_LockAccount(t *testing.T) {
	req := LockAccountRequest{
		Address: TestLocalAddress,
	}
	resp, err := admin.LockAccount(req)
	if err != nil {
		t.Error("TestAdmin_LockAccount failed")
	} else {
		t.Log("TestAdmin_LockAccount Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_SendTransaction(t *testing.T) {
	//contract := ContractRequest{
	//	Function: "getCreatedList",
	//	Args:     string(utils.EncodeToJsonBytes([]string{TestNetAccountAddress})),
	//}

	req := TransactionRequest{
		From:     TestLocalAddress,
		To:       "n1GmkKH6nBMw4rrjt16RrJ9WcgvKUtAZP1s",
		Value:    "100",
		Nonce:    3,
		GasPrice: "1000000",
		GasLimit: "2000000",
		//Contract: &contract,
	}

	resp, err := admin.SendTransaction(req)
	if err != nil {
		t.Error("TestAdmin_SendTransaction failed")
	} else {
		t.Log("TestAdmin_SendTransaction Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_SignHash(t *testing.T) {
	req := SignHashRequest{
		Address: TestLocalAddress,
		Hash:    "W+rOKNqs/tlvz02ez77yIYMCOr2EubpuNh5LvmwceI0=",
		Alg:     1,
	}

	resp, err := admin.SignHash(req)
	if err != nil {
		t.Error("TestAdmin_SignHash failed")
	} else {
		t.Log("TestAdmin_SignHash Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_SignTransactionWithPassphrase(t *testing.T) {
	tx := TransactionRequest{
		From:     TestLocalAddress,
		To:       "n1GmkKH6nBMw4rrjt16RrJ9WcgvKUtAZP1s",
		Value:    "100",
		Nonce:    3,
		GasPrice: "1000000",
		GasLimit: "2000000",
	}

	req := SignTransactionPassphraseRequest{
		Transaction: &tx,
		Passphrase:  "123456",
	}
	resp, err := admin.SignTransactionWithPassphrase(req)
	if err != nil {
		t.Error("TestAdmin_SignTransactionWithPassphrase failed")
	} else {
		t.Log("TestAdmin_SignTransactionWithPassphrase Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_SendTransactionWithPassphrase(t *testing.T) {
	tx := TransactionRequest{
		From:     TestLocalAddress,
		To:       "n1GmkKH6nBMw4rrjt16RrJ9WcgvKUtAZP1s",
		Value:    "100",
		Nonce:    3,
		GasPrice: "1000000",
		GasLimit: "2000000",
	}

	req := SendTransactionPassphraseRequest{
		Transaction: &tx,
		Passphrase:  "123456",
	}
	resp, err := admin.SendTransactionWithPassphrase(req)
	if err != nil {
		t.Error("TestAdmin_SendTransactionWithPassphrase failed")
	} else {
		t.Log("TestAdmin_SendTransactionWithPassphrase Resp:", string(utils.EncodeToJsonBytes(resp)))
	}

}

func TestAdmin_StartPprof(t *testing.T) {
	req := PprofRequest{
		Listen: "0.0.0.0:10086",
	}
	resp, err := admin.StartPprof(req)
	if err != nil {
		t.Error("TestAdmin_StartPprof failed")
	} else {
		t.Log("TestAdmin_StartPprof Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}

func TestAdmin_GetConfig(t *testing.T) {
	resp, err := admin.GetConfig()
	if err != nil {
		t.Error("TestAdmin_GetConfig failed")
	} else {
		t.Log("TestAdmin_GetConfig Resp:", string(utils.EncodeToJsonBytes(resp)))
	}
}
