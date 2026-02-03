package entity

type Role uint8

const (
	userRole  Role = iota + 1
	adminRole Role = 2
)

func (r Role) String() string {
	switch r {
	case userRole:
		return "user"
	case adminRole:
		return "admin"
	}
	return ""
}
