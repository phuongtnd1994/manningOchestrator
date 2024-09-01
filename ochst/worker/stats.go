package worker

import "github.com/c9s/goprocinfo/linux"

type Stats struct {
	MemStats  *linux.MemInfo
	DiskStats *linux.Disk
	CpuStats  *linux.CPUStat
	LoadStats *linux.LoadAvg
}

func (s *Stats) MemTotalKb() uint64 {
	return s.MemStats.MemTotal
}
