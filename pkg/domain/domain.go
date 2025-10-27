package domain

type Log struct {
	TS      int64 `json:"ts"` // unix ms
	Level   string `json:"level"`
	Message string `json:"message"`
	Service string `json:"service"`
}
