package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/pkg/errors"
)

//Organziation is used for model capturingorg details
type Organziation struct {
	DocType                  string `json:"docType"`
	OrganizationID           string `json:"organizationID" validate:"required"`
	OrganizationName         string `json:"organizationName" validate:"required"`
	OrganizationCode         string `json:"organizationCode" validate:"required" `
	OrganizationType         string `json:"organizationType,omitempty"`
	ParentOrganizationID     string `json:"parentOrganizationID,omitempty"`
	ContactPerson            string `json:"contactPerson,omitempty"`
	AddressLine1             string `json:"addressLine1,omitempty"`
	AddressLine2             string `json:"addressLine2,omitempty"`
	CityName                 string `json:"cityName,omitempty"`
	StateCode                string `json:"stateCode,omitempty"`
	CountyCode               string `json:"countyCode,omitempty"`
	CountryCode              string `json:"countryCode,omitempty"`
	ZipCode                  string `json:"zipCode,omitempty"`
	ContactPhone             string `json:"contactPhone,omitempty"`
	EmailID                  string `json:"emailID,omitempty"`
	Fax                      string `json:"fax,omitempty"`
	WebsiteURL               string `json:"websiteURL,omitempty"`
	TIN                      string `json:"TIN,omitempty"`
	CCN                      string `json:"CCN,omitempty"`
	SSN                      string `json:"SSN,omitempty"`
	NPI                      string `json:"NPI,omitempty"`
	EIN                      string `json:"EIN,omitempty"`
	IsInternal               string `json:"isInternal,omitempty"`
	IsActive                 string `json:"isActive,omitempty"`
	ProviderCommercialNumber string `json:"providerCommercialNumber,omitempty"`
	LocationNumber           string `json:"locationNumber,omitempty"`
	NCPDPNumber              string `json:"NCPDPNumber,omitempty"`
	OtherIdentifier          string `json:"otherIdentifier,omitempty"`
	ProviderUPINNumber       string `json:"providerUPINNumber,omitempty"`
	StateLicenseNumber       string `json:"stateLicenseNumber,omitempty"`
}

//AddOrg is used to register the organization into network
func (t *Organziation) fnAddOrg(ctx contractapi.TransactionContextInterface, actiontype string) error {
	err := FnValidationCheckStruct(t)
	if err != nil {
		return err
	}
	orgBytes, _ := ctx.GetStub().GetState(t.OrganizationID)
	fmt.Println(orgBytes, "orgBytes")
	switch strings.ToLower(actiontype) {
	case "add":
		if orgBytes != nil {
			return fmt.Errorf("%s,%s", Keyerror, t.OrganizationID)
		}
		break
	case "update":
		if err != nil {
			fmt.Println("Got issue for update funtionlaity")
			return fmt.Errorf("%s,%s", IDDosenotExists, t.OrganizationID)
		}
		break
	default:
		return fmt.Errorf("No Action Found for Organziation")
	}
	orgBytes, _ = json.Marshal(t)
	return ctx.GetStub().PutState(t.OrganizationID, orgBytes)
}

//GetOrg is used to get organization details
func GetOrg(ctx contractapi.TransactionContextInterface, orgID string) ([]byte, error) {
	if len(orgID) == 0 {
		return nil, errors.Errorf("Excepting orgID but found nil")
	}
	return ctx.GetStub().GetState(orgID)
}
