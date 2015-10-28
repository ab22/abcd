;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('HomeCtrl', ['$scope','$location',
		function($scope,$location) {

			$scope.signOut = function(){
				$location.path('/login');
			};

		}
	]);
})(angular);
