;(function(angular) {
	'use strict';
		angular.module('app.controllers', []);
	var app = angular.module('app', [
		'ngRoute',
		'ui.router',
		'app.controllers'
	]);

	app.config(function($interpolateProvider){
	    $interpolateProvider.startSymbol('[[');
	    $interpolateProvider.endSymbol(']]');
	});

})(angular);

