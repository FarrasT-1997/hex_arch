package main

import (
	"hex/config/connectDB"
	"hex/config/setting"
)

func main() {
	config := setting.GetConfig()

	db := connectDB.Connection(config)

}
