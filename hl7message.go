package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//Response ..
type Response struct {
	TxID string `json:"txID"`
	Key  string `json:"Key"`
}

//Hl7Msh Main Message type
type Hl7Msh struct {
	DocType                                               string `json:"docType"`
	MSHKEY                                                string `json:"MSHKEY" validate:"required"`
	SendingApplicationNamespaceID                         string `json:"SendingApplicationNamespaceID,omitempty"`
	SendingApplicationUniversalID                         string `json:"SendingApplicationUniversalID,omitempty"`
	SendingApplicationUniversalIDType                     string `json:"SendingApplicationUniversalIDType,omitempty"`
	SendingFacilityNamespaceID                            string `json:"SendingFacilityNamespaceID,omitempty"`
	SendingFacilityUniversalID                            string `json:"SendingFacilityUniversalID,omitempty"`
	SendingFacilityUniversalIDType                        string `json:"SendingFacilityUniversalIDType,omitempty"`
	ReceivingApplicationNamespaceID                       string `json:"ReceivingApplicationNamespaceID,omitempty"`
	ReceivingApplicationUniversalID                       string `json:"ReceivingApplicationUniversalID ,omitempty"`
	ReceivingApplicationUniversalIDType                   string `json:"ReceivingApplicationUniversalIDType,omitempty"`
	ReceivingFacilityNamespaceID                          string `json:"ReceivingFacilityNamespaceID,omitempty"`
	ReceivingFacilityUniversalID                          string `json:"ReceivingFacilityUniversalID,omitempty"`
	ReceivingFacilityUniversalIDType                      string `json:"ReceivingFacilityUniversalIDType,omitempty"`
	MessageTime                                           string `json:"MessageTime,omitempty"`
	MessageYear                                           string `json:"MessageYear,omitempty"`
	MessageMonth                                          string `json:"MessageMonth,omitempty"`
	MessageDay                                            string `json:"MessageDay,omitempty"`
	MessageHours                                          string `json:"MessageHours,omitempty"`
	MessageMinutes                                        string `json:"MessageMinutes,omitempty"`
	MessageSeconds                                        string `json:"MessageSeconds,omitempty"`
	MessageMills                                          string `json:"MessageMills,omitempty"`
	MessageGmtOffset                                      string `json:"MessageGmtOffset,omitempty"`
	Precision                                             string `json:"Precision,omitempty"`
	Security                                              string `json:"Security,omitempty"`
	MessageType                                           string `json:"MessageType,omitempty"`
	TriggerEvent                                          string `json:"TriggerEvent,omitempty"`
	MessageStructure                                      string `json:"MessageStructure,omitempty"`
	MessageControlID                                      string `json:"MessageControlID,omitempty"`
	ProcessingID                                          string `json:"ProcessingID,omitempty"`
	ProcessingMode                                        string `json:"ProcessingMode,omitempty"`
	VersionID                                             string `json:"VersionID,omitempty"`
	VersionIdentifier                                     string `json:"VersionIdentifier,omitempty"`
	VersionText                                           string `json:"VersionText,omitempty"`
	VersionNameOfCodingSystem                             string `json:"VersionNameOfCodingSystem,omitempty"`
	VersionAlternateIdentifier                            string `json:"VersionAlternateIdentifier,omitempty"`
	VersionAlternateText                                  string `json:"VersionAlternateText,omitempty"`
	VersionNameOfAlternateCodingSystem                    string `json:"VersionNameOfAlternateCodingSystem,omitempty"`
	InternationalVersionIIdentifier                       string `json:"InternationalVersionIIdentifier,omitempty"`
	InternationalVersionIText                             string `json:"InternationalVersionIText,omitempty"`
	InternationalVersionINameOfCodingSystem               string `json:"InternationalVersionINameOfCodingSystem,omitempty"`
	InternationalVersionIAlternateIdentifier              string `json:"InternationalVersionIAlternateIdentifier,omitempty"`
	InternationalVersionIAlternateText                    string `json:"InternationalVersionIAlternateText,omitempty"`
	InternationalVersionINameOfAlternateCodingSystem      string `json:"InternationalVersionINameOfAlternateCodingSystem,omitempty"`
	SequenceNumber                                        string `json:"SequenceNumber,omitempty"`
	ContinuationPointer                                   string `json:"ContinuationPointer,omitempty"`
	AcceptAcknowledgementType                             string `json:"AcceptAcknowledgementType,omitempty"`
	ApplicationAcknowledgementType                        string `json:"ApplicationAcknowledgementType,omitempty"`
	CountryCode                                           string `json:"CountryCode,omitempty"`
	PrincipalLanguageOfMessageIdentifier                  string `json:"PrincipalLanguageOfMessageIdentifier,omitempty"`
	PrincipalLanguageOfMessageText                        string `json:"PrincipalLanguageOfMessageText,omitempty"`
	PrincipalLanguageOfMessageNameOfCodingSystem          string `json:"PrincipalLanguageOfMessageNameOfCodingSystem,omitempty"`
	PrincipalLanguageOfMessageAlternateIdentifier         string `json:"PrincipalLanguageOfMessageAlternateIdentifier,omitempty"`
	PrincipalLanguageOfMessageAlternateText               string `json:"PrincipalLanguageOfMessageAlternateText,omitempty"`
	PrincipalLanguageOfMessageNameOfAlternateCodingSystem string `json:"PrincipalLanguageOfMessageNameOfAlternateCodingSystem,omitempty"`
	AlternateCharacterSetHandlingScheme                   string `json:"AlternateCharacterSetHandlingScheme,omitempty"`
}

