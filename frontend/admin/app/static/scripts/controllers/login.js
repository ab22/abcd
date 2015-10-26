
;(function(angular) {
	'use strict';
angular.module('app.controllers').controller('LoginCtrl', ['$scope','$location',
		function($scope,$location) {
			$scope.credential = {
					username: "",
					password: ""
			};
			$scope.signIn=function(){


				//Call backendService
			}

}]);
})(angular);

