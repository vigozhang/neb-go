package rpc

import (
	"encoding/json"
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
		logError("NodeInfo", string(resp), err)
		return nil, err
	}

	var response NodeInfoResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("NodeInfo", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) Accounts() (*AccountsResponse, error) {
	resp, err := admin.HttpRequest.Get("/admin/accounts", nil)
	if err != nil {
		logError("Accounts", string(resp), err)
		return nil, err
	}

	var response AccountsResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("Accounts", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) NewAccount(req NewAccountRequest) (*NewAccountResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/account/new", req)
	if err != nil {
		logError("NewAccount", string(resp), err)
		return nil, err
	}

	var response NewAccountResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("NewAccount", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) UnlockAccount(req UnlockAccountRequest) (*UnlockAccountResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/account/unlock", req)
	if err != nil {
		logError("UnlockAccount", string(resp), err)
		return nil, err
	}

	var response UnlockAccountResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("UnlockAccount", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) LockAccount(req LockAccountRequest) (*LockAccountResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/account/lock", req)
	if err != nil {
		logError("LockAccount", string(resp), err)
		return nil, err
	}

	var response LockAccountResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("LockAccount", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) SendTransaction(req TransactionRequest) (*SendTransactionResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/transaction", req)
	if err != nil {
		logError("SendTransaction", string(resp), err)
		return nil, err
	}

	var response SendTransactionResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("SendTransaction", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) SignHash(req SignHashRequest) (*SignHashResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/sign/hash", req)
	if err != nil {
		logError("SignHash", string(resp), err)
		return nil, err
	}

	var response SignHashResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("SignHash", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) SignTransactionWithPassphrase(req SignTransactionPassphraseRequest) (*SignTransactionPassphraseResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/sign", req)
	if err != nil {
		logError("SignTransactionWithPassphrase", string(resp), err)
		return nil, err
	}

	var response SignTransactionPassphraseResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("SignTransactionWithPassphrase", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) SendTransactionWithPassphrase(req SendTransactionPassphraseRequest) (*SendTransactionResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/transactionWithPassphrase", req)
	if err != nil {
		logError("SendTransactionWithPassphrase", string(resp), err)
		return nil, err
	}

	var response SendTransactionResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("SendTransactionWithPassphrase", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) StartPprof(req PprofRequest) (*PprofResponse, error) {
	resp, err := admin.HttpRequest.Post("/admin/pprof", req)
	if err != nil {
		logError("StartPprof", string(resp), err)
		return nil, err
	}

	var response PprofResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("StartPprof", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (admin *Admin) GetConfig() (*GetConfigResponse, error) {
	resp, err := admin.HttpRequest.Get("/admin/getConfig", nil)
	if err != nil {
		logError("GetConfig", string(resp), err)
		return nil, err
	}

	var response GetConfigResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetConfig", string(resp), err)
		return nil, err
	}
	return &response, nil
}
