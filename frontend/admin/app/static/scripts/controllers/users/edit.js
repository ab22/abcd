;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('EditUserCtrl', ['$scope', '$stateParams',
		function($scope, $stateParams) {
			$scope.user = {
				id: parseInt($stateParams.userId) || 0,
				username: '',
				password: '',
				firstName: '',
				lastName: '',
				email: ''
			};
		}
	]);
})(angular);
