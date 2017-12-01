package main

import "github.com/hyperledger/fabric/core/chaincode/shim"

type HlfCarProvider struct {
	stub shim.ChaincodeStubInterface
}

func NewHlfCarProvider(stub shim.ChaincodeStubInterface) HlfCarProvider {
	return HlfCarProvider{
		stub: stub,
	}
}

func (p *HlfCarProvider) SaveCar(car string) error {
	return nil
}

func (p *HlfCarProvider) GetCar(id string) (string, error) {
	return "", nil
}
