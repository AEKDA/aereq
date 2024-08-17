package logger

type Config struct {
	Level     string `json:"level"`
	StructLog bool   `json:"structure_log" default:"false"`
}
