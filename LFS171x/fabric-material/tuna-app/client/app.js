// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.user_id = "";
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

	$scope.recordTuna = function(){

		appFactory.recordTuna($scope.tuna, function(data){
			$scope.create_tuna = data;
			$("#success_create").show();
		});
	}

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

	$scope.goToMyProfile = function(){
		window.location.href = "./my_profile.html?id=" + $scope.user_id;
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

	$scope.recordDoctor = function(){
		debugger
		appFactory.recordDoctor($scope.doctor, function(data){
			$scope.create_doctor = data;
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
});

// Angular Factory
app.factory('appFactory', function($http){
	
	var factory = {};

    factory.queryAllTuna = function(callback){

    	$http.get('/get_all_tuna/').success(function(output){
			callback(output)
		});
	}

	factory.queryTuna = function(id, callback){
    	$http.get('/get_tuna/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordTuna = function(data, callback){

		data.location = data.longitude + ", "+ data.latitude;

		var tuna = data.id + "-" + data.location + "-" + data.timestamp + "-" + data.holder + "-" + data.vessel;

    	$http.get('/add_tuna/'+tuna).success(function(output){
			callback(output)
		});
	}

	factory.changeHolder = function(data, callback){

		var holder = data.id + "-" + data.name;

    	$http.get('/change_holder/'+holder).success(function(output){
			callback(output)
		});
	}

	factory.addDoctorToPatient = function(data, callback){

		var doctor = data.patientId + "-" + data.doctorId;

    	$http.get('/add_doctor_to_patient/'+doctor).success(function(output){
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

	factory.recordDoctor = function(data, callback){

		var doctor = data.id + "-" + data.crm + "-" + data.cpf + "-" + data.name + "-" + data.phone + "-" + data.email

    	$http.get('/add_doctor/'+doctor).success(function(output){
			callback(output)
		});
	}

	return factory;
});


