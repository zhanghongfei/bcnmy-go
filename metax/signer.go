package metax

import (
	"crypto/ecdsa"
	"io/ioutil"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

type Signer struct {
	Address common.Address
	key     *ecdsa.PrivateKey
}

func NewSigner(privKey string) (*Signer, error) {
	key, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}
	return &Signer{
		key:     key,
		Address: crypto.PubkeyToAddress(key.PublicKey),
	}
}

func NewSignerFromPath(privPath string) (*Signer, error) {
	f, err := ioutil.ReadFile(privPath)
	if err != nil {
		return nil, err
	}
	return NewSigner(strings.TrimSpace(string(f)))
}

func (s *Signer) GetPublicKey() []byte {
	return crypto.FromECDSAPub(&s.key.PublicKey)
}

func (s *Signer) SignTypedData(typedData apitypes.TypedData) ([]byte, error) {
	hashData, rawData, err := apitypes.TypedDataAndHash(typedData)
	if err != nil {
		return nil, err
	}
	sig, err := crypto.Sign(hash, s.key)
	if err != nil {
		return nil, err
	}
	if sig[64] == 0 || sig[64] == 1 {
		sig[64] += 27
	}
	return sig, nil
}