//UnmarshalJSON .. custome unmarsahl
// func (s *Hl7Msh) UnmarshalJSON(data []byte) error {
// 	err := json.Unmarshal(data, &s)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (s *Hl7Msh) fnAddMSH(ctx contractapi.TransactionContextInterface) error {

	err := FnValidationCheckStruct(s)
	if err != nil {
		return err
	}
	fmt.Println("validation is complete ")
	mesBytes, err := ctx.GetStub().GetState(s.MSHKEY)
	if err != nil {
		return err
	}
	fmt.Println(mesBytes, "bytesvalue is ")
	if mesBytes != nil {
		return fmt.Errorf("%s,%s", Keyerror, s.MSHKEY)
	}
	s.DocType = "MessageKey"
	mesBytes, _ = json.Marshal(s)
	fmt.Println("marsahlling value is done")
	return ctx.GetStub().PutState(s.MSHKEY, mesBytes)
}

func fnResponse(key string, TxID string) *Response {
	r := Response{}
	r.TxID = TxID
	r.Key = key
	return &r
}

// func fnsetMessagekeyStateBasedEndorsement(ctx contractapi.TransactionContextInterface, MSHKEY string, orgToEndorse string) error {

// 	endorsementPolicy, err := statebased.NewStateEP(nil)
// 	err = endorsementPolicy.AddOrgs(statebased.RoleTypePeer, orgToEndorse)
// 	if err != nil {
// 		return fmt.Errorf("failed to add org to endorsement policy: %s", err.Error())
// 	}
// 	epBytes, err := endorsementPolicy.Policy()
// 	if err != nil {
// 		return fmt.Errorf("failed to create endorsement policy bytes from org: %s", err.Error())
// 	}
// 	err = ctx.GetStub().SetStateValidationParameter(MSHKEY, epBytes)
// 	if err != nil {
// 		return fmt.Errorf("failed to set validation parameter on marble: %s", err.Error())
// 	}
// 	return nil
// }

// func fnUpdateSatteBasedEndorsement(ctx contractapi.TransactionContextInterface, MSHKEY, AddOrgtoList string) error {

// 	epBytes, err := ctx.GetStub().GetStateValidationParameter(MSHKEY)
// 	if epBytes == nil && err != nil {
// 		return fmt.Errorf("Cannot Update stateLevel Endorsement")
// 	}
// 	endorsement, err := statebased.NewStateEP(epBytes)
// 	if err != nil {
// 		return fmt.Errorf("Cannot Update with new Endorsment Policy")
// 	}
// 	err = endorsement.AddOrgs(statebased.RoleTypePeer, AddOrgtoList)
// 	if err != nil {
// 		return fmt.Errorf("Cannot Update stateLevel Endorsement")
// 	}
// 	epBytes, err = endorsement.Policy()
// 	if err != nil {
// 		return fmt.Errorf("failed to create endorsement policy bytes from org: %s", err.Error())
// 	}
// 	err = ctx.GetStub().SetStateValidationParameter(MSHKEY, epBytes)
// 	return err

// }

// func fngetClientOrgID(ctx contractapi.TransactionContextInterface, VerifyItAlso bool) (string, error) {

// 	ClientOrgID, err := ctx.GetClientIdentity().GetMSPID()
// 	if err != nil {
// 		return "", fmt.Errorf("failed getting client's orgID: %s", err.Error())
// 	}
// 	if VerifyItAlso {
// 		peerOrgID, err := shim.GetMSPID()
// 		if err != nil {
// 			return "", fmt.Errorf("failed getting peer's orgID: %s", err.Error())
// 		}
// 		if ClientOrgID != peerOrgID {
// 			return "", fmt.Errorf("MSHKEY Can be read write have only Specific Peer %s,%s", ClientOrgID, peerOrgID)
// 		}
// 	}
// 	return ClientOrgID, nil

// }
