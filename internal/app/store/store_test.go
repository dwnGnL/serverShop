package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		databaseURL = "host=localhost dbname=serverShop_test password=1234 sslmode=disable"
	}

	os.Exit(m.Run())
}
