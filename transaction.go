package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//RegisterOrg ..
func (hc *HieCC) RegisterOrg(ctx contractapi.TransactionContextInterface, args string, actionType string) (*Response, error) {
	fmt.Println("value is ", args)
	t := Organziation{}
	err := json.Unmarshal([]byte(args), &t)
	if err != nil {
		return nil, err
	}
	err = t.fnAddOrg(ctx, actionType)
	if err != nil {
		return nil, err
	}
	return fnResponse(t.OrganizationID, ctx.GetStub().GetTxID()), nil
}

//TransferOrgMSH ..
func (hc *HieCC) TransferOrgMSH(ctx contractapi.TransactionContextInterface, args string) (*Response, error) {
	fmt.Println("value is ", args)
	t := TransferDetail{}
	err := json.Unmarshal([]byte(args), &t)
	if err != nil {
		return nil, err
	}
	err = t.fnTransferFile(ctx)
	if err != nil {
		return nil, err
	}
	return fnResponse(t.MSHID, ctx.GetStub().GetTxID()), nil
}

func fnMessageRegisterion(ctx contractapi.TransactionContextInterface, hl7format, messageType string) (string, error) {
	var err error
	var key string
	switch strings.ToLower(messageType) {
	case "msh":
		t := Hl7Msh{}
		err := json.Unmarshal([]byte(hl7format), &t)
		if err != nil {
			return "", err
		}
		err = t.fnAddMSH(ctx)
		key = t.MSHKEY
		break
	case "obr":
		mo := Hl7ObrFormat{}
		err = json.Unmarshal([]byte(hl7format), &mo)
		if err != nil {
			return "", err
		}
		key = mo.MSHOBR
		err = mo.fnAddObr(ctx)
		fmt.Println("process completted")
		break
	case "obx":
		mo := Hl7ObxFormat{}
		err = json.Unmarshal([]byte(hl7format), &mo)
		if err != nil {
			return "", err
		}
		key = mo.MSHOBX
		err = mo.fnAddObx(ctx)
		break
	case "pid":
		mo := Hl7PidFormat{}
		err = json.Unmarshal([]byte(hl7format), &mo)
		if err != nil {
			return "", err
		}
		key = mo.MSHPID
		err = mo.fnAddPid(ctx)
		break
	case "pv1":
		mo := Hl7Pv1Format{}
		err = json.Unmarshal([]byte(hl7format), &mo)
		if err != nil {
			return "", err
		}
		key = mo.PV1KEY
		err = mo.fnAddPv1(ctx)
	case "orc":
		mo := Hl7OrcFormat{}
		err = json.Unmarshal([]byte(hl7format), &mo)
		if err != nil {
			return "", err
		}
		key = mo.MSHORC
		fmt.Println(key, "key is added to network")
		err = mo.fnAddOrc(ctx)
		break
	default:
		key = ""
		err = fmt.Errorf("Cannot find Message type")
		break
	}
	fmt.Println("waiting for process")
	return key, err
}

//RegisterMSHMessage ..
func (hc *HieCC) RegisterMSHMessage(ctx contractapi.TransactionContextInterface, hl7format string, MessageType string) (*Response, error) {

	fmt.Println(hl7format, MessageType)
	id, err := fnMessageRegisterion(ctx, hl7format, MessageType)
	fmt.Println(err, "error is ")
	if err != nil {
		fmt.Println("erro occured in this ")
		return nil, err
	}
	fnsetEvent(ctx, id, MessageType)
	response := fnResponse(id, ctx.GetStub().GetTxID())
	return response, err
}

func fnsetEvent(ctx contractapi.TransactionContextInterface, messageKey string, messageType string) {
	fmt.Println("process started")
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"transfer\",\"MSHKEY\":\"%s\",\"isActive\":true},\"fields\":[\"desOrganizationID\"]}", messageKey)
	queryiterater, err := ctx.GetStub().GetQueryResult(queryString)
	if err != nil || queryiterater == nil {
		return
	}
	for queryiterater.HasNext() {
		queryResponse, err := queryiterater.Next()
		if err != nil {
			return
		}
		s := strings.Split(queryResponse.Key, "`")
		v := struct {
			DestID      string `json:"destID"`
			MSHKEY      string `json:"MSHKEY"`
			MessageType string `json:"messageType"`
		}{
			DestID:      s[1],
			MSHKEY:      messageKey,
			MessageType: messageType,
		}
		eventBytes, _ := json.Marshal(v)
		ctx.GetStub().SetEvent("imessages", eventBytes)
	}

}
