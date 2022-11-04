package obs

type Subject interface {
	Subscribe(Users)
	UnSubscribe(Users)
	Notify()
}
