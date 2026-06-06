package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type (
	Config struct {
		GRPC struct {
			Port        string `env:"GRPC_PORT"`
			GatewayPort string `env:"GRPC_GATEWAY_PORT"`
		}

		PG struct {
			URL      string
			Host     string `env:"POSTGRES_HOST"`
			Port     string `env:"POSTGRES_PORT"`
			DB       string `env:"POSTGRES_DB"`
			User     string `env:"POSTGRES_USER"`
			Password string `env:"POSTGRES_PASSWORD"`
			MaxConn  string `env:"POSTGRES_MAX_CONN"`
		}

		Outbox struct {
			Enabled         bool          `env:"OUTBOX_ENABLED"`
			Workers         int           `env:"OUTBOX_WORKERS"`
			BatchSize       int           `env:"OUTBOX_BATCH_SIZE"`
			WaitTimeMS      time.Duration `env:"OUTBOX_WAIT_TIME_MS"`
			InProgressTTLMS time.Duration `env:"OUTBOX_IN_PROGRESS_TTL_MS"`
			AuthorSendURL   string        `env:"OUTBOX_AUTHOR_SEND_URL"`
			BookSendURL     string        `env:"OUTBOX_BOOK_SEND_URL"`
		}

		Observability struct {
			MetricsPort  string `env:"METRICS_PORT"`
			JaegerURL    string `env:"JAEGER_URL"`
			ServiceName  string `env:"SERVICE_NAME"`
			PyroscopeURL string `env:"PYROSCOPE_URL"`
			LogFile      string `env:"LOG_FILE"`
		}
	}
)

func NewConfig() (*Config, error) {
	loadEnvFile()

	cfg := &Config{}

	cfg.GRPC.Port = os.Getenv("GRPC_PORT")
	cfg.GRPC.GatewayPort = os.Getenv("GRPC_GATEWAY_PORT")

	cfg.PG.Host = os.Getenv("POSTGRES_HOST")
	cfg.PG.Port = os.Getenv("POSTGRES_PORT")
	cfg.PG.DB = os.Getenv("POSTGRES_DB")
	cfg.PG.User = os.Getenv("POSTGRES_USER")
	cfg.PG.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.PG.MaxConn = os.Getenv("POSTGRES_MAX_CONN")
	if cfg.PG.MaxConn == "" {
		cfg.PG.MaxConn = "10"
	}

	cfg.PG.URL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&pool_max_conns=%s",
		cfg.PG.User,
		cfg.PG.Password,
		cfg.PG.Host,
		cfg.PG.Port,
		cfg.PG.DB,
		cfg.PG.MaxConn,
	)

	cfg.Outbox.AuthorSendURL = os.Getenv("OUTBOX_AUTHOR_SEND_URL")
	cfg.Outbox.BookSendURL = os.Getenv("OUTBOX_BOOK_SEND_URL")
	cfg.Outbox.Enabled = os.Getenv("OUTBOX_ENABLED") == "true"

	cfg.Observability.MetricsPort = os.Getenv("METRICS_PORT")
	cfg.Observability.JaegerURL = os.Getenv("JAEGER_URL")
	cfg.Observability.ServiceName = os.Getenv("SERVICE_NAME")
	cfg.Observability.PyroscopeURL = os.Getenv("PYROSCOPE_URL")

	if w := os.Getenv("OUTBOX_WORKERS"); w != "" {
		cfg.Outbox.Workers, _ = strconv.Atoi(w)
	}

	if w := os.Getenv("OUTBOX_BATCH_SIZE"); w != "" {
		cfg.Outbox.BatchSize, _ = strconv.Atoi(w)
	}

	cfg.Outbox.WaitTimeMS = ParseOutboxDurationMS(os.Getenv("OUTBOX_WAIT_TIME_MS"))
	cfg.Outbox.InProgressTTLMS = ParseOutboxDurationMS(os.Getenv("OUTBOX_IN_PROGRESS_TTL_MS"))

	if cfg.Observability.MetricsPort == "" {
		cfg.Observability.MetricsPort = "9000"
	}
	if cfg.Observability.JaegerURL == "" {
		cfg.Observability.JaegerURL = "jaeger:4318"
	}
	if cfg.Observability.ServiceName == "" {
		cfg.Observability.ServiceName = "library-service"
	}
	if cfg.Observability.PyroscopeURL == "" {
		cfg.Observability.PyroscopeURL = "http://pyroscope:4040"
	}
	if cfg.Observability.LogFile == "" {
		if root := findModuleRoot(); root != "" {
			cfg.Observability.LogFile = filepath.Join(root, "logs", "library.log")
		} else {
			cfg.Observability.LogFile = "logs/library.log"
		}
	}

	return cfg, nil
}
