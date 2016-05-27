;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('EditUserCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'User',
		function($scope, $stateParams, $location, ngToast, User) {
			$scope.userNotFound = false;
			$scope.statuses = User.statuses;
			$scope.originalUsername = '';
			$scope.user = {
				id: parseInt($stateParams.userId) || 0,
				username: '',
				firstName: '',
				lastName: '',
				email: '',
				status: 0,
				isAdmin: false,
				isTeacher: false
			};

			User.findById($scope.user.id).success(function(response) {
				$scope.user = response;
				$scope.originalUsername = $scope.user.username;
			}).error(function(response) {
				ngToast.create({
					className: 'danger',
					content: 'No se encontró el usuario!',
					dismissButton: true
				});
				$scope.userNotFound = response.status === 404;
			});

			$scope.editUser = function() {
				User.edit($scope.user).success(function(response) {
					if (!response.success) {
						ngToast.create({
							className: 'danger',
							content: response.errorMessage,
							dismissButton: true
						});
						return;
					}

					ngToast.create('Se actualizó el usuario!');
					$location.path('/main/users/all');
				});
			};

			$scope.onUsernameChange = function() {
				if ($scope.user.username === $scope.originalUsername) {
					$scope.userForm.username.$setValidity('available', true);
					return;
				}

				User.findByUsername($scope.user.username).success(function() {
					$scope.userForm.username.$setValidity('available', false);
				}).error(function(response, status) {
					$scope.userForm.username.$setValidity('available', status === 404);
				});
			};
		}
	]);
})(angular);
