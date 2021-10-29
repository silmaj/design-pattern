package create_pattern

import (
	"context"
	"sync"
)

type ConnectionPool interface {
	connect(ctx context.Context) error
}

type pool struct {}

func (p *pool) connect(ctx context.Context) error {
	return nil
}

var (
	p  	 *pool
	once sync.Once
)

func GetConnectionPool() ConnectionPool {
	once.Do(func() {
		p = &pool{}
	})
	return p
}
