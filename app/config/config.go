package config

import (
	"errors"
	"os"

	dotenv "github.com/golobby/dotenv"
)

func Init() (Conf, error) {
	conf := Conf{}
	//Load Environment file
	file, err := os.Open(".env")
	if err != nil {
		return conf, errors.New("error loading .env file")
	}
	err = dotenv.NewDecoder(file).Decode(&conf)
	if err != nil {
		return conf, errors.New("cannot decode .env file")
	}
	return conf, err
}
