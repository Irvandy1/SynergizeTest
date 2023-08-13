package config

import (
	"fmt"
)

func Init() (err error) {
	err = initPostgres()
	if err != nil {
		fmt.Println(err)
	}

	initRedis()

	return
}
