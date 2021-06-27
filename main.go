package main

import (
	"SimpleTool/cmd"
	"log"
)

func main()  {
	err := cmd.Execute()
	if err != nil{
		log.Fatal("cmd.Execute err: %v",err)
	}
}