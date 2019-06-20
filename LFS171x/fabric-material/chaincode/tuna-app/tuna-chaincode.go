// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario
  This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
*/

package main

/* Imports
* 4 utility libraries for handling bytes, reading and writing JSON,
formatting, and string manipulation
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts
*/
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Exame structure, with 4 properties.
Structure tags are used by encoding/json library
*/
type Exame struct {
	Paciente  string `json:"paciente"`
	Cpf 	  string    `json:"cpf"`
	Medico    string `json:"medico"`
	Crm 	  string 	 `json:"crm"`
}
	//	Laudo     string `json:"laudo"`

/*
 * The Init method *
 called when the Smart Contract "Exame-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function
 -- see initLedger()
*/
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "Exame-chaincode"
 The app also specifies the specific smart contract function to call with args
*/
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryExame" {
		return s.queryExame(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordExame" {
		return s.recordExame(APIstub, args)
	} else if function == "queryAllExames" {
		return s.queryAllExames(APIstub)
	} else if function == "changeExameCrm" {
		return s.changeExameCrm(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryExame method *
Used to view the records of one particular Exame
It takes one argument -- the key for the Exame in question
*/
func (s *SmartContract) queryExame(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ExameAsBytes, _ := APIstub.GetState(args[0])
	if ExameAsBytes == nil {
		return shim.Error("Could not locate Exame")
	}
	return shim.Success(ExameAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 Exame catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	exames := []Exame{
		Exame{Paciente: "923F", Cpf: "67000676", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "M83T", Cpf: "91.39594", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "T012", Cpf: "58.04891", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "P490", Cpf: "-45.0949", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "S439", Cpf: "-107.603", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "J205", Cpf: "-155.223", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "S22L", Cpf: "103.8877", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "EI89", Cpf: "-1326983", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "129R", Cpf: "153.0529", Medico: "Carlos", Crm: "541"},
		Exame{Paciente: "49W4", Cpf: "51.95435", Medico: "Carlos", Crm: "541"},
	}

	i := 0
	for i < len(exames) {
		fmt.Println("i is ", i)
		examesAsBytes, _ := json.Marshal(exames[i])
		APIstub.PutState(strconv.Itoa(i+1), examesAsBytes)
		fmt.Println("Added", exames[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordExame method *
Fisherman like Sarah would use to record each of her Exame catches.
This method takes in five arguments (attributes to be saved in the ledger).
*/
func (s *SmartContract) recordExame(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	var Exame = Exame{Paciente: args[1], Cpf: args[2], Medico: args[3], Crm: args[4]}

	ExameAsBytes, _ := json.Marshal(Exame)
	err := APIstub.PutState(args[0], ExameAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record Exame catch: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The queryAllExame method *
allows for assessing all the records added to the ledger(all Exame catches)
This method does not take any arguments. Returns JSON string containing results.
*/
func (s *SmartContract) queryAllExames(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "0"
	endKey := "999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllExames:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changeExameCrm method *
The data in the world state can be updated with who has possession.
This function takes in 2 arguments, Exame id and new Crm name.
*/
func (s *SmartContract) changeExameCrm(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	ExameAsBytes, _ := APIstub.GetState(args[0])
	if ExameAsBytes == nil {
		return shim.Error("Could not locate Exame")
	}
	Exame := Exame{}

	json.Unmarshal(ExameAsBytes, &Exame)
	// Normally check that the specified argument is a valid Crm of Exame
	// we are skipping this check for this example
	Exame.Crm = args[1]

	ExameAsBytes, _ = json.Marshal(Exame)
	err := APIstub.PutState(args[0], ExameAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change Exame Crm: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * main function *
calls the Start function
The main function starts the chaincode in the container during instantiation.
*/
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
