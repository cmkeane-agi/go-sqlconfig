//------------------------------------------------------------
// SQL File Loader
//------------------------------------------------------------

package sqlconfig

import (
	"fmt"
	"os"
	"strings"
	"time"
)

type Query struct {
	Name     string
	SQL      string
	File     string
	Line     int
	LoadedAt time.Time
}

type Store struct {
	Queries map[string]Query
}

// ------------------------------------------------------------------------------------
func LoadSQL(path string) (*Store, error) {
	store := &Store{
		Queries: make(map[string]Query),
	}

	if err := loadSingleFile(store, path); err != nil {
		return nil, err
	}

	if len(store.Queries) == 0 {
		return nil, fmt.Errorf("no SQL queries found in %s", path)
	}

	return store, nil
}

// ------------------------------------------------------------------------------------
func loadSingleFile(store *Store, path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	var (
		currentName string
		startLine   int
		buf         []string
	)

	flush := func(lineNo int) error {
		if currentName == "" {
			return nil
		}

		sql := strings.TrimSpace(strings.Join(buf, "\n"))
		if sql == "" {
			return fmt.Errorf("%s:%d query %q is empty", path, startLine, currentName)
		}

		if _, exists := store.Queries[currentName]; exists {
			return fmt.Errorf("duplicate query name %q (found again at %s:%d)", currentName, path, startLine)
		}

		store.Queries[currentName] = Query{
			Name:     currentName,
			SQL:      sql,
			File:     path,
			Line:     startLine,
			LoadedAt: time.Now(),
		}

		buf = nil
		currentName = ""
		return nil
	}

	for i, line := range lines {
		trim := strings.TrimSpace(line)

		if strings.HasPrefix(trim, "-- name:") {
			if err := flush(i + 1); err != nil {
				return err
			}

			currentName = strings.TrimSpace(strings.TrimPrefix(trim, "-- name:"))
			if currentName == "" {
				return fmt.Errorf("%s:%d missing query name", path, i+1)
			}

			startLine = i + 1
			continue
		}

		if currentName != "" {
			buf = append(buf, line)
		}
	}

	return flush(len(lines))
}

// ------------------------------------------------------------------------------------
func (s *Store) Must(name string) string {
	q, ok := s.Queries[name]
	if !ok {
		panic(fmt.Sprintf("missing SQL query %q", name))
	}
	return q.SQL
}

// ------------------------------------------------------------------------------------
func (s *Store) Get(name string) (Query, bool) {
	q, ok := s.Queries[name]
	return q, ok
}
