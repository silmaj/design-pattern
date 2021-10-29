package struct_pattern

type Chains interface {
	SendTx(tx string) error
}

type Ethereum struct {}

// factory function of chains
func NewChain() Chains {
	return &Ethereum{}
}

func (e *Ethereum) SendTx(tx string) error {
	return nil
}

type EthAltChain interface {
	BroadCast(tx string) error
}

type BSC struct {}

func NewAltChain() EthAltChain {
	return &BSC{}
}

func (bsc *BSC) BroadCast(tx string) error {
	return nil
}

// BSCAdapter 适配 Chains 接口到 EthAltChain 接口
type BSCAdapter struct {
	BSC
}

func (bsc *BSCAdapter) SendTx(tx string) error {
	return bsc.BroadCast(tx)
}
