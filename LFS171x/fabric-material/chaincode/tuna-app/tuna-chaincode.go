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

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

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
	} else if function == "recordDoctor" {
		return s.recordDoctor(APIstub, args)
	} else if function == "queryDoctor" {
		return s.queryDoctor(APIstub, args)
	} else if function == "addDoctorToPatient" {
		return s.addDoctorToPatient(APIstub, args)
	} else if function == "removeDoctorFromPatient" {
		return s.removeDoctorFromPatient(APIstub, args)
	} else if function == "recordEnterprise" {
		return s.recordEnterprise(APIstub, args)
	} else if function == "addEnterpriseToPatient" {
		return s.addEnterpriseToPatient(APIstub, args)
	} else if function == "addExam" {
		return s.addExam(APIstub, args)
	}

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
 * The queryDoctor method *
Used to view the records of one particular Doctor
It takes one argument -- the key for the Patient in question
*/
func (s *SmartContract) queryDoctor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

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
Will add test data (10 Exame catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {

	patients := []Patient{
		Patient{Id: "1", CPF: "84532652098", Name: "Pedro Sousa Meireles", Sex: "M", Phone: "21933000494", Email: "psmeireles25@gmail.com", Height: "175", Weight: "61", Age: "22", BloodType: "A+", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
		Patient{Id: "2", CPF: "29847293956", Name: "José da Silva", Sex: "M", Phone: "2112345678", Email: "josesilva@gmail.com", Height: "185", Weight: "91", Age: "65", BloodType: "A+", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
		Patient{Id: "3", CPF: "23485292659", Name: "Rafael Cabral", Sex: "M", Phone: "2294874938", Email: "rafarubim@gmail.com", Height: "175", Weight: "71", Age: "21", BloodType: "A+", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
		Patient{Id: "4", CPF: "98436939758", Name: "Luiza Lima", Sex: "F", Phone: "11938427583", Email: "lulima@gmail.com", Height: "160", Weight: "51", Age: "36", BloodType: "AB+", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
		Patient{Id: "5", CPF: "64934875378", Name: "Antônia Meira", Sex: "F", Phone: "21929384938", Email: "tonya@gmail.com", Height: "165", Weight: "70", Age: "52", BloodType: "O-", Doctors: []string{}, Exams: []string{}, Enterprises: []string{}},
	}

	doctors := []Doctor{
		Doctor{Id: "82029156787", CRM: "512974", CPF: "82029156787", Name: "Carla Eliane Carvalho de Sousa", Phone: "21999839210", Email: "carla.sousa@uol.com.br", Patients: []string{}, Exams: []string{}},
		Doctor{Id: "82029156788", CRM: "640816", CPF: "82029156788", Name: "Claudio Luiz Bastos Bragança", Phone: "21999839210", Email: "claudio.braganca@uol.com.br", Patients: []string{}, Exams: []string{}},
	}

	i := 0
	for i < len(patients) {
		fmt.Println("i is ", i)
		patientAsBytes, _ := json.Marshal(patients[i])
		APIstub.PutState(patients[i].Id, patientAsBytes)
		fmt.Println("Added", patients[i])
		i = i + 1
	}

	i = 0
	for i < len(doctors) {
		fmt.Println("i is ", i)
		doctorAsBytes, _ := json.Marshal(doctors[i])
		APIstub.PutState(doctors[i].Id, doctorAsBytes)
		fmt.Println("Added", doctors[i])
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

	var Patient = Patient{Id: args[0], CPF: args[1], Name: args[2], Sex: args[3], Phone: args[4], Email: args[5], Height: args[6], Weight: args[7], Age: args[8], BloodType: args[9], Doctors: []string{}, Exams: []string{}}

	patientAsBytes, _ := json.Marshal(Patient)
	err := APIstub.PutState(args[0], patientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record patient: %s", args[0]))
	}

	return shim.Success(nil)
}

/*
 * The recordDoctor method *
 */
func (s *SmartContract) recordDoctor(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var Doctor = Doctor{Id: args[0], CRM: args[1], CPF: args[2], Name: args[3], Phone: args[4], Email: args[5], Patients: []string{}, Exams: []string{}}

	doctorAsBytes, _ := json.Marshal(Doctor)
	err := APIstub.PutState(args[0], doctorAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record doctor: %s", args[0]))
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
 * The addDoctorToPatient method *
 */
func (s *SmartContract) addDoctorToPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	if patientAsBytes == nil {
		return shim.Error("Could not locate patient")
	}
	patient := Patient{}

	json.Unmarshal(patientAsBytes, &patient)

	doctorAsBytes, _ := APIstub.GetState(args[1])
	if doctorAsBytes == nil {
		return shim.Error("Could not locate doctor")
	}
	doctor := Doctor{}

	json.Unmarshal(doctorAsBytes, &doctor)

	// Normally check that the specified argument is a valid holder of tuna
	// we are skipping this check for this example
	patient.Doctors = append(patient.Doctors, args[1])
	doctor.Patients = append(doctor.Patients, args[0])
	doctor.Exams = append(doctor.Exams, patient.Exams...)

	patientAsBytes, _ = json.Marshal(patient)
	err := APIstub.PutState(args[0], patientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add doctor to patient: %s", args[0]))
	}

	doctorAsBytes, _ = json.Marshal(doctor)
	err = APIstub.PutState(args[1], doctorAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add patient to doctor: %s", args[1]))
	}

	return shim.Success(nil)
}

/*
 * The removeDoctorFromPatient method *
 */
func (s *SmartContract) removeDoctorFromPatient(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	if patientAsBytes == nil {
		return shim.Error("Could not locate patient")
	}
	patient := Patient{}

	json.Unmarshal(patientAsBytes, &patient)

	doctorAsBytes, _ := APIstub.GetState(args[1])
	if doctorAsBytes == nil {
		return shim.Error("Could not locate doctor")
	}
	doctor := Doctor{}

	json.Unmarshal(doctorAsBytes, &doctor)
	// Normally check that the specified argument is a valid holder of tuna
	// we are skipping this check for this example

	doctors := []string{}
	for i := 0; i < len(patient.Doctors); i++ {
		if patient.Doctors[i] != args[1] {
			doctors = append(doctors, patient.Doctors[i])
		}
	}
	patient.Doctors = doctors

	patients := []string{}
	for i := 0; i < len(doctor.Patients); i++ {
		if doctor.Patients[i] != args[0] {
			patients = append(patients, doctor.Patients[i])
		}
	}
	doctor.Patients = patients

	exams := []string{}
	for i := 0; i < len(doctor.Exams); i++ {
		add := true
		for j := 0; j < len(patient.Exams); j++ {
			if doctor.Exams[i] == patient.Exams[j] {
				add = false
			}
		}

		if add == true {
			exams = append(exams, doctor.Exams[i])
		}
	}

	doctor.Exams = exams

	patientAsBytes, _ = json.Marshal(patient)
	err := APIstub.PutState(args[0], patientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to remove doctor from patient: %s", args[0]))
	}

	doctorAsBytes, _ = json.Marshal(doctor)
	err = APIstub.PutState(args[1], doctorAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to remove patient from doctor: %s", args[1]))
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
 * The addExam method *
 */
func (s *SmartContract) addExam(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	patientAsBytes, _ := APIstub.GetState(args[0])
	if patientAsBytes == nil {
		return shim.Error("Could not locate patient")
	}
	patient := Patient{}

	json.Unmarshal(patientAsBytes, &patient)

	doctorAsBytes, _ := APIstub.GetState(args[1])
	if doctorAsBytes == nil {
		return shim.Error("Could not locate doctor")
	}
	doctor := Doctor{}

	json.Unmarshal(doctorAsBytes, &doctor)

	// Normally check that the specified argument is a valid holder of tuna
	// we are skipping this check for this example
	patient.Exams = append(patient.Exams, args[2])
	doctor.Exams = append(doctor.Exams, args[2])

	patientAsBytes, _ = json.Marshal(patient)
	err := APIstub.PutState(args[0], patientAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add exam to patient: %s", args[0]))
	}

	doctorAsBytes, _ = json.Marshal(doctor)
	err = APIstub.PutState(args[1], doctorAsBytes)
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to add exam to doctor: %s", args[1]))
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
