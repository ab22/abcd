;(function(angular) {
	'use strict';

	angular.module('app.controllers', []);
	angular.module('app.services', []);
	angular.module('app.directives', []);

	var app = angular.module('app', [
		'app.controllers',
		'app.services',
		'app.directives',

		'smart-table',
		'ngToast',
		'ngRoute',
		'ngSanitize',
		'ui.router',
		'ui.bootstrap',
		'ngLoadingSpinner'
	]);

	app.factory('httpInterceptor', ['$rootScope', '$q', '$location', 'ngToast', function($rootScope, $q, $location, ngToast) {
		var onRequest = function(config) {
			return config;
		};

		var onResponse = function(response) {
			return response;
		};

		var onError = function(rejection) {
			var status = rejection.status;

			// If user needs to authenticate.
			if (status === 401) {
				$location.path('login');
				return $q.reject(rejection);
			}

			if (status === 403) {
				ngToast.create({
					className: 'danger',
					content: 'No tienes acceso a esta función!',
					dismissButton: true
				});

				return $q.reject(rejection);
			}

			if (status === 404) {
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

