package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	metax "github.com/oblzh/bcnmy-go/metax"
	"github.com/stretchr/testify/assert"
)

func TestCheckLimits(t *testing.T) {
	b, _ := metax.NewBcnmy(os.Getenv("httpRpc"), os.Getenv("apiKey"), time.Second*10)
	resp, err := b.CheckLimits("0x96774c64dc3f46f64d17034ce6cf7b2ef31da56a", "transfer")
	assert.Nil(t, err)
	fmt.Println(resp)
}
