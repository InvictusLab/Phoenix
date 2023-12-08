package main

import "phoenix/router"

func main() {
	r := router.App()
	r.Run(":8081")
}
