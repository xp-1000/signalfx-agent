// Code generated by monitor-code-gen. DO NOT EDIT.

package cgroups

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "cgroups"

const (
	groupCPU           = "cpu"
	groupCPUCfs        = "cpu-cfs"
	groupCpuacct       = "cpuacct"
	groupCpuacctPerCPU = "cpuacct-per-cpu"
	groupMemory        = "memory"
)

var groupSet = map[string]bool{
	groupCPU:           true,
	groupCPUCfs:        true,
	groupCpuacct:       true,
	groupCpuacctPerCPU: true,
	groupMemory:        true,
}

const (
	cgroupCPUCfsPeriodUs                    = "cgroup.cpu_cfs_period_us"
	cgroupCPUCfsQuotaUs                     = "cgroup.cpu_cfs_quota_us"
	cgroupCPUShares                         = "cgroup.cpu_shares"
	cgroupCPUStatNrPeriods                  = "cgroup.cpu_stat_nr_periods"
	cgroupCPUStatNrThrottled                = "cgroup.cpu_stat_nr_throttled"
	cgroupCPUStatThrottledTime              = "cgroup.cpu_stat_throttled_time"
	cgroupCpuacctUsageNs                    = "cgroup.cpuacct_usage_ns"
	cgroupCpuacctUsageNsPerCPU              = "cgroup.cpuacct_usage_ns_per_cpu"
	cgroupCpuacctUsageSystemNs              = "cgroup.cpuacct_usage_system_ns"
	cgroupCpuacctUsageSystemNsPerCPU        = "cgroup.cpuacct_usage_system_ns_per_cpu"
	cgroupCpuacctUsageUserNs                = "cgroup.cpuacct_usage_user_ns"
	cgroupCpuacctUsageUserNsPerCPU          = "cgroup.cpuacct_usage_user_ns_per_cpu"
	cgroupMemoryFailcnt                     = "cgroup.memory_failcnt"
	cgroupMemoryLimitInBytes                = "cgroup.memory_limit_in_bytes"
	cgroupMemoryStatActiveAnon              = "cgroup.memory_stat_active_anon"
	cgroupMemoryStatActiveFile              = "cgroup.memory_stat_active_file"
	cgroupMemoryStatCache                   = "cgroup.memory_stat_cache"
	cgroupMemoryStatDirty                   = "cgroup.memory_stat_dirty"
	cgroupMemoryStatHierarchicalMemoryLimit = "cgroup.memory_stat_hierarchical_memory_limit"
	cgroupMemoryStatInactiveAnon            = "cgroup.memory_stat_inactive_anon"
	cgroupMemoryStatInactiveFile            = "cgroup.memory_stat_inactive_file"
	cgroupMemoryStatMappedFile              = "cgroup.memory_stat_mapped_file"
	cgroupMemoryStatPgfault                 = "cgroup.memory_stat_pgfault"
	cgroupMemoryStatPgmajfault              = "cgroup.memory_stat_pgmajfault"
	cgroupMemoryStatPgpgin                  = "cgroup.memory_stat_pgpgin"
	cgroupMemoryStatPgpgout                 = "cgroup.memory_stat_pgpgout"
	cgroupMemoryStatRss                     = "cgroup.memory_stat_rss"
	cgroupMemoryStatRssHuge                 = "cgroup.memory_stat_rss_huge"
	cgroupMemoryStatShmem                   = "cgroup.memory_stat_shmem"
	cgroupMemoryStatTotalActiveAnon         = "cgroup.memory_stat_total_active_anon"
	cgroupMemoryStatTotalActiveFile         = "cgroup.memory_stat_total_active_file"
	cgroupMemoryStatTotalCache              = "cgroup.memory_stat_total_cache"
	cgroupMemoryStatTotalDirty              = "cgroup.memory_stat_total_dirty"
	cgroupMemoryStatTotalInactiveAnon       = "cgroup.memory_stat_total_inactive_anon"
	cgroupMemoryStatTotalInactiveFile       = "cgroup.memory_stat_total_inactive_file"
	cgroupMemoryStatTotalMappedFile         = "cgroup.memory_stat_total_mapped_file"
	cgroupMemoryStatTotalPgfault            = "cgroup.memory_stat_total_pgfault"
	cgroupMemoryStatTotalPgmajfault         = "cgroup.memory_stat_total_pgmajfault"
	cgroupMemoryStatTotalPgpgin             = "cgroup.memory_stat_total_pgpgin"
	cgroupMemoryStatTotalPgpgout            = "cgroup.memory_stat_total_pgpgout"
	cgroupMemoryStatTotalRss                = "cgroup.memory_stat_total_rss"
	cgroupMemoryStatTotalRssHuge            = "cgroup.memory_stat_total_rss_huge"
	cgroupMemoryStatTotalShmem              = "cgroup.memory_stat_total_shmem"
	cgroupMemoryStatTotalUnevictable        = "cgroup.memory_stat_total_unevictable"
	cgroupMemoryStatTotalWriteback          = "cgroup.memory_stat_total_writeback"
	cgroupMemoryStatUnevictable             = "cgroup.memory_stat_unevictable"
	cgroupMemoryStatWriteback               = "cgroup.memory_stat_writeback"
)

