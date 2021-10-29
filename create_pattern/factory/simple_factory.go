package factory

import "fmt"

type SimpleChain interface {
	GetChainID() string
}

type ChainFactory struct {}

type BitcoinContext struct {}

func NewBitcoinContext() *BitcoinContext {
	return &BitcoinContext{}
}

func (c *BitcoinContext) GetChainID() string {
	return "main"
}

type EthereumContext struct {}

func NewEthereumContext() *EthereumContext {
	return &EthereumContext{}
}

func (c *EthereumContext) GetChainID() string {
	return "37"
}

func SimpleFactory(ty string) (SimpleChain, error) {
	switch ty {
	case "bitcoin": // convert to enum
		return NewBitcoinContext(), nil
	case "ethereum":
		return NewEthereumContext(), nil
	default:
		return nil, fmt.Errorf("unsupport chain type: %s", ty)
	}
}
