package domain

import "errors"

type Log struct {
	TS      int64  `json:"ts"` // unix ms
	Level   string `json:"level"`
	Message string `json:"message"`
	Service string `json:"service"`
}

func (l *Log) San() error {
	if l.TS <= 0 {
		return errors.New("ts is required")
	}
	if l.Service == "" {
		return errors.New("service is required")
	}
	if l.Message == "" {
		return errors.New("message is required")
	}
	switch l.Level {
	case "info", "warn", "error", "debug", "trace":
		// ok
	default:
		return errors.New("invalid log level, use (info, warn, error, debug or trace)")
	}
	return nil
}
