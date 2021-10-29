package factory

import "context"

type Client interface {
	Connect() error
	QueryTx(tx string) (interface{}, error)
	Broadcast() error
}

type TxBuilder interface {
	BuildTx() (string, error)
	SignTx() (string, error)
}

type Manager interface {
	GetClient() Client
	GetTxBuilder() TxBuilder
}

type BitcoinClient struct {
	ctx context.Context
	tx 	TxBuilder
}

func (c *BitcoinClient) Connect() error {
	return nil
}

func (c *BitcoinClient) QueryTx(tx string) (interface{}, error) {
	return nil, nil
}

func (c *BitcoinClient) Broadcast() error {
	return nil
}

type EthereumClient struct {}

func (c *EthereumClient) Connect() error {
	return nil
}

func (c *EthereumClient) QueryTx(tx string) (interface{}, error) {
	return nil, nil
}

func (c *EthereumClient) Broadcast() error {
	return nil
}

type BitcoinTxBuilder struct {}

func (b *BitcoinTxBuilder) BuildTx() (string, error) {
	return "", nil
}

func (b *BitcoinTxBuilder) SignTx() (string, error) {
	return "", nil
}

type EthereumTxBuilder struct {}

func (b *EthereumTxBuilder) BuildTx() (string, error) {
	return "", nil
}

func (b *EthereumTxBuilder) SignTx() (string, error) {
	return "", nil
}

type BitcoinManager struct {}

func (m *BitcoinManager) GetClient() Client {
	return &BitcoinClient{}
}

func (m *BitcoinManager) GetTxBuilder() TxBuilder {
	return &BitcoinTxBuilder{}
}

type EthereumManager struct {}

func (m *EthereumManager) GetClient() Client {
	return &EthereumClient{}
}

func (m *EthereumManager) GetTxBuilder() TxBuilder {
	return &EthereumTxBuilder{}
}
