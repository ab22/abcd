;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('confirmationModalCtrl', ['$scope', '$modalInstance', 'data',
		function($scope, $modalInstance, data) {
			$scope.modal = data;

			$scope.accept = function() {
				$modalInstance.close();
			};

			$scope.cancel = function() {
				$modalInstance.dismiss();
			};
		}
	]);
})(angular);
