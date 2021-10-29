package factory

type MethodChain interface {
	GetChainID() 	string
	GetChainType() 	string
	Connect() 		error
}

type MethodFactory interface {
	Create()
}

type BaseChain struct {}

// 抽象 connect 可能并不贴切
func (c *BaseChain) Connect() error {
	return nil
}

// 继承 base chain
type BitcoinChain struct {
	BaseChain
}

func (c *BitcoinChain) GetChainID() string {
	return "main"
}

func (c *BitcoinChain) GetChainType() string {
	return "bitcoin"
}

type EthereumChain struct {
	BaseChain
}

func (c *EthereumChain) GetChainID() string {
	return "37"
}

func (c *EthereumChain) GetChainType() string {
	return "ethereum"
}

type BitcoinFactory struct {}

func (c *BitcoinFactory) Create() MethodChain {
	return &BitcoinChain{}
}

type EthereumFactory struct {}

func (c *EthereumFactory) Create() MethodChain {
	return &EthereumChain{}
}
