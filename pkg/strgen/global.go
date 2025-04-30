package strgen

import "sync"

var Global *Generator = nil

func InitGlobal(opts ...Option) {
	once := sync.Once{}
	once.Do(func() {
		Global = newGenerator(opts...)
	})
}
