;(function(angular) {
	'use strict';
	angular.module('app').service('LoginService', ['$http', function($http, HostFactory){

	var loginService = {};

	loginService.login = function(user){
		//Call Backend
	};
	return  loginService;
}]);
})(angular);
