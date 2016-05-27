;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('AllUsersCtrl', ['$scope', '$modal', 'ngToast', 'User',
		function($scope, $modal, ngToast, User) {
			$scope.users = null;
			$scope.rowCollection = [];

			$scope.statusToString = User.statusToString;

			function requestUsers() {
				User.findAll().success(function(response) {
					$scope.users = response;
					$scope.rowCollection = [];

					$scope.users.forEach(function(user) {
						$scope.rowCollection.push(user);
					});
				});
			}

			function removeUser(userId) {
				var users = $scope.users;

				for (var i in users) {
					var user = users[i];

					if (user.id === userId) {
						users.splice(i, 1);
					}
				}
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
				var modalData = {
					title: 'Eliminar usuario: ' + username,
					body: '¿Estás seguro de eliminar el usuario <strong>' + username + '</strong>?'
				};

				var modalInstance = $modal.open({
					animation: true,
					controller: 'confirmationModalCtrl',
					templateUrl: 'static/views/common_modals/confirmation_modal.html',
					resolve: {
						data: function() {
							return modalData;
						}
					}
				});

				modalInstance.result.then(function() {
					User.delete(userId).success(function() {
						ngToast.create('Se eliminó el usuario!');
						removeUser(userId);
					});
				});
			};

			$scope.reloadUsers = function() {
				requestUsers();
			};

			requestUsers();
		}
	]);
})(angular);
