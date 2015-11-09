;(function(angular) {
	'use strict';

	angular.module('app').config(['$stateProvider', '$urlRouterProvider',
		function($stateProvider, $urlRouterProvider) {
			var viewsPath = 'static/views/';
			$urlRouterProvider.otherwise('/main/home');

			$stateProvider.
				state('main', {
					url: '/main',
					templateUrl: viewsPath + 'layout.html',
					controller:'MainLayoutCtrl',
					requiresAuthentication: true
				}).
				state('main.home', {
					url: '/home',
					templateUrl: viewsPath + 'home.html',
					controller:'HomeCtrl',
					requiresAuthentication: true
				}).state('login', {
					url: '/login',
					templateUrl: viewsPath + 'login.html',
					controller:'LoginCtrl',
					requiresAuthentication: false
				}).state('main.users', {
					url: '/users',
					templateUrl: viewsPath + 'users/layout.html',
					requiresAuthentication: true
				}).state('main.users.all', {
					url: '/all',
					templateUrl: viewsPath + 'users/all.html',
					controller: 'AllUsersCtrl',
					requiresAuthentication: true
				}).state('main.users.edit', {
					url: '/edit/{userId}',
					templateUrl: viewsPath + 'users/edit.html',
					controller: 'EditUserCtrl',
					requiresAuthentication: true,
					resolve: {
						roles: ['Role', function(Role) {
							return Role.findAll().then(function(response) {
								return response.data;
							}, function() {
								return [];
							});
						}]
					}
				}).state('main.users.create', {
					url: '/create',
					templateUrl: viewsPath + 'users/create.html',
					controller: 'CreateUserCtrl',
					requiresAuthentication: true
				});
		}
	]);
})(angular);
