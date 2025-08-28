package main

import (
	"awtrix3ctl/cmd"
	"log"
)

func main() {
	rootCmd := cmd.NewRootCmd()
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal(err)
	}
}
