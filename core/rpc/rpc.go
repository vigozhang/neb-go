package rpc

type GetNebStateResponse struct {
	// Block chain id
	ChainId uint32 `json:"chain_id,omitempty"`
	// Current neb tail hash
	Tail string `json:"tail,omitempty"`
	// Current neb lib hash
	Lib string `json:"lib,omitempty"`
	// Current neb tail block height
	Height uint64 `json:"height,string,omitempty"`
	// The current neb protocol version.
	ProtocolVersion string `json:"protocol_version,omitempty"`
	// The peer sync status.
	Synchronized bool `json:"synchronized,omitempty"`
	// neb version
	Version string `json:"version,omitempty"`
}

type GetAccountStateRequest struct {
	Address string `json:"address,omitempty"`
	Height  uint64 `json:"height,omitempty"`
}

type GetAccountStateResponse struct {
	// Current balance in unit of 1/(10^18) nas.
	Balance string `json:"balance,omitempty"`
	// Current transaction count.
	Nonce uint64 `json:"nonce,string,omitempty"`
	// Account type
	Type uint32 `json:"type,omitempty"`
}

type CallResponse struct {
	// result of smart contract method call.
	Result string `json:"result,omitempty"`
	// execute error
	ExecuteErr string `json:"execute_err,omitempty"`
	// estimate gas used
	EstimateGas string `json:"estimate_gas,omitempty"`
}

type SendRawTransactionRequest struct {
	Data string `json:"data"`
}

type GetBlockByHashRequest struct {
	// Hex string of block hash.
	Hash string `json:"hash,omitempty"`
	// If true it returns the full transaction objects, if false only the hashes of the transactions.
	FullFillTransaction bool `json:"full_fill_transaction,omitempty"`
}

type GetBlockByHeightRequest struct {
	// block height.
	Height uint64 `json:"height,omitempty"`
	// If true it returns the full transaction objects, if false only the hashes of the transactions.
	FullFillTransaction bool `json:"full_fill_transaction,omitempty"`
}

type BlockResponse struct {
	// Hex string of block hash.
	Hash string `json:"hash,omitempty"`
	// Hex string of block parent hash.
	ParentHash string `json:"parent_hash,omitempty"`
	// block height
	Height uint64 `json:"height,string,omitempty"`
	// block nonce
	Nonce uint64 `json:"nonce,string,omitempty"`
	// Hex string of coinbase address.
	Coinbase string `json:"coinbase,omitempty"`
	// block timestamp.
	Timestamp int64 `json:"timestamp,string,omitempty"`
	// block chain id
	ChainId uint32 `json:"chain_id,omitempty"`
	// Hex string of state root.
	StateRoot string `json:"state_root,omitempty"`
	// Hex string of txs root.
	TxsRoot string `json:"txs_root,omitempty"`
	// Hex string of event root.
	EventsRoot string `json:"events_root,omitempty"`
	// Hex string of consensus root.
	ConsensusRoot *ConsensusRoot `json:"consensus_root,omitempty"`
	// Miner
	Miner string `json:"miner,omitempty"`
	// is finaliy
	IsFinality bool `json:"is_finality,omitempty"`
	// transaction slice
	Transactions []*TransactionResponse `json:"transactions,omitempty"`
}

type ConsensusRoot struct {
	Timestamp   int64  `json:"timestamp,string,omitempty"`
	Proposer    []byte `json:"proposer,omitempty"`
	DynastyRoot []byte `json:"dynasty_root,omitempty"`
}

