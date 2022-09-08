package enums

// DatabaseDriver ...
type DatabaseDriver int

const (
	// Postgres SQL ...
	PostgresSQL DatabaseDriver = iota + 1
	// SQL Server ...
	SQLServer
)
