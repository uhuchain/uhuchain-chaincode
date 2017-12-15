package main

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/uhuchain/uhuchain-core/models"
)

// HlfCarProvider implments car provider using hlf
type HlfCarProvider struct {
	stub shim.ChaincodeStubInterface
}

// NewHlfCarProvider creates an HlfCarProvider
func NewHlfCarProvider(stub shim.ChaincodeStubInterface) HlfCarProvider {
	return HlfCarProvider{
		stub: stub,
	}
}

// SaveCar implements the the provider using hlf
func (p *HlfCarProvider) SaveCar(car models.Car) error {
	carValue, err := car.MarshalBinary()
	if err != nil {
		return err
	}
	err = p.stub.PutState(strconv.FormatInt(car.ID, 10), carValue)
	if err != nil {
		return err
	}
	return nil
}

// GetCar implements the the provider using hlf
func (p *HlfCarProvider) GetCar(id int64) (models.Car, error) {
	car := models.Car{}
	carValue, err := p.stub.GetState(strconv.FormatInt(id, 10))
	if err != nil {
		return car, err
	}
	if len(carValue) == 0 {
		return car, fmt.Errorf("Code 404 - No entry found for key %d", id)
	}
	car.UnmarshalBinary(carValue)
	return car, nil
}
