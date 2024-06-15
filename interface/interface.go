package _interface

import "github.com/redis-cli/pkg/models"

type Runner interface {
	// Close closes all connections in the pool
	Close()
	// TestConnect test if the db connection is ready
	TestConnect() error
	// Exec returns the number of rows affected by the SQL execution
	Exec(sql *string, args ...any) (*models.QueryResult, error)
	// Query returns the results of the SQL query
	Query(sql *string, args ...any) (*models.QueryResult, error)
}
