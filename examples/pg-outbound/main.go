package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
	"github.com/fermyon/spin-go-sdk/pg"
)

type Pet struct {
	ID        int64
	Name      string
	Prey      *string // nullable field must be a pointer
	IsFinicky bool
}

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

		// addr is the environment variable set in `spin.toml` that points to the
		// address of the Mysql server.
		addr := os.Getenv("DB_URL")

		db := pg.Open(addr)
		defer db.Close()

		_, err := db.Query("INSERT INTO pets VALUES ($1, 'Maya', $2, $3);", int32(4), "bananas", true)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		rows, err := db.Query("SELECT * FROM pets")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var pets []*Pet
		for rows.Next() {
			var pet Pet
			if err := rows.Scan(&pet.ID, &pet.Name, &pet.Prey, &pet.IsFinicky); err != nil {
				fmt.Println(err)
			}
			pets = append(pets, &pet)
		}
		json.NewEncoder(w).Encode(pets)
	})
}

func main() {}
