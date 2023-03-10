package server

import (
	"calc-bot/config"
	"calc-bot/internal/calc-bot/storage/sqlite"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Info struct {
	Count int `json:"count"`
}

func Start(config *config.Config, storage *sqlite.Storage) error {
	infoHandler := func(w http.ResponseWriter, r *http.Request) {
		res, _ := storage.GetCount(context.TODO())
		json.NewEncoder(w).Encode(Info{Count: res})
	}
	http.HandleFunc("/info", infoHandler)
	address := fmt.Sprintf("%s:%s", config.HTTP.Host, config.HTTP.Port)
	return http.ListenAndServe(address, nil)
}
