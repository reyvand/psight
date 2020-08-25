package main

import (
	"flag"
	"fmt"
	"os"
	"psight"
	"reflect"
	"strconv"
	"strings"
)

func help() {
	banner :=
		`
   __   __        _       _     _   
  / /   \ \      (_)     | |   | |  
 | | _ __| |  ___ _  __ _| |__ | |_ 
/ / | '_ \\ \/ __| |/ _' | '_ \| __|
\ \ | |_) / /\__ \ | (_| | | | | |_ 
 | || .__/ | |___/_|\__, |_| |_|\__|
  \_\ | /_/          __/ |          
    |_|             |___/      v.0.1
`
	banner += "\nUsage:\n\t-t IP\tTarget IP\nOptional:\n\t-p PORT\tTarget Port\n\nExamples:\n\n- Discover all open port from target [default]\n  $ psight -t 127.0.0.1\n- Check single port from target\n  $ psight -t 127.0.0.1 -p 22\n- Check multiple port from target\n  $ psight -t 127.0.0.1 -p 21,22\n- Check range of port from target\n  $ psight -t 127.0.0.1 -p 22-25\n"
	fmt.Println(banner)
}

func main() {
	if len(os.Args) == 1 {
		help()
	} else {
		destIP := flag.String("t", "", "IP target")
		destPort := flag.String("p", "", "Port")
		flag.Parse()
		if len(*destIP) == 0 {
			fmt.Println("Invalid command: missing IP target")
		} else {
			if len(*destPort) == 0 {
				psight.ScanAll(*destIP)
			} else {
				if strings.Contains(*destPort, ",") {
					rport := strings.Split(*destPort, ",")
					for _, p := range rport {
						scan := psight.Connect(*destIP, p)
						result := reflect.ValueOf(&scan).Elem()
						ip := result.FieldByName("Ip")
						port := result.FieldByName("Rport")
						status := result.FieldByName("Status")
						fmt.Printf("[%s] %s/%s\n", status, ip, port)
					}
				} else if strings.Contains(*destPort, "-") {
					rport := strings.Split(*destPort, "-")
					low, _ := strconv.Atoi(rport[0])
					high, _ := strconv.Atoi(rport[1])
					for p := low; p <= high; p++ {
						scan := psight.Connect(*destIP, strconv.Itoa(p))
						result := reflect.ValueOf(&scan).Elem()
						ip := result.FieldByName("Ip")
						port := result.FieldByName("Rport")
						status := result.FieldByName("Status")
						fmt.Printf("[%s] %s/%s\n", status, ip, port)
					}
				} else {
					scan := psight.Connect(*destIP, *destPort)
					result := reflect.ValueOf(&scan).Elem()
					ip := result.FieldByName("Ip")
					port := result.FieldByName("Rport")
					status := result.FieldByName("Status")
					fmt.Printf("[%s] %s/%s\n", status, ip, port)
				}

			}
		}
	}
}
