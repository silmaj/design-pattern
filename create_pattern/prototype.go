package create_pattern

type Chain interface {
	Detail()
}

type ChainManager struct {
	chains map[string]Chain
}

func NewChainManger() *ChainManager {
	return &ChainManager{
		chains: make(map[string]Chain),
	}
}

func (m *ChainManager) Get(name string) Chain {
	chain, ok := m.chains[name]
	if !ok {
		return nil
	}
	return chain
}

func (m *ChainManager) Set(name string, chain Chain) {
	m.chains[name] = chain
}
