package create_pattern

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"time"
)

type TxBuilder interface {
	UnpackParams(input string) error
	BuildTx() string
	SignTx(privkey []byte) error
}

type BitcoinBuilder struct {
	from 	string
	to 		string
	amount 	uint64
	payload string
	sign   	string
}

func (b *BitcoinBuilder) UnpackParams(input string) error {
	err := json.Unmarshal([]byte(input), b)
	if err != nil {
		return err
	}
	return nil
}

func (b *BitcoinBuilder) BuildTx() string {
	res := make([]byte, 0)
	res = append(res, byte(len(b.from)))
	res = append(res, []byte(b.from)...)
	res = append(res, byte(len(b.to)))
	res = append(res, []byte(b.to)...)

	return hex.EncodeToString(res)
}

func (b *BitcoinBuilder) SignTx(privkey []byte) error {
	bz, err := hex.DecodeString(b.payload)
	if err != nil {
		return nil
	}
	reader := bytes.NewReader(privkey)
	pk, err := ecdsa.GenerateKey(elliptic.P256(), reader)
	if err != nil {
		return nil
	}
	bz, err = pk.Sign(rand.New(rand.NewSource(time.Now().UnixNano())), bz, nil)
	if err != nil {
		return nil
	}
	b.sign = hex.EncodeToString(bz)
	return nil
}

type EthereumBuilder struct {
	from 	string
	to 		string
	amount 	uint64
}

func (b *EthereumBuilder) UnpackParams(input string) error {
	return nil
}

func (b *EthereumBuilder) BuildTx() string {
	return ""
}

func (b *EthereumBuilder) SignTx(privkey []byte) error {
	return nil
}
