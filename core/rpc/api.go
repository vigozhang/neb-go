package rpc

import (
	"encoding/json"
	"log"
	"bytes"
	"bufio"
	"net/http"

	"github.com/vigozhang/neb-go/utils/httprequest"
)

type Api struct {
	HttpRequest *httprequest.HttpRequest
}

func NewApi(neb *Neb) *Api {
	return &Api{neb.HttpRequest}
}

func (api *Api) SetRequest(request *httprequest.HttpRequest) {
	api.HttpRequest = request
}

func (api *Api) GetNebState() (*GetNebStateResponse, error) {
	resp, err := api.HttpRequest.Get("/user/nebstate", nil)
	if err != nil {
		logError("GetNebState", string(resp), err)
		return nil, err
	}

	var response GetNebStateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetNebState", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) LatestIrreversibleBlock() (*BlockResponse, error) {
	resp, err := api.HttpRequest.Get("/user/lib", nil)
	if err != nil {
		logError("LatestIrreversibleBlock", string(resp), err)
		return nil, err
	}

	var response BlockResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("LatestIrreversibleBlock", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetAccountState(req GetAccountStateRequest) (*GetAccountStateResponse, error) {
	resp, err := api.HttpRequest.Post("/user/accountstate", req)
	if err != nil {
		logError("GetAccountState", string(resp), err)
		return nil, err
	}

	var response GetAccountStateResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetAccountState", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) Call(req TransactionRequest) (*CallResponse, error) {
	resp, err := api.HttpRequest.Post("/user/call", req)
	if err != nil {
		logError("Call", string(resp), err)
		return nil, err
	}

	var response CallResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("Call", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) SendRawTransaction(req SendRawTransactionRequest) (*SendTransactionResponse, error) {
	resp, err := api.HttpRequest.Post("/user/rawtransaction", req)
	if err != nil {
		logError("SendRawTransaction", string(resp), err)
		return nil, err
	}

	var response SendTransactionResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("SendRawTransaction", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetBlockByHash(req GetBlockByHashRequest) (*BlockResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getBlockByHash", req)
	if err != nil {
		logError("GetBlockByHash", string(resp), err)
		return nil, err
	}

	var response BlockResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetBlockByHash", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetBlockByHeight(req GetBlockByHeightRequest) (*BlockResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getBlockByHeight", req)
	if err != nil {
		logError("GetBlockByHeight", string(resp), err)
		return nil, err
	}

	var response BlockResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetBlockByHeight", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetTransactionReceipt(req HashRequest) (*TransactionResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getTransactionReceipt", req)
	if err != nil {
		logError("GetTransactionReceipt", string(resp), err)
		return nil, err
	}

	var response TransactionResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetTransactionReceipt", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetTransactionByContract(req GetTransactionByContractRequest) (*TransactionResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getTransactionByContract", req)
	if err != nil {
		logError("GetTransactionByContract", string(resp), err)
		return nil, err
	}

	var response TransactionResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetTransactionByContract", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) Subscribe(req SubscribeRequest, callback func(response *SubscribeResponse)) (error) {
	err := postStreamForSubscribe(api.HttpRequest, "/user/subscribe", req, callback)
	if err != nil {
		log.Printf("Subscribe error:%s", err)
		return err
	}
	return nil
}

func (api *Api) GasPrice() (*GasPriceResponse, error) {
	resp, err := api.HttpRequest.Get("/user/getGasPrice", nil)
	if err != nil {
		logError("GasPrice", string(resp), err)
		return nil, err
	}

	var response GasPriceResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GasPrice", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) EstimateGas(req TransactionRequest) (*GasResponse, error) {
	resp, err := api.HttpRequest.Post("/user/estimateGas", req)
	if err != nil {
		logError("EstimateGas", string(resp), err)
		return nil, err
	}

	var response GasResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("EstimateGas", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetEventsByHash(req HashRequest) (*EventsResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getEventsByHash", req)
	if err != nil {
		logError("GetEventsByHash", string(resp), err)
		return nil, err
	}

	var response EventsResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetEventsByHash", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func (api *Api) GetDynasty(req ByBlockHeightRequest) (*GetDynastyResponse, error) {
	resp, err := api.HttpRequest.Post("/user/dynasty", req)
	if err != nil {
		logError("GetDynasty", string(resp), err)
		return nil, err
	}

	var response GetDynastyResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		logError("GetDynasty", string(resp), err)
		return nil, err
	}
	return &response, nil
}

func logError(method string, resp string, err error) {
	log.Printf("%s error:%s, response is %s", method, err, resp)
}

func postStreamForSubscribe(req *httprequest.HttpRequest, api string, reqBody interface{}, callback func(response *SubscribeResponse)) (error) {
	url := req.CreateUrl(api)
	contentType := "application/json"

	jsonReqBody, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	response, err := http.Post(url, contentType, bytes.NewBuffer(jsonReqBody))
	defer response.Body.Close()

	if err != nil {
		return err
	}

	reader := bufio.NewReader(response.Body)
	for {
		line, err := reader.ReadBytes('\n')

		if err != nil {
			break
		}

		var resp SubscribeResponse
		err = json.Unmarshal(line, &resp)
		if err != nil {
			log.Printf("%s error:%s, response is %s", "Subscribe", err, resp)
			break
		}

		callback(&resp)
	}

	if err != nil {
		return err
	}

	return nil
}
