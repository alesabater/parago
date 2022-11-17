package para

import (
	"os"
	"path"
)

type Client interface {

	// Init will initializer PARA directory structure or return error
	Init() error
}

func CreateConfig(str string) error {
	_, err := os.Stat(str)

	if os.IsNotExist(err) {
		if err = os.MkdirAll(path.Dir(str), 0700); err != nil {
			return err
		}
		_, err = os.Create(str)
		if err != nil {
			return err
		}
	}

	if err != nil {
		return err
	}
	return nil
}
