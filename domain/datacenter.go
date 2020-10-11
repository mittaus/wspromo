package domain

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type Datacenter struct {
	PromovarCode        string `json:"promovarCode"`
	PromovarDescription string `json:"promovarDescription"`
	Type                string `json:"type"`
}

func (Datacenter) TableName() string {
	return "datacenters"
}

func (s *Datacenter) BeforeCreate(scope *gorm.Scope) {
	uuid := uuid.NewV4()
	scope.SetColumn("ID", uuid.String())
	scope.SetColumn("IsDelete", false)
}
