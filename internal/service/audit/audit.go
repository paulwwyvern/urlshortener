package audit

import (
	"errors"

	"github.com/paulwwyvern/urlshortener/internal/model"
)

type observer interface {
	Update(event *model.AuditEvent) error
}

type Publisher struct {
	observers []observer
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

func (p *Publisher) Register(observer observer) {
	p.observers = append(p.observers, observer)
}

func (p *Publisher) Notify(event *model.AuditEvent) error {
	errs := make([]error, 0)
	for _, o := range p.observers {
		if err := o.Update(event); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.Join(errs...)
}

func (p *Publisher) Update(event *model.AuditEvent) error {
	return p.Notify(event)
}
