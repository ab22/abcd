;(function(angular) {
	'use strict';

	var app= angular.module('app', ['ui.router']);

	app.config(function($interpolateProvider){
	    $interpolateProvider.startSymbol('[[');
	    $interpolateProvider.endSymbol(']]');
	});

})(angular);

