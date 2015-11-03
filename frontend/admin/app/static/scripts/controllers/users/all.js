;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('AllUsersCtrl', ['$scope', 'User',
		function($scope, User) {
			$scope.users = null;

			$scope.statusToString = User.statusToString;

			function requestUsers() {
				User.findAll().success(function(response) {
					$scope.users = response;
				});
			}

			requestUsers();
		}
	]);
})(angular);
