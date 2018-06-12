package model

import (
	"time"

	"github.com/go-xorm/xorm"
)

type MedicalItem struct {
	Id   int64  `json:"id" xorm:"pk"`
	Name string `json:"name" xorm:"varchar(250)"`

	ObjectType       string `json:"docType"`
	Title            string `json:"title"`
	Quantity         string `json:"quantity"`
	Price            string `json:"price"`
	SupplierID       string `json:"supplierID"`       // 生产厂家
	BarCode          string `json:"barCode"`          // 条码
	BatchNumber      string `json:"batchNumber"`      //批号
	PermissionNumber string `json:"permissionNumber"` // 批准文号
	ProductionDate   string `json:"productionDate"`
	ExpiredDate      string `json:"expiredDate"`

	Created     time.Time `xorm:"-"`
	CreatedUnix int64
	Updated     time.Time `xorm:"-"`
	UpdatedUnix int64
}

func (me *MedicalItem) BeforeInsert() {
	me.CreatedUnix = time.Now().Unix()
	me.UpdatedUnix = me.CreatedUnix
}

func (me *MedicalItem) BeforeUpdate() {
	me.UpdatedUnix = time.Now().Unix()
}

func (me *MedicalItem) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "created_unix":
		me.Created = time.Unix(me.CreatedUnix, 0).Local()
	case "updated_unix":
		me.Updated = time.Unix(me.UpdatedUnix, 0)
	}
}
