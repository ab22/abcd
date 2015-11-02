;(function(angular) {
	'use strict';

	angular.module('app.services').factory('Auth', ['$http', 'Api',
		function($http, Api) {
			var authService = {};

			authService.login = function(credentials) {
				return $http({
					url: Api.getRoute('auth/login/'),
					method: 'POST',
					data: credentials
				});
			};

			authService.checkAuthentication = function() {
				return $http({
					url: Api.getRoute('auth/checkAuthentication/'),
					method: 'POST'
				});
			};

			authService.logout = function() {
				return $http({
					url: Api.getRoute('auth/logout/'),
					method: 'POST'
				});
			};

			return authService;
		}
	]);
})(angular);
