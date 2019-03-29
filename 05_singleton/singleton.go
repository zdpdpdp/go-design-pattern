package _5_singleton

import "sync"

type Singleton struct {
}

var instance *Singleton

func clear() {
	instance = nil
}

//EasyInstance 懒汉单例模式  线程不安全
func EasyInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}

	return instance
}

var lock sync.Mutex

//加锁下的单例
func CheckInstance() *Singleton {
	if instance != nil {
		return instance
	}
	//如果有两个线程同时进来,同时检测到 nil,那么就会同时创建了,所以还是不安全
	lock.Lock()
	instance = &Singleton{}
	lock.Unlock()
	return instance
}

//DoubleCheckInstance 双重检查机制
func DoubleCheckInstance() *Singleton {
	if instance != nil {
		return instance
	}
	lock.Lock()
	defer lock.Unlock()
	if instance != nil { //获得锁以后再检查多一次
		return instance
	}
	instance = &Singleton{}
	return instance
}

var once sync.Once

//OnceInstance 使用 go 的 sync.Once 来保证只执行一次
func OnceInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}
