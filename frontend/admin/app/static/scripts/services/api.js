;(function(angular) {
	'use strict';

	angular.module('app.services').factory('Api', [function() {
			var api = {};

			api.getRoute = function(route) {
				var url = 'api/';

				return url + route;
			};

			return api;
		}
	]);
})(angular);
