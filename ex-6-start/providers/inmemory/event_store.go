package inmemory

import (
	"bus/cqrs"
	"github.com/satori/go.uuid"
)

type eventStore struct {
	aggregates     map[uuid.UUID][]*cqrs.DomainEvent
	eventPublisher cqrs.EventPublisher
}

func NewEventStore(eventPublisher cqrs.EventPublisher) *eventStore {
	return &eventStore{
		aggregates:     map[uuid.UUID][]*cqrs.DomainEvent{},
		eventPublisher: eventPublisher,
	}
}

func (s *eventStore) GetEvents(aggregateType string, aggregateId uuid.UUID) ([]*cqrs.DomainEvent, error) {
	events, found := s.aggregates[aggregateId]
	if !found {
		return nil, cqrs.AggregateNotFound
	}
	return events, nil
}

func (s *eventStore) SaveEvents(aggregateType string, aggregateId uuid.UUID, events []*cqrs.DomainEvent) error {
	s.aggregates[aggregateId] = events
	for _, e := range events {
		if err := s.eventPublisher.PublishEvent(e); err != nil {
			return err
		}
	}
	return nil
}