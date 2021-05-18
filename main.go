package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

//HieCC is contract
type HieCC struct {
	contractapi.Contract
}

//GetTimeStampforTXID ...
func fnGetTimeStampforTXID(ctx contractapi.TransactionContextInterface) int64 {
	sec, _ := ctx.GetStub().GetTxTimestamp()
	return sec.GetSeconds()
}

func main() {
	hieCc, err := contractapi.NewChaincode(new(HieCC))
	if err != nil {
		fmt.Println(ChaincodeErr, err.Error())
		return
	}
	if err = hieCc.Start(); err != nil {
		fmt.Println(ChaincodeErr, err.Error())
	}
}
