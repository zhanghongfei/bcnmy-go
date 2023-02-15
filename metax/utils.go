package metax

import (
	"bytes"

	apitypes "github.com/ethereum/go-ethereum/signer/core/apitypes"
)

func formatTypedData(typedData apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := typedData.HashStruct(EIP712Domain, typedData.Domain.Map())
	if err != nil {
		return nil, err
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return nil, err
	}
	return []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash))), nil
}
