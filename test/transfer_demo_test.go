package test

import (
	"fmt"
	"math/big"
	//"encoding/hex"

	"os"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	"github.com/stretchr/testify/assert"

	demo "github.com/oblzh/bcnmy-go/abi/demo"
	metax "github.com/oblzh/bcnmy-go/metax"
)

func buildBcnmy() *metax.Bcnmy {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), time.Second*10)
	b = b.WithAuthToken(os.Getenv("authToken"))
	b = b.WithFieldTimeout(time.Second * 60)
	return b
}

// Finished https://mumbai.polygonscan.com/tx/0x39b3ed93123d9c45583cd6c68c72943fb13c8f72d489deb00b96a02a8fd21745
func TestTransferDemo(t *testing.T) {
	b := buildBcnmy()
	b.WithDapp(demo.TransferDemoABI, common.HexToAddress("0x56B71565F6e7f9dE4c3217A6E5d4133bc7fc67EB"))

	metaTxMessage := &metax.MetaTxMessage{
		From:          common.HexToAddress("0xD1cc56810a3947d1D8b05448afB9889c6cFCF0F1"),
		To:            common.HexToAddress("0x56B71565F6e7f9dE4c3217A6E5d4133bc7fc67EB"),
		Token:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
		TxGas:         150000,
		TokenGasPrice: "0",
		BatchId:       big.NewInt(0),
		BatchNonce:    big.NewInt(13),
		Deadline:      big.NewInt(1683630490),
		Data:          "0x71234eb00000000000000000000000006a22dda833c14ca6189f32e0dbcdf41ac2a3c951000000000000000000000000c015fb756fd4d49c6280eca2d47df30e8f6d083100000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000000000000000000000000000000900000000000000000000000000000000000000000000000000000000645a299a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001bad55851ffb8bc0549bb7b6abf5ffa1676a0049eb7d8550d702a2d14f6e4522ef583bd957b462b550621bc762b52df9a7cbb48b36b4ff4da7dc168e3c447dff56",
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
	typedDataHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	signature := hexutil.MustDecode("0x6038d766c89900874f7a25cfa590ef305fcd2dd3d6985a4efdb8d4f4204cbb02596d00c9534bfa2f8572638076ca36f97f6b6ca57176a4af65c858c4e228ecb71b")
	fmt.Println(signature)
	resp, txn, _, err := b.EnhanceTransact(
		common.HexToAddress("0xD1cc56810a3947d1D8b05448afB9889c6cFCF0F1").Hex(),
		"permitEIP2612AndTransfer",
		signature,
		metaTxMessage,
		typedDataHash.String(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.NotNil(t, txn)
	fmt.Println(txn)
}

//// Finished
//func TestTransferDemo(t *testing.T) {
//b := buildBcnmy()
//b.WithDapp(demo.TransferDemoABI, common.HexToAddress("0x56B71565F6e7f9dE4c3217A6E5d4133bc7fc67EB"))

//metaTxMessage := &metax.MetaTxMessage{
//From:          common.HexToAddress("0x8595492B1195Dd553f3B87114C2C8c900e8cdCcF"),
//To:            common.HexToAddress("0x56B71565F6e7f9dE4c3217A6E5d4133bc7fc67EB"),
//Token:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
//TxGas:         4000000,
//TokenGasPrice: "0",
//BatchId:       big.NewInt(0),
//BatchNonce:    big.NewInt(1),
//Deadline:      big.NewInt(1681393939),
//Data:          "0xbeabacc80000000000000000000000003d27d0f52803886b43002b8845b27f3738c85b41000000000000000000000000c6972a28ddf68c75b93eea7b771464612ec8f9990000000000000000000000000000000000000000000000008ac7230489e80000",
//}

//typedData := apitypes.TypedData{
//Types:       metax.SignedTypes,
//PrimaryType: metax.ForwardRequestType,
//Domain: apitypes.TypedDataDomain{
//Name:              metax.ForwardRequestName,
//Version:           metax.Version,
//VerifyingContract: common.HexToAddress("0x69015912AA33720b842dCD6aC059Ed623F28d9f7").Hex(),
//Salt:              hexutil.Encode(common.LeftPadBytes(big.NewInt(80001).Bytes(), 32)),
//},
//Message: metaTxMessage.TypedData(),
//}
//typedDataHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
//signature := hexutil.MustDecode("0x13f2a7ce9bbb28988afa78475d8a4cc4ab4fe3e3b1e79b0c56fbfd9875b3069928a5166afccacb8a320acc7502535851bdaafd8899d2e21ebf508fad4452f3c01b")
//fmt.Println(signature)
//txn, _, err := b.EnhanceTransact(
//common.HexToAddress("0x8595492B1195Dd553f3B87114C2C8c900e8cdCcF").Hex(),
//"transfer",
//signature,
//metaTxMessage,
//typedDataHash.String(),
//)
//assert.Nil(t, err)
//assert.NotNil(t, txn)
//}
