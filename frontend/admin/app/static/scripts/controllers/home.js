
;(function(angular) {
	'use strict';
angular.module('app.controllers').controller('HomeCtrl', ['$scope','$location',
		function($scope,$location) {
			$scope.changeLoginView=function(){
				$location.path('/home/login');
			}

   }
]);
})(angular);
