package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/blobtrtl3/trtl3-logpool/internal/infra"
	"github.com/blobtrtl3/trtl3-logpool/internal/usecase"
	"github.com/blobtrtl3/trtl3-logpool/internal/workers"
	"github.com/blobtrtl3/trtl3-logpool/pkg/domain"
)

// TODO: load balance
func main() {
	var ctx = context.Background()
	var wg sync.WaitGroup

	redis := infra.NewRedistClient(ctx)
	logs := usecase.NewLogsUseCase()

	workers.LogQueueWorkers(ctx, &wg, redis, logs)

	http.HandleFunc("POST /logs", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "failed to request body", http.StatusBadRequest)
			return
		}

		var log domain.Log
		if err := json.Unmarshal(b, &log); err != nil {
			http.Error(w, "failed to parse body", http.StatusBadRequest)
			return
		}

		if err := log.San(); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := redis.LPush(ctx, "logs.queue", b).Err(); err != nil {
			http.Error(w, "failed to enqueue log", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	// TODO: cache with redis
	http.HandleFunc("GET /logs", func(w http.ResponseWriter, r *http.Request) {
		// create filters
		// fast query with filters
		logs.Take()
	})

	// TODO: generate a html page to view logs

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error when create http server: %v", err)
	}
}
