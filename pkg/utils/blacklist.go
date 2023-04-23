package utils

import "sync"

type BlackList struct {
	sync.Map
}

func (bl *BlackList) Add(name string) {
	bl.Store(name, 1)
}

func (bl *BlackList) Get(name string) bool {
	_, ok := bl.Load(name)
	return ok
}

func (bl *BlackList) Clear() {
	bl.Range(func(key, value any) bool {
		bl.Delete(key)
		return true
	})
}

var bl = new(BlackList)

func GetBlackList() *BlackList {
	return bl
}
