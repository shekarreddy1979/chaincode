package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockStub struct {
	shim.ChaincodeStubInterface
	mock.Mock
}

type MockClientidentity struct {
	cid.ClientIdentity
	mock.Mock
}

type MockContract struct {
	contractapi.TransactionContextInterface
	mock.Mock
}

func (ms *MockStub) GetState(key string) ([]byte, error) {
	args := ms.Called(key)

	return args.Get(0).([]byte), args.Error(1)
}

func (ms *MockStub) PutState(key string, value []byte) error {
	args := ms.Called(key, value)
	return args.Error(0)
}

func (ms *MockStub) DelState(key string) error {
	args := ms.Called(key)
	return args.Error(0)
}

// func (ms *MockStub) SetEvent(key string, value []byte) error {
// 	args := ms.Called(key, value)
// 	return args.Error(0)
// }

func (ms *MockStub) GetMSPID() (string, error) {
	args := ms.Called()

	return args.Get(0).(string), args.Error(0)
}

func (ms *MockStub) GetTxID() string {
	args := ms.Called()
	return args.Get(0).(string)
}

// func (ms *MockStub) GetQueryResult(value string) shim.MockQueryIteratorInterface {
// 	args := ms.Called()
// 	return
// }

func (mc *MockContract) GetStub() shim.ChaincodeStubInterface {
	args := mc.Called()
	return args.Get(0).(*MockStub)
}

func (mid *MockClientidentity) GetMSPID() (string, error) {
	args := mid.Called()
	return args.Get(0).(string), args.Error(1)
}

func (mc *MockContract) GetClientIdentity() cid.ClientIdentity {
	args := mc.Called()
	return args.Get(0).(*MockClientidentity)

}

