package storage

type Storage interface {
	Save(u *User) error
	Remove(u *User) error
	Get(id string) (User, error)
	IsExists(u *User) (bool, error)
}

type User struct {
	Id       int
	Username string
	Meta     interface{}
}
