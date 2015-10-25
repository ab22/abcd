;(function(angular) {
'use strict';

	angular.module('app')
	.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider){
			var viewsPath = '/static/views/';
			$urlRouterProvider.otherwise('/home/login');

			$stateProvider.
				state('home',{
					url: '/home',
					templateUrl: viewsPath+'home.html'
			}).
				state('home.login',{
					url: '/login',
					templateUrl: viewsPath+'login.html',
					requiresAuthentication: false

				});

}]);

})(angular);
