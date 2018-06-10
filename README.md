# neb-go

Go SDK for NAS API

neb.js API Reference: [https://nebulasio.github.io/neb.js/index.html](https://nebulasio.github.io/neb.js/index.html)



## Installation

`go get github.com/vigozhang/neb-go`



## Usage



### Account

```go
// create a new account
acc := account.NewAccount()

// set private key
accPriv, _ := big.NewInt(0).SetString("ac3773e06ae74c0fa566b0e421d4e391333f31aef90b383f0c0e83e4873609d6", 16)
acc.SetPrivateKey(accPriv.Bytes())

// export account
salt, _ := big.NewInt(0).SetString("acded3c4c3f7655cbcd38eea82fdc939251354f21c14cd7faec05c73aa9f70e3", 16)
iv, _ := big.NewInt(0).SetString("e90bd4e307153dc61632ce428606620c", 16)
id, _ := uuid.FromString("0c560f8c-4c0e-4584-9a26-fb57d1dfd0e5")
keyopts := account.KeyOptions{
	Salt: salt.Bytes(),
	Iv:   iv.Bytes(),
	Uuid: id.Bytes(),
}
key, err := acc.ToKey("passphrase", &keyopts)
keystring, err := acc.ToKeyString("passphrase", &keyopts)

// load account
accLoaded, err := acc.FromKey(keystring, "passphrase", true)
```



### API

```go
httpreq := httprequest.NewHttpRequest(httprequest.TestNet, httprequest.APIVersion1)
neb := rpc.NewNeb(httpreq)
api := neb.Api

// GetNebState
respNeb, err := api.GetNebState()

// LatestIrreversibleBlock
respBlock, err := api.LatestIrreversibleBlock()
```



### Transaction

```go
httpreq := httprequest.NewHttpRequest(httprequest.TestNet, httprequest.APIVersion1)
neb := rpc.NewNeb(httpreq)
api := neb.Api

getAccountReq := rpc.GetAccountStateRequest{
	Address: "n1UHqTFvng8vXbcoWxYECwc4shXKnrcXwdz",
}
respAccState, err := api.GetAccountState(getAccountReq)
if err != nil {
	log.Print("get account error:", err)
}

// load account
acc := loadAccount()

// build transaction
contract := transaction.Contract{
	Function: "getCreatedList",
	Args:     string(utils.EncodeToJsonBytes([]string{"n1UHqTFvng8vXbcoWxYECwc4shXKnrcXwdz"})),
	}

txOpts := transaction.TransactionOptions{
	ChainID:  1001,
	From:     acc,
	To:       "n1f5rgBtVKVEjxPBwrDcaV8H8QqxniFyhPk",
	Value:    0,
	Nonce:    respAccState.Nonce + 1,
	GasPrice: 1000000,
	GasLimit: 2000000,
	Contract: &contract,
}

tx := transaction.NewTransaction(txOpts)
tx.SignTransaction()
raw, _ := tx.ToProtoString()

// send transaction
req := rpc.SendRawTransactionRequest{
	Data: raw,
}
resp, _ := api.SendRawTransaction(req)

jsonBytes, _ := json.Marshal(resp)
log.Println(string(jsonBytes))
```



### Call Contract

```go
httpreq := httprequest.NewHttpRequest(httprequest.TestNet, httprequest.APIVersion1)
neb := rpc.NewNeb(httpreq)
api := neb.Api

contract := rpc.ContractRequest{
	Function: "getCreatedList",
	Args:     string(utils.EncodeToJsonBytes([]string{"n1UHqTFvng8vXbcoWxYECwc4shXKnrcXwdz"})),
}

req := rpc.TransactionRequest{
	From:     "n1UHqTFvng8vXbcoWxYECwc4shXKnrcXwdz",
	To:       "n1f5rgBtVKVEjxPBwrDcaV8H8QqxniFyhPk",
	Value:    "0",
	Nonce:    3,
	GasPrice: "1000000",
	GasLimit: "2000000",
	Contract: &contract,
}

resp, err := api.Call(req)
if err != nil {
	log.Println("TestApi_Call failed")
} else {
	log.Println("TestApi_Call Resp:", string(utils.EncodeToJsonBytes(resp)))
}
```



