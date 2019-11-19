package main

import "github.com/yaien/mask/cmd"

func main() {
	root := cmd.Root()
	root.AddCommand(cmd.Decrypt())
	root.AddCommand(cmd.Encrypt())
	root.Execute()
}
