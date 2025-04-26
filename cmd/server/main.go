package main

import (
	"go-social-network/internal/routers"
)

func main() {
	r := routers.NewRouter()

	err := r.Run(":9090")
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:9090 (for windows "localhost:9090")
}
