package chain_test

import (
	"fmt"
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

func TestWalletSuite(t *testing.T) {
	suite.Run(t, new(WalletSuite))
}
