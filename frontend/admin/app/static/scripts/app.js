;(function(angular) {
	'use strict';

	angular.module('app.controllers', []);
	angular.module('app.services', []);

	var app = angular.module('app', [
		'app.controllers',
		'app.services',

		'ngToast',
		'ngRoute',
		'ngSanitize',
		'ui.router',
		'ui.bootstrap'
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

	app.run(['$rootScope', '$state', 'Auth', function($rootScope, $state, Auth) {
		$rootScope.$on('$stateChangeStart', function(event, toState) {
			if (toState.url === '/login') {
				return;
			}

			if (!toState.requiresAuthentication) {
				return;
			}

			Auth.checkAuthentication().error(function(response, status) {
				if (status === 401) {
					if (toState.requiresAuthentication) {
						$state.go('login');
					}
				}
			});
		});
	}]);

})(angular);

