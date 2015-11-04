;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('EditUserCtrl', ['$scope', '$stateParams', 'User',
		function($scope, $stateParams, User) {
			$scope.statuses = User.statuses;
			$scope.user = {
				id: parseInt($stateParams.userId) || 0,
				username: '',
				firstName: '',
				lastName: '',
				email: '',
				status: 0
			};

			User.findById($scope.user.id).success(function(response) {
				$scope.user = response;
			}).error(function() {
			});

			$scope.editUser = function() {
			}
		}
	]);
})(angular);
