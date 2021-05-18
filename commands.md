


#Creating PAckage for new lifecycle chaincode 
peer lifecycle chaincode package mycc.tar.gz --path github.com/chaincode/ --lang golang --label mycc_5




# InStallation of chaincode 
peer lifecycle chaincode install mycc1.tar.gz

##EXport Evironmentvariable 
<h1>After succefully installation of chaincode</h1>
export    CC_PACKAGE_ID=mycc_6:56c3ea50d40cd00693e8e3b0454e40fca1ef0e78c8c7e5dfe7b1e16213e918c4

##APPROVE  ORG DEFENIATIon
peer lifecycle chaincode approveformyorg -o orderer.example.com:7050 --channelID mychannel --name mycc --version 1.0 --package-id $CC_PACKAGE_ID --sequence 7
#CC_PACKAGE_ID we need to provide after installation 
#--init-required is not need for the high level api chaincode 
# sequence indicates version of chaincode 
#version is also indicates how many times approve request been risied


## commit chaincode in to network after successfully installed and collected approved from orgs
peer lifecycle chaincode commit -o orderer.example.com:7050 --channelID mychannel --name mycc --version 1.0 --sequence 7

#INVOKE CHAINCODE  FOR REGISTERMSH function

#RegisterMSH
peer chaincode invoke -o orderer.example.com:7050  -C mychannel  -n mycc  -c '{"Args":["RegisterMSH","{\"MSHKEY\":\"12356\"}"]}'

#exceptedOutPUT
{\"txID\":\"5d05ddfea311899e907d020751735f2320a677213242d9f8199451022432ea4b\",\"Key\":\"12346\"}

#RegisterORG
<h5>sorucrORG<h5>
peer chaincode invoke -o orderer.example.com:7050  -C mychannel  -n mycc  -c '{"Args":["RegisterOrg","{\"OrganizationID\":\"789456\",\"OrganizationName\":\"apollo\",\"OrganizationCode\":\"asdfg\"}","add"]}'

peer chaincode invoke -o orderer.example.com:7050  -C mychannel  -n mycc  -c '{"Args":["RegisterOrg","{\"OrganizationID\":\"654987\",\"OrganizationName\":\"care\",\"OrganizationCode\":\"asdfg\"}","add"]}'

#TRansferORg
peer chaincode invoke -o orderer.example.com:7050  -C mychannel  -n mycc  -c '{"Args":["TransferOrgMSH","{\"MSHKEY\":\"12356\",\"SrcOrganizationID\":\"789456\",\"DestOrganizationID\":\"654987\"}"]}'


peer chaincode invoke -o orderer.example.com:7050  -C mychannel  -n mycc  -c '{"Args":["RegisterMSHMessage","{\"MSHKEY\":\"123456\",\"MSHOBX\":\"123456_OBR\"}","OBR"]}'