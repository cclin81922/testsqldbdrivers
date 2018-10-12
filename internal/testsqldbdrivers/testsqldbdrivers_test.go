package testsqldbdrivers

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

func TestSQLite(t *testing.T) {
	db, err := sql.Open("sqlite3", "./dummy.db")

	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatal(err)
	}

	db.Close()
}

func TestMySQL(t *testing.T) {
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/dummy")

	if err != nil {
		t.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		t.Fatal(err)
	}

	db.Close()
}
