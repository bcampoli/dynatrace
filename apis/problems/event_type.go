package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// The type of an event.
type eventType struct {
	fmt.Stringer
	name string
}

type elasticLoadBalancerEventType struct {
	HighUnhealthyHostRate  eventType
	HighFailureRate        eventType
	HighBackendFailureRate eventType
}

type webCheckEventType struct {
	GlobalOutage eventType
	LocalOutage  eventType
}

type pgiEventType struct {
	PGIOfServiceUnavailable eventType
	HAProxy                 pgiHAProxyEventType
	MySQL                   pgiMySQLEventType
	RMQ                     pgiRMQEventType
}

type pgiRMQEventType struct {
	HighMemUsage      eventType
	HighFileDescUsage eventType
	HighSocketsUsage  eventType
	HighProcessUsage  eventType
	LowDiskSpace      eventType
}

type pgiMySQLEventType struct {
	SlowQueriesRateHigh eventType
}

type pgiHAProxyEventType struct {
	QueuedRequestsHigh eventType
	SessionUsageHigh   eventType
}

type customEventType struct {
	Annotation    eventType
	Deployment    eventType
	Configuration eventType
	Info          eventType
	Alert         eventType
}

type processEventType struct {
	Crashed     eventType
	Unavailable eventType
	Custom      processCustomEventType
	Log         processLogEventType
}

type processCustomEventType struct {
	Availability eventType
	Error        eventType
	Performance  eventType
}

type processLogEventType struct {
	Availability eventType
	Error        eventType
	Performance  eventType
}

type hostEventType struct {
	ConnectionLost        eventType
	ConnectionFailed      eventType
	Maintenance           eventType
	NoConnection          eventType
	GracefullyShutdown    eventType
	Timeout               eventType
	Shutdown              eventType
	DataStoreLowDiskSpace eventType
	DiskLowInodes         eventType
	Log                   hostLogEventType
}

type hostLogEventType struct {
	Availability eventType
	Error        eventType
	Performance  eventType
	Matched      eventType
}

type logEventType struct {
	Availability eventType
	Error        eventType
	Performance  eventType
	Matched      eventType
}

type osiEventType struct {
	Docker osiDockerEventType
}

type osiDockerEventType struct {
	DeviceMapper osiDockerDeviceMapperEventType
}

type osiDockerDeviceMapperEventType struct {
	LowDataSpace     eventType
	LowMetaDataSpace eventType
}

type syntheticEventType struct {
	Slowdown     eventType
	Availability eventType
}

type openStackEventType struct {
	KeyStone openStackKeyStoneEventType
}

type openStackKeyStoneEventType struct {
	Unhealthy eventType
	Slow      eventType
}

type eventTypes struct {
	CPUSaturated                   eventType
	MemorySaturated                eventType
	SlowDisk                       eventType
	HighLatency                    eventType
	InsufficientDiskQueueDepth     eventType
	HighGCActivity                 eventType
	MemoryResourcesExhausted       eventType
	ThreadsResourcesExhausted      eventType
	HighConnectivityFailures       eventType
	HighNetworkLossRate            eventType
	OverloadedStorage              eventType
	HighDroppedPacketsRate         eventType
	HighNetworkErrorRate           eventType
	HighNetworkUtilization         eventType
	LowDiskSpace                   eventType
	ConnectionLost                 eventType
	ServiceResponseTimeDegraded    eventType
	FailureRateIncreased           eventType
	UnexpectedHighLoad             eventType
	UnexpectedLowLoad              eventType
	UserActionDurationDegredation  eventType
	JavaScriptErrorRateIncreased   eventType
	ESXIStart                      eventType
	VirtualMachineShutdown         eventType
	LowStorageSpace                eventType
	EBSVolumeHighLatency           eventType
	MobileAppCrashRateIncreased    eventType
	CustomAppCrashRateIncreased    eventType
	HostOfServiceUnavilable        eventType
	DockerMemorySaturation         eventType
	RDSOfServiceUnavailable        eventType
	RDSRestartSequence             eventType
	LambdaFunctionHighErrorRate    eventType
	ProcessGroupLowInstanceCount   eventType
	PerformanceEvent               eventType
	ErrorEvent                     eventType
	AvailabilityEvent              eventType
	ResourceContention             eventType
	DatabaseConnectionFailure      eventType
	ApplicationJSFrameworkDetected eventType
	MonitoringUnavailable          eventType
	OpenStack                      openStackEventType
	Synthetic                      syntheticEventType
	OSI                            osiEventType
	Host                           hostEventType
	Process                        processEventType
	Custom                         customEventType
	ElasticLoadBalancer            elasticLoadBalancerEventType
	WebCheck                       webCheckEventType
	PGI                            pgiEventType
	Log                            logEventType
}

