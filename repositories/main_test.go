package repositories_test

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"os/exec"
	"testing"
)

var testDB *sql.DB

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setup.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setup.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func setUp() error {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setUp()
	if err != nil {
		os.Exit(1)
	}
	m.Run()
	teardown()
}
