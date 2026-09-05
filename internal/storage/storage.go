//Package storage saves uploaded originals to the local filesystem using names the server generates itself, never names given by the client.

package storage

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type OriginalStore struct {
	Dir string
}

// NewOriginalStore ensures Dir exists and returns a store rooted there
func NewOriginalStore(dir string) (*OriginalStore, error) {
	if err := os.MkdirAll(dir, 0o755); //creates all missing parent directories recursively.
	err != nil {
		return nil, fmt.Errorf("creating uploade directory %q: %w", dir, err)
	}
	return &OriginalStore{Dir: dir}, nil
}

// randomFilename builds a server-controlled name
func randomFilename(ext string) (string, error) {
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return "", err
	}
	return hex.EncodeToString(buf) + ext, nil
}

// SaveOriginal copies "r" into a new file under Dir and returns the stored filename
func (s *OriginalStore) SaveOriginal(r io.Reader, ext string) (storedFilename string, sizeBytes int64, err error) {
	name, err := randomFilename(ext)
	if err != nil {
		return "", 0, err
	}
	dst := filepath.Join(s.Dir, name)

	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644) // O_EXCL refuses to silently overwrite an existing file
	if err != nil {
		return "", 0, err
	}
	defer f.Close()

	written, err := io.Copy(f, r)
	if err != nil {
		os.Remove(dst)
		return "", 0, err
	}
	return name, written, nil
}

// Path rejoins a stored filename with this store's root directory
func (s *OriginalStore) Path(storedFilename string) string {
	return filepath.Join(s.Dir, storedFilename)
}
