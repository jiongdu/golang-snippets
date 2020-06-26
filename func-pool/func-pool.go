package pool

import (
	"github.com/gomodule/redigo/redis"
	"github.com/hashicorp/go-multierror"
)

type Pool interface {
	Get() redis.Conn
}

type ActFunc func(Pool) (bool, error)

type Manager struct {
	pools []Pool
}

func (m *Manager) actOnPoolsAsync(actFn ActFunc) (int, error) {
	type result struct {
		Status bool
		Err    error
	}
	ch := make(chan result)
	for _, pool := range m.pools {
		go func(pool Pool) {
			r := result{}
			r.Status, r.Err = actFn(pool)
			ch <- r
		}(pool)
	}
	n := 0
	var err error
	for range m.pools {
		r := <-ch
		if r.Status {
			n++
		} else if r.Err != nil {
			err = multierror.Append(err, r.Err)
		}
	}
	return n, err
}
