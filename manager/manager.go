package manager

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	"github.com/urfave/cli"
	"github.com/xuyuntech/register-visit/blockchain"
	"github.com/xuyuntech/register-visit/model"
)

type Manager interface {
	ChainQuery(queryString string) (string, error)
	ChainSetupChannel() error
	InstallAndInstantiateCC() error
}

type DefaultManager struct {
	engine      *xorm.Engine
	fabricSetup *blockchain.FabricSetup
}

func NewManager(c *cli.Context) (Manager, error) {
	fSetup := &blockchain.FabricSetup{
		// Channel parameters
		ChannelID:     "chainhero",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/xuyuntech/register-visit/fabric/artifacts/chainhero.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "heroes-service",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "github.com/xuyuntech/register-visit/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Org1",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}
	if err := fSetup.Initialize(); err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("Unable to initialize the Fabric SDK: %v", err))
	}
	engine, err := model.NewEngine(c.String("database-datasource"), []interface{}{new(model.MedicalItem)})
	if err != nil {
		return nil, err
	}
	return &DefaultManager{
		engine:      engine,
		fabricSetup: fSetup,
	}, nil
}
