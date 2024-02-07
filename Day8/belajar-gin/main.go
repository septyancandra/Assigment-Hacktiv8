package main

import (
	"belajar-gin/routers"

	_ "github.com/lib/pq"
)

func main() {

	var PORT = ":3000"

	routers.StartServer().Run(PORT)
}
