package rpc

import (
	"encoding/json"
	"log"

	"github.com/vigozhang/neb-go/utils/httprequest"
)

type Admin struct {
	HttpRequest *httprequest.HttpRequest
}

func NewAdmin(neb *Neb) *Admin {
	return &Admin{neb.HttpRequest}
}

func (admin *Admin) SetRequest(request *httprequest.HttpRequest) {
	admin.HttpRequest = request
}

func (admin *Admin) NodeInfo() (*NodeInfoResponse, error) {
	resp, err := admin.HttpRequest.Get("/admin/nodeinfo", nil)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*NodeInfoResponse)
	mapResult["result"] = new(NodeInfoResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("nodeinfo error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) Accounts() (*AccountsResponse, error) {
	resp, err := admin.HttpRequest.Get("/admin/accounts", nil)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*AccountsResponse)
	mapResult["result"] = new(AccountsResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("Accounts error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) NewAccount(req NewAccountRequest) (*NewAccountResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/account/new", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*NewAccountResponse)
	mapResult["result"] = new(NewAccountResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("NewAccount error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) UnlockAccount(req UnlockAccountRequest) (*UnlockAccountResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/account/unlock", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*UnlockAccountResponse)
	mapResult["result"] = new(UnlockAccountResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("UnlockAccount error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) LockAccount(req LockAccountRequest) (*LockAccountResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/account/lock", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*LockAccountResponse)
	mapResult["result"] = new(LockAccountResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("LockAccount error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) SendTransaction(req TransactionRequest) (*SendTransactionResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/transaction", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*SendTransactionResponse)
	mapResult["result"] = new(SendTransactionResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Printf("SendTransaction error:%s,response is %s", err, string(resp))
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) SignHash(req SignHashRequest) (*SignHashResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/sign/hash", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*SignHashResponse)
	mapResult["result"] = new(SignHashResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Printf("SignHash error:%s, response is %s", err, string(resp))
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) SignTransactionWithPassphrase(req SignTransactionPassphraseRequest) (*SignTransactionPassphraseResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/sign", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*SignTransactionPassphraseResponse)
	mapResult["result"] = new(SignTransactionPassphraseResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("SignTransactionWithPassphrase error:", err)
		return nil, err
	}
	return mapResult["result"], err
}

func (admin *Admin) SendTransactionWithPassphrase(req SendTransactionPassphraseRequest) (*SendTransactionResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/transactionWithPassphrase", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*SendTransactionResponse)
	mapResult["result"] = new(SendTransactionResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("SendTransactionWithPassphrase error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) StartPprof(req PprofRequest) (*PprofResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/pprof", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*PprofResponse)
	mapResult["result"] = new(PprofResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("StartPprof error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (admin *Admin) GetConfig() (*GetConfigResponse, error) {
	resp, err := admin.HttpRequest.Get("/admin/getConfig", nil)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*GetConfigResponse)
	mapResult["result"] = new(GetConfigResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetConfig error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}
