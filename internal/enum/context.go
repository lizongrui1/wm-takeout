package enum

const (
	CurrentId   = "currentId"
	CurrentName = "currentName"
)

type PageNum = int

const (
	MaxPageSize PageNum = 30
	MinPageSize PageNum = 10
)

type CommonInt = int

const (
	MaxUrl CommonInt = 1
)
