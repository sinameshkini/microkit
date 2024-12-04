package models

type Operation string

const (
	EQUAL       Operation = "et"
	NOTEQUAL    Operation = "net"
	GRATER      Operation = "gt"
	LOWER       Operation = "lt"
	EQUALGRATER Operation = "egt"
	EQUALLOWER  Operation = "elt"
	IN          Operation = "in"
	NOTIN       Operation = "nin"
	LIKE        Operation = "like"
)

type SortDirection string

const (
	ASC  SortDirection = "asc"
	DESC SortDirection = "desc"
)
