package main

import router "github.com/ghjan/golesson/route"

func main() {
	r := router.Router

	router.SetRouter()

	r.Run(":1234")
}
