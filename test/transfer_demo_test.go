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
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), time.Second*100)
	b = b.WithAuthToken(os.Getenv("authToken"))
	b = b.WithFieldTimeout(time.Second * 60)
	return b
}

// Finished https://mumbai.polygonscan.com/tx/0x39b3ed93123d9c45583cd6c68c72943fb13c8f72d489deb00b96a02a8fd21745
// Latest Update finish bsctest https://testnet.bscscan.com/tx/0x109c20a18e95afd8d8a6502f54d8788fddc1d1ae0013e9aa4b97718a5b0c049b
//func TestTransferDemo(t *testing.T) {
//b := buildBcnmy()
//b.WithDapp(demo.TransferDemoABI, common.HexToAddress("0x26F9A493149d0518B48f0cC72F510d4CDe628181"))

//metaTxMessage := &metax.MetaTxMessage{
//From:          common.HexToAddress("0xEcA4844265429C34A8ceD84128523cA6574f7a90"),
//To:            common.HexToAddress("0x26F9A493149d0518B48f0cC72F510d4CDe628181"),
//Token:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
//TxGas:         150000,
//TokenGasPrice: "0",
//BatchId:       big.NewInt(0),
//BatchNonce:    big.NewInt(0),
//Deadline:      big.NewInt(1684116992),
//Data:          "0x71234eb000000000000000000000000067697359f94663c7b842ef1ebb9802af8146f585000000000000000000000000c015fb756fd4d49c6280eca2d47df30e8f6d083100000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000646196000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c5c8f5ee33c626a04d4ce8ec6407533b675ab8669b5668c322762f9045103b6f6667561777b2b034e2895b9396857bd93936fe63228731c7214a454deadf969cc",
//}

//typedData := apitypes.TypedData{
//Types:       metax.SignedTypes,
//PrimaryType: metax.ForwardRequestType,
//Domain: apitypes.TypedDataDomain{
//Name:              metax.ForwardRequestName,
//Version:           metax.Version,
//VerifyingContract: common.HexToAddress("0x61456BF1715C1415730076BB79ae118E806E74d2").Hex(),
//Salt:              hexutil.Encode(common.LeftPadBytes(big.NewInt(97).Bytes(), 32)),
//},
//Message: metaTxMessage.TypedData(),
//}
//typedDataHash, _ := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
//signature := hexutil.MustDecode("0xebae05b9ae439a94ff869bcdd33ac6f403f377d3b48a4f208bc30049d4203f5e4cb85b3d1849ecb559d5352e694ad3290a0535b8a23b7855c0386f42a62f9c9b1c")
//fmt.Println(signature)
//resp, txn, _, err := b.EnhanceTransact(
//common.HexToAddress("0xEcA4844265429C34A8ceD84128523cA6574f7a90").Hex(),
//"permitEIP2612AndTransfer",
//signature,
//metaTxMessage,
//typedDataHash.String(),
//)
//assert.Nil(t, err)
//assert.NotNil(t, resp)
//assert.NotNil(t, txn)
//fmt.Println(txn)
//}

// Finished
func TestTransferDemo(t *testing.T) {
	b := buildBcnmy()
	b.WithDapp(demo.TransferDemoABI, common.HexToAddress("0x56b71565f6e7f9de4c3217a6e5d4133bc7fc67eb"))

	metaTxMessage := &metax.MetaTxMessage{
		From:          common.HexToAddress("0xD1cc56810a3947d1D8b05448afB9889c6cFCF0F1"),
		To:            common.HexToAddress("0x56b71565f6e7f9de4c3217a6e5d4133bc7fc67eb"),
		Token:         common.HexToAddress("0x0000000000000000000000000000000000000000"),
		TxGas:         150000,
		TokenGasPrice: "0",
		BatchId:       big.NewInt(0),
		BatchNonce:    big.NewInt(16),
		Deadline:      big.NewInt(1684487752),
		Data:          "0x71234eb00000000000000000000000006a22dda833c14ca6189f32e0dbcdf41ac2a3c951000000000000000000000000c015fb756fd4d49c6280eca2d47df30e8f6d083100000000000000000000000000000000000000000000000000000000000186a000000000000000000000000000000000000000000000000000000000000186a0000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000000000000000000000000000000000064673e480000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001c37d998cc391ab69268976c49d13f81b59b4ec1a95a5a7ea61a792df08e2948604df00bd905299c22be3f8e4279663a95c142ae74e347d1b54392451318c00b7a",
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
	signature := hexutil.MustDecode("0x74d2bebb662b388748a0a8ccce91c0bc2b998a15cb5b0069c6c5a8aec0c0673f4a4d0fc552814bd44c85a46d983f9a1873dff01348e0739b31a5602ebd4842561b")
	fmt.Println(signature)
	_, txn, _, err := b.EnhanceTransact(
		common.HexToAddress("0xD1cc56810a3947d1D8b05448afB9889c6cFCF0F1").Hex(),
		"permitEIP2612AndTransfer",
		signature,
		metaTxMessage,
		typedDataHash.String(),
	)
	assert.Nil(t, err)
	assert.NotNil(t, txn)
}
