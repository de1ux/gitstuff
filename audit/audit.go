package audit

import (
	"fmt"
	"os"
	"path/filepath"
)

const auditFile = ".gitstuff.audit"

func Write(msg string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil
	}
	path := filepath.Join(home, auditFile)
	// create file if it doesnt exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
		if _, err := os.Create(path); err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
	}

	// open for append
	f, err :=
		os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer f.Close()

	// write the audit message
	if _, err := f.WriteString(msg + "\n"); err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}
