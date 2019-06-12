// SPDX-License-Identifier: Apache-2.0

'use strict';

var app = angular.module('application', []);

// Angular Controller
app.controller('appController', function($scope, appFactory){

	$("#success_holder").hide();
	$("#success_create").hide();
	$("#error_holder").hide();
	$("#error_query").hide();
	
	$scope.registerDoc = function(){

		appFactory.registerDoc($scope.tuna, function(data){

			/* Verify if its a valid doc */
			
			//register_doc is the function that will invoke the chainCode
			$scope.register_doc = data;
			$("#success_reg").show();
		});
	}

	return factory;
});


