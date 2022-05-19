package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type Config struct {
	host     string
	port     int
	user     string
	password string
	name     string
}

type DbConfig struct {
	host string
	name string
	Port int
	Db   *sql.DB
}

type any interface{}

func (cfg Config) connect() *DbConfig {
	db_conn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s",
		cfg.host,
		cfg.user,
		cfg.password,
		cfg.name)
	if db, err := sql.Open("postgres", db_conn); err == nil {
		fmt.Printf("[XAPI INFO] Connected to Postgres Db %s/%s\n", cfg.host, cfg.name)
		return &DbConfig{
			cfg.host,
			cfg.name,
			cfg.port,
			db,
		}
	} else {
		fmt.Printf("[XAPI WARN] Database connection error: %s", err)
		return nil
	}
}

func EnvConnect() *DbConfig {
	godotenv.Load("../../.env")
	if port, err := strconv.Atoi(os.Getenv("BACKEND_PORT")); err == nil {
		cfg := Config{
			os.Getenv("HOST"),
			port,
			os.Getenv("DB_USERNAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		}
		return cfg.connect()
	} else {
		fmt.Printf("[XAPI INFO] Environment variables not set. Defaulting to postgres/postgres, %v", err)
		return nil
	}
}

func (db_cfg *DbConfig) Query(qStr string, args ...any) *sql.Rows {
	var rows *sql.Rows
	var err error
	if len(args) > 0 {
		rows, err = db_cfg.Db.Query(qStr, pq.Array(args))
	} else {
		rows, err = db_cfg.Db.Query(qStr)
	}
	if err == nil {
		return rows
	} else {
		fmt.Printf("[XAPI WARN] Error executing query %s: %s\n", qStr, err)
		return nil
	}
}

func (db_cfg *DbConfig) QueryOne(qStr string, args ...any) *sql.Row {
	var row *sql.Row
	if len(args) > 1 {
		row = db_cfg.Db.QueryRow(qStr, args)
	} else if len(args) == 1 {
		row = db_cfg.Db.QueryRow(qStr, args[0])
	} else {
		row = db_cfg.Db.QueryRow(qStr)
	}
	if row != nil {
		return row
	} else {
		fmt.Printf("[XAPI WARN] Query returned 0 rows.\n")
		return nil
	}
}

func (db_cfg *DbConfig) Close() {
	defer db_cfg.Db.Close()
}
