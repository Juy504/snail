package main

import (
	"snail/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}