type TransactionResponse struct {
	// Hex string of tx hash.
	Hash    string `json:"hash,omitempty"`
	ChainId uint32 `json:"chainId,omitempty"`
	// Hex string of the sender account addresss.
	From string `json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To    string `json:"to,omitempty"`
	Value string `json:"value,omitempty"`
	// Transaction nonce.
	Nonce           uint64 `json:"nonce,string,omitempty"`
	Timestamp       int64  `json:"timestamp,string,omitempty"`
	Type            string `json:"type,omitempty"`
	Data            []byte `json:"data,omitempty"`
	GasPrice        string `json:"gas_price,omitempty"`
	GasLimit        string `json:"gas_limit,omitempty"`
	ContractAddress string `json:"contract_address,omitempty"`
	// transaction status 0 failed, 1 success, 2 pending
	Status int32 `json:"status,omitempty"`
	// transaction gas used
	GasUsed string `json:"gas_used,omitempty"`
	// contract execute error
	ExecuteError string `json:"execute_error,omitempty"`
	// contract execute result
	ExecuteResult string `json:"execute_result,omitempty"`
}

type GetTransactionReceiptReq struct {
	Hash string `json:"hash"`
}

type GetTransactionByContractRequest struct {
	// string of contract address.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

type SubscribeRequest struct {
	Topics []string `protobuf:"bytes,1,rep,name=topics" json:"topics,omitempty"`
}

type SubscribeResponse struct {
	Topic string `json:"topic,omitempty"`
	Data  string `json:"data,omitempty"`
}

type GasPriceResponse struct {
	GasPrice string `json:"gas_price,omitempty"`
}

type GasResponse struct {
	Gas string `json:"gas,omitempty"`
	Err string `json:"err,omitempty"`
}

type HashRequest struct {
	// Hex string of block/transaction hash.
	Hash string `json:"hash,omitempty"`
}

type EventsResponse struct {
	Events []*Event `json:"events,omitempty"`
}

type Event struct {
	Topic string `json:"topic,omitempty"`
	Data  string `json:"data,omitempty"`
}

type ByBlockHeightRequest struct {
	Height uint64 `json:"height,omitempty"`
}

type GetDynastyResponse struct {
	Miners []string `json:"miners,omitempty"`
}

// Response message of node info.
type NodeInfoResponse struct {
	// the node ID.
	Id string `json:"id,omitempty"`
	// the block chainID.
	ChainId uint32 `json:"chain_id,omitempty"`
	// coinbase
	Coinbase string `json:"coinbase,omitempty"`
	// Number of peers currenly connected.
	PeerCount uint32 `json:"peer_count,omitempty"`
	// the node synchronized status.
	Synchronized bool `json:"synchronized,omitempty"`
	// the node route table bucket size.
	BucketSize int32 `json:"bucket_size,omitempty"`
	// the network protocol version.
	ProtocolVersion string        `json:"protocol_version,omitempty"`
	RouteTable      []*RouteTable `json:"route_table,omitempty"`
}

type AccountsResponse struct {
	// Account list
	Addresses []string `json:"addresses,omitempty"`
}

type RouteTable struct {
	Id      string   `json:"id,omitempty"`
	Address []string `json:"address,omitempty"`
}

type NewAccountRequest struct {
	Passphrase string `json:"passphrase"`
}

type NewAccountResponse struct {
	Address string `json:"address,omitempty"`
}

type UnlockAccountRequest struct {
	Address    string `json:"address"`
	Passphrase string `json:"passphrase"`
	Duration   uint64 `json:"duration"`
}

type UnlockAccountResponse struct {
	Result bool `json:"result,omitempty"`
}

type LockAccountRequest struct {
	Address string `json:"address"`
}

type LockAccountResponse struct {
	Result bool `json:"result,omitempty"`
}

type SignHashRequest struct {
	// sign address
	Address string `json:"address,omitempty"`
	// sign msg
	Hash string `json:"hash,omitempty"`
	// sign algorithm
	Alg uint32 `json:"alg,omitempty"`
}

type SignHashResponse struct {
	Data []byte `json:"data,omitempty"`
}

type TransactionRequest struct {
	// Hex string of the sender account addresss.
	From string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	// Hex string of the receiver account addresss.
	To string `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	// Amount of value sending with this transaction.
	Value string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	// Transaction nonce.
	Nonce uint64 `protobuf:"varint,4,opt,name=nonce,proto3" json:"nonce,omitempty"`
	// gasPrice sending with this transaction.
	GasPrice string `protobuf:"bytes,5,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price,omitempty"`
	// gasLimit sending with this transaction.
	GasLimit string `protobuf:"bytes,6,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit,omitempty"`
	// contract sending with this transaction
	Contract *ContractRequest `protobuf:"bytes,7,opt,name=contract" json:"contract,omitempty"`
	// binary data for transaction
	Binary []byte `protobuf:"bytes,10,opt,name=binary,proto3" json:"binary,omitempty"`
	// transaction payload type, enum:binary, deploy, call
	Type string `protobuf:"bytes,20,opt,name=type,proto3" json:"type,omitempty"`
}

type ContractRequest struct {
	// contract source code.
	Source string `protobuf:"bytes,1,opt,name=source,proto3" json:"source,omitempty"`
	// contract source type, support javascript and typescript
	SourceType string `protobuf:"bytes,2,opt,name=source_type,json=sourceType,proto3" json:"source_type,omitempty"`
	// call contract function name
	Function string `protobuf:"bytes,3,opt,name=function,proto3" json:"function,omitempty"`
	// the params of contract.
	Args string `protobuf:"bytes,4,opt,name=args,proto3" json:"args,omitempty"`
}

