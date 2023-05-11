package test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	metax "github.com/oblzh/bcnmy-go/metax"
)

func TestGetDappAPI(t *testing.T) {
	b, _ := metax.NewBcnmy(
		os.Getenv("httpRpc"),
		os.Getenv("apiKey"),
		10*time.Second,
	)

	resp, _ := b.GetDappAPI(context.Background())
	fmt.Printf("%+v\n", resp.Dapp)
}
