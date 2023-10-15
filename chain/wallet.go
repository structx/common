// Package chain provides a blockchain interface.
package chain

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.dedis.ch/kyber/v3"
	"go.dedis.ch/kyber/v3/group/edwards25519"

	"golang.org/x/crypto/sha3"
)

// Wallet is a wallet interface.
type Wallet struct {
	PublicKey  kyber.Point
	privateKey kyber.Scalar
}

type exportWallet struct {
	PublicKey  []byte `json:"public_key"`
	PrivateKey []byte `json:"private_key"`
}

// A basic, verifiable signature
type basicSign struct {
	C kyber.Scalar // challenge
	R kyber.Scalar // response
}

// NewWallet returns a new Wallet instance.
func NewWallet(suite *edwards25519.SuiteEd25519) *Wallet {

	a := suite.Scalar().Pick(suite.RandomStream()) // private key
	A := suite.Point().Mul(a, nil)                 // public key

	return &Wallet{
		privateKey: a,
		PublicKey:  A,
	}
}

// NewWalletFromFile returns a new Wallet instance from a file.
func NewWalletFromFile(suite *edwards25519.SuiteEd25519) (*Wallet, error) {

	path := os.Getenv("WALLET_PATH")
	if path == "" {
		return nil, fmt.Errorf("WALLET_PATH is not set")
	}

	filebytes, err := os.ReadFile(filepath.Clean(path))
	if err != nil && err == os.ErrNotExist {
		return nil, os.ErrNotExist
	} else if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var wallet exportWallet
	if err := json.Unmarshal(filebytes, &wallet); err != nil {
		return nil, fmt.Errorf("failed to unmarshal wallet: %w", err)
	}

	a := suite.Scalar()
	A := suite.Point()

	err = a.UnmarshalBinary(wallet.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal private key: %w", err)
	}

	err = A.UnmarshalBinary(wallet.PublicKey)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal public key: %w", err)
	}

	return &Wallet{
		privateKey: a,
		PublicKey:  A,
	}, nil
}

// GetAddress returns the wallet address.
func (w *Wallet) GetAddress() (string, error) {

	publicbytes, err := w.PublicKey.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	h := sha3.New224()
	_, err = h.Write(publicbytes)
	if err != nil {
		return "", fmt.Errorf("failed to write public key: %w", err)
	}

	hash := h.Sum(nil)

	return hex.EncodeToString(hash), nil
}

// Signature returns the wallet signature.
func (w *Wallet) Signature(suite *edwards25519.SuiteEd25519) (string, error) {

	publicbytes, err := w.PublicKey.MarshalBinary()
	if err != nil {
		return "", fmt.Errorf("failed to marshal public key: %w", err)
	}

	rand := suite.XOF(publicbytes)

	v := suite.Scalar().Pick(rand)
	T := suite.Point().Mul(v, nil)

	addr, err := w.GetAddress()
	if err != nil {
		return "", fmt.Errorf("failed to get address: %w", err)
	}

	c, err := hashSchnorr(suite, []byte(addr), T)
	if err != nil {
		return "", fmt.Errorf("failed to hash schnorr: %w", err)
	}

	r := suite.Scalar()
	r.Mul(w.privateKey, c).Sub(v, r)

	buf := bytes.Buffer{}
	sig := &basicSign{C: c, R: r}
	err = suite.Write(&buf, &sig)
	if err != nil {
		return "", fmt.Errorf("failed to write signature: %w", err)
	}

	return hex.EncodeToString(buf.Bytes()), nil
}

// MarshalToFile marshals the wallet to a file.
func (w *Wallet) MarshalToFile(path string) error {

	publicbytes, err := w.PublicKey.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal public key: %w", err)
	}

	privatebytes, err := w.privateKey.MarshalBinary()
	if err != nil {
		return fmt.Errorf("failed to marshal private key: %w", err)
	}

	wallet := exportWallet{
		PublicKey:  publicbytes,
		PrivateKey: privatebytes,
	}

	filebytes, err := json.Marshal(wallet)
	if err != nil {
		return fmt.Errorf("failed to marshal wallet: %w", err)
	}

	err = os.WriteFile(filepath.Clean(path), filebytes, 0600)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func hashSchnorr(suite *edwards25519.SuiteEd25519, message []byte, p kyber.Point) (kyber.Scalar, error) {

	pb, err := p.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal public key: %w", err)
	}

	c := suite.XOF(pb)
	_, err = c.Write(message)
	if err != nil {
		return nil, fmt.Errorf("failed to write message: %w", err)
	}

	return suite.Scalar().Pick(c), nil
}
