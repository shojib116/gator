package main

import (
	"fmt"

	"github.com/shojib116/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	cfg.CurrentUserName = "shojib116"
	cfg.SetUser()

	newCfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(newCfg)
}
