package utilities

import (
	"swapper/package/twitter"
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Accounts                 struct {
		AccountToRecieveUsername twitter.Account
		AccountWithRareUsername  twitter.Account
	}
}

func ReadConfig(path string) Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	config := Config{}

	err = json.Unmarshal([]byte(file), &config)
	if err != nil {
		panic(err)
	}

	return config
}
