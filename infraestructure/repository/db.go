package repository

import (
	"fmt"

	"dev.azure.com/spsa/wspromo/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func NewDb(c *config.GORMConfig) ([]*gorm.DB, error) {

	arrayConnections := make([]*gorm.DB, 0)

	var db *gorm.DB

	var err error

	// the follows connections where injected for Hygen

	if len(c.DB.Mysql) > 0 {
		for _, v := range c.DB.Mysql {
			connStr := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				v.Username,
				v.Password,
				v.Host,
				v.Port,
				v.Database)

			db, err = gorm.Open("mysql", connStr)

			if err != nil {
				fmt.Printf(" err = %v", err)
			}

			arrayConnections = append(arrayConnections, db)

			db = &gorm.DB{}

		}
	}

	return arrayConnections, nil
}
