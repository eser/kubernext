package entities

import (
	"github.com/eser/go-service/pkg/shared"
)

// definition

type EntitiesService struct {
	repo IRepository
}

func NewEntitiesService() *EntitiesService {
	repo := NewRepository()

	return &EntitiesService{repo}
}

// get record

func (s *EntitiesService) GetRecord(id string, audit *shared.AuditRecord) (*Model, error) {
	return s.repo.Get(id, audit)
}

// create record

type CreateRecordDto struct {
	Fullname string
}

func (s *EntitiesService) CreateRecord(dto CreateRecordDto, audit *shared.AuditRecord) error {
	record := Model{
		Id:       shared.GenerateUniqueId(),
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

func (s *EntitiesService) UpdateRecord(dto UpdateRecordDto, audit *shared.AuditRecord) error {
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

func (s *EntitiesService) RemoveRecord(dto RemoveRecordDto, audit *shared.AuditRecord) error {
	_, err := s.repo.Remove(dto.Id, audit)

	return err
}
