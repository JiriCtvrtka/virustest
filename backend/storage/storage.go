package storage

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func SaveData(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connectToDB()
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	query := `INSERT INTO virus_data (data, ip_address) VALUES ($1, $2)`
	_, err = db.Exec(query, string(b), r.RemoteAddr)
	if err != nil {
		log.Println(err)
		return
	}
}

func proceedEnvVariable(key string) string {
	envVar := os.Getenv(key)
	if envVar == "" {
		log.Panicf("%s env variable is not defined", key)
	}

	return envVar
}

func connectToDB() (*sql.DB, error) {
	port, err := strconv.Atoi(proceedEnvVariable("PORT"))
	if err != nil {
		log.Panic(err)
	}

	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		proceedEnvVariable("HOST"), port, proceedEnvVariable("USER"), proceedEnvVariable("PASS"), proceedEnvVariable("DBNAME"))

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
