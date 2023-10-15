package chain_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"go.dedis.ch/kyber/v3/group/edwards25519"

	"github.com/structx/common/chain"
)

type WalletSuite struct {
	suite.Suite
	wallet *chain.Wallet
}

func (s *WalletSuite) SetupTest() {

	suite := edwards25519.NewBlakeSHA256Ed25519()

	s.wallet = chain.NewWallet(suite)
}

func (s *WalletSuite) TestAddress() {

	assert := s.Assert()

	addr, err := s.wallet.GetAddress()
	assert.NoError(err)

	fmt.Println(addr)
}

func (s *WalletSuite) TestSignature() {

	assert := s.Assert()

	suite := edwards25519.NewBlakeSHA256Ed25519()

	addr, err := s.wallet.Signature(suite)
	assert.NoError(err)

	fmt.Println(addr)
}

func (s *WalletSuite) TestMarshalToFile() {

	assert := s.Assert()

	err := s.wallet.MarshalToFile("./testfiles/wallet.json")
	assert.NoError(err)
}

func (s *WalletSuite) TestNewFromFile() {

	assert := s.Assert()

	_ = os.Setenv("WALLET_PATH", "./testfiles/wallet.json")

	suite := edwards25519.NewBlakeSHA256Ed25519()

	_, err := chain.NewWalletFromFile(suite)
	assert.NoError(err)
}

func TestWalletSuite(t *testing.T) {
	suite.Run(t, new(WalletSuite))
}
