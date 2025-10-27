package cmd

import (
	"bytes"
	"context"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/blobtrtl3/trtl3-logpool/internal/infra"
	"github.com/blobtrtl3/trtl3-logpool/internal/usecase"
)

// TODO: load balance
func main() {
	var ctx = context.Background()

	redis := infra.NewRedistClient(ctx)
	logs := usecase.NewLogsUseCase()

	http.HandleFunc("POST /logs", func(w http.ResponseWriter, r *http.Request) {
		b, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		i := bytes.Index(b, []byte(`"ts":`))
		if i == -1 {
			http.Error(w, "invalid data", http.StatusBadRequest)
			return
		}

		i += len(`"ts":`)

		for i < len(b) && (b[i] == ' ' || b[i] == '\t' || b[i] == '\n') { // jump spaces, taps and line breaks
			i++
		}

		j := i // j is ts start
		for i < len(b) && b[i] >= '0' && b[i] <= '9' { // loops to the next , verifying the elements
			i++
		}

		ts, err := strconv.ParseInt(string(b[j:i]), 10, 64)
		if err != nil {
			http.Error(w, "invalid timestamp", http.StatusBadRequest)
			return
		}

		// verify timestamp age
		now := time.Now().Unix()
		if ts > now {
			http.Error(w, "invalid timestamp", http.StatusBadRequest)
			return
		}

		var retriesErr error
		for atp := 0; atp <= 1; atp++ { // 2 retries
			if retriesErr = redis.LPush(ctx, "logs.queue", b).Err(); retriesErr == nil {
				break // success
			}
		}

		if retriesErr != nil {
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
