// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.queryAllTuna = function(){

		appFactory.queryAllTuna(function(data){
			var array = [];
			for (var i = 0; i < data.length; i++){
				parseInt(data[i].Key);
				data[i].Record.Key = parseInt(data[i].Key);
				array.push(data[i].Record);
			}
			array.sort(function(a, b) {
			    return parseFloat(a.Key) - parseFloat(b.Key);
			});
			$scope.all_tuna = array;
		});
	}

	$scope.queryTuna = function(){

		var id = $scope.tuna_id;

		appFactory.queryTuna(id, function(data){
			$scope.query_tuna = data;

			if ($scope.query_tuna == "Could not locate tuna"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.user_id = "";

	$scope.changeHolder = function(){

		appFactory.changeHolder($scope.holder, function(data){
			$scope.change_holder = data;
			if ($scope.change_holder == "Error: no tuna catch found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});
	}

	$scope.goToMyProfile = function(){

		appFactory.queryPatient(id, function(data){
			$scope.goToMyProfile = data;
			if ($scope.goToMyProfile == "Could not locate patient")
			{
				appFactory.queryDoctor(id, function(data){
					$scope.goToMyProfile = data;
					if ($scope.goToMyProfile == "Could not locate doctor"){
						type = "Enterprise"
					} else{
						type = "Doc"
					}
				});
			}
			else{ 
				type = "User"
			}
		});
		window.location.href = "./my" + type + "_profile.html?id=" + $scope.user_id;
	}

	$scope.addDoctorToPatient = function(){

		debugger
		var doctor = {}
		doctor.doctorId = $scope.doctorToAdd
		doctor.patientId = $scope.user.id

		appFactory.addDoctorToPatient(doctor, function(data){
			$scope.doctor_added = data;
			if ($scope.doctor_added == "Error: no patient found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});

	}
	$scope.addEnterpriseToPatient = function(){
		debugger
		var doctor = {}
		enterprise.enterpriseId = $scope.enterpriseToAdd
		enterprise.patientId = $scope.user.id

		appFactory.addEnterpriseToPatient(enterprise, function(data){
			$scope.enterprise_added = data;
			if ($scope.enterprise_added == "Error: no patient found"){
				$("#error_holder").show();
				$("#success_holder").hide();
			} else{
				$("#success_holder").show();
				$("#error_holder").hide();
			}
		});

	}

	$scope.queryPatient = function(){

		var id = $scope.patient_id;

		appFactory.queryPatient(id, function(data){
			$scope.query_patient = data;
			if ($scope.query_patient == "Could not locate patient"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordPatient = function(){

		appFactory.recordPatient($scope.patient, function(data){
			$scope.create_patient = data;
			$("#success_create").show();
		});
	}

	$scope.queryDoctor = function(){

		var id = $scope.doctor_id;

		appFactory.queryDoctor(id, function(data){
			$scope.query_doctor = data;
			if ($scope.query_doctor == "Could not locate doctor"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.queryEnterprise = function(){

		var id = $scope.enterprise_id;

		appFactory.queryEnterprise(id, function(data){
			$scope.query_enterprise = data;
			if ($scope.query_doctor == "Could not locate enterprise"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
		});
	}

	$scope.recordDoctor = function(){
		debugger
		appFactory.recordDoctor($scope.doctor, function(data){
			$scope.create_doctor = data;
			$("#success_create").show();
		});
	}

	$scope.recordEnterprise = function(){
		debugger
		appFactory.recordEnterprise($scope.enterprise, function(data){
			$scope.create_enterprise = data;
			$("#success_create").show();
		});
	}

	$scope.getUserData = function(){
		var id = window.location.href.split('=')[1]
		appFactory.queryPatient(id, function(data){
			$scope.user = data;
			if ($scope.user == "Could not locate patient"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}

			var doctors = $scope.user.doctors
			$scope.user.doctors = []
			for(var i = 0; i < doctors.length; i++){
				appFactory.queryDoctor(doctors[i], function(data){
					$scope.user.doctors.push(data)
					if ($scope.user == "Could not locate doctor"){
						console.log()
						$("#error_query").show();
					} else{
						$("#error_query").hide();
					}
				})
			}
		});
	}

	$scope.getEnterpriseData = function(){
		var id = window.location.href.split('=')[1]
		appFactory.queryEnterprise(id, function(data){
			$scope.enterprise = data;
			if ($scope.user == "Could not locate enterprise"){
				console.log()
				$("#error_query").show();
			} else{
				$("#error_query").hide();
			}
			var name = $scope.enterprise.name
			var id = $scope.enterprise.id
			var users = $scope.enterprise.patients
			$scope.enterprise.patients = []
			for(var i = 0; i < users.length; i++){
				appFactory.queryPatient(patients[i], function(data){
					$scope.enterprise.patients.push(data)
					if ($scope.enterprise == "Could not locate patient"){
						console.log()
						$("#error_query").show();
					} else{
						$("#error_query").hide();
					}
				})
			}
		});
	}

});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

	factory.addDoctorToPatient = function(data, callback){

		var doctor = data.patientId + "-" + data.doctorId;

    	$http.get('/add_doctor_to_patient/'+doctor).success(function(output){
			callback(output)
		});
	}
	factory.addDEnterpriseToPatient = function(data, callback){

		var enterprise = data.patientId + "-" + data.enterpriseId;

    	$http.get('/add_enterprise_to_patient/'+doctor).success(function(output){
			callback(output)
		});
	}


	factory.queryPatient = function(id, callback){
    	$http.get('/get_patient/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordPatient = function(data, callback){

		var patient = data.id + "-" + data.cpf + "-" + data.name + "-" + data.sex + "-" + data.phone 
		+ "-" + data.email + "-" + data.height + "-" + data.weight + "-" + data.age + "-" + data.bloodType;

    	$http.get('/add_patient/'+patient).success(function(output){
			callback(output)
		});
	}

	factory.queryDoctor = function(id, callback){
    	$http.get('/get_doctor/'+id).success(function(output){
			callback(output)
		});
	}

	factory.queryEnterprise = function(id, callback){
    	$http.get('/get_enterprise/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordDoctor = function(data, callback){

		var doctor = data.id + "-" + data.crm + "-" + data.cpf + "-" + data.name + "-" + data.phone + "-" + data.email

    	$http.get('/add_doctor/'+doctor).success(function(output){
			callback(output)
		});
	}

	factory.recordEnterprise = function(data, callback){

		var enterprise = data.id + "-" + data.cnpj + "-" + data.name + "-" + data.phone + "-" + data.email 

    	$http.get('/add_enterprise/'+enterprise).success(function(output){
			callback(output)
		});
	}
	return factory; 
});