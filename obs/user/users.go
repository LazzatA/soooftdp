package obs

type Users interface {
	Update(pubName string)
	GetName() string
}
