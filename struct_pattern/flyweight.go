package struct_pattern

type EVMChainFactory struct {
	chains map[string]*EVMChainImpl
}

var chainFactory *EVMChainFactory

func NewChainFactory() *EVMChainFactory {
	if chainFactory == nil {
		chainFactory = &EVMChainFactory{
			chains: make(map[string]*EVMChainImpl, 0),
		}
	}
	return chainFactory
}

func (f *EVMChainFactory) Get(ty string) {
	impl := f.chains[ty]
	if impl == nil {
		impl = NewEVMChainImpl(ty)
		f.chains[ty] = impl
	}
}

type EVMChainImpl struct {
	chainID uint32
	url 	string
	symbol 	string
	explorerLink string
}

func NewEVMChainImpl(ty string) *EVMChainImpl {
	switch ty {
	case "eth":
		return &EVMChainImpl{

		}
	case "bsc":
		return &EVMChainImpl{}
	default:
		return nil
	}
}

type EthereumChainConfig struct {
	EVMChainImpl
}

type BSCChainConfig struct {
	EVMChainImpl
}

