package main

import (
	"ValorantApi/code/server"
	"ValorantApi/code/test"
)

func main() {
	test.RajoutDonneesJson()
	server.GenererPagesPersonnages()
	server.Server()

}
