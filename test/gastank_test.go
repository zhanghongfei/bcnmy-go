package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	metax "github.com/oblzh/bcnmy-go/metax"
	"github.com/stretchr/testify/assert"
)

func TestGasTankBalance(t *testing.T) {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), time.Second*10)
	err := b.WithBackend(os.Getenv("email"), os.Getenv("passwd"), time.Second*10)
	assert.Nil(t, err)

	resp, err := b.GetGasTankEffectiveBalance()
	assert.Nil(t, err)
	fmt.Println(resp)
}
