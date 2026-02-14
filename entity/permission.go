package entity

type Permission struct {
	ID    uint
	Title string
}

type PermissionsTitle string

const (
	UserListPermission   = PermissionsTitle("user-list")
	UserDeletePermission = PermissionsTitle("user-delete")
)
