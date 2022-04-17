package main

import (
	utils "people-api/utils"
)

func main() {
	utils.InitLogging()
	router := utils.CreateRouter()

	router.Run(":8080")
}
