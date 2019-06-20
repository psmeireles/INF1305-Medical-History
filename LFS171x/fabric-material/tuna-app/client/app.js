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

	$scope.goToMyProfile = function(){
		window.location.href = "./my_profile.html?id=" + $scope.user_id;
	}

	$scope.queryPatient = function(){

		var id = $scope.patient_id;

		appFactory.queryPatient(id, function(data){
			$scope.query_patient = data;
			debugger
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

	factory.queryPatient = function(id, callback){
    	$http.get('/get_patient/'+id).success(function(output){
			callback(output)
		});
	}

	factory.recordPatient = function(data, callback){

		data.location = data.longitude + ", "+ data.latitude;

		var patient = data.id + "-" + data.cpf + "-" + data.name + "-" + data.sex + "-" + data.phone 
		+ "-" + data.email + "-" + data.height + "-" + data.weight + "-" + data.age + "-" + data.bloodType;

    	$http.get('/add_patient/'+patient).success(function(output){
			callback(output)
		});
	}

	return factory;
});


