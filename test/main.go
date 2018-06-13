package main

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/pkg/config"
	fab "github.com/hyperledger/fabric-sdk-go/api/apifabclient"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/hyperledger/fabric-sdk-go/api/apiconfig"
	"github.com/pkg/errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabric-client/chconfig"
)

func main(){
	logrus.SetLevel(logrus.DebugLevel)
	if err := run(); err != nil {
		logrus.Fatal(err)
	}
}

func run() error {
	sdk, err := fabsdk.New(config.FromFile("config.yaml"))
	if err != nil {
		return fmt.Errorf("failed to create sdk: %v", err)
	}

	clientCtx := sdk.NewClient(fabsdk.WithUser("Admin"), fabsdk.WithOrg("Org1"))

	identity, err := clientCtx.Session()
	if err != nil {
		return err
	}

	channel, err := GetChannel(sdk, identity, sdk.Config(), chconfig.NewChannelCfg("mychannel"), []string{"Org1"})
	if err != nil {
		return errors.Wrapf(err, "create channel (%s) failed: %v", "mychannel")
	}

	logrus.Debugf("channel.PrimaryPeer(): %+v", channel.PrimaryPeer())

	sdk.FabricProvider().CreateResourceClient(identity)

	return nil
}


// GetChannel initializes and returns a channel based on config
func GetChannel(sdk *fabsdk.FabricSDK, ic fab.IdentityContext, config apiconfig.Config, chCfg fab.ChannelCfg, orgs []string) (fab.Channel, error) {

	channel, err := sdk.FabricProvider().CreateChannelClient(ic, chCfg)
	if err != nil {
		return nil, errors.WithMessage(err, "NewChannel failed")
	}

	for _, org := range orgs {
		peerConfig, err := config.PeersConfig(org)
		if err != nil {
			return nil, errors.WithMessage(err, "reading peer config failed")
		}
		for _, p := range peerConfig {
			endorser, err := sdk.FabricProvider().CreatePeerFromConfig(&apiconfig.NetworkPeer{PeerConfig: p})
			if err != nil {
				return nil, errors.WithMessage(err, "NewPeer failed")
			}
			err = channel.AddPeer(endorser)
			if err != nil {
				return nil, errors.WithMessage(err, "adding peer failed")
			}
		}
	}

	return channel, nil
}