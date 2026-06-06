package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// loadEnvFile подгружает .env из корня go-модуля (для локального запуска из IDE).
// Уже заданные переменные окружения не перезаписываются.
func loadEnvFile() {
	root := findModuleRoot()
	if root == "" {
		return
	}

	_ = parseEnvFile(filepath.Join(root, ".env"), false)
	_ = parseEnvFile(filepath.Join(root, ".env.local"), true)
}

const libraryModule = "module github.com/project/library"

func findModuleRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	for {
		data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
		if err == nil && strings.Contains(string(data), libraryModule) {
			return dir
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

func parseEnvFile(path string, override bool) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		key, val, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}

		key = strings.TrimSpace(key)
		val = strings.TrimSpace(val)
		val = strings.Trim(val, `"'`)

		if key == "" {
			continue
		}
		if override || os.Getenv(key) == "" {
			_ = os.Setenv(key, val)
		}
	}

	return scanner.Err()
}
