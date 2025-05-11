# Logger Package
This package wraps zerolog to simplify configuration and logging in Go applications.

## Features
- Configurable log levels: `Debug`, `Info`, `Warn`, `Error`, `Fatal`, `Panic`, `Dpanic`
- JSON or human-readable log formats (Console).
- Flexible log output (stdout, file, or other destinations).
- Timestamps in RFC3339 format.
- Development (Dpanic) and production-friendly logging.
- Log level filtering and conditional logging.

## Usage
Here’s an example of how to use the logger in your application.
```go
package main

import (
    "github.com/DucTran999/shared-pkg/logger"
    "os"
)

func main() {
	conf := logger.Config{
		Environment: logger.Production,
		LogToFile:   true,
		FilePath:    "logs/app.log",
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		log.Fatalln("Init logger ERR", err)
	}
	defer logInst.Sync()

    // Log at different levels
    logInst.Debug("Debug log")
    logInst.Info("Info log")
    logInst.Warn("Warning log")
    logInst.Error("Error log")
    logInst.Fatal("Fatal error log")  // Exits program after logging
    logInst.Panic("Panic log")        // Panics after logging
}
```

## Configuration
The logger can be configured with the following options:

- Environment: Set to logger.Development or logger.Production depending on the mode.
- LogToFile: Set to true to enable logging to a file. The file path is specified by FilePath.
- FilePath: The location where the log file will be saved.

Example configuration:
```go
conf := logger.Config{
    Environment: logger.Production,  // Switch to Development for dev mode
    LogToFile:   true,
    FilePath:    "logs/app.log",
}
```