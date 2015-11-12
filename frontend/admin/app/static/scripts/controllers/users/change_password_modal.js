;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('changePasswordModalCtrl', ['$scope', '$modalInstance', 'user',
		function($scope, $modalInstance, user) {
			var self = this;
			$scope.user = user;
			$scope.password = '';
			$scope.passwordRepeat = '';
			$scope.showError = false;

			self.clearPasswordFields = function() {
				$scope.password = '';
				$scope.passwordRepeat = '';
			};

			$scope.changePassword = function() {
				$scope.resetPassForm.$submitted = true;

				if ($scope.resetPassForm.$invalid) {
					return;
				}

				if ($scope.password !== $scope.passwordRepeat) {
					$scope.showError = true;
					self.clearPasswordFields();
					return;
				}

				$scope.showError = false;
				$modalInstance.close($scope.password);
			};

			$scope.cancel = function() {
				$modalInstance.dismiss();
			};
		}
	]);
})(angular);
