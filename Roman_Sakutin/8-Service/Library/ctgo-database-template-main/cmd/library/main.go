package main

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/project/library/config"
	"github.com/project/library/internal/app"
)

func main() {
	cfg, err := config.NewConfig()

	if err != nil {
		log.Fatalf("can not get application config: %s", err)
	}

	logger, err := NewFileLogger(cfg.Observability.LogFile)

	if err != nil {
		fmt.Println(err.Error()) // todo
		log.Fatalf("can not initialize logger: %s", err)
	}

	app.Run(logger, cfg)
}

func NewFileLogger(logFile string) (*zap.Logger, error) {
	if err := os.MkdirAll(filepath.Dir(logFile), 0755); err != nil {
		return nil, fmt.Errorf("create log dir: %w", err)
	}

	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	if err != nil {
		return nil, err
	}

	writeSyncer := zapcore.AddSync(file)
	encoderCfg := zap.NewProductionEncoderConfig()
	encoder := zapcore.NewJSONEncoder(encoderCfg)

	core := zapcore.NewCore(encoder, writeSyncer, zap.InfoLevel)

	return zap.New(core), nil
}
