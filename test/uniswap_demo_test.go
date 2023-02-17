package test

import (
	"math/big"
	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"

	demo "github.com/oblzh/bcnmy-go/abi/demo"
	metax "github.com/oblzh/bcnmy-go/metax"
)

func buildBcnmy() *metax.Bcnmy {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"))
	b = b.WithAuthToken(os.Getenv("authToken"))
	b = b.WithFieldTimeout(time.Second * 60)
	return b
}

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

	txn, _, err := b.RawTransact(signer, "exactInputSingle", params...)
	assert.Nil(t, err)
	assert.NotNil(t, txn)
}
