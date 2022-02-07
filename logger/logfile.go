package logger

import (
	"log"
	"os"
)

// logging tools to stdout
var (
	Info  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	Error = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	Success = log.New(os.Stdout, "SUCCESS: ", log.Ldate|log.Ltime)
)