// EventTypes TODO: documentation
var EventTypes = eventTypes{
	ApplicationJSFrameworkDetected: eventType{name: "APPLICATION_JS_FRAMEWORK_DETECTED"},
	AvailabilityEvent:              eventType{name: "AVAILABILITY_EVENT"},
	ConnectionLost:                 eventType{name: "CONNECTION_LOST"},
	CPUSaturated:                   eventType{name: "CPU_SATURATED"},
	Custom: customEventType{
		Alert:         eventType{name: "CUSTOM_ALERT"},
		Annotation:    eventType{name: "CUSTOM_ANNOTATION"},
		Configuration: eventType{name: "CUSTOM_CONFIGURATION"},
		Deployment:    eventType{name: "CUSTOM_DEPLOYMENT"},
		Info:          eventType{name: "CUSTOM_INFO"},
	},
	CustomAppCrashRateIncreased: eventType{name: "CUSTOM_APP_CRASH_RATE_INCREASED"},
	DatabaseConnectionFailure:   eventType{name: "DATABASE_CONNECTION_FAILURE"},
	DockerMemorySaturation:      eventType{name: "DOCKER_MEMORY_SATURATION"},
	EBSVolumeHighLatency:        eventType{name: "EBS_VOLUME_HIGH_LATENCY"},
	ElasticLoadBalancer: elasticLoadBalancerEventType{
		HighBackendFailureRate: eventType{name: "ELASTIC_LOAD_BALANCER_HIGH_BACKEND_FAILURE_RATE"},
		HighFailureRate:        eventType{name: "ELASTIC_LOAD_BALANCER_HIGH_FAILURE_RATE"},
		HighUnhealthyHostRate:  eventType{name: "ELASTIC_LOAD_BALANCER_HIGH_UNHEALTHY_HOST_RATE"},
	},
	ErrorEvent:               eventType{name: "ERROR_EVENT"},
	ESXIStart:                eventType{name: "ESXI_START"},
	FailureRateIncreased:     eventType{name: "FAILURE_RATE_INCREASED"},
	HighConnectivityFailures: eventType{name: "HIGH_CONNECTIVITY_FAILURES"},
	HighDroppedPacketsRate:   eventType{name: "HIGH_DROPPED_PACKETS_RATE"},
	HighGCActivity:           eventType{name: "HIGH_GC_ACTIVITY"},
	HighLatency:              eventType{name: "HIGH_LATENCY"},
	HighNetworkErrorRate:     eventType{name: "HIGH_NETWORK_ERROR_RATE"},
	HighNetworkLossRate:      eventType{name: "HIGH_NETWORK_LOSS_RATE"},
	HighNetworkUtilization:   eventType{name: "HIGH_NETWORK_UTILIZATION"},
	Host: hostEventType{
		ConnectionFailed:      eventType{name: "HOST_CONNECTION_FAILED"},
		ConnectionLost:        eventType{name: "HOST_CONNECTION_LOST"},
		DataStoreLowDiskSpace: eventType{name: "HOST_DATASTORE_LOW_DISK_SPACE"},
		GracefullyShutdown:    eventType{name: "HOST_GRACEFULLY_SHUTDOWN"},
		DiskLowInodes:         eventType{name: "HOST_DISK_LOW_INODES"},
		Log: hostLogEventType{
			Availability: eventType{name: "HOST_LOG_AVAILABILITY"},
			Error:        eventType{name: "HOST_LOG_ERROR"},
			Matched:      eventType{name: "HOST_LOG_MATCHED"},
			Performance:  eventType{name: "HOST_LOG_PERFORMANCE"},
		},
		Maintenance:  eventType{name: "HOST_MAINTENANCE"},
		NoConnection: eventType{name: "HOST_NO_CONNECTION"},
		Shutdown:     eventType{name: "HOST_SHUTDOWN"},
		Timeout:      eventType{name: "HOST_TIMEOUT"},
	},
	HostOfServiceUnavilable:      eventType{name: "HOST_OF_SERVICE_UNAVAILABLE"},
	InsufficientDiskQueueDepth:   eventType{name: "INSUFFICIENT_DISK_QUEUE_DEPTH"},
	JavaScriptErrorRateIncreased: eventType{name: "JAVASCRIPT_ERROR_RATE_INCREASED"},
	LambdaFunctionHighErrorRate:  eventType{name: "LAMBDA_FUNCTION_HIGH_ERROR_RATE"},
	Log: logEventType{
		Availability: eventType{name: "LOG_AVAILABILITY"},
		Error:        eventType{name: "LOG_ERROR"},
		Matched:      eventType{name: "LOG_MATCHED"},
		Performance:  eventType{name: "LOG_PERFORMANCE"},
	},
	LowDiskSpace:                eventType{name: "LOW_DISK_SPACE"},
	LowStorageSpace:             eventType{name: "LOW_STORAGE_SPACE"},
	MemoryResourcesExhausted:    eventType{name: "MEMORY_RESOURCES_EXHAUSTED"},
	MemorySaturated:             eventType{name: "MEMORY_SATURATED"},
	MobileAppCrashRateIncreased: eventType{name: "MOBILE_APP_CRASH_RATE_INCREASED"},
	MonitoringUnavailable:       eventType{name: "MONITORING_UNAVAILABLE"},
	OpenStack: openStackEventType{
		KeyStone: openStackKeyStoneEventType{
			Slow:      eventType{name: "OPENSTACK_KEYSTONE_SLOW"},
			Unhealthy: eventType{name: "OPENSTACK_KEYSTONE_UNHEALTHY"},
		},
	},
	OSI: osiEventType{
		Docker: osiDockerEventType{
			DeviceMapper: osiDockerDeviceMapperEventType{
				LowDataSpace:     eventType{name: "OSI_DOCKER_DEVICEMAPPER_LOW_DATA_SPACE"},
				LowMetaDataSpace: eventType{name: "OSI_DOCKER_DEVICEMAPPER_LOW_METADATA_SPACE"},
			},
		},
	},
	OverloadedStorage: eventType{name: "OVERLOADED_STORAGE"},
	PerformanceEvent:  eventType{name: "PERFORMANCE_EVENT"},
	PGI: pgiEventType{
		HAProxy: pgiHAProxyEventType{
			QueuedRequestsHigh: eventType{name: "PGI_HAPROXY_QUEUED_REQUESTS_HIGH"},
			SessionUsageHigh:   eventType{name: "PGI_HAPROXY_SESSION_USAGE_HIGH"},
		},
		MySQL: pgiMySQLEventType{
			SlowQueriesRateHigh: eventType{name: "PGI_MYSQL_SLOW_QUERIES_RATE_HIGH"},
		},
		PGIOfServiceUnavailable: eventType{name: "PGI_OF_SERVICE_UNAVAILABLE"},
		RMQ: pgiRMQEventType{
			HighFileDescUsage: eventType{name: "PGI_RMQ_HIGH_FILE_DESC_USAGE"},
			HighMemUsage:      eventType{name: "PGI_RMQ_HIGH_MEM_USAGE"},
			HighSocketsUsage:  eventType{name: "PGI_RMQ_HIGH_SOCKETS_USAGE"},
			HighProcessUsage:  eventType{name: "PGI_RMQ_HIGH_PROCESS_USAGE"},
			LowDiskSpace:      eventType{name: "PGI_RMQ_LOW_DISK_SPACE"},
		},
	},
	Process: processEventType{
		Crashed: eventType{name: "PROCESS_CRASHED"},
		Custom: processCustomEventType{
			Availability: eventType{name: "PROCESS_CUSTOM_AVAILABILITY"},
			Error:        eventType{name: "PROCESS_CUSTOM_ERROR"},
			Performance:  eventType{name: "PROCESS_CUSTOM_PERFORMANCE"},
		},
		Log: processLogEventType{
			Availability: eventType{name: "PROCESS_LOG_AVAILABILITY"},
			Error:        eventType{name: "PROCESS_LOG_ERROR"},
			Performance:  eventType{name: "PROCESS_LOG_PERFORMANCE"},
		},
		Unavailable: eventType{name: "PROCESS_UNAVAILABLE"},
	},
	ProcessGroupLowInstanceCount: eventType{name: "PROCESS_GROUP_LOW_INSTANCE_COUNT"},
	ResourceContention:           eventType{name: "RESOURCE_CONTENTION"},
	RDSOfServiceUnavailable:      eventType{name: "RDS_OF_SERVICE_UNAVAILABLE"},
	RDSRestartSequence:           eventType{name: "RDS_RESTART_SEQUENCE"},
	ServiceResponseTimeDegraded:  eventType{name: "SERVICE_RESPONSE_TIME_DEGRADED"},
	SlowDisk:                     eventType{name: "SLOW_DISK"},
	Synthetic: syntheticEventType{
		Availability: eventType{name: "SYNTHETIC_AVAILABILITY"},
		Slowdown:     eventType{name: "SYNTHETIC_SLOWDOWN"},
	},
	ThreadsResourcesExhausted:     eventType{name: "THREADS_RESOURCES_EXHAUSTED"},
	UnexpectedHighLoad:            eventType{name: "UNEXPECTED_HIGH_LOAD"},
	UnexpectedLowLoad:             eventType{name: "UNEXPECTED_LOW_LOAD"},
	UserActionDurationDegredation: eventType{name: "USER_ACTION_DURATION_DEGRADATION"},
	VirtualMachineShutdown:        eventType{name: "VIRTUAL_MACHINE_SHUTDOWN"},
	WebCheck: webCheckEventType{
		GlobalOutage: eventType{name: "WEB_CHECK_GLOBAL_OUTAGE"},
		LocalOutage:  eventType{name: "WEB_CHECK_LOCAL_OUTAGE"},
	},
}

func (eventType *eventType) String() string {
	return eventType.name
}

func (eventType *eventType) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(eventType.name)), nil
}

func (eventType *eventType) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	eventType.name = quoted
	return err
}

func (eventType *eventType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: eventType.name}, nil
}

func (eventType *eventType) UnmarshalXMLAttr(attr xml.Attr) error {
	eventType.name = attr.Value
	return nil
}
