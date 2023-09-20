package entities

import (
	"github.com/eser/go-service/pkg/shared"
	"github.com/oklog/ulid/v2"
)

// definition

type IService interface {
	GetRecord(id string, audit *shared.AuditRecord) (*Model, error)
	CreateRecord(dto CreateRecordDto, audit *shared.AuditRecord) error
	UpdateRecord(dto UpdateRecordDto, audit *shared.AuditRecord) error
	RemoveRecord(dto RemoveRecordDto, audit *shared.AuditRecord) error
}

type Service struct {
	repo IRepository
}

func NewService() IService {
	repo := NewRepository()

	return &Service{repo}
}

// get record

func (s *Service) GetRecord(id string, audit *shared.AuditRecord) (*Model, error) {
	return s.repo.Get(id, audit)
}

// create record

type CreateRecordDto struct {
	Fullname string
}

func (s *Service) CreateRecord(dto CreateRecordDto, audit *shared.AuditRecord) error {
	record := Model{
		Id:       ulid.MustNew(ulid.Now(), nil).String(),
		Fullname: dto.Fullname,
	}

	_, err := s.repo.Insert(record, audit)

	return err
}

// update record

type UpdateRecordDto struct {
	Id       string
	Fullname string
}

func (s *Service) UpdateRecord(dto UpdateRecordDto, audit *shared.AuditRecord) error {
	record := Model{
		Id:       dto.Id, // FIXME(@eser) ensure that existing id is passed
		Fullname: dto.Fullname,
	}

	_, err := s.repo.Update(dto.Id, record, audit)

	return err
}

// remove record

type RemoveRecordDto struct {
	Id string
}

func (s *Service) RemoveRecord(dto RemoveRecordDto, audit *shared.AuditRecord) error {
	_, err := s.repo.Remove(dto.Id, audit)

	return err
}
