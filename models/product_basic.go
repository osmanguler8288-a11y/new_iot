package models

import "gorm.io/gorm"

type ProductBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(50);" json:"identity"`
	Name     string `gorm:"column:name; type:varchar(50);" json:"name"`
	Key      string `gorm:"column:key; type:varchar(50);" json:"key"`
	Secret   string `gorm:"column:secret; type:varchar(50);" json:"secret"`
}

func (table ProductBasic) TableName() string {
	return "product_basic"
}

func ProductList(name string) *gorm.DB {
	tx := DB.Debug().Model(new(ProductBasic)).Select("identity, name, `desc`, `key`, created_at")
	if name != "" {
		tx.Where("name LIKE ?", "%"+name+"%")
	}
	return tx
}
