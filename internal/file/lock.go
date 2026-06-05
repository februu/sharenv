package file

import (
	"fmt"
	"os"

	"github.com/februu/sharenv/internal/models"
	"github.com/pelletier/go-toml/v2"
)

func ReadLockFile() (map[string]models.Entry, error) {
	var values map[string]models.Entry
	data, err := os.ReadFile(".sharenv")
	if err != nil {
		return nil, err
	}
	err = toml.Unmarshal(data, &values)
	if err != nil {
		return nil, err
	}
	if values == nil {
		values = make(map[string]models.Entry)
	}
	return values, nil
}

func WriteLockFile(values map[string]models.Entry) error {
	data, err := toml.Marshal(values)
	if err != nil {
		return err
	}
	return os.WriteFile(".sharenv", data, 0600)
}

func CreateLockFile() error {
	f, err := os.OpenFile(".sharenv", os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if os.IsExist(err) {
		return fmt.Errorf("lock file already exists, is sharenv already initialized?")
	}
	if err != nil {
		return err
	}
	return f.Close()
}
