;(function(angular) {
	'use strict';

	angular.module('app').config(['$stateProvider', '$urlRouterProvider',
		function($stateProvider, $urlRouterProvider) {
			var viewsPath = '/static/views/';
			$urlRouterProvider.otherwise('/home');

			$stateProvider.
				state('home', {
					url: '/home',
					templateUrl: viewsPath + 'home.html',
					controller:'HomeCtrl',
					requiresAuthentication: true
				}).state('login', {
					url: '/login',
					templateUrl: viewsPath + 'login.html',
					controller:'LoginCtrl',
					requiresAuthentication: false
				}).state('home.users', {
					url: '/users',
					templateUrl: viewsPath + 'users/layout.html',
					requiresAuthentication: true
				}).state('home.users.all', {
					url: '/all',
					templateUrl: viewsPath + 'users/all.html',
					controller: 'AllUsersCtrl',
					requiresAuthentication: true
				});
		}
	]);
})(angular);
