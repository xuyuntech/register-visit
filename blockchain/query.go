package blockchain

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/api/apitxn/chclient"
	"errors"
)

// QueryUser query the chaincode to get the state of hello
func (setup *FabricSetup) Query(queryString string) (string, error) {

	// Prepare arguments
	var args []string
	args = append(args, "query")
	args = append(args, queryString)
	fmt.Println("queryString:", queryString)
	if setup.client == nil {
		return "", errors.New("setup client is nil")
	}

	response, err := setup.client.Query(chclient.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return "", fmt.Errorf("failed to query: %v", err)
	}

	return string(response.Payload), nil
}
