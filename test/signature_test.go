package test

import (
	"fmt"
	"math/big"
	"testing"

	//ethereum "github.com/ethereum/go-ethereum"
	//"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	//"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	metax "github.com/oblzh/bcnmy-go/metax"
)

func TestSignature(t *testing.T) {
	metaTxMessage := &metax.MetaTxMessage{
		From:          common.HexToAddress("0x9ac13c0f3D120dAFf4777BbDE2C87AdD30F5F8eE"),
		To:            common.HexToAddress("0x5aF72FB56a7c0E9173B8682ca615F0290D627E59"),
		Token:         common.HexToAddress("0x0"),
		TxGas:         21000,
		TokenGasPrice: "0",
		BatchId:       big.NewInt(0),
		BatchNonce:    big.NewInt(0),
		Deadline:      big.NewInt(1681294917075),
		Data:          "0x",
	}

	typedData := apitypes.TypedData{
		Types:       metax.SignedTypes,
		PrimaryType: metax.ForwardRequestType,
		Domain: apitypes.TypedDataDomain{
			Name:              metax.ForwardRequestName,
			Version:           metax.Version,
			VerifyingContract: common.HexToAddress("0x69015912AA33720b842dCD6aC059Ed623F28d9f7").Hex(),
			Salt:              hexutil.Encode(common.LeftPadBytes(big.NewInt(80001).Bytes(), 32)),
		},
		Message: metaTxMessage.TypedData(),
	}

	hash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	fmt.Println(hash)
}
