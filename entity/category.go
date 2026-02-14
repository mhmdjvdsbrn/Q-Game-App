package entity

type Category string

// TODO - add category to db
const (
	FootballCategory = "football"
)

func (c Category) IsValid() bool {
	switch c {
	case FootballCategory:
		return true
	}
	return false
}
