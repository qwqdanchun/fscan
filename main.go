package main

import (
	"fmt"
	"example.com/fxscan/Plugins"
	"example.com/fxscan/common"
	"time"
)

func main() {
	start := time.Now()
	var Info common.HostInfo
	common.Flag(&Info)
	common.Parse(&Info)
	Plugins.Scan(Info)
	t := time.Now().Sub(start)
	fmt.Printf("[*] 扫描已结束,耗时: %s\n", t)
}
