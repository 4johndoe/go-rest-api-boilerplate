package entity

type User struct {
	ID   string
	Name string
}

func (u User) GetID() string {
	return u.ID
}

func (u User) GetName() string {
	return u.Name
}
