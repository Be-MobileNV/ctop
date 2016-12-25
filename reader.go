package main

import (
	"github.com/fsouza/go-dockerclient"
)

type StatReader struct {
	CPUUtil  int
	MemUsage int64
	MemLimit int64
	//MemPercent int64
	lastCpu    float64
	lastSysCpu float64
}

func (s *StatReader) Read(stats *docker.Stats) {
	s.ReadCPU(stats)
	s.ReadMem(stats)
}

func (s *StatReader) ReadCPU(stats *docker.Stats) {
	ncpus := float64(len(stats.CPUStats.CPUUsage.PercpuUsage))
	total := float64(stats.CPUStats.CPUUsage.TotalUsage)
	system := float64(stats.CPUStats.SystemCPUUsage)

	cpudiff := total - s.lastCpu
	syscpudiff := system - s.lastSysCpu
	s.CPUUtil = round((cpudiff / syscpudiff * 100) * ncpus)
	s.lastCpu = total
	s.lastSysCpu = system
}

func (s *StatReader) ReadMem(stats *docker.Stats) {
	s.MemUsage = int64(stats.MemoryStats.Usage)
	s.MemLimit = int64(stats.MemoryStats.Limit)
	//s.MemPercent = round((float64(cur) / float64(limit)) * 100)
}