package problems

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type eventSeverityContext struct {
	fmt.Stringer
	name string
}

type eventSeverities struct {
	CrashRate                  eventSeverityContext
	PageFaults                 eventSeverityContext
	CommandAbort               eventSeverityContext
	PgAvailable                eventSeverityContext
	FailureRate                eventSeverityContext
	ResponseTime50thPercentile eventSeverityContext
	ResponseTime90thPercentile eventSeverityContext
	CPU                        cpuEventSeverities
	Memory                     memoryEventSeverities
	Network                    networkEventSeverities
	Hypervisor                 hypervisorEventSeverities
}

type hypervisorEventSeverities struct {
	PacketsDropped hypervisorPacketsDroppedEventSeverities
}

type hypervisorPacketsDroppedEventSeverities struct {
	Received    eventSeverityContext
	Transmitted eventSeverityContext
}

type memoryEventSeverities struct {
	SwapInRate        eventSeverityContext
	SwapOutRate       eventSeverityContext
	CompressionRate   eventSeverityContext
	DecompressionRate eventSeverityContext
	Usage             eventSeverityContext
}

type cpuEventSeverities struct {
	Usage     eventSeverityContext
	ReadyTime eventSeverityContext
}

type networkEventSeverities struct {
	ErrorRate           networkErrorRateEventSeverities
	PacketsDropped      networkPacketsDroppedEventSeverities
	HighUtilizationRate networkHighUtilizationRateEventSeverities
}

type networkErrorRateEventSeverities struct {
	Received    eventSeverityContext
	Transmitted eventSeverityContext
}

type networkPacketsDroppedEventSeverities struct {
	Received    eventSeverityContext
	Transmitted eventSeverityContext
}

type networkHighUtilizationRateEventSeverities struct {
	Received    eventSeverityContext
	Transmitted eventSeverityContext
}

// EventSeverities TODO: documentation
var EventSeverities = eventSeverities{
	CPU: cpuEventSeverities{
		Usage:     eventSeverityContext{name: "CPU_USAGE"},
		ReadyTime: eventSeverityContext{name: "CPU_READY_TIME"},
	},
	Network: networkEventSeverities{
		ErrorRate: networkErrorRateEventSeverities{
			Received:    eventSeverityContext{name: "NETWORK_RECEIVED_ERROR_RATE"},
			Transmitted: eventSeverityContext{name: "NETWORK_TRANSMITTED_ERROR_RATE"},
		},
		PacketsDropped: networkPacketsDroppedEventSeverities{
			Received:    eventSeverityContext{name: "NETWORK_PACKETS_RECEIVED_DROPPED"},
			Transmitted: eventSeverityContext{name: "NETWORK_PACKETS_TRANSMITTED_DROPPED"},
		},
		HighUtilizationRate: networkHighUtilizationRateEventSeverities{
			Received:    eventSeverityContext{name: "NETWORK_HIGH_RECEIVED_UTILIZATION_RATE"},
			Transmitted: eventSeverityContext{name: "NETWORK_HIGH_TRANSMITTED_UTILIZATION_RATE"},
		},
	},
	Memory: memoryEventSeverities{
		Usage:             eventSeverityContext{name: "MEMORY_USAGE"},
		SwapInRate:        eventSeverityContext{name: "MEMORY_SWAP_IN_RATE"},
		SwapOutRate:       eventSeverityContext{name: "MEMORY_SWAP_OUT_RATE"},
		CompressionRate:   eventSeverityContext{name: "MEMORY_COMPRESSION_RATE"},
		DecompressionRate: eventSeverityContext{name: "MEMORY_DECOMPRESSION_RATE"},
	},
	CrashRate:    eventSeverityContext{name: "CRASH_RATE"},
	PageFaults:   eventSeverityContext{name: "PAGE_FAULTS"},
	CommandAbort: eventSeverityContext{name: "COMMAND_ABORT"},
	Hypervisor: hypervisorEventSeverities{
		PacketsDropped: hypervisorPacketsDroppedEventSeverities{
			Received:    eventSeverityContext{name: "HYPERVISOR_PACKETS_RECEIVED_DROPPED"},
			Transmitted: eventSeverityContext{name: "HYPERVISOR_PACKETS_TRANSMITTED_DROPPED"},
		},
	},
	PgAvailable:                eventSeverityContext{name: "PG_AVAILABLE"},
	FailureRate:                eventSeverityContext{name: "FAILURE_RATE"},
	ResponseTime50thPercentile: eventSeverityContext{name: "RESPONSE_TIME_50TH_PERCENTILE"},
	ResponseTime90thPercentile: eventSeverityContext{name: "RESPONSE_TIME_90TH_PERCENTILE"},
}

func (context *eventSeverityContext) String() string {
	return context.name
}

func (context *eventSeverityContext) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(context.name)), nil
}

func (context *eventSeverityContext) UnmarshalJSON(data []byte) error {
	quoted, err := strconv.Unquote(string(data))
	context.name = quoted
	return err
}

func (context *eventSeverityContext) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{Name: name, Value: context.name}, nil
}

func (context *eventSeverityContext) UnmarshalXMLAttr(attr xml.Attr) error {
	context.name = attr.Value
	return nil
}
