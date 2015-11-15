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
					User.changePassword($scope.user.id, newPassword).success(function() {
						ngToast.create('Se cambió la clave del usuario!');
					});
				});
			};

		}
	]);
})(angular);
