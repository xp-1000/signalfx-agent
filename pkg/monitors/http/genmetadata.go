// Code generated by monitor-code-gen. DO NOT EDIT.

package http

import (
	"github.com/signalfx/golib/v3/datapoint"
	"github.com/signalfx/signalfx-agent/pkg/monitors"
)

const monitorType = "http"

var groupSet = map[string]bool{}

const (
	httpCertExpiry    = "http.cert_expiry"
	httpCertValid     = "http.cert_valid"
	httpCodeMatched   = "http.code_matched"
	httpContentLength = "http.content_length"
	httpRegexMatched  = "http.regex_matched"
	httpResponseTime  = "http.response_time"
	httpStatusCode    = "http.status_code"
)

var metricSet = map[string]monitors.MetricInfo{
	httpCertExpiry:    {Type: datapoint.Gauge},
	httpCertValid:     {Type: datapoint.Gauge},
	httpCodeMatched:   {Type: datapoint.Gauge},
	httpContentLength: {Type: datapoint.Gauge},
	httpRegexMatched:  {Type: datapoint.Gauge},
	httpResponseTime:  {Type: datapoint.Gauge},
	httpStatusCode:    {Type: datapoint.Gauge},
}

var defaultMetrics = map[string]bool{
	httpCertExpiry:    true,
	httpCertValid:     true,
	httpCodeMatched:   true,
	httpContentLength: true,
	httpRegexMatched:  true,
	httpResponseTime:  true,
	httpStatusCode:    true,
}

var groupMetricsMap = map[string][]string{}

var monitorMetadata = monitors.Metadata{
	MonitorType:     "http",
	DefaultMetrics:  defaultMetrics,
	Metrics:         metricSet,
	SendUnknown:     false,
	Groups:          groupSet,
	GroupMetricsMap: groupMetricsMap,
	SendAll:         false,
}
