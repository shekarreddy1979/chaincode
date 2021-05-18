package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//QueryMessages ..
type QueryMessages struct {
	Hl7Msh       *Hl7Msh       `json:"Hl7Msh"`
	Hl7ObrFormat *Hl7ObrFormat `json:"Hl7ObrFormat"`
	Hl7ObxFormat *Hl7ObxFormat `json:"Hl7ObxFormat"`
	Hl7OrcFormat *Hl7OrcFormat `json:"Hl7OrcFormat"`
	Hl7PidFormat *Hl7PidFormat `json:"Hl7PidFormat"`
	Hl7Pv1Format *Hl7Pv1Format `json:"Hl7Pv1Format"`
}

//QueryMSDetails ..
func QueryMSDetails(ctx contractapi.TransactionContextInterface, mshkey string) (*QueryMessages, error) {

	desID, err := GetAttirbuteValue(ctx)
	if err != nil || len(desID) > 0 {
		return nil, fmt.Errorf("%s,%s", err.Error(), "Cannot Find Attribute Value")
	}
	desBytes, err := ctx.GetStub().GetState(mshkey + "`" + desID)
	if desBytes == nil || err != nil {
		return nil, fmt.Errorf("Message key and Org Id Mismatch Combination")
	}
	mshBytes, err := ctx.GetStub().GetState(mshkey)
	qlist := QueryMessages{}
	err = json.Unmarshal(mshBytes, qlist.Hl7Msh)
	if err != nil {
		return nil, err
	}

	if obrBytes, _ := ctx.GetStub().GetState(mshkey + "_" + "OBR"); obrBytes != nil {
		json.Unmarshal(obrBytes, qlist.Hl7ObrFormat)
	}
	if orcBytes, _ := ctx.GetStub().GetState(mshkey + "_" + "ORC"); orcBytes != nil {
		json.Unmarshal(orcBytes, qlist.Hl7OrcFormat)
	}
	if pidBytes, _ := ctx.GetStub().GetState(mshkey + "_" + "PID"); pidBytes != nil {
		json.Unmarshal(pidBytes, qlist.Hl7PidFormat)
	}
	if pv1Bytes, _ := ctx.GetStub().GetState(mshkey + "_" + "PV1"); pv1Bytes != nil {
		json.Unmarshal(pv1Bytes, qlist.Hl7Pv1Format)
	}
	if obxBytes, _ := ctx.GetStub().GetState(mshkey + "_" + "PV1"); obxBytes != nil {
		json.Unmarshal(obxBytes, qlist.Hl7ObxFormat)
	}
	return &qlist, nil
}

//QuerySpecificmessage ..
func QuerySpecificmessage(ctx contractapi.TransactionContextInterface, MessageKey string, messageType string) ([]byte, error) {

	fmt.Println("query function got called")
	return ctx.GetStub().GetState(MessageKey + "_" + messageType)
}

//GetAttirbuteValue client Identity
func GetAttirbuteValue(ctx contractapi.TransactionContextInterface) (string, error) {

	desID, checkBool, err := ctx.GetClientIdentity().GetAttributeValue("organizationID")
	if checkBool && err == nil {
		return desID, nil
	}
	return desID, err
}
