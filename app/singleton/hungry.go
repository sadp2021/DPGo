package singleton

var hungrySingleton *Singleton

func init() {
	hungrySingleton = &Singleton{}
}

//GetHungryInstance 得到饿汉单例
func GetHungryInstance() *Singleton {
	return hungrySingleton
}
