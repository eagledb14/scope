package main

import (
	// "fmt"
	// "net"
	// "os"
	"fmt"

	"github.com/spf13/pflag"
)

func main() {
	var file string
	var ip string

	pflag.StringVarP(&file, "file", "f", "", "CIDR File Path")
	pflag.StringVar(&ip, "ip", "", "IP file path")

	pflag.Parse()
	pflag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("\tscope [ip] [CIDR]")
		fmt.Println("\tscope [ip] [CIDR]")
		fmt.Println("\tscope [ip] -f [CIDR File]")
		fmt.Println("\tscope -ip [ip] -f [CIDR File]")
		fmt.Println("Options:")
		pflag.PrintDefaults()
	}

	args := pflag.Args()

	// if len(args) == 
	fmt.Println(args)

	if len(args) == 2 {
		pflag.Usage()
	} else if len(args) == 1 {

	} 
}




