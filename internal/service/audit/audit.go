package audit

import (
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
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
	errors := make([]error, 0)
	for _, o := range p.observers {
		if err := o.Update(event); err != nil {
			errors = append(errors, err)
		}
	}
	if len(errors) == 0 {
		return nil
	}
	return errs.NewErrAuditNotify(errors...)
}

func (p *Publisher) Update(event *model.AuditEvent) error {
	return p.Notify(event)
}
