package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
)

// TransferDetail : transfering of message key from one org to another  org for msh
type TransferDetail struct {
	DocType           string `json:"docType"`
	MSHID             string `json:"MSHID"`
	MSHKEY            string `json:"MSHKEY" validate:"required"`
	SrcOrganizationID string `json:"srcOrganizationID" validate:"required"`
	DesOrganizationID string `json:"desOrganizationID" validate:"required"`
	TransferDate      int64  `json:"transferDate"`
	IsActive          bool   `json:"isActive"`
}

func (t *TransferDetail) fnTransferFile(ctx contractapi.TransactionContextInterface) error {

	err := FnValidationCheckStruct(t)
	if err != nil {
		return err
	}
	t.DocType = "transfer"
	t.TransferDate = fnGetTimeStampforTXID(ctx)
	t.IsActive = true
	t.MSHID = t.MSHKEY + "`" + t.DesOrganizationID
	imesBytes, err := json.Marshal(t)
	if err != nil {
		return errors.Wrapf(err, "Got error while UnMarsahlling")
	}
	err = ctx.GetStub().PutState(t.MSHID, imesBytes)
	if err != nil {
		return errors.Wrapf(err, "Cannot insert Data")
	}
	v := struct {
		DestID      string `json:"destID"`
		MSHKEY      string `json:"MSHKEY"`
		MessageType string `json:"messageType"`
	}{
		DestID:      t.DesOrganizationID,
		MSHKEY:      t.MSHKEY,
		MessageType: "MSH",
	}
	eventBytes, _ := json.Marshal(v)
	ctx.GetStub().SetEvent("imessages", eventBytes)
	return nil
}
