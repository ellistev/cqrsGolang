package cqrs

import "github.com/satori/go.uuid"

type Event interface{}

type DomainEvent struct {
	aggregateId uuid.UUID
	id          uuid.UUID
	version     int
	event       Event
}

type EventPublisher interface {
	PublishEvent(event *DomainEvent) error
}

type EventHandler interface {
	HandleEvent(*DomainEvent) error
}

func NewDomainEvent(
	aggregateId uuid.UUID,
	id uuid.UUID,
	version int,
	event Event,
) *DomainEvent {
	return &DomainEvent{
		aggregateId: aggregateId,
		id:          id,
		version:     version,
		event:       event,
	}
}

func (e *DomainEvent) AggregateId() uuid.UUID { return e.aggregateId }

func (e *DomainEvent) Id() uuid.UUID { return e.id }

func (e *DomainEvent) Version() int { return e.version }

func (e *DomainEvent) Event() Event { return e.event }

type EventHandlerFunc func(*DomainEvent) error

func (f EventHandlerFunc) HandleEvent(evt *DomainEvent) error {
	return f(evt)
}

type EventFactory interface {
	CreateEvent(name string) (Event, error)
	GetEventType(Event) (string, error)
}