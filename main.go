package main

import (
	"app/config"
	"app/routers"
)

func main() {
	conf := config.NewConfig()
	router := routers.NewRouter(conf)
	router.RouterAuth()
	
	router.Run()
}
