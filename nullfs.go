package nullfs

import (
	"os"
	"time"

	"github.com/spf13/afero"
)

type Fs struct{}

func (c Fs) Create(string) (afero.File, error)  { return nil, nil }
func (c Fs) Mkdir(string, os.FileMode) error    { return nil }
func (c Fs) MkdirAll(string, os.FileMode) error { return nil }
func (c Fs) Open(string) (afero.File, error)    { return nil, nil }
func (c Fs) OpenFile(string, int, os.FileMode) (afero.File, error) {
	return nil, nil
}
func (c Fs) Remove(string) error              { return nil }
func (c Fs) RemoveAll(string) error           { return nil }
func (c Fs) Rename(string, string) error      { return nil }
func (c Fs) Stat(string) (os.FileInfo, error) { return nil, nil }

func (c Fs) Chmod(string, os.FileMode) error            { return nil }
func (c Fs) Chtimes(string, time.Time, time.Time) error { return nil }
func (c Fs) Name() string                               { return "" }
