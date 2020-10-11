package repository

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"dev.azure.com/spsa/wspromo/domain"
	"dev.azure.com/spsa/wspromo/shared"
	"github.com/jinzhu/gorm"
)

type DatacenterRepository struct {
	db        []*gorm.DB
	connOne   *gorm.DB
	connTwo   *gorm.DB
	connThree *gorm.DB
	connFour  *gorm.DB
}

func NewDatacenterRepository(db []*gorm.DB) domain.IDatacenterRepository {
	log.Println("para ejecutar sp 0: ")
	return &DatacenterRepository{db: db, connOne: db[0], connTwo: db[0], connThree: db[0], connFour: db[0]}
}

func (s *DatacenterRepository) GetAll(p *shared.Pagination) ([]*domain.Datacenter, error) {
	log.Println("para ejecutar sp 1: ")
	atoiLimit, _ := strconv.Atoi(p.Limit)
	atoiOffset, _ := strconv.Atoi(p.OffSet)
	datacenters := make([]*domain.Datacenter, 0)
	if s.connTwo.Order("name").Limit(atoiLimit).Offset(atoiOffset).Where("is_delete = ?", false).Find(&datacenters).Error != nil {
		return nil, errors.New("Datacenter not found")
	}
	fmt.Printf("datacenters = %v \n", datacenters)
	return datacenters, nil
}

func (s *DatacenterRepository) GetByID(id string) (*domain.Datacenter, error) {
	// db := &util.ConnStruct{}
	log.Println("para ejecutar sp 2: ")
	datacenter := &domain.Datacenter{}
	log.Println("para ejecutar sp GetByID: ", datacenter)
	//if s.connTwo.Where("datacenter_id = ? and is_delete = ?", id, false).First(&datacenter).Error != nil {
	if s.connOne.Exec("select * from attributes") != nil {
		return nil, errors.New("Datacenter not found")
	}
	return datacenter, nil
}

func (s *DatacenterRepository) Store(srv *domain.Datacenter) error {
	// db := &util.ConnStruct{}
	resp := s.connTwo.Create(&srv)
	log.Println("para ejecutar sp 3: ")
	if resp.Error != nil {
		edv := shared.ConvertErrorDvToErrorDbStrcut(resp.Attrs().Error)
		messageError := strings.ReplaceAll(edv.Message, "\"", "")
		var auxInterface interface{}
		return errors.New(shared.ResponseService("Failure", edv.Code, messageError, auxInterface))
	} else {
		return nil
	}
}

func (s *DatacenterRepository) DeleteByID(id string) error {
	// db := &util.ConnStruct{}
	log.Println("para ejecutar sp 4: ")
	datacenter := &domain.Datacenter{}
	resp := s.connTwo.Model(datacenter).Where("datacenter_id = ?", id).Update("is_delete", true)
	if resp.Error != nil {
		edv := shared.ConvertErrorDvToErrorDbStrcut(resp.Attrs().Error)
		messageError := strings.ReplaceAll(edv.Message, "\"", "")
		var auxInterface interface{}
		return errors.New(shared.ResponseService("Failure", edv.Code, messageError, auxInterface))
	} else {
		return nil
	}
}

func (s *DatacenterRepository) UpdateByID(srv *domain.Datacenter) error {
	log.Println("para ejecutar sp 5: ")
	/*resp := s.connTwo.Model(&srv).Updates(datacenter.Datacenter{Name: srv.Name, Type: srv.Type, Location: srv.Location})
	if resp.Error != nil {
		edv := util.ConvertErrorDvToErrorDbStrcut(resp.Attrs().Error)
		messageError := strings.ReplaceAll(edv.Message, "\"", "")
		var auxInterface interface{}
		return errors.New(util.ResponseService("Failure", edv.Code, messageError, auxInterface))
	} else {
		return nil
	}*/
	return nil
}
