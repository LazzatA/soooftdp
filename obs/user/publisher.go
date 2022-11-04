package obs

type Publisher struct {
	Name  string
	Users map[string]Users
}

func (pub *Publisher) Subscribe(users Users) {
	pub.Users[users.GetName()] = users
}

func (pub *Publisher) UnSubscribe(users Users) {
	delete(pub.Users, users.GetName())
}

func (pub *Publisher) Notify(users Users) {
	for _, users := range pub.Users {
		users.Update(pub.Name)
	}
}
