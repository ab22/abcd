;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('changeEmailModalCtrl', ['$scope', '$modalInstance', 'user',
		function($scope, $modalInstance, user) {
			$scope.user = user;
			$scope.email = user.email;

			$scope.accept = function() {
				$scope.changeEmailForm.$submitted = true;

				if ($scope.changeEmailForm.$invalid) {
					return;
				}

				$modalInstance.close($scope.email);
			};

			$scope.cancel = function() {
				$modalInstance.dismiss();
			};
		}
	]);
})(angular);