func TestRegisterMSH(t *testing.T) {

	dummymshkey := Response{TxID: "456789", Key: "123456"}
	msk := Hl7Msh{MSHKEY: "123456", DocType: "MessageKey"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456").Return(asd, nil)
	t.Log(mshBytes)
	stub.On("PutState", msk.MSHKEY, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData, err := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\"}", "msh")
	fmt.Println(jsonData, "cfgvhbnjkm")
	assert.Equal(t, err, nil, "value is null ")
	assert.Equal(t, jsonData, &dummymshkey, "value is same")

}

func TestRegisterOBR(t *testing.T) {

	dummymshkey := Response{TxID: "456789", Key: "123456_OBR"}
	msk := Hl7ObrFormat{MSHKEY: "123456", MSHOBR: "123456_OBR"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456_OBR").Return(asd, nil)
	stub.On("PutState", msk.MSHOBR, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData, err := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\",\"MSH_OBR\":\"123456_OBR\"}", "obr")
	assert.Equal(t, err, nil, "value is null")
	assert.Equal(t, jsonData, &dummymshkey, "value is same")
	t.Logf("TestRegisterOBR Got Success")

}

func TestRegisterORC(t *testing.T) {

	dummymshkey := Response{TxID: "456789", Key: "123456_ORC"}
	msk := Hl7OrcFormat{MSHKEY: "123456", MSHORC: "123456_ORC"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456_ORC").Return(asd, nil)
	stub.On("PutState", msk.MSHORC, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData, err := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\",\"MSH_ORC\":\"123456_ORC\"}", "orc")
	assert.Equal(t, err, nil, "value is null")
	assert.Equal(t, jsonData, &dummymshkey, "value is same")
	t.Logf("TestRegisterORC Got Success")

}

func TestRegisterOBX(t *testing.T) {

	dummymshkey := Response{TxID: "456789", Key: "123456_OBX"}
	msk := Hl7ObxFormat{MSHKEY: "123456", MSHOBX: "123456_OBX"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456_OBX").Return(asd, nil)
	stub.On("PutState", msk.MSHOBX, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData, err := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\",\"MSH_OBX\":\"123456_OBX\"}", "OBX")
	assert.Equal(t, err, nil, "value is null")
	assert.Equal(t, jsonData, &dummymshkey, "value is same")
	t.Logf("TestRegisterOBX Got Success")

}

func TestRegisterPID(t *testing.T) {

	dummymshkey := Response{TxID: "456789", Key: "123456_PID"}
	msk := Hl7PidFormat{MSHKEY: "123456", MSHPID: "123456_PID"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456_PID").Return(asd, nil)
	stub.On("PutState", msk.MSHPID, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData, err := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\",\"MSH_PID\":\"123456_PID\"}", "PID")
	assert.Equal(t, err, nil, "value is null")
	assert.Equal(t, jsonData, &dummymshkey, "value is same")
	t.Logf("TestRegisterPID Got Success")

}

func TestRegisterPV1(t *testing.T) {

	dummymshkey := Response{TxID: "456789", Key: "123456_PV1"}
	msk := Hl7Pv1Format{MSHKEY: "123456", MSHPV1: "123456_PV1"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456_PV1").Return(asd, nil)
	stub.On("PutState", msk.MSHPV1, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData, err := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\",\"MSH_PV1\":\"123456_PV1\"}", "PV1")
	assert.Equal(t, err, nil, "value is null")
	assert.NotEqual(t, jsonData, &dummymshkey, "value is same")
	t.Logf("TestRegisterPV1 Got Success")

}

func TestUnwantedCase(t *testing.T) {
	dummymshkey := Response{TxID: "", Key: ""}
	ctx := new(MockContract)
	stub := new(MockStub)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData1, err1 := contract.RegisterMSHMessage(ctx, "{\"MSHKEY\":\"123456\",\"MSH_PV1\":\"123456_PV1\"}", "qwe")
	assert.Equal(t, err1.Error(), "Cannot find Message type", "value is null")
	assert.NotEqual(t, jsonData1, &dummymshkey, "value is same")
	t.Logf("TestRegisterunwanted code Got Success")
}

func TestRegisterOrg(t *testing.T) {
	dummymshkey := Response{TxID: "456789", Key: "123456"}
	msk := Organziation{DocType: "", OrganizationID: "123456", OrganizationName: "asdf", OrganizationCode: "qwerty"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	var asd []byte
	stub.On("GetState", "123456").Return(asd, nil)
	stub.On("PutState", msk.OrganizationID, mshBytes).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData1, err1 := contract.RegisterOrg(ctx, "{\"OrganizationID\":\"123456\",\"OrganizationName\":\"asdf\",\"OrganizationCode\":\"qwerty\"}", "add")
	assert.Equal(t, err1, nil, "value is null")
	assert.Equal(t, jsonData1, &dummymshkey, "value is same")
	t.Logf("TestRegisterunwanted code Got Success")
}

func TestUpdateOrg(t *testing.T) {
	dummymshkey := Response{TxID: "456789", Key: "123456"}
	msk := Organziation{DocType: "", OrganizationID: "123456", OrganizationName: "asdf", OrganizationCode: "qwerty"}
	mshBytes, _ := json.Marshal(msk)
	stub := new(MockStub)
	msk1 := Organziation{DocType: "", OrganizationID: "123456", OrganizationName: "asdf", OrganizationCode: "qwerty", OtherIdentifier: "78945"}
	//var asd []byte
	mshBytes1, _ := json.Marshal(msk1)
	stub.On("GetState", "123456").Return(mshBytes, nil)
	stub.On("PutState", msk.OrganizationID, mshBytes1).Return(nil)
	stub.On("GetTxID").Return("456789")
	ctx := new(MockContract)
	ctx.On("GetStub").Return(stub)
	contract := new(HieCC)
	jsonData1, err1 := contract.RegisterOrg(ctx, "{\"OrganizationID\":\"123456\",\"OrganizationName\":\"asdf\",\"OrganizationCode\":\"qwerty\",\"OtherIdentifier\":\"78945\"}", "update")
	assert.Equal(t, err1, nil, "value is null")
	assert.Equal(t, jsonData1, &dummymshkey, "value is same")
	t.Logf("TestRegisterunwanted code Got Success")
}

// func TestUpdateOrg1(t *testing.T) {
// 	dummymshkey := Response{TxID: "456789", Key: "123456"}
// 	// msk := Organziation{DocType: "", OrganizationID: "123456", OrganizationName: "asdf", OrganizationCode: "qwerty"}
// 	// mshBytes, _ := json.Marshal(msk)
// 	stub := new(MockStub)
// 	msk1 := Organziation{DocType: "", OrganizationID: "123456", OrganizationName: "asdf", OrganizationCode: "qwerty", OtherIdentifier: "78945"}
// 	var asd []byte
// 	mshBytes1, _ := json.Marshal(msk1)
// 	stub.On("GetState", "123456").Return(asd, fmt.Errorf("cannot parse"))
// 	stub.On("PutState", msk1.OrganizationID, mshBytes1).Return(nil)
// 	stub.On("GetTxID").Return("456789")
// 	ctx := new(MockContract)
// 	ctx.On("GetStub").Return(stub)
// 	contract := new(HieCC)
// 	jsonData1, err1 := contract.RegisterOrg(ctx, "{\"OrganizationID\":\"123456\",\"OrganizationName\":\"asdf\",\"OrganizationCode\":\"qwerty\",\"OtherIdentifier\":\"78945\"}", "update")
// 	assert.Equal(t, err1.Error(), fmt.Sprintf("%s,%s", IDDosenotExists, msk1.OrganizationID), "value is not null")
// 	assert.NotEqual(t, jsonData1, &dummymshkey, "value is same")
// 	t.Logf("TestRegisterunwanted code Got Success")
// }

// func TestRegisterTransferOrgMSH(t *testing.T) {

// 	msk := TransferDetail{MSHKEY: "123456", SrcOrganizationID: "654987", DesOrganizationID: "789456"}
// 	dummymshkey := Response{TxID: "456789", Key: "123456`789456"}
// 	mshBytes, _ := json.Marshal(msk)
// 	v := struct {
// 		DestID      string `json:"destID"`
// 		MSHKEY      string `json:"MSHKEY"`
// 		MessageType string `json:"messageType"`
// 	}{
// 		DestID:      "789456",
// 		MSHKEY:      "123456",
// 		MessageType: "MSH",
// 	}
// 	eventBytes, _ := json.Marshal(v)
// 	stub := new(MockStub)
// 	var asd []byte
// 	stub.On("GetState", "123456`789456").Return(asd, nil)
// 	stub.On("PutState", msk.MSHID, mshBytes).Return(nil)
// 	stub.On("GetTxID").Return("456789")
// 	stub.On("SetEvent", "imessages", eventBytes).Return(nil)
// 	ctx := new(MockContract)
// 	ctx.On("GetStub").Return(stub)
// 	contract := new(HieCC)
// 	jsonData, err := contract.TransferOrgMSH(ctx, "{\"MSHKEY\":\"123456\",\"SrcOrganizationID\":\"654987\",\"DesOrganizationID\":\"789456\"}")
// 	assert.Equal(t, err, nil, "value is null")
// 	assert.Equal(t, jsonData, &dummymshkey, "value is same")
// 	t.Logf("TestRegisterORGtransfer Got Success")
// }
