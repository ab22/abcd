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
					requiresAuthentication: true,
					resolve: {
						privileges: ['User', function(User) {
							return User.current.getPrivileges().then(function(response) {
								return response.data;
							},

							function() {
								return {};
							});
						}]
					}
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
				}).state('main.profile', {
					url: '/profile',
					templateUrl: viewsPath + 'my_profile.html',
					controller:'MyProfileCtrl',
					requiresAuthentication: true
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
					requiresAuthentication: true
				}).state('main.users.create', {
					url: '/create',
					templateUrl: viewsPath + 'users/create.html',
					controller: 'CreateUserCtrl',
					requiresAuthentication: true
				}).state('main.students', {
					url: '/students',
					templateUrl: viewsPath + 'students/layout.html',
					requiresAuthentication: true
				}).state('main.students.all', {
					url: '/all',
					templateUrl: viewsPath + 'students/all.html',
					controller: 'AllStudentsCtrl',
					requiresAuthentication: true
				});
		}
	]);
})(angular);
