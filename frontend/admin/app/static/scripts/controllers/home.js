;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('HomeCtrl', ['$scope','$location', 'Auth',
		function($scope, $location, Auth) {

			$scope.signOut = function() {
				Auth.logout().success(function() {
					$location.path('/login');
				});
			};

		}
	]);
})(angular);
