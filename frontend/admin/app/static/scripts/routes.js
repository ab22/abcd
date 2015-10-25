;(function(angular) {
'use strict';

	angular.module('app')
	.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider){
			var viewsPath = 'static/views/';
			$urlRouterProvider.otherwise('/home');

			$stateProvider.
				state('home',{
					url: '/home',
					templateUrl: viewsPath+'home.html'
			});

}]);

})(angular);
