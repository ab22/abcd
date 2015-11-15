;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('MyProfileCtrl', ['$scope', 'User',
		function($scope, User) {
			$scope.user = null;
			$scope.isLoading = true;
			$scope.statusToString = User.statusToString;

			User.getProfile().success(function(response) {
				$scope.user = response;
			}).finally(function() {
				$scope.isLoading = false;
			});

		}
	]);
})(angular);
