package singleton

import "sync"

var (
	lazyInstance *Singleton
	once         = &sync.Once{}
)

//GetLazyInstance 得到懒汉单例
func GetLazyInstance() *Singleton {
	if lazyInstance == nil {
		once.Do(func() {
			lazyInstance = &Singleton{}
		})
	}

	return lazyInstance
}
