package psight

import (
	"fmt"
	"net"
	"reflect"
	"strconv"
	"time"
)

type ScanRes struct {
	Ip, Rport, Status string
}

func Connect(ip, port string) ScanRes {
	target := fmt.Sprintf("%s:%s", ip, port)
	conn, error := net.DialTimeout("tcp", target, 500*time.Millisecond)
	if error != nil {
		return ScanRes{ip, port, "closed"}
	}
	conn.Close()
	return ScanRes{ip, port, "open"}
}

func ScanAll(ip string) {
	slice := make([]int, 0)
	for i := 1; i <= 65535; i++ {
		slice = append(slice, i)
	}
	batch := 1024
	for i := 0; i < len(slice); i += batch {
		j := i + batch
		if j > len(slice) {
			j = len(slice)
		}
		for k := i; k <= j; k++ {
			scan := Connect(ip, strconv.Itoa(k))
			result := reflect.ValueOf(&scan).Elem()
			ip := result.FieldByName("Ip")
			port := result.FieldByName("Rport")
			status := result.FieldByName("Status").Interface()
			if status == "open" {
				fmt.Printf("[%s] %s/%s\n", status, ip, port)
			}
		}
	}
}
