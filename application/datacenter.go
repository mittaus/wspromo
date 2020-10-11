package application

import (
	"errors"

	"dev.azure.com/spsa/wspromo/domain"
	"dev.azure.com/spsa/wspromo/shared"
	uuid "github.com/satori/go.uuid"
)

type Datacenter struct {
	repository domain.IDatacenterRepository
}

func NewDatacenter(repository domain.IDatacenterRepository) *Datacenter {
	return &Datacenter{
		repository: repository,
	}
}

func (dc Datacenter) GetAll(p *shared.Pagination) (result []*domain.Datacenter, err error) {
	result, err = dc.repository.GetAll(p)
	return
}

func (dc Datacenter) GetByID(id string) (result *domain.Datacenter, err error) {

	if _, err := uuid.FromString(id); err != nil {
		err = errors.New("Error trying convert uuid to string")
		return nil, err
	}

	datacenter, err := dc.repository.GetByID(id)
	return datacenter, err
}

func (dc Datacenter) Store(datacenter *domain.Datacenter) error {

	err := dc.repository.Store(datacenter)
	return err
}

func (dc Datacenter) DeleteByID(id string) error {
	if _, err := uuid.FromString(id); err != nil {
		err = errors.New("Error trying convert uuid to string")
		return err
	}

	err := dc.repository.DeleteByID(id)

	return err
}

func (dc Datacenter) UpdateByID(id string, datacenter *domain.Datacenter) error {
	if _, err := uuid.FromString(id); err != nil {
		err = errors.New("Error trying convert uuid to string")
		return err
	}

	err := dc.repository.UpdateByID(datacenter)
	return err
}
