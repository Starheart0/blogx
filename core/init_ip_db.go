package core

import (
	"fmt"
	"strings"

	iputils "blogx_server/utils/ip"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"github.com/sirupsen/logrus"
)

var searcher *xdb.Searcher

func InitIPDB() {
	var dbPath = "init/ip2region.xdb"
	_searcher, err := xdb.NewWithFileOnly(xdb.IPv4, dbPath)
	if err != nil {
		fmt.Printf("ip database failed to create searcher: %s\n", err.Error())
		return
	}
	searcher = _searcher
}

func GetIpAddr(ip string) (addr string) {
	if iputils.HasLocalIPAddr(ip) {
		return "内网"
	}
	region, err := searcher.SearchByStr(ip)
	if err != nil {
		return
	}
	_addrList := strings.Split(region, "|")
	if len(_addrList) != 5 {
		logrus.Warnf("unknown ip addr %s", ip)
		return "unknown ip addr"
	}
	country := _addrList[0]
	province := _addrList[1]
	city := _addrList[2]
	if province != "0" && city != "0" {
		return fmt.Sprintf("%s·%s", province, city)
	}
	if country != "0" && province != "0" {
		return fmt.Sprintf("%s·%s", country, province)
	}
	if country != "0" {
		return country
	}
	return region
	// 中国|广东省|广州市|中国教育网|CN
}
