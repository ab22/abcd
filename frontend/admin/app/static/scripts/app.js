;(function(angular) {
	'use strict';

	angular.module('app.controllers', []);

	var app = angular.module('app', [
		'ngToast',
		'ngRoute',
		'ui.router',
		'app.controllers',
		'app.services'
	]);

	app.factory('httpInterceptor', ['$rootScope', '$q', 'ngToast', function($rootScope, $q, ngToast) {
		var onRequest = function(config) {
			return config;
		};

		var onResponse = function(response) {
			return response;
		};

		var onError = function(rejection) {
			var status = rejection.status;

			if (status === 401 || status === 404) {
				return $q.reject(rejection);
			}

			var message = 'Ocurrió un error al procesar la solicitud! :(';

			ngToast.create({
				className: 'danger',
				content: message,
				dismissButton: true
			});

			return $q.reject(rejection);
		};

		return {
			request: onRequest,
			response: onResponse,
			responseError: onError,
			requestError: onError
		};
	}]);

	app.config(['$httpProvider', function($httpProvider) {
		$httpProvider.interceptors.push('httpInterceptor');
	}]);

})(angular);

