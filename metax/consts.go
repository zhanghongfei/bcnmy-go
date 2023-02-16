package metax

import (
	"github.com/ethereum/go-ethereum/common"
	apitypes "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

const (
	MetaAPIURL                 = "https://api.biconomy.io/api/v1/meta-api"
	MetaTxNativeURL            = "https://api.biconomy.io/api/v2/meta-tx/native"
	CreateDappPublicURL        = "https://api.biconomy.io/api/v1/dapp/public-api/create-dapp"
	AddContractURL             = "https://api.biconomy.io/api/v1/smart-contract/public-api/addContract"
	AddMethodURL               = "https://api.biconomy.io/api/v1/meta-api/public-api/addMethod"
	DeleteContractURL          = "https://api.biconomy.io/api/v1/smart-contract/public-api/deleteContract"
	DeleteMethodURL            = "https://api.biconomy.io/api/v1/meta-api/public-api/deleteMethod"
	AddDestinationAddressesURL = "https://api.biconomy.io/api/v1/dapp/whitelist/destination"
	ProxyContractsURL          = "https://api.biconomy.io/api/v1/dapp/whitelist/proxy-contracts"
	UniqueUserDataURL          = "https://data.biconomy.io/api/v1/dapp/uniqueUserData"
	UserLimitURL               = "https://data.biconomy.io/api/v1/dapp/user-limit"
	GasTankBalanceURL          = "https://data.biconomy.io/api/v1/dapp/gas-tank-balance"
)

const (
	SignatureEIP712Type = "EIP712_SIGN"
	EIP712DomainType    = "EIP712Domain"
	ForwardRequestType  = "ERC20ForwardRequest"
	ForwardRequestName  = "Biconomy Forwarder"
	Version             = "1"
)

var SignedTypes = apitypes.Types{
	"EIP712Domain": []apitypes.Type{
		{Name: "name", Type: "string"},
		{Name: "version", Type: "string"},
		{Name: "verifyingContract", Type: "address"},
		{Name: "salt", Type: "bytes32"},
	},
	"ERC20ForwardRequest": []apitypes.Type{
		{Name: "from", Type: "address"},
		{Name: "to", Type: "address"},
		{Name: "token", Type: "address"},
		{Name: "txGas", Type: "uint256"},
		{Name: "tokenGasPrice", Type: "uint256"},
		{Name: "batchId", Type: "uint256"},
		{Name: "batchNonce", Type: "uint256"},
		{Name: "deadline", Type: "uint256"},
		{Name: "data", Type: "bytes"},
	},
}

var ForwarderAddressMap = map[string]common.Address{
	"1":      common.HexToAddress("0x84a0856b038eaAd1cC7E297cF34A7e72685A8693"), // Ethereum mainnet
	"3":      common.HexToAddress("0x3D1D6A62c588C1Ee23365AF623bdF306Eb47217A"), // Ropsten testnet
	"4":      common.HexToAddress("0xFD4973FeB2031D4409fB57afEE5dF2051b171104"), // Rinkeby testnet
	"5":      common.HexToAddress("0xE041608922d06a4F26C0d4c27d8bCD01daf1f792"), // Goerli testnet
	"42":     common.HexToAddress("0xF82986F574803dfFd9609BE8b9c7B92f63a1410E"), // Kovan testnet
	"56":     common.HexToAddress("0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8"), // BSC mainnet
	"97":     common.HexToAddress("0x61456BF1715C1415730076BB79ae118E806E74d2"), // Binance testnet
	"100":    common.HexToAddress("0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8"), // xDai mainnet
	"137":    common.HexToAddress("0x86C80a8aa58e0A4fa09A69624c31Ab2a6CAD56b8"), // Matic mainnet
	"250":    common.HexToAddress("0x64CD353384109423a966dCd3Aa30D884C9b2E057"), // Fantom mainnet
	"1287":   common.HexToAddress("0x3AF14449e18f2c3677bFCB5F954Dc68d5fb74a75"), // Moonbeam alpha testnet
	"4002":   common.HexToAddress("0x69FB8Dca8067A5D38703b9e8b39cf2D51473E4b4"), // Fantom testnet
	"80001":  common.HexToAddress("0x9399BB24DBB5C4b782C70c2969F58716Ebbd6a3b"), // Matic testnet mumbai
	"42161":  common.HexToAddress("0xfe0fa3C06d03bDC7fb49c892BbB39113B534fB57"), // Arbitrum mainnet
	"421611": common.HexToAddress("0x67454E169d613a8e9BA6b06af2D267696EAaAf41"), // Arbitrum test

}
