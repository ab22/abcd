;(function(angular) {
	'use strict';

	angular.module('app.services').factory('User', ['$http', 'Api',
		function($http, Api) {
			var userService = {};

			userService.findAll = function() {
				return $http({
					url: Api.getRoute('user/findAll'),
					method: 'GET'
				});
			};

			userService.statusToString = function(statusId) {
				switch (statusId){
					case 0:
						return 'Activo';
					case 1:
						return 'Deshabilitado';
					default:
						return '';
				}
			};

			return userService;
		}
	]);
})(angular);
