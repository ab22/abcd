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
					templateUrl: 'static/views/users/changePasswordModal.html',
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

			requestUsers();
		}
	]);
})(angular);
