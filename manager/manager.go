package manager

import (
	"github.com/go-xorm/xorm"
	"github.com/urfave/cli"
	"github.com/xuyuntech/register-visit/model"
)

type Manager interface {
}

type DefaultManager struct {
	engine *xorm.Engine
}

func NewManager(c *cli.Context) (Manager, error) {
	engine, err := model.NewEngine(c.String("database-datasource"), []interface{}{new(model.MedicalItem)})
	if err != nil {
		return nil, err
	}
	return &DefaultManager{
		engine: engine,
	}, nil
}