type SendTransactionResponse struct {
	// Hex string of transaction hash.
	Txhash string `protobuf:"bytes,1,opt,name=txhash,proto3" json:"txhash,omitempty"`
	// Hex string of contract address if transaction is deploy type
	ContractAddress string `protobuf:"bytes,2,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty"`
}

type SignTransactionPassphraseRequest struct {
	// transaction struct
	Transaction *TransactionRequest `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	// from account passphrase
	Passphrase string `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
}

type SignTransactionPassphraseResponse struct {
	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

type SendTransactionPassphraseRequest struct {
	// transaction struct
	Transaction *TransactionRequest `protobuf:"bytes,1,opt,name=transaction" json:"transaction,omitempty"`
	// from account passphrase
	Passphrase string `protobuf:"bytes,2,opt,name=passphrase,proto3" json:"passphrase,omitempty"`
}

type PprofRequest struct {
	Listen string `json:"listen,omitempty"`
}

type PprofResponse struct {
	Result bool `json:"result,omitempty"`
}

type GetConfigResponse struct {
	// Config
	Config Config `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
}

type Config struct {
	// Network config.
	Network *NetworkConfig `protobuf:"bytes,1,opt,name=network" json:"network"`
	// Chain config.
	Chain *ChainConfig `protobuf:"bytes,2,opt,name=chain" json:"chain"`
	// RPC config.
	Rpc *RPCConfig `protobuf:"bytes,3,opt,name=rpc" json:"rpc"`
	// Stats config.
	Stats *StatsConfig `protobuf:"bytes,100,opt,name=stats" json:"stats"`
	// Misc config.
	Misc *MiscConfig `protobuf:"bytes,101,opt,name=misc" json:"misc"`
	// App Config.
	App *AppConfig `protobuf:"bytes,102,opt,name=app" json:"app"`
}

type NetworkConfig struct {
	// Neb seed node address.
	Seed []string `protobuf:"bytes,1,rep,name=seed" json:"seed"`
	// Listen addresses.
	Listen []string `protobuf:"bytes,2,rep,name=listen" json:"listen"`
	// Network node privateKey address. If nil, generate a new node.
	PrivateKey string `protobuf:"bytes,3,opt,name=private_key,json=privateKey,proto3" json:"private_key"`
	// Network ID
	NetworkId            uint32 `protobuf:"varint,4,opt,name=network_id,json=networkId,proto3" json:"network_id"`
	StreamLimits         int32  `protobuf:"varint,5,opt,name=stream_limits,json=streamLimits,proto3" json:"stream_limits"`
	ReservedStreamLimits int32  `protobuf:"varint,6,opt,name=reserved_stream_limits,json=reservedStreamLimits,proto3" json:"reserved_stream_limits"`
}

type ChainConfig struct {
	// ChainID.
	ChainId uint32 `protobuf:"varint,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id"`
	// genesis conf file path
	Genesis string `protobuf:"bytes,2,opt,name=genesis,proto3" json:"genesis"`
	// Data dir.
	Datadir string `protobuf:"bytes,11,opt,name=datadir,proto3" json:"datadir"`
	// Key dir.
	Keydir string `protobuf:"bytes,12,opt,name=keydir,proto3" json:"keydir"`
	// Start mine at launch
	StartMine bool `protobuf:"varint,20,opt,name=start_mine,json=startMine,proto3" json:"start_mine"`
	// Coinbase.
	Coinbase string `protobuf:"bytes,21,opt,name=coinbase,proto3" json:"coinbase"`
	// Miner.
	Miner string `protobuf:"bytes,22,opt,name=miner,proto3" json:"miner"`
	// Passphrase.
	Passphrase string `protobuf:"bytes,23,opt,name=passphrase,proto3" json:"passphrase"`
	// Enable remote sign server
	EnableRemoteSignServer bool `protobuf:"varint,24,opt,name=enable_remote_sign_server,json=enableRemoteSignServer,proto3" json:"enable_remote_sign_server"`
	// Remote sign server
	RemoteSignServer string `protobuf:"bytes,25,opt,name=remote_sign_server,json=remoteSignServer,proto3" json:"remote_sign_server"`
	// Lowest GasPrice.
	GasPrice string `protobuf:"bytes,26,opt,name=gas_price,json=gasPrice,proto3" json:"gas_price"`
	// Max GasLimit.
	GasLimit string `protobuf:"bytes,27,opt,name=gas_limit,json=gasLimit,proto3" json:"gas_limit"`
	// Supported signature cipher list. ["ECC_SECP256K1"]
	SignatureCiphers   []string `protobuf:"bytes,28,rep,name=signature_ciphers,json=signatureCiphers" json:"signature_ciphers"`
	SuperNode          bool     `protobuf:"varint,30,opt,name=super_node,json=superNode,proto3" json:"super_node"`
	UnsupportedKeyword string   `protobuf:"bytes,31,opt,name=unsupported_keyword,json=unsupportedKeyword,proto3" json:"unsupported_keyword"`
}

type RPCConfig struct {
	// RPC listen addresses.
	RpcListen []string `protobuf:"bytes,1,rep,name=rpc_listen,json=rpcListen" json:"rpc_listen"`
	// HTTP listen addresses.
	HttpListen []string `protobuf:"bytes,2,rep,name=http_listen,json=httpListen" json:"http_listen"`
	// Enabled HTTP modules.["api", "admin"]
	HttpModule       []string `protobuf:"bytes,3,rep,name=http_module,json=httpModule" json:"http_module"`
	ConnectionLimits int32    `protobuf:"varint,4,opt,name=connection_limits,json=connectionLimits,proto3" json:"connection_limits"`
	HttpLimits       int32    `protobuf:"varint,5,opt,name=http_limits,json=httpLimits,proto3" json:"http_limits"`
	// HTTP CORS allowed origins
	HttpCors []string `protobuf:"bytes,6,rep,name=http_cors,json=httpCors" json:"http_cors"`
}

type StatsConfig struct {
	// Enable metrics or not.
	EnableMetrics   bool                          `protobuf:"varint,1,opt,name=enable_metrics,json=enableMetrics,proto3" json:"enable_metrics"`
	ReportingModule []StatsConfig_ReportingModule `protobuf:"varint,2,rep,packed,name=reporting_module,json=reportingModule,enum=nebletpb.StatsConfig_ReportingModule" json:"reporting_module"`
	// Influxdb config.
	Influxdb    *InfluxdbConfig `protobuf:"bytes,11,opt,name=influxdb" json:"influxdb"`
	MetricsTags []string        `protobuf:"bytes,12,rep,name=metrics_tags,json=metricsTags" json:"metrics_tags"`
}

type MiscConfig struct {
	// Default encryption ciper when create new keystore file.
	DefaultKeystoreFileCiper string `protobuf:"bytes,1,opt,name=default_keystore_file_ciper,json=defaultKeystoreFileCiper,proto3" json:"default_keystore_file_ciper"`
}

type AppConfig struct {
	LogLevel string `protobuf:"bytes,1,opt,name=log_level,json=logLevel,proto3" json:"log_level"`
	LogFile  string `protobuf:"bytes,2,opt,name=log_file,json=logFile,proto3" json:"log_file"`
	// log file age, unit is s.
	LogAge            uint32 `protobuf:"varint,3,opt,name=log_age,json=logAge,proto3" json:"log_age"`
	EnableCrashReport bool   `protobuf:"varint,4,opt,name=enable_crash_report,json=enableCrashReport,proto3" json:"enable_crash_report"`
	CrashReportUrl    string `protobuf:"bytes,5,opt,name=crash_report_url,json=crashReportUrl,proto3" json:"crash_report_url"`
	// pprof config
	Pprof   *PprofConfig `protobuf:"bytes,6,opt,name=pprof" json:"pprof"`
	Version string       `protobuf:"bytes,100,opt,name=version,proto3" json:"version"`
}

type StatsConfig_ReportingModule int32

type InfluxdbConfig struct {
	// Host.
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host"`
	// Port.
	Port uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port"`
	// Database name.
	Db string `protobuf:"bytes,3,opt,name=db,proto3" json:"db"`
	// Auth user.
	User string `protobuf:"bytes,4,opt,name=user,proto3" json:"user"`
	// Auth password.
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password"`
}

type PprofConfig struct {
	// pprof listen address, if not configured, the function closes.
	HttpListen string `protobuf:"bytes,1,opt,name=http_listen,json=httpListen,proto3" json:"http_listen"`
	// cpu profiling file, if not configured, the profiling not start
	Cpuprofile string `protobuf:"bytes,2,opt,name=cpuprofile,proto3" json:"cpuprofile"`
	// memory profiling file, if not configured, the profiling not start
	Memprofile string `protobuf:"bytes,3,opt,name=memprofile,proto3" json:"memprofile"`
}
