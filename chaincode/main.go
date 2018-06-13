package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	functionName, args := stub.GetFunctionAndParameters()
	fmt.Sprintf("CC Invoke: %s", functionName)
	switch functionName {
	case "query":
		return s.query(stub, args)
	}
	return shim.Error(fmt.Sprintf("Invalid Smart Contract function name(%s).", functionName))
}

func (s *SmartContract) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	return shim.Success([]byte("query func."))
}

func main() {}
