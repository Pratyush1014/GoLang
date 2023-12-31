package db

import (
	"context"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/jackc/pgx/v5"
)

const (                                                     
    dbDriver = "postgres"
    dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries 

func TestMain (m *testing.M) {                              
    conn, err := pgx.Connect(context.Background(), dbSource)
    if err != nil {
        log.Fatal("cannot connect to db:", err)
    }

    testQueries = New(conn)

    os.Exit(m.Run())
}