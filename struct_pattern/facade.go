package struct_pattern

const (
	bitcoin  = "bitcoin"
	ethereum = "ethereum"
)

type BackendInfo interface {
	GetBackend(ty string) Backend
	ShowBackend(ty string) string
}

type Backend interface {
	Show() string
}

type BitcoinBackend struct {
	BestHeight uint64
	BestHash   string
}

func (b *BitcoinBackend) Show() string {
	return ""
}

type EthereumBackend struct {
	BestHeight uint64
	BestHash   string
}

func (b *EthereumBackend) Show() string {
	return ""
}

type BackendManager struct {
	backends map[string]Backend
}

func (m *BackendManager) GetBackend(ty string) Backend {
	b, ok := m.backends[ty]
	if !ok {
		return nil
	}
	return b
}

func (m *BackendManager) ShowBackend(ty string) string {
	b, ok := m.backends[ty]
	if !ok {
		return ""
	}
	return b.Show()
}
