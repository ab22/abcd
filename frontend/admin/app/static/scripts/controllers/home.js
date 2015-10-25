
;(function(angular) {
	'use strict';
angular.module('app.controllers').controller('HomeCtrl', ['$scope','$location'
		function($scope,$location) {
			$scopre.changeLoginView=function(){
				$location.path('/home/login');
			}


})(angular);
