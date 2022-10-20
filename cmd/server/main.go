package main

import "github.com/misael28/GoApi/configs"

func main(){
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}