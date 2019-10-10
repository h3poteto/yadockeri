package db

import (
	"database/sql"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v2"
)

// Database has database connection object.
type Database struct {
	Connection *sql.DB
}

var sharedInstance = newDBConnection()

func newDBConnection() *Database {
	env := os.Getenv("ECHO_ENV")
	path := "db/dbconf.yml"
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	driver := m[env].(map[interface{}]interface{})["driver"].(string)
	open := m[env].(map[interface{}]interface{})["open"].(string)
	pool := m[env].(map[interface{}]interface{})["pool"].(int)
	open = os.ExpandEnv(open)

	db, err := sql.Open(driver, open)
	if err != nil {
		panic(err)
	}

	// MaxIdle: The Limit of connection pool which is held while there is no access to database.
	// MaxOpen: The limit of idle + active connection pool.
	db.SetMaxIdleConns(pool)
	db.SetMaxOpenConns(pool)

	return &Database{
		Connection: db,
	}
}

// SharedInstance return database connection object.
func SharedInstance() *Database {
	return sharedInstance
}

// Close database connection.
func (d *Database) Close() error {
	return d.Connection.Close()
}
