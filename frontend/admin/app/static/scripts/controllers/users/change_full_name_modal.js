;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('changeFullNameModalCtrl', ['$scope', '$modalInstance', 'user',
		function($scope, $modalInstance, user) {
			$scope.user = user;
			$scope.firstName = '';
			$scope.lastName = '';

			$scope.accept = function() {
				$scope.changeFullNameForm.$submitted = true;

				if ($scope.changeFullNameForm.$invalid) {
					return;
				}

				$modalInstance.close({
					firstName: $scope.firstName,
					lastName: $scope.lastName
				});
			};

			$scope.cancel = function() {
				$modalInstance.dismiss();
			};
		}
	]);
})(angular);
