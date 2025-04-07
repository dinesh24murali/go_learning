package common

// UserStatus Enum

type UserStatus uint8

const (
	Active UserStatus = iota
	Inactive
	Invited
	Blocked
)

var stateName = map[UserStatus]string{
	Active:   "active",
	Inactive: "inactive",
	Invited:  "invited",
	Blocked:  "blocked",
}

func (ss UserStatus) String() string {
	return stateName[ss]
}

// UserRole Enum

type UserRole uint8

const (
	AdminRole UserRole = iota
	CustomerRole
)

var roleName = map[UserRole]string{
	AdminRole:    "admin",
	CustomerRole: "customer",
}

func (ss UserRole) String() string {
	return roleName[ss]
}

// SalesStatus Enum

type SalesStatus uint8

const (
	Open SalesStatus = iota
	Fulfilled
	Cancelled
)

var salesStatusName = map[SalesStatus]string{
	Open:      "open",
	Fulfilled: "fulfilled",
	Cancelled: "cancelled",
}

func (ss SalesStatus) String() string {
	return salesStatusName[ss]
}
