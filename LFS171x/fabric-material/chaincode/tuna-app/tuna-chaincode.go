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

<<<<<<< Updated upstream
/* Define Tuna structure, with 4 properties.
Structure tags are used by encoding/json library
*/
type Tuna struct {
	Vessel    string `json:"vessel"`
	Timestamp string `json:"timestamp"`
	Location  string `json:"location"`
	Holder    string `json:"holder"`
=======
/* Define structures
 */

type Patient struct {
	Id          string   `json:"id"`
	CPF         string   `json:"cpf"`
	Name        string   `json:"name"`
	Sex         string   `json:"sex"`
	Phone       string   `json:"phone"`
	Email       string   `json:"email"`
	Height      string   `json:"height"`
	Weight      string   `json:"weight"`
	Age         string   `json:"age"`
	BloodType   string   `json:"bloodType"`
	Doctors     []string `json:"doctors"`
	Exams       []string `json:"exams"`
	Enterprises []string `json:"enterprises"`
}

type Doctor struct {
	Id       string   `json:"id"`
	CRM      string   `json:"crm"`
	CPF      string   `json:"cpf"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Patients []string `json:"patients"`
	Exams    []string `json:"exams"`
}

type Enterprise struct {
	Id       string   `json:"id"`
	CNPJ     string   `json:"cnpj"`
	Name     string   `json:"name"`
	Phone    string   `json:"phone"`
	Email    string   `json:"email"`
	Patients []string `json:"patients"`
	Doctors  []string `json:"doctors"`
	Exams    []string `json:"exams"`
}
type Exam struct {
	PatientId string `json:"patientId"`
	DoctorId  string `json:"doctorId"`
	ExamId    string `json:"examId"`
>>>>>>> Stashed changes
}

/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function
 -- see initLedger()
*/
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
*/
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "queryTuna" {
		return s.queryTuna(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
<<<<<<< Updated upstream
	} else if function == "recordTuna" {
		return s.recordTuna(APIstub, args)
	} else if function == "queryAllTuna" {
		return s.queryAllTuna(APIstub)
	} else if function == "changeTunaHolder" {
		return s.changeTunaHolder(APIstub, args)
	}
=======
	} else if function == "recordPatient" {
		return s.recordPatient(APIstub, args)
	} else if function == "queryPatient" {
		return s.queryPatient(APIstub, args)
	} else if function == "recordDoctor" {
		return s.recordDoctor(APIstub, args)
	} else if function == "queryDoctor" {
		return s.queryDoctor(APIstub, args)
	} else if function == "addDoctorToPatient" {
		return s.addDoctorToPatient(APIstub, args)
	} else if function == "recordEnterprise" {
		return s.recordEnterprise(APIstub, args)
	}

	//  else if function == "recordExam" {
	// 	return s.recordExam(APIstub, args)
	// } else if function == "queryAllExams" {
	// 	return s.queryAllExams(APIstub)
	// }
	//  else if function == "recordEnterprise" {
	// 	return s.recordEnterprise(APIstub, args)
	// } else if function == "queryEnterprise" {
	// 	return s.queryEnterprise(APIstub, args)
	// }

>>>>>>> Stashed changes
	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The queryTuna method *
Used to view the records of one particular tuna
It takes one argument -- the key for the tuna in question
*/
func (s *SmartContract) queryTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	tunaAsBytes, _ := APIstub.GetState(args[0])
	if tunaAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	return shim.Success(tunaAsBytes)
}

/*
 * The queryEnterprise method *
Used to view the records of one particular Enterprise
It takes one argument -- the key for the Enterprise in question
*/
func (s *SmartContract) queryEnterprise(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	doctorAsBytes, _ := APIstub.GetState(args[0])
	if doctorAsBytes == nil {
		return shim.Error("Could not locate Doctor")
	}
	return shim.Success(doctorAsBytes)
}

/*
 * The initLedger method *
Will add test data (10 tuna catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
<<<<<<< Updated upstream
	tuna := []Tuna{
		Tuna{Vessel: "923F", Location: "67.0006, -70.5476", Timestamp: "1504054225", Holder: "Miriam"},
		Tuna{Vessel: "M83T", Location: "91.2395, -49.4594", Timestamp: "1504057825", Holder: "Dave"},
		Tuna{Vessel: "T012", Location: "58.0148, 59.01391", Timestamp: "1493517025", Holder: "Igor"},
		Tuna{Vessel: "P490", Location: "-45.0945, 0.7949", Timestamp: "1496105425", Holder: "Amalea"},
		Tuna{Vessel: "S439", Location: "-107.6043, 19.5003", Timestamp: "1493512301", Holder: "Rafa"},
		Tuna{Vessel: "J205", Location: "-155.2304, -15.8723", Timestamp: "1494117101", Holder: "Shen"},
		Tuna{Vessel: "S22L", Location: "103.8842, 22.1277", Timestamp: "1496104301", Holder: "Leila"},
		Tuna{Vessel: "EI89", Location: "-132.3207, -34.0983", Timestamp: "1485066691", Holder: "Yuan"},
		Tuna{Vessel: "129R", Location: "153.0054, 12.6429", Timestamp: "1485153091", Holder: "Carlo"},
		Tuna{Vessel: "49W4", Location: "51.9435, 8.2735", Timestamp: "1487745091", Holder: "Fatima"},
=======

	patients := []Patient{
		Patient{Id: "1", CPF: "1", Name: "Pedro", Sex: "M", Phone: "123", Email: "a@a.a", Height: "175", Weight: "61", Age: "22", BloodType: "A+", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
		Patient{Id: "2", CPF: "2", Name: "José", Sex: "M", Phone: "123", Email: "a@a.a", Height: "175", Weight: "61", Age: "22", BloodType: "A+", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
	}

	doctors := []Doctor{
		Doctor{Id: "82029156787", CRM: "512974", CPF: "82029156787", Name: "Carla", Phone: "21999839210", Email: "carla.sousa@uol.com.br", Patients: []string{}, Exams: []string{}},
		Doctor{Id: "82029156788", CRM: "512975", CPF: "82029156788", Name: "Claudio", Phone: "21999839210", Email: "carla.sousa@uol.com.br", Patients: []string{}, Exams: []string{}},
>>>>>>> Stashed changes
	}

	i := 0
	for i < len(tuna) {
		fmt.Println("i is ", i)
		tunaAsBytes, _ := json.Marshal(tuna[i])
		APIstub.PutState(strconv.Itoa(i+1), tunaAsBytes)
		fmt.Println("Added", tuna[i])
		i = i + 1
	}

	return shim.Success(nil)
}

/*
 * The recordTuna method *
Fisherman like Sarah would use to record each of her tuna catches.
This method takes in five arguments (attributes to be saved in the ledger).
*/
func (s *SmartContract) recordTuna(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

<<<<<<< Updated upstream
	var tuna = Tuna{Vessel: args[1], Location: args[2], Timestamp: args[3], Holder: args[4]}
=======
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

	var Patient = Patient{Id: args[0], CPF: args[1], Name: args[2], Sex: args[3], Phone: args[4], Email: args[5], Height: args[6], Weight: args[7], Age: args[8], BloodType: args[9], Doctors: []string{}, Exams: []string{}, Enterprises: []string{}}
>>>>>>> Stashed changes

	tunaAsBytes, _ := json.Marshal(tuna)
	err := APIstub.PutState(args[0], tunaAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record tuna catch: %s", args[0]))
	}
	return shim.Success(nil)
}

/*
 * The queryAllTuna method *
allows for assessing all the records added to the ledger(all tuna catches)
This method does not take any arguments. Returns JSON string containing results.
*/
func (s *SmartContract) queryAllTuna(APIstub shim.ChaincodeStubInterface) sc.Response {

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

	fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

/*
 * The changeTunaHolder method *
The data in the world state can be updated with who has possession.
This function takes in 2 arguments, tuna id and new holder name.
*/
func (s *SmartContract) changeTunaHolder(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	tunaAsBytes, _ := APIstub.GetState(args[0])
	if tunaAsBytes == nil {
		return shim.Error("Could not locate tuna")
	}
	tuna := Tuna{}

	json.Unmarshal(tunaAsBytes, &tuna)
	// Normally check that the specified argument is a valid holder of tuna
	// we are skipping this check for this example
	tuna.Holder = args[1]

	tunaAsBytes, _ = json.Marshal(tuna)
	err := APIstub.PutState(args[0], tunaAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to change tuna holder: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The addDoctorToPatient method *
 */
func (s *SmartContract) addEnterpriseToPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	if patientAsBytes == nil {
		return shim.Error("Could not locate patient")
	}
	patient := Patient{}

	json.Unmarshal(patientAsBytes, &patient)
	// Normally check that the specified argument is a valid holder of tuna
	// we are skipping this check for this example
	patient.Enterprises = append(patient.Enterprises, args[1])
	patientAsBytes, _ = json.Marshal(patient)
	err := APIstub.PutState(args[0], patientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add enterprise to patient: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The recordEnterprise method *
 */
func (s *SmartContract) recordEnterprise(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments in recordEnterprise. Expecting 5")
	}

	var Enterprise = Enterprise{Id: args[0], CNPJ: args[1], Name: args[2], Phone: args[3], Email: args[4], Patients: []string{}, Exams: []string{}, Doctors: []string{}}

	enterpriseAsBytes, _ := json.Marshal(Enterprise)
	err := APIstub.PutState(args[0], enterpriseAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record enterprise: %s", args[0]))
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
