package cqrs

type AggregateRepository struct {
	storage EventStore
}

func NewAggregateRepository(storage EventStore) *AggregateRepository {
	return &AggregateRepository{storage}
}

func (r *AggregateRepository) Load(a *AggregateRoot) error {
	events, err := r.storage.GetEvents(a.name, a.id)
	if err != nil {
		return err
	}
	a.LoadHistory(events)
	return nil
}

func (r *AggregateRepository) Save(a *AggregateRoot) error {
	if err := r.storage.SaveEvents(a.name, a.id, a.GetUncommittedChanges()); err != nil {
		return err
	}
	a.MarkChangesAsCommitted()
	return nil
}