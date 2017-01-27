package nullfs_test

import (
	"errors"
	"testing"
	"time"

	nullfs "github.com/jncornett/afero-nullfs"
)

func TestNullFS(t *testing.T) {
	valTests := []func(c nullfs.Fs) error{
		func(c nullfs.Fs) error {
			_, err := c.Create("")
			return err
		},
		func(c nullfs.Fs) error { return c.Mkdir("", 0) },
		func(c nullfs.Fs) error { return c.MkdirAll("", 0) },
		func(c nullfs.Fs) error {
			_, err := c.Open("")
			return err
		},
		func(c nullfs.Fs) error {
			_, err := c.OpenFile("", 0, 0)
			return err
		},
		func(c nullfs.Fs) error { return c.Remove("") },
		func(c nullfs.Fs) error { return c.RemoveAll("") },
		func(c nullfs.Fs) error { return c.Rename("", "") },
		func(c nullfs.Fs) error {
			_, err := c.Stat("")
			return err
		},
		func(c nullfs.Fs) error { return c.Chmod("", 0) },
		func(c nullfs.Fs) error { return c.Chtimes("", time.Time{}, time.Time{}) },
		func(c nullfs.Fs) error {
			s := c.Name()
			if s != "" {
				return errors.New("expected an empty string")
			}
			return nil
		},
	}

	for _, fn := range valTests {
		t.Run("", func(t *testing.T) {
			err := fn(nullfs.Fs{})
			if err != nil {
				t.Error(err)
			}
		})
	}

	ptrTests := []func(c *nullfs.Fs) error{
		func(c *nullfs.Fs) error {
			_, err := c.Create("")
			return err
		},
		func(c *nullfs.Fs) error { return c.Mkdir("", 0) },
		func(c *nullfs.Fs) error { return c.MkdirAll("", 0) },
		func(c *nullfs.Fs) error {
			_, err := c.Open("")
			return err
		},
		func(c *nullfs.Fs) error {
			_, err := c.OpenFile("", 0, 0)
			return err
		},
		func(c *nullfs.Fs) error { return c.Remove("") },
		func(c *nullfs.Fs) error { return c.RemoveAll("") },
		func(c *nullfs.Fs) error { return c.Rename("", "") },
		func(c *nullfs.Fs) error {
			_, err := c.Stat("")
			return err
		},
		func(c *nullfs.Fs) error { return c.Chmod("", 0) },
		func(c *nullfs.Fs) error { return c.Chtimes("", time.Time{}, time.Time{}) },
		func(c *nullfs.Fs) error {
			s := c.Name()
			if s != "" {
				return errors.New("expected an empty string")
			}
			return nil
		},
	}

	for _, fn := range ptrTests {
		t.Run("", func(t *testing.T) {
			err := fn(&nullfs.Fs{})
			if err != nil {
				t.Error(err)
			}
		})
	}
}
