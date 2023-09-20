package entities

import (
	"time"

	"github.com/eser/go-service/pkg/shared"
)

// definition

type IRepository interface {
	Get(id string, audit *shared.AuditRecord) (*Model, error)
	Insert(record Model, audit *shared.AuditRecord) (*Model, error)
	Update(id string, record Model, audit *shared.AuditRecord) (*Model, error)
	Remove(id string, audit *shared.AuditRecord) (*Model, error)
}

type Repository struct {
}

func NewRepository() IRepository {
	return &Repository{}
}

// get record

func (r *Repository) Get(id string, audit *shared.AuditRecord) (*Model, error) {
	// TODO(@eser) implement
	record := Model{
		Id: id,

		Fullname: "Jane Doe",

		CreatedAt: audit.RequestedAt,
		CreatedBy: audit.RequestedBy,
		UpdatedAt: time.Time{},
		UpdatedBy: "",
		RemovedAt: time.Time{},
		RemovedBy: "",
	}

	return &record, nil
}

// insert record

func (r *Repository) Insert(record Model, audit *shared.AuditRecord) (*Model, error) {
	// TODO(@eser) implement
	alteredModel := record

	alteredModel.CreatedAt = audit.RequestedAt
	alteredModel.CreatedBy = audit.RequestedBy

	return &alteredModel, nil
}

// update record

func (r *Repository) Update(id string, record Model, audit *shared.AuditRecord) (*Model, error) {
	// TODO(@eser) implement
	alteredModel, err := r.Get(id, audit)

	if err != nil {
		return nil, err
	}

	alteredModel.UpdatedAt = audit.RequestedAt
	alteredModel.UpdatedBy = audit.RequestedBy

	return alteredModel, nil
}

// remove record

func (r *Repository) Remove(id string, audit *shared.AuditRecord) (*Model, error) {
	// TODO(@eser) implement
	alteredModel, err := r.Get(id, audit)

	if err != nil {
		return nil, err
	}
	alteredModel.RemovedAt = audit.RequestedAt
	alteredModel.RemovedBy = audit.RequestedBy

	return alteredModel, nil
}
