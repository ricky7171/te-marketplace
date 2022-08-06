package main

import "github.com/ricky7171/te-marketplace/internal/injector"

func main() {

	// router := router.Router{}
	// router.InitGin()

	router := injector.InitializedRouter()
	router.Run()

	//accountrouter.init()
	// router.InitApi()
	// router.Run()

}
