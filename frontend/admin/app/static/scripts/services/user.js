;(function(angular) {
	'use strict';

	angular.module('app.services').factory('User', ['$http', 'Api',
		function($http, Api) {
			var userService = {};

			userService.statuses = {
				0: 'Activo',
				1: 'Deshabilitado'
			};

			userService.findAll = function() {
				return $http({
					url: Api.getRoute('user/findAll/'),
					method: 'GET'
				});
			};

			userService.statusToString = function(statusId) {
				var statuses = userService.statuses;

				for (var i in statuses) {
					if (i === statusId) {
						return statuses[i];
					}
				}

				return '';
			};

			userService.findById = function(userId) {
				return $http({
					url: Api.getRoute('user/findById/'),
					method: 'POST',
					data: {
						userId: userId
					}
				});
			};

			return userService;
		}
	]);
})(angular);
