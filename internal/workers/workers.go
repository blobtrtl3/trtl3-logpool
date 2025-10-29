package workers

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/blobtrtl3/trtl3-logpool/internal/usecase"
	"github.com/blobtrtl3/trtl3-logpool/pkg/domain"
	"github.com/redis/go-redis/v9"
)

func LogQueueWorkers(ctx context.Context, wg *sync.WaitGroup, redis *redis.Client, logs *usecase.LogsUseCase) {
	for i := 0; i < 6; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				values, err := redis.BLPop(ctx, 0, "logs.queue").Result()
				if err != nil {
					continue
				}

				var log domain.Log
				if err := json.Unmarshal([]byte(values[1]), &log); err != nil {
					continue
				}

				if err := logs.Create(&log); err != nil {
					continue
				}
			}
		}(i+1)
	}
}

