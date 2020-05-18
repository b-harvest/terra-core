// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/terra-project/core/x/wasm/internal/types/
// ALIASGEN: github.com/terra-project/core/x/wasm/internal/keeper/
package wasm

import (
	"github.com/terra-project/core/x/wasm/internal/keeper"
	"github.com/terra-project/core/x/wasm/internal/types"
)

const (
	ModuleName                = types.ModuleName
	StoreKey                  = types.StoreKey
	TStoreKey                 = types.TStoreKey
	QuerierRoute              = types.QuerierRoute
	RouterKey                 = types.RouterKey
	MaxWasmSize               = types.MaxWasmSize
	WasmMsgParserRouteBank    = types.WasmMsgParserRouteBank
	WasmMsgParserRouteStaking = types.WasmMsgParserRouteStaking
	WasmMsgParserRouteWasm    = types.WasmMsgParserRouteWasm
	DefaultParamspace         = types.DefaultParamspace
	QueryGetByteCode          = types.QueryGetByteCode
	QueryGetCodeInfo          = types.QueryGetCodeInfo
	QueryGetContractInfo      = types.QueryGetContractInfo
	QueryRawStore             = types.QueryRawStore
	QueryContractStore        = types.QueryContractStore
	WasmQueryRouteBank        = types.WasmQueryRouteBank
	WasmQueryRouteStaking     = types.WasmQueryRouteStaking
	WasmQueryRouteWasm        = types.WasmQueryRouteWasm
)

var (
	// functions aliases
	RegisterCodec                 = types.RegisterCodec
	DefaultWasmConfig             = types.DefaultWasmConfig
	EncodeSdkCoin                 = types.EncodeSdkCoin
	EncodeSdkCoins                = types.EncodeSdkCoins
	FeatureStaking                = types.FeatureStaking
	ParseResult                   = types.ParseResult
	ParseToCoin                   = types.ParseToCoin
	ParseToCoins                  = types.ParseToCoins
	NewCodeInfo                   = types.NewCodeInfo
	NewContractInfo               = types.NewContractInfo
	NewWasmAPIParams              = types.NewWasmAPIParams
	NewWasmCoins                  = types.NewWasmCoins
	NewGenesisState               = types.NewGenesisState
	DefaultGenesisState           = types.DefaultGenesisState
	ValidateGenesis               = types.ValidateGenesis
	GetCodeInfoKey                = types.GetCodeInfoKey
	GetContractInfoKey            = types.GetContractInfoKey
	GetContractStoreKey           = types.GetContractStoreKey
	NewMsgStoreCode               = types.NewMsgStoreCode
	NewMsgInstantiateContract     = types.NewMsgInstantiateContract
	NewMsgExecuteContract         = types.NewMsgExecuteContract
	DefaultParams                 = types.DefaultParams
	ParamKeyTable                 = types.ParamKeyTable
	NewQueryCodeIDParams          = types.NewQueryCodeIDParams
	NewQueryContractAddressParams = types.NewQueryContractAddressParams
	NewQueryRawStoreParams        = types.NewQueryRawStoreParams
	NewQueryContractParams        = types.NewQueryContractParams
	NewWasmMsgParser              = keeper.NewWasmMsgParser
	NewWasmQuerier                = keeper.NewWasmQuerier
	NewKeeper                     = keeper.NewKeeper
	NewQuerier                    = keeper.NewQuerier
	CreateTestInput               = keeper.CreateTestInput

	// variable aliases
	ModuleCdc                    = types.ModuleCdc
	ErrStoreCodeFailed           = types.ErrStoreCodeFailed
	ErrAccountExists             = types.ErrAccountExists
	ErrInstantiateFailed         = types.ErrInstantiateFailed
	ErrExecuteFailed             = types.ErrExecuteFailed
	ErrGasLimit                  = types.ErrGasLimit
	ErrInvalidGenesis            = types.ErrInvalidGenesis
	ErrNotFound                  = types.ErrNotFound
	ErrInvalidMsg                = types.ErrInvalidMsg
	ErrNoRegisteredQuerier       = types.ErrNoRegisteredQuerier
	ErrNoRegisteredParser        = types.ErrNoRegisteredParser
	LastCodeIDKey                = types.LastCodeIDKey
	LastInstanceIDKey            = types.LastInstanceIDKey
	CodeKey                      = types.CodeKey
	ContractInfoKey              = types.ContractInfoKey
	ContractStoreKey             = types.ContractStoreKey
	ParamStoreKeyMaxContractSize = types.ParamStoreKeyMaxContractSize
	ParamStoreKeyMaxContractGas  = types.ParamStoreKeyMaxContractGas
	ParamStoreKeyGasMultiplier   = types.ParamStoreKeyGasMultiplier
	DefaultMaxContractSize       = types.DefaultMaxContractSize
	DefaultMaxContractGas        = types.DefaultMaxContractGas
	DefaultGasMultiplier         = types.DefaultGasMultiplier
)

type (
	WasmConfig                 = types.WasmConfig
	WasmWrapper                = types.WasmWrapper
	Model                      = types.Model
	CodeInfo                   = types.CodeInfo
	ContractInfo               = types.ContractInfo
	AccountKeeper              = types.AccountKeeper
	BankKeeper                 = types.BankKeeper
	GenesisState               = types.GenesisState
	Code                       = types.Code
	Contract                   = types.Contract
	MsgStoreCode               = types.MsgStoreCode
	MsgInstantiateContract     = types.MsgInstantiateContract
	MsgExecuteContract         = types.MsgExecuteContract
	WasmMsgParserInterface     = types.WasmMsgParserInterface
	WasmCustomMsg              = types.WasmCustomMsg
	MsgParser                  = types.MsgParser
	Params                     = types.Params
	QueryCodeIDParams          = types.QueryCodeIDParams
	QueryContractAddressParams = types.QueryContractAddressParams
	QueryRawStoreParams        = types.QueryRawStoreParams
	QueryContractParams        = types.QueryContractParams
	WasmQuerierInterface       = types.WasmQuerierInterface
	Querier                    = types.Querier
	WasmCustomQuery            = types.WasmCustomQuery
	Keeper                     = keeper.Keeper
	InitMsg                    = keeper.InitMsg
)