var metricSet = map[string]monitors.MetricInfo{
	cgroupCPUCfsPeriodUs:                    {Type: datapoint.Gauge, Group: groupCPUCfs},
	cgroupCPUCfsQuotaUs:                     {Type: datapoint.Gauge, Group: groupCPUCfs},
	cgroupCPUShares:                         {Type: datapoint.Gauge, Group: groupCPU},
	cgroupCPUStatNrPeriods:                  {Type: datapoint.Counter, Group: groupCPUCfs},
	cgroupCPUStatNrThrottled:                {Type: datapoint.Counter, Group: groupCPUCfs},
	cgroupCPUStatThrottledTime:              {Type: datapoint.Counter, Group: groupCPUCfs},
	cgroupCpuacctUsageNs:                    {Type: datapoint.Counter, Group: groupCpuacct},
	cgroupCpuacctUsageNsPerCPU:              {Type: datapoint.Counter, Group: groupCpuacctPerCPU},
	cgroupCpuacctUsageSystemNs:              {Type: datapoint.Counter, Group: groupCpuacct},
	cgroupCpuacctUsageSystemNsPerCPU:        {Type: datapoint.Counter, Group: groupCpuacctPerCPU},
	cgroupCpuacctUsageUserNs:                {Type: datapoint.Counter, Group: groupCpuacct},
	cgroupCpuacctUsageUserNsPerCPU:          {Type: datapoint.Counter, Group: groupCpuacctPerCPU},
	cgroupMemoryFailcnt:                     {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryLimitInBytes:                {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatActiveAnon:              {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatActiveFile:              {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatCache:                   {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatDirty:                   {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatHierarchicalMemoryLimit: {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatInactiveAnon:            {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatInactiveFile:            {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatMappedFile:              {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatPgfault:                 {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatPgmajfault:              {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatPgpgin:                  {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatPgpgout:                 {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatRss:                     {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatRssHuge:                 {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatShmem:                   {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalActiveAnon:         {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalActiveFile:         {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalCache:              {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalDirty:              {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalInactiveAnon:       {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalInactiveFile:       {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalMappedFile:         {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalPgfault:            {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatTotalPgmajfault:         {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatTotalPgpgin:             {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatTotalPgpgout:            {Type: datapoint.Counter, Group: groupMemory},
	cgroupMemoryStatTotalRss:                {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalRssHuge:            {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalShmem:              {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalUnevictable:        {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatTotalWriteback:          {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatUnevictable:             {Type: datapoint.Gauge, Group: groupMemory},
	cgroupMemoryStatWriteback:               {Type: datapoint.Gauge, Group: groupMemory},
}

var defaultMetrics = map[string]bool{
	cgroupCpuacctUsageNs: true,
}

var groupMetricsMap = map[string][]string{
	groupCPU: []string{
		cgroupCPUShares,
	},
	groupCPUCfs: []string{
		cgroupCPUCfsPeriodUs,
		cgroupCPUCfsQuotaUs,
		cgroupCPUStatNrPeriods,
		cgroupCPUStatNrThrottled,
		cgroupCPUStatThrottledTime,
	},
	groupCpuacct: []string{
		cgroupCpuacctUsageNs,
		cgroupCpuacctUsageSystemNs,
		cgroupCpuacctUsageUserNs,
	},
	groupCpuacctPerCPU: []string{
		cgroupCpuacctUsageNsPerCPU,
		cgroupCpuacctUsageSystemNsPerCPU,
		cgroupCpuacctUsageUserNsPerCPU,
	},
	groupMemory: []string{
		cgroupMemoryFailcnt,
		cgroupMemoryLimitInBytes,
		cgroupMemoryStatActiveAnon,
		cgroupMemoryStatActiveFile,
		cgroupMemoryStatCache,
		cgroupMemoryStatDirty,
		cgroupMemoryStatHierarchicalMemoryLimit,
		cgroupMemoryStatInactiveAnon,
		cgroupMemoryStatInactiveFile,
		cgroupMemoryStatMappedFile,
		cgroupMemoryStatPgfault,
		cgroupMemoryStatPgmajfault,
		cgroupMemoryStatPgpgin,
		cgroupMemoryStatPgpgout,
		cgroupMemoryStatRss,
		cgroupMemoryStatRssHuge,
		cgroupMemoryStatShmem,
		cgroupMemoryStatTotalActiveAnon,
		cgroupMemoryStatTotalActiveFile,
		cgroupMemoryStatTotalCache,
		cgroupMemoryStatTotalDirty,
		cgroupMemoryStatTotalInactiveAnon,
		cgroupMemoryStatTotalInactiveFile,
		cgroupMemoryStatTotalMappedFile,
		cgroupMemoryStatTotalPgfault,
		cgroupMemoryStatTotalPgmajfault,
		cgroupMemoryStatTotalPgpgin,
		cgroupMemoryStatTotalPgpgout,
		cgroupMemoryStatTotalRss,
		cgroupMemoryStatTotalRssHuge,
		cgroupMemoryStatTotalShmem,
		cgroupMemoryStatTotalUnevictable,
		cgroupMemoryStatTotalWriteback,
		cgroupMemoryStatUnevictable,
		cgroupMemoryStatWriteback,
	},
}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "cgroups",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         true,
}
