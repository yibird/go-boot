package sys

import (
	"context"
	"os"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/docker"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/shirou/gopsutil/v3/winservices"
)

type MonitorApi struct{}

func init() {
	// 获取cpu信息
	cpu.Info()
	// 获取内存信息
	mem.VirtualMemory()
	// 获取磁盘信息
	disk.Partitions(true)
	// 获取网络信息
	net.Interfaces()
	// 获取当前进程信息
	process.NewProcess(int32(os.Getpid()))
	// 获取Windows服务信息
	winservices.NewService("go-boot")
	// 获取host信息
	host.Info()
	// 获取Docker信息
	docker.GetDockerStatWithContext(context.Background())
}

// 获取监控信息,包括CPU、内存、磁盘、网络、当前进程信息
func (s *MonitorApi) getMonitorInfo() {

}
