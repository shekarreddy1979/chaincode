package main

import (
	"github.com/pkg/errors"               //used to get errors
	"gopkg.in/go-playground/validator.v9" //Used to  validate the any struct
)

const (
	//Keyerror is common error
	Keyerror = "Key Already Exists in the Network"
	//IncorrectArgs ..
	IncorrectArgs = "In Correct Number of Arguments"
	//ChaincodeErr ...
	ChaincodeErr = "Error Creating Chaincode"
	//IDDosenotExists ..
	IDDosenotExists = "Key Dosenot exists in the network"
)

//FnValidationCheckStruct is used to verify the required feilds in strcut
func FnValidationCheckStruct(t interface{}) error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return errors.Errorf("Excepting Particular feild but seems to empty %s", e.Field())
		}
	}
	return nil
}
