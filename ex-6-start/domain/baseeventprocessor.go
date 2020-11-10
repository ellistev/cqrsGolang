package domain

import (
	"github.com/avast/retry-go"
	"github.com/jdextraze/go-gesclient/client"
)

type EventProcessor interface {
	ProcessEvent(event *client.ResolvedEvent) error

	GetProcessorName() string

}

type BaseEventProcessor struct {
	Name string   // EventProcessor name
	EventTypes []string
	RetryConfig retry.Config // retry config, used by ProcessEventWithRetry
	ProcessorName string
	LastEventProcessed int
}

func (p *BaseEventProcessor) ProcessEvent(event *client.ResolvedEvent) error {
	panic("Base class, method not implemented")
}