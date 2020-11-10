package inmemory

import (
	"bus/cqrs"
	"database/sql"
)

type Bus struct {
	commandFactory  cqrs.CommandFactory
	commandHandlers map[string]cqrs.CommandHandler
	eventFactory    cqrs.EventFactory
	eventHandlers   map[string][]cqrs.EventHandler
}

func NewBus(
	commandFactory cqrs.CommandFactory,
	eventFactory cqrs.EventFactory,
) *Bus {
	return &Bus{
		commandFactory:  commandFactory,
		commandHandlers: map[string]cqrs.CommandHandler{},
		eventFactory:    eventFactory,
		eventHandlers:   map[string][]cqrs.EventHandler{},
	}
}

func (b *Bus) RegisterCommandHandler(
	cmd cqrs.Command,
	handler cqrs.CommandHandler,
) error {
	name, err := b.commandFactory.GetCommandType(cmd)
	if err != nil {
		return err
	}
	if _, found := b.commandHandlers[name]; found {
		return cqrs.CommandHandlerAlreadyRegistered
	}
	b.commandHandlers[name] = handler
	return nil
}

func (b *Bus) RegisterEventHandler(
	evt cqrs.Event,
	handler cqrs.EventHandler,
) error {
	name, err := b.eventFactory.GetEventType(evt)
	if err != nil {
		return err
	}
	if _, found := b.eventHandlers[name]; !found {
		b.eventHandlers[name] = []cqrs.EventHandler{}
	}
	b.eventHandlers[name] = append(b.eventHandlers[name], handler)
	return nil
}

func (b *Bus) SendCommand(cmd cqrs.Command, db *sql.DB) error {
	name, err := b.commandFactory.GetCommandType(cmd)
	if err != nil {
		return err
	}
	handler := b.commandHandlers[name]
	return handler.HandleCommand(cmd, db)
}

func (b *Bus) PublishEvent(evt *cqrs.DomainEvent) error {
	name, err := b.eventFactory.GetEventType(evt.Event())
	if err != nil {
		return err
	}
	handlers := b.eventHandlers[name]
	for _, h := range handlers {
		if err := h.HandleEvent(evt); err != nil {
			return err
		}
	}
	return nil
}