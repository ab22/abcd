;(function(angular) {
	'use strict';

	var app= angular.module('app', ['ngRoute']);

	app.config(function($interpolateProvider){
	    $interpolateProvider.startSymbol('[[');
	    $interpolateProvider.endSymbol(']]');
	});

})(angular);

