package main_test

import (
	"database/sql"
	"frame_up_bot/config"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDbConnection(t *testing.T) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		t.Errorf("Error DB connection %s", err.Error())

	}
	defer database.Close()

}
