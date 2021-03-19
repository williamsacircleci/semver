package utils

type PR struct {
	Repo          string
	Type          PRType
	Author        string
	Name          string
	ReleaseNumber string
	ReleaseNotes  string
}

type PRType int

const (
	Frontend = PRType(0)
	Backend  = PRType(1)
	DevOps   = PRType(2)
	Graphql  = PRType(3)
)
