package metax

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"

	"github.com/oblzh/bcnmy-go/abi/forwarder"
)

type Bcnmy struct {
	ctx    context.Context
	logger *logrus.Entry

	ethClient  *ethclient.Client
	httpClient *http.Client

	// DAPP abi and address
	abi     abi.ABI
	address common.Address

	authToken string
	apiKey    string
	/// method apiID
	apiID map[string]struct {
		ID              string
		ContractAddress string
	}

	batchId *big.Int
	chainId *big.Int

	trustedForwarder struct {
		Address  common.Address
		Contract *forwarder.Forwarder
	}
}

func NewBcnmy(httpRpc string, apiKey string) (*Bcnmy, error) {
	var err error
	bcnmy := &Bcnmy{
		ctx:    context.Background(),
		logger: logrus.WithField("metax", "bcnmy"),
		apiKey: apiKey,
		apiID: make(map[string]struct {
			ID              string
			ContractAddress string
		}),
		batchId:    big.NewInt(0),
		httpClient: &http.Client{Timeout: 10 * time.Second},
	}
	bcnmy.ethClient, err = ethclient.DialContext(bcnmy.ctx, httpRpc)
	if err != nil {
		bcnmy.logger.WithError(err).Error("DialContext ethclient failed")
		return nil, err
	}
	bcnmy.chainId, err = bcnmy.ethClient.ChainID(bcnmy.ctx)
	if err != nil {
		bcnmy.logger.WithError(err).Error("ethClient getchainId failed")
		return nil, err
	}

	forwarderAddress, ok := ForwarderAddressMap[bcnmy.chainId.String()]
	if !ok {
		err = fmt.Errorf("Chain ID not supported: %v", bcnmy.chainId)
		bcnmy.logger.Error(err.Error())
		return nil, err
	}

	forwarderContract, err := forwarder.NewForwarder(forwarderAddress, bcnmy.ethClient)
	if err != nil {
		bcnmy.logger.WithError(err).Error("Load Forwarder Contract failed")
		return nil, err
	}

	bcnmy.trustedForwarder = struct {
		Address  common.Address
		Contract *forwarder.Forwarder
	}{
		Address:  forwarderAddress,
		Contract: forwarderContract,
	}
	resp, err := bcnmy.GetMetaAPI(bcnmy.ctx)
	if err != nil {
		bcnmy.logger.WithError(err).Error(err.Error())
		return nil, err
	}
	for _, info := range resp.ListAPI {
		// filter non contractAddress
		if common.IsHexAddress(info.ContractAddress) {
			bcnmy.apiID[info.Method] = struct {
				ID              string
				ContractAddress string
			}{
				ID:              info.ID,
				ContractAddress: info.ContractAddress,
			}
		}
	}
	return bcnmy, nil
}

func (b *Bcnmy) WithDapp(jsonABI string, dappAddress common.Address) (*Bcnmy, error) {
	var err error
	b.address = dappAddress
	b.abi, err = abi.JSON(strings.NewReader(jsonABI))
	if err != nil {
		b.logger.WithError(err).Error("jsonABI parse failed")
		return nil, err
	}
	return b, nil
}

func (b *Bcnmy) WithAuthToken(authToken string) *Bcnmy {
	b.authToken = authToken
	return b
}

func (b *Bcnmy) WithFieldTimeout(timeout time.Duration) *Bcnmy {
	b.httpClient = &http.Client{Timeout: timeout}
	return b
}

func (b *Bcnmy) GetAuthorization() string {
	return fmt.Sprintf("User %s", b.authToken)
}
