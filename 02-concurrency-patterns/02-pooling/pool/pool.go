package pool

import (
	"errors"
	"fmt"
	"io"
	"sync"
)

var ErrNegativePoolSize = errors.New("negative pool size")
var ErrPoolClosed = errors.New("pool is closed")

type Pool struct {
	resources chan io.Closer
	mutex     sync.Mutex
	factory   func() (io.Closer, error)
	closed    bool
}

func New(factoryFn func() (io.Closer, error), size int) (*Pool, error) {
	if size <= 0 {
		return nil, ErrNegativePoolSize
	}

	return &Pool{
		factory:   factoryFn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (p *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-p.resources:
		if !ok {
			return nil, ErrPoolClosed
		}
		fmt.Println("Acquire : From pool")
		return r, nil
	default:
		fmt.Println("Acquire : New Resource from factory")
		return p.factory()
	}
}

func (p *Pool) Release(resource io.Closer) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return ErrPoolClosed
	}

	select {
	case p.resources <- resource:
		fmt.Println("Release : To pool")
		return nil
	default:
		fmt.Println("Release : Close resource")
		return resource.Close()
	}
}

func (p *Pool) Close() {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	if p.closed {
		return
	}

	p.closed = true

	close(p.resources)
	for r := range p.resources {
		r.Close()
	}
}
