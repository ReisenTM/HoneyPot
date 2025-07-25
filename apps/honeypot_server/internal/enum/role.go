package enum

type Role int8

const (
	RoleUser Role = iota + 1
	RoleAdmin
)
