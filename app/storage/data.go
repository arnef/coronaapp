package storage

import (
	"io/ioutil"
	"os"
	"path"

	log "github.com/sirupsen/logrus"
)

const AppName string = "coronaapp.de.arnef"

func UserConfigDir() (string, error) {
	config, err := os.UserConfigDir()
	if err != nil {
		log.Errorln(err)
		return "", err
	}
	appConfigPath := path.Join(config, AppName)
	if err := os.Mkdir(appConfigPath, 0700); !os.IsExist(err) {
		log.Errorln(err)
		return "", err
	}
	return appConfigPath, nil
}

func WriteFile(filename string, data []byte) error {
	dir, err := UserConfigDir()
	if err != nil {
		log.Errorln(err)
		return err
	}

	return ioutil.WriteFile(path.Join(dir, filename), data, 0600)
}

func ReadFile(file string) ([]byte, error) {
	dir, err := UserConfigDir()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	return ioutil.ReadFile(path.Join(dir, file))
}

func ReadDir(name string) ([]string, error) {
	dir, err := UserConfigDir()
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	fullPath := dir
	if len(name) > 0 {
		fullPath = path.Join(dir, name)
	}

	outFiles, err := ioutil.ReadDir(fullPath)
	if err != nil {
		log.Errorln(err)
		return nil, err
	}
	files := make([]string, len(outFiles))

	for i := range outFiles {

		files[i] = outFiles[i].Name()
	}

	return files, nil
}
