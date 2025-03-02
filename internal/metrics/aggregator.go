package metrics

type MetricsAggregator struct {
	nodeMetrics   map[string]*NodeMetrics
	podMetrics    map[string]*PodMetrics
	retentionDays int
}

type NodeMetrics struct {
	CPU    []MetricPoint
	Memory []MetricPoint
	Disk   []MetricPoint
}

type PodMetrics struct {
	CPU    []MetricPoint
	Memory []MetricPoint
}

type MetricPoint struct {
	Timestamp int64
	Value     float64
}
