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

/* Define structures
 */

type Patient struct {
	Id        string `json:"id"`
	CPF       string `json:"cpf"`
	Name      string `json:"name"`
	Sex       string `json:"sex"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	Height    string `json:"height"`
	Weight    string `json:"weight"`
	Age       string `json:"age"`
	BloodType string `json:"bloodType"`
}

type Doctor struct {
	Id    string `json:"id"`
	CRM   string `json:"crm"`
	CPF   string `json:"cpf"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

type Enterprise struct {
	Id    string `json:"id"`
	CNPJ  string `json:"cnpj"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}
type Exam struct {
	PatientId string `json:"patientId"`
	DoctorId  string `json:"doctorId"`
	ExamId    string `json:"examId"`
}

/*
 * The Init method *
 called when the Smart Contract "Exam-chaincode" is instantiated by the network
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
	// if function == "queryExam" {
	// 	return s.queryExam(APIstub, args)
	// } else
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordPatient" {
		return s.recordPatient(APIstub, args)
	} else if function == "queryPatient" {
		return s.queryPatient(APIstub, args)
	}
	//  else if function == "recordExam" {
	// 	return s.recordExam(APIstub, args)
	// } else if function == "queryAllExams" {
	// 	return s.queryAllExams(APIstub)
	// }
	// else if function == "recordDoctor" {
	// 	return s.recordDoctor(APIstub, args)
	// } else if function == "queryDoctor" {
	// 	return s.queryDoctor(APIstub, args)
	// } else if function == "recordEnterprise" {
	// 	return s.recordEnterprise(APIstub, args)
	// } else if function == "queryEnterprise" {
	// 	return s.queryEnterprise(APIstub, args)
	// }

	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryExame method *
Used to view the records of one particular Exam
It takes one argument -- the key for the Exam in question
*/
func (s *SmartContract) queryExame(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	ExamAsBytes, _ := APIstub.GetState(args[0])
	if ExamAsBytes == nil {
		return shim.Error("Could not locate Exame")
	}
	return shim.Success(ExamAsBytes)
}

/*
 * The queryPatient method *
Used to view the records of one particular Patient
It takes one argument -- the key for the Patient in question
*/
func (s *SmartContract) queryPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	if patientAsBytes == nil {
		return shim.Error("Could not locate Patient")
	}
	return shim.Success(patientAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 Exame catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	// exams := []Exam{
	// 	Exam{Paciente: "923F", Cpf: "67000676", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "M83T", Cpf: "91.39594", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "T012", Cpf: "58.04891", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "P490", Cpf: "-45.0949", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "S439", Cpf: "-107.603", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "J205", Cpf: "-155.223", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "S22L", Cpf: "103.8877", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "EI89", Cpf: "-1326983", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "129R", Cpf: "153.0529", Medico: "Carlos", Crm: "541"},
	// 	Exam{Paciente: "49W4", Cpf: "51.95435", Medico: "Carlos", Crm: "541"},
	// }

	patients := []Patient{
		Patient{Id: "1", CPF: "1", Name: "Pedro", Sex: "M", Phone: "123", Email: "a@a.a", Height: "175", Weight: "61", Age: "22", BloodType: "A+"},
		Patient{Id: "2", CPF: "2", Name: "José", Sex: "M", Phone: "123", Email: "a@a.a", Height: "175", Weight: "61", Age: "22", BloodType: "A+"},
	}

	i := 0
	for i < len(patients) {
		fmt.Println("i is ", i)
		patientAsBytes, _ := json.Marshal(patients[i])
		APIstub.PutState(strconv.Itoa(i+1), patientAsBytes)
		fmt.Println("Added", patients[i])
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

	var Exam = Exam{ExamId: args[0], PatientId: args[1], DoctorId: args[2]}

	ExamAsBytes, _ := json.Marshal(Exam)
	err := APIstub.PutState(args[0], ExamAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record Exam: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The recordPatient method *
 */
func (s *SmartContract) recordPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}

	var Patient = Patient{Id: args[0], CPF: args[1], Name: args[2], Sex: args[3], Phone: args[4], Email: args[5], Height: args[6], Weight: args[7], Age: args[8], BloodType: args[9]}

	patientAsBytes, _ := json.Marshal(Patient)
	err := APIstub.PutState(args[0], patientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record patient: %s", args[0]))
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
