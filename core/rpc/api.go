package rpc

import (
	"encoding/json"
	"log"

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
		return nil, err
	}
	mapResult := make(map[string]*GetNebStateResponse)
	mapResult["result"] = new(GetNebStateResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetNebState error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) LatestIrreversibleBlock() (*BlockResponse, error) {
	resp, err := api.HttpRequest.Get("/user/lib", nil)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*BlockResponse)
	mapResult["result"] = new(BlockResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetNebState error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GetAccountState(req GetAccountStateRequest) (*GetAccountStateResponse, error) {
	resp, err := api.HttpRequest.Post("/user/accountstate", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*GetAccountStateResponse)
	mapResult["result"] = new(GetAccountStateResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetAccountState error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) Call(req TransactionRequest) (*CallResponse, error) {
	resp, err := api.HttpRequest.Post("/user/call", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*CallResponse)
	mapResult["result"] = new(CallResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("Call error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) SendRawTransaction(req SendRawTransactionRequest) (*SendTransactionResponse, error) {
	resp, err := api.HttpRequest.Post("/user/rawtransaction", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*SendTransactionResponse)
	mapResult["result"] = new(SendTransactionResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Printf("SendRawTransaction error:%s, response is %s", err, string(resp))
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GetBlockByHash(req GetBlockByHashRequest) (*BlockResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getBlockByHash", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*BlockResponse)
	mapResult["result"] = new(BlockResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetBlockByHashRequest error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GetBlockByHeight(req GetBlockByHeightRequest) (*BlockResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getBlockByHeight", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*BlockResponse)
	mapResult["result"] = new(BlockResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetBlockByHashRequest error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GetTransactionByContract(req GetTransactionByContractRequest) (*TransactionResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getTransactionByContract", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*TransactionResponse)
	mapResult["result"] = new(TransactionResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetTransactionByContract error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) Subscribe(req SubscribeRequest) (*SubscribeResponse, error) {
	// TODO: has problem
	resp, err := api.HttpRequest.Post("/user/subscribe", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*SubscribeResponse)
	mapResult["result"] = new(SubscribeResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Printf("Subscribe error:%s, response is %s", err, string(resp))
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GasPrice() (*GasPriceResponse, error) {
	resp, err := api.HttpRequest.Get("/user/getGasPrice", nil)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*GasPriceResponse)
	mapResult["result"] = new(GasPriceResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GasPrice error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) EstimateGas(req TransactionRequest) (*GasResponse, error) {
	resp, err := api.HttpRequest.Post("/user/estimateGas", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*GasResponse)
	mapResult["result"] = new(GasResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("EstimateGas error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GetEventsByHash(req HashRequest) (*EventsResponse, error) {
	resp, err := api.HttpRequest.Post("/user/getEventsByHash", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*EventsResponse)
	mapResult["result"] = new(EventsResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetEventsByHash error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}

func (api *Api) GetDynasty(req ByBlockHeightRequest) (*GetDynastyResponse, error) {
	resp, err := api.HttpRequest.Post("/user/dynasty", req)
	if err != nil {
		return nil, err
	}
	mapResult := make(map[string]*GetDynastyResponse)
	mapResult["result"] = new(GetDynastyResponse)
	err = json.Unmarshal(resp, &mapResult)
	if err != nil {
		log.Println("GetDynasty error:", err)
		return nil, err
	}
	return mapResult["result"], nil
}
