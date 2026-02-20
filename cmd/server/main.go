package main

import (
	"gin-app/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
