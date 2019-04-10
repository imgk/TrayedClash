package clash

var instance, _ = NewClash()

func GetInstance() *Clash {
	return instance
}
