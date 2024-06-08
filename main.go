package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
)

type Tenant struct {
	ID          uuid.UUID
	Name        string
	Company     string
	Status      *string
	IsDedicated bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
	DeletedAt   *time.Time
}

func main() {
	envType := flag.String("APP_ENV", "development", "set APP_ENV")
	flag.Parse()

	log.Println("envType: ", *envType)

	if *envType == "development" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal("Could not load Env file.", err.Error())
		}
	}

	dbpool, err := SetupDbConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer dbpool.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, "Server is running good.")
	})

	mux.HandleFunc("/test", func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var tenant Tenant
		err = dbpool.QueryRow(context.Background(), "select * from tenant Limit 1").Scan(&tenant.ID, &tenant.Name, &tenant.Company, &tenant.Status, &tenant.IsDedicated, &tenant.CreatedAt, &tenant.UpdatedAt, &tenant.DeletedAt)
		if err != nil {
			log.Println("Query Error: ", err)
		}

		log.Println("tenant: ", tenant)

		fmt.Fprint(w, tenant)
	})

	server := &http.Server{
		Handler:        mux,
		Addr:           ":3000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(server.ListenAndServe())
}
