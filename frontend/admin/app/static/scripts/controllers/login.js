;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('LoginCtrl', ['$scope','$location',
		function($scope,$location) {

			$scope.authenticate = function(){
				$location.path('/home');
			};

		}
	]);
})(angular);
