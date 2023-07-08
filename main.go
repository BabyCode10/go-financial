package main

import (
	"go-financial/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.Api()
}
