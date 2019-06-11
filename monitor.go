package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

type data struct {
	loadCPU5 string
	memUsed  string
	memFree  string
	diskUsed string
}

func getInfoHW() data {
	time.Sleep(100 * time.Millisecond) //time of each request information
	m, _ := mem.VirtualMemory()
	l, _ := load.Avg()
	d, _ := disk.Usage("/")
	var dataHW data
	dataHW.memFree = fmt.Sprintf("%v", m.Free)
	dataHW.memUsed = fmt.Sprintf("%f", m.UsedPercent)
	dataHW.loadCPU5 = fmt.Sprintf("%f", l.Load5)
	dataHW.diskUsed = fmt.Sprintf("%f", d.UsedPercent)
	fmt.Println(dataHW)
	return dataHW
}

func main() {
	for {
		getInfoHW()
	}
}
