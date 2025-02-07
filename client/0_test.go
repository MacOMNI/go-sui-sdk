package client

import (
	"os"
	"testing"

	"github.com/coming-chat/go-sui/account"
	"github.com/stretchr/testify/require"
)

const DevnetRpcUrl = "https://fullnode.devnet.sui.io"
const TestnetRpcUrl = "https://fullnode.testnet.sui.io"

var M1Mnemonic = os.Getenv("WalletSdkTestM1")

func TestnetClient(t *testing.T) *Client {
	c, err := Dial(TestnetRpcUrl)
	require.Nil(t, err)
	return c
}

func DevnetClient(t *testing.T) *Client {
	c, err := Dial(DevnetRpcUrl)
	require.Nil(t, err)
	return c
}

func M1Account(t *testing.T) *account.Account {
	a, err := account.NewAccountWithMnemonic(M1Mnemonic)
	require.Nil(t, err)
	return a
}
