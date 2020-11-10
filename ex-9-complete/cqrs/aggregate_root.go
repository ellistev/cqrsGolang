package cqrs

import (
	"github.com/satori/go.uuid"
)

type ApplyHandlerFunc func(Event)

type AggregateRoot struct {
	name    string
	id      uuid.UUID
	version int
	applyer ApplyHandlerFunc
	changes []*DomainEvent
}

func NewAggregateRoot(name string, id uuid.UUID, apply ApplyHandlerFunc) *AggregateRoot {
	return &AggregateRoot{
		name:    name,
		id:      id,
		applyer: apply,
		version: -1,
		changes: []*DomainEvent{},
	}
}

func (a *AggregateRoot) Name() string { return a.name }

func (a *AggregateRoot) Id() uuid.UUID { return a.id }

func (a *AggregateRoot) Version() int { return a.version }

func (a *AggregateRoot) LoadHistory(events []*DomainEvent) {
	for _, e := range events {
		a.applyer(e.event)
		a.version = e.version
	}
}

func (a *AggregateRoot) Apply(event Event) error {
	a.applyer(event)
	a.version += 1
	guid, _ :=uuid.NewV4()
	a.changes = append(a.changes, &DomainEvent{
		aggregateId: a.id,
		id:          guid,
		version:     a.version,
		event:       event,
	})
	return nil
}

func (a *AggregateRoot) GetUncommittedChanges() []*DomainEvent {
	return a.changes
}

func (a *AggregateRoot) MarkChangesAsCommitted() {
	a.changes = []*DomainEvent{}
}