package database

import (
	"context"
	"errors"
	"fmt"

	"net/url"
	"reflect"
	"strings"

	"encore.dev/rlog"
	"github.com/jmoiron/sqlx"
)

// Config is the database configuration
// Config is the required properties to use the database.
type Config struct {
	User         string
	Password     string
	Host         string
	Name         string
	MaxIdleConns int
	MaxOpenConns int
	DisableTLS   bool
}

// Open opens the database connection
//
//	@param config - Config
//	@return sqlx.DB
//	@return error
func Open(config Config) (*sqlx.DB, error) {
	// set the TLS mode
	tlsMode := "require"
	if config.DisableTLS {
		tlsMode = "disable"
	}

	// set the connection string
	userCredentials := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(config.User, config.Password),
		Host:   config.Host,
		Path:   config.Name,
		RawQuery: (&url.Values{
			"sslmode":  []string{tlsMode},
			"timezone": []string{"utc"},
		}).Encode(),
	}
	// create the database connection
	db, err := sqlx.Open("postgres", userCredentials.String())
	if err != nil {
		return nil, err
	}

	// set the maximum number of idle connections
	db.SetMaxIdleConns(config.MaxIdleConns)

	// set the maximum number of open connections
	db.SetMaxOpenConns(config.MaxOpenConns)

	// return the database connection
	return db, nil
}

// Status - check the database connection status and return an error if the connection is not available.
//
//	 @param ctx - context
//		@param db - database connection
//		@return error - error if any
func Status(ctx context.Context, db *sqlx.DB) error {
	// Ping the database to check the connection.
	var pingErr error

	// Ping the database.
	for attempt := 1; ; attempt++ {
		pingErr = db.Ping()

		// If there is no error, break out of the loop.
		if pingErr == nil {
			break
		}

		// If the context is done, return the error.
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}

	// check for timeout
	if ctx.Err() != nil {
		return ctx.Err()
	}

	// run a query to check the connection
	q := "SELECT 1"

	// Run the query.
	var result int

	// Scan the result into the variable.
	return db.QueryRowContext(ctx, q).Scan(&result)
}

// queryString - pretty prints the query string
//
//	@param query - query to execute
//	@param data - data to bind to the query
//	@return string - query string
func queryString(query string, args ...interface{}) string {
	// takes up the query and returns it with the params
	query, params, err := sqlx.Named(query, args)
	if err != nil {
		return err.Error()
	}

	// loop through the parameters
	for _, param := range params {
		var value string
		switch v := param.(type) {
		case string:
			value = fmt.Sprintf("%q", v)
		case []byte:
			value = fmt.Sprintf("%q", string(v))
		default:
			value = fmt.Sprintf("%v", v)
		}
		query = strings.Replace(query, "?", value, 1)
	}

	query = strings.ReplaceAll(query, "\t", "")
	query = strings.ReplaceAll(query, "\n", " ")

	return strings.Trim(query, " ")
}

// NamedExecQuery - helper function for executing queries that return no rows.
// Most of the time, this will be used for INSERT, UPDATE, and DELETE queries.
//
//	@param ctx - context
//	@param db - database connection
//	@param query - query to execute
//	@param data - data to bind to the query
//	@return error - error if any
func NamedExecQuery(ctx context.Context, db *sqlx.DB, query string, data interface{}) error {
	q := queryString(query, data)
	rlog.Info("database.NamedExecQuery", "query", q)

	// Execute the query.
	_, err := db.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	return nil
}

// NamedSliceQuery - helper function for executing queries that return a slice of rows.
// Most of the time, this will be used for SELECT queries.
//
//	@param ctx - context
//	@param db - database connection
//	@param query - query to execute
//	@param data - data to bind to the query
//	@param dest - destination to scan the rows into
//	@return error - error if any
func NamedSliceQuery(ctx context.Context, db *sqlx.DB, query string, data interface{}, dest interface{}) error {
	// get formated query string
	q := queryString(query, data)
	// log query info
	rlog.Info("database.NamedSliceQuery", "query", q)

	// get value of the slice
	val := reflect.ValueOf(dest)
	// check if the value is a pointer to a slice
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return errors.New("must provide a pointer to a slice")
	}

	// Execute the query.
	rows, err := db.NamedQueryContext(ctx, query, data)
	if err != nil {
		return err
	}

	// get the next row
	slice := val.Elem()
	// loop through the rows
	for rows.Next() {
		v := reflect.New(slice.Type().Elem())
		// Scan the row into the destination.
		if err := rows.StructScan(v.Interface()); err != nil {
			return err
		}
		// append the row to the slice
		slice.Set(reflect.Append(slice, v.Elem()))
	}

	return nil
}

// NamedStructQuery - helper function for executing queries that return a single row.
// Most of the time, this will be used for SELECT queries.
//
//	@param ctx - context
//	@param db - database connection
//	@param query - query to execute
//	@param data - data to bind to the query
//	@param dest - destination to scan the row into
//	@return error - error if any
func NamedStructQuery(ctx context.Context, db *sqlx.DB, query string, data interface{}, dest interface{}) error {
	q := queryString(query, data)
	rlog.Info("database.NamedStructQuery", "query", q)

	// Execute the query.
	rows, err := db.NamedQueryContext(ctx, query, data)
	if err != nil {
		return err
	}

	// If there are no rows, return an error.
	if !rows.Next() {
		return ErrNotFound
	}

	// Scan the row into the destination.
	if err := rows.StructScan(dest); err != nil {
		return err
	}

	return nil
}

// NamedCountQuery is a helper function for executing queries that return a count.
//
//	@param ctx - context
//	@param db - database connection
//	@param query - query to execute
//	@param data - data to bind to the query
//	@return int - integer value returned from the query
//	@return error - error if any
func NamedCountQuery(ctx context.Context, db *sqlx.DB, query string, data interface{}) (int, error) {
	q := queryString(query, data)
	rlog.Info("database.NamedQueryCount", "query", q)

	// Execute the query.
	rows, err := db.NamedQueryContext(ctx, query, data)
	if err != nil {
		return 0, err
	}

	// If there are no rows, return an error.
	if !rows.Next() {
		return 0, ErrNotFound
	}

	// Declare a variable to hold the count.
	var count int
	// Scan the count into the variable.
	if err := rows.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
