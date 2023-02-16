package test

import (
	"math/big"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	demo "github.com/oblzh/bcnmy-go/abi/demo"
	//token "github.com/oblzh/bcnmy-go/abi/token"
	metax "github.com/oblzh/bcnmy-go/metax"
)

func buildBcnmy() *metax.Bcnmy {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"))
	b = b.WithAuthToken(os.Getenv("authToken"))
	return b
}

//// Finished
//func TestDeleteContract(t *testing.T) {
//b := buildBcnmy()

//data := &metax.DeleteContractRequest{
//ContractAddress: "0x0a364431476a8d1dd475590b0a028b40686ce542",
//ContractType:    "SC",
//}

//resp, err := b.DeleteContract(data)
//assert.Nil(t, err)
//assert.NotNil(t, resp)
//assert.Equal(t, resp.Code, 143)
//}

//// Finished
//func TestDeleteMethod(t *testing.T) {
//b := buildBcnmy()

//data := &metax.DeleteMethodRequest{
//ContractAddress: "0xa6b71e26c5e0845f74c812102ca7114b6a896ab2",
//Method:          "createProxyWithNonce",
//}

//resp, err := b.DeleteMethod(data)
//assert.Nil(t, err)
//assert.NotNil(t, resp)
//assert.Equal(t, resp.Code, 143)
//}

//// Finished
//func TestAddContract(t *testing.T) {
//b := buildBcnmy()
//assert.NotNil(t, b)

//data := &metax.AddContractRequest{
//ContractName:        "TestToken",
//ContractAddress:     "0xeaC94633FFf8C65aD9EFdCF237741D931fa995Cd",
//ContractType:        "SC",
//WalletType:          "",
//MetaTransactionType: "DEFAULT",
//ABI:                 token.TestTokenABI,
//}

//resp, err := b.AddContract(data)
//assert.Nil(t, err)
//assert.NotNil(t, resp)
//assert.Equal(t, resp.Code, 200)
//}

//// Finished
//func TestAddMethod(t *testing.T) {
//b := buildBcnmy()
//assert.NotNil(t, b)

//data := &metax.AddMethodRequest{
//ContractAddress: "0xeaC94633FFf8C65aD9EFdCF237741D931fa995Cd",
//ApiType:         "custom",
//Name:            "mintTo",
//MethodType:      "write",
//Method:          "mintTo",
//}

//resp, err := b.AddMethod(data)
//assert.Nil(t, err)
//assert.NotNil(t, resp)
//assert.Equal(t, resp.Code, 200)
//}

//// Finished
//func TestCreateDapp(t *testing.T) {
//b := buildBcnmy()
//assert.NotNil(t, b)

//data := &metax.CreateDappRequest{
//DappName:             "test-create",
//NetworkId:            "5",
//EnableBiconomyWallet: false,
//}

//resp, err := b.CreateDapp(data)
//assert.Nil(t, err)
//assert.NotNil(t, resp)
//assert.Equal(t, resp.Code, 200)
//}

// Finished
func TestUniswapDemo(t *testing.T) {
	b := buildBcnmy()
	b.WithDapp(demo.UniswapDemoABI, common.HexToAddress("0x9F4E38ce68a6e09E985F03B3De9126fc55d5642f"))
	signer, err := metax.NewSigner(os.Getenv("pk5"))
	assert.Nil(t, err)

	amountIn, _ := new(big.Int).SetString("1000000000000000", 10)
	tokenIn := "0xb4fbf271143f4fbf7b91a5ded31805e42b2208d6"
	tokenOut := "0x8da8c41181084508d8e1e2c4e7652c705ba65339"
	fee := big.NewInt(3000)
	amountOutMinimum := big.NewInt(0)
	sqrtPriceLimitX96 := big.NewInt(0)

	inputSingleParams := demo.ISwapRouterV3ExactInputSingleParams{
		TokenIn:           common.HexToAddress(tokenIn),
		TokenOut:          common.HexToAddress(tokenOut),
		Fee:               fee,
		Recipient:         signer.Address,
		AmountIn:          amountIn,
		AmountOutMinimum:  amountOutMinimum,
		SqrtPriceLimitX96: sqrtPriceLimitX96,
	}
	params := []interface{}{
		signer.Address,
		inputSingleParams,
	}
	packed, err := b.Pack("exactInputSingle", params...)
	assert.Nil(t, err)
	assert.NotNil(t, packed)

	txn, err := b.RawTransact(signer, "exactInputSingle", params...)
	assert.Nil(t, err)
	assert.NotNil(t, txn)
}
