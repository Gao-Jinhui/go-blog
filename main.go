package main

import (
	"go-blog/routers"
)

func main() {
	router := routers.InitRouter()
	router.Run()
}
