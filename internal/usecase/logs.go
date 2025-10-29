package usecase

import (
	"fmt"

	"github.com/blobtrtl3/trtl3-logpool/pkg/domain"
)

type LogsUseCase struct {
}

func NewLogsUseCase() *LogsUseCase {
	return &LogsUseCase{}
}

func (logs *LogsUseCase) Create(log *domain.Log) error {
	fmt.Println("saving log")
	return nil
}

func (logs *LogsUseCase) Take() {
}
