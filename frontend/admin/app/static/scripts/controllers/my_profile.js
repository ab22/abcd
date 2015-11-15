;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('MyProfileCtrl', ['$scope', '$modal', 'ngToast', 'User',
		function($scope, $modal, ngToast, User) {
			$scope.user = null;
			$scope.isLoading = true;
			$scope.statusToString = User.statusToString;

			User.getProfile().success(function(response) {
				$scope.user = response;
			}).finally(function() {
				$scope.isLoading = false;
			});

			$scope.openChangePasswordModal = function() {
				var modalInstance = $modal.open({
					animation: true,
					controller: 'changePasswordModalCtrl',
					templateUrl: 'static/views/users/change_password_modal.html',
					resolve: {
						user: function() {
							return {
								id: $scope.user.id,
								username: $scope.user.username
							};
						}
					}
				});

				modalInstance.result.then(function(newPassword) {
					User.current.changePassword(newPassword).success(function() {
						ngToast.create('Se actualizó tu clave!');
					});
				});
			};

			$scope.openChangeEmailModal = function() {
				var modalInstance = $modal.open({
					animation: true,
					controller: 'changeEmailModalCtrl',
					templateUrl: 'static/views/users/change_email_modal.html',
					resolve: {
						user: function() {
							return {
								id: $scope.user.id,
								username: $scope.user.username,
								email: $scope.user.email
							};
						}
					}
				});

				modalInstance.result.then(function(newEmail) {
					User.current.changeEmail(newEmail).success(function() {
						ngToast.create('Se actualizó tu correo!');
						$scope.user.email = newEmail;
					});
				});
			};

			$scope.openChangeFullNameModal = function() {
				var modalInstance = $modal.open({
					animation: true,
					controller: 'changeFullNameModalCtrl',
					templateUrl: 'static/views/users/change_full_name_modal.html',
					resolve: {
						user: function() {
							return {
								id: $scope.user.id,
								username: $scope.user.username,
								firstName: $scope.user.firstName,
								lastName: $scope.user.lastName
							};
						}
					}
				});

				modalInstance.result.then(function(newInfo) {
					User.current.changeFullName(newInfo.firstName, newInfo.lastName).success(function() {
						$scope.user.firstName = newInfo.firstName;
						$scope.user.lastName = newInfo.lastName;
						ngToast.create('Se actualizó tu nombre completo!');
					});
				});
			};
		}
	]);
})(angular);
