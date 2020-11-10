package cqrs

import (
	"errors"
	"github.com/satori/go.uuid"
)

type EventStore interface {
	GetEvents(aggregateType string, aggregateId uuid.UUID) ([]*DomainEvent, error)
	SaveEvents(aggregateType string, aggregateId uuid.UUID, events []*DomainEvent) error
}

var AggregateNotFound = errors.New("Aggregate not found")