package main

import (
	"blogx_server/core"
	"fmt"
)

func main() {
	ip2region()
}
func ip2region() {
	core.InitIPDB()
	fmt.Println(core.GetIpAddr("192.168.0.1"))
	fmt.Println(core.GetIpAddr("202.116.31.153"))
	fmt.Println(core.GetIpAddr("12.135.135.1"))
	fmt.Println(core.GetIpAddr("145.15.43.15"))
}
