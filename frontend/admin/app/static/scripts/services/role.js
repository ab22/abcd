;(function(angular) {
	'use strict';

	angular.module('app.services').factory('Role', ['$http', 'Api',
		function($http, Api) {
			var roleService = {};

			roleService.findAll = function() {
				return $http({
					url: Api.getRoute('role/findAll/'),
					method: 'GET'
				});
			};

			return roleService;
		}
	]);
})(angular);
