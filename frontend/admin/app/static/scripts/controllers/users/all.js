;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('AllUsersCtrl', ['$scope', '$modal', 'ngToast', 'User',
		function($scope, $modal, ngToast, User) {
			$scope.users = null;

			$scope.statusToString = User.statusToString;

			function requestUsers() {
				User.findAll().success(function(response) {
					$scope.users = response;
				});
			}

			$scope.openChangePasswordModal = function(userId, username) {
				var modalInstance = $modal.open({
					animation: true,
					controller: 'changePasswordModalCtrl',
					templateUrl: 'static/views/users/change_password_modal.html',
					resolve: {
						user: function() {
							return {
								id: userId,
								username: username
							};
						}
					}
				});

				modalInstance.result.then(function(newPassword) {
					User.changePassword(userId, newPassword).success(function() {
						ngToast.create('Se cambió la clave del usuario!');
					});
				});
			};

			$scope.deleteUser = function(userId, username) {
				User.delete(userId).success(function() {
					ngToast.create('Se eliminó el usuario!');
				});
			};

			requestUsers();
		}
	]);
})(angular);
