package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/spf13/pflag"
)

func main() {
	var file string
	var ipFileName string

	pflag.StringVarP(&file, "file", "f", "", "CIDR File Path")
	pflag.StringVar(&ipFileName, "ip", "", "IP file path")

	pflag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("\tscope [ip] [CIDR]")
		fmt.Println("\tscope [ip] -f [CIDR File]")
		fmt.Println("\tscope --ip [ip] -f [CIDR File]")
		fmt.Println("Options:")
		pflag.PrintDefaults()
	}
	pflag.Parse()

	args := pflag.Args()

	if len(args) == 2 {
		if ipInScope(args[0], args[1]) {
			fmt.Println("[IN SCOPE]", args[0])
		} else {
			fmt.Println("[OUT OF SCOPE]", args[1])
		}
	} else if len(args) == 1 {
		// check if file is checked

		if file == "" {
			pflag.Usage()
			return
		}

		cidrFile, err := os.ReadFile(file)
		if err != nil {
			pflag.Usage()
			return
		}

		cidrList := strings.Split(string(cidrFile), "\n")
		cidrList = filterEmptyString(cidrList)

		for _, cidr := range cidrList {
			if ipInScope(args[0], cidr) {
				fmt.Println("[IN SCOPE]", args[0])
				return
			}
		}

		fmt.Println("[OUT OF SCOPE]", args[0])

	} else if len(args) == 0 {
		// both flags have to be set

		if file == "" || ipFileName == "" {
			pflag.Usage()
			return
		}

		cidrFile, err := os.ReadFile(file)
		if err != nil {
			pflag.Usage()
			return
		}

		ipFile, err := os.ReadFile(ipFileName)
		if err != nil {
			pflag.Usage()
			return
		}

		cidrList := strings.Split(string(cidrFile), "\n")
		ipList := strings.Split(string(ipFile), "\n")

		cidrList = filterEmptyString(cidrList)
		ipList = filterEmptyString(ipList)

		for _, ip := range ipList {
			scope := false
			for _, cidr := range cidrList {
				if ipInScope(ip, cidr) {
					scope = true
					break
				}
			}

			if scope == true {
				fmt.Println("[IN SCOPE]", ip)
			} else {
				fmt.Println("[OUT OF SCOPE]", ip)
			}
		}

	} else {
		pflag.Usage()
	}
}

func ipInScope(ip string, cidr string) bool {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return false
	}

	parseIp := net.ParseIP(ip)
	if parseIp == nil {
		return false
	}

	return ipNet.Contains(parseIp)
}

func filterEmptyString(input []string) []string {
	result := []string{}
	for _, str := range input {
		if str != "" {
			result = append(result, str)
		}
	}

	return result
}
