;(function(angular) {
	'use strict';

	var app = angular.module('app', [
		'ngRoute',
		'ui.router'
	]);

	app.config(function($interpolateProvider){
	    $interpolateProvider.startSymbol('[[');
	    $interpolateProvider.endSymbol(']]');
	});

})(angular);

