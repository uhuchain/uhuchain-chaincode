/*
Copyright Uhuchain. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/uhuchain/uhuchain-core/models"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/uhuchain/uhuchain-core/usecases"
)

// CarChaincode example simple Chaincode implementation
type CarChaincode struct {
}

// Init car chaincode
func (t *CarChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	log.Println("Init car chaincode")
	_, args := stub.GetFunctionAndParameters()
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	Aval, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B = args[2]
	Bval, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Invoke a function from the chaincode
func (t *CarChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("car chaincode Invoke")
	function, args := stub.GetFunctionAndParameters()
	carProvider := NewHlfCarProvider(stub)
	carUsecase := usecases.NewCarUsecase(&carProvider)
	fmt.Printf("requested chaincode function %s", function)
	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	} else if function == "saveCar" {
		return t.saveCar(stub, carUsecase, args)
	} else if function == "getCar" {
		return t.saveCar(stub, carUsecase, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Deletes an entity from state
func (t *CarChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// Write a car with a given ID onto the ledger
func (t *CarChaincode) saveCar(stub shim.ChaincodeStubInterface, usecase *usecases.CarUsecase, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ")
	}
	carValue := []byte(args[1])
	car := models.Car{}
	err := car.UnmarshalBinary(carValue)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = usecase.SaveCar(car)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *CarChaincode) getCar(stub shim.ChaincodeStubInterface, usecase *usecases.CarUsecase, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting ")
	}
	carID, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		return shim.Error(err.Error())
	}
	car, err := usecase.GetCar(carID)
	if err != nil {
		return shim.Error(err.Error())
	}
	carValue, err := car.MarshalBinary()
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(carValue)
}

// query callback representing the query of a chaincode
func (t *CarChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]
	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(CarChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
