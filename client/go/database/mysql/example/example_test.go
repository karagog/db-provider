// Package example provides a basic working example of how to instantiate
// a new Mysql database and use it in a test.
//
// Note: The instance provider service must already be running in order for
// this test to work. See README for details.
package example

import (
	"context"
	"flag"
	"testing"

	"github.com/karagog/db-provider/client/go/database"
	"github.com/karagog/db-provider/client/go/database/mysql"
)

func TestMysqlDatabase(t *testing.T) {
	// Instantiate a fresh new database.
	i := database.NewFromFlags(context.Background())

	// Connect as the administrative user in order to create tables.
	rootDB := mysql.ConnectOrDie(i.Info.RootConn)
	if _, err := rootDB.Exec(`
		CREATE TABLE foo (
			id INT NOT NULL AUTO_INCREMENT,
			greeting VARCHAR(50),
			PRIMARY KEY (id)
		)`); err != nil {
		t.Fatal(err)
	}

	// Connect as the application user to do CRUD operations.
	db := mysql.ConnectOrDie(i.Info.AppConn)
	if _, err := db.Exec(`
		INSERT INTO foo (greeting)
		VALUES ('Hello World!')
	`); err != nil {
		t.Fatal(err)
	}

	row := db.QueryRow(`SELECT * FROM foo`)
	var id int64
	var greeting string
	if err := row.Scan(&id, &greeting); err != nil {
		t.Fatal(err)
	}
	if got, want := id, int64(1); got != want {
		t.Errorf("Got id %v, want %v", got, want)
	}
	if got, want := greeting, "Hello World!"; got != want {
		t.Errorf("Got greeting %v, want %v", got, want)
	}
}

func TestPanicOnError(t *testing.T) {
	// Set the address flag to an invalid value to see if we panic.
	flag.Set("db_instance_provider_address", "")

	// This function detects whether the database instantiation panicked or not.
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Did not panic, want panic")
			}
		}()
		database.NewFromFlags(context.Background())
	}()
}
