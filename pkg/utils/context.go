package utils

import (
	"context"
	"sync"
)

var ctx context.Context

// GetCTX 懒汉式生成context
func GetCTX() context.Context {
	once := sync.Once{}
	once.Do(func() {
		ctx = context.Background()
	})
	return ctx
}
