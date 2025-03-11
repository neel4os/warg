package storage

import (
	"github.com/neel4os/warg/libs/boilerplate"
	"github.com/neel4os/warg/libs/boilerplate/clients"
)

type StorageDomain struct {
	repo *StorageRepo
	deps map[string]boilerplate.Dependent
}

func NewStorageDomain(deps map[string]boilerplate.Dependent) *StorageDomain {
	_postgresconIf := deps["postgres"]
	_postgrescon := _postgresconIf.(*clients.PgClient)
	_repo := NewStorageRepo(_postgrescon.Dbcon)
	return &StorageDomain{
		repo: _repo,
	}
}

func (s *StorageDomain) Create(sin *Storage) (*Storage, error) {
	sin, err := s.repo.Create(sin)
	if err != nil {
		return nil, err
	}
	return sin, nil
}
