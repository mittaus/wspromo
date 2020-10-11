package domain

import "dev.azure.com/spsa/wspromo/shared"

type IDatacenterRepository interface {
	GetAll(p *shared.Pagination) ([]*Datacenter, error)
	GetByID(id string) (*Datacenter, error)
	Store(u *Datacenter) error
	DeleteByID(id string) error
	UpdateByID(u *Datacenter) error
}
