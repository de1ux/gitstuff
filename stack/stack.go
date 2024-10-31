package stack

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const configFile = ".gitstuff.gob"

type Stack struct {
	Stack  []string
	Cursor int
}

func Load() (*Stack, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	path := filepath.Join(home, configFile)
	b, err := os.ReadFile(path)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return &Stack{
			Stack:  make([]string, 0),
			Cursor: 0,
		}, nil
	}
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return decode(b)

}

func decode(data []byte) (*Stack, error) {
	s := &Stack{
		Stack:  make([]string, 0),
		Cursor: 0,
	}
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(s); err != nil {
		return nil, fmt.Errorf("error gob decoding: %w", err)
	}
	return s, nil
}

func encode(s *Stack) ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *Stack) Save() {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	path := filepath.Join(home, configFile)
	b, err := encode(s)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(path, b, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Stack) Push(branch string) {
	if len(s.Stack) == 0 {
		s.Stack = []string{branch}
		return
	}
	n := make([]string, 0)
	for i := 0; i <= s.Cursor; i++ {
		if s.Stack[i] == branch {
			continue
		}
		n = append(n, s.Stack[i])
	}
	s.Stack = append(n, branch)
	for i := 0; i < len(s.Stack); i++ {
		if s.Stack[i] == branch {
			s.Cursor = i
			break
		}
	}
}

func (s *Stack) Forward() (string, error) {
	if s.Cursor+1 >= len(s.Stack) {
		return "", errors.New("failed to go forward: already at newest entry")
	}
	s.Cursor++
	return s.Stack[s.Cursor], nil
}

func (s *Stack) Back() (string, error) {
	if s.Cursor-1 < 0 {
		return "", errors.New("failed to go back: already at oldest entry")
	}
	s.Cursor--
	return s.Stack[s.Cursor], nil
}
