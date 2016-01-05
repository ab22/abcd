;(function(angular) {
	'use strict';

	angular.module('app.services').factory('User', ['$http', 'Api',
		function($http, Api) {
			var userService = {
				// Current contains all methods that don't require any user id
				// to pass as parameters. All of these functions modify the
				// logged user's data, so the user id is taken from the session
				// cookie.
				current: {}
			};

			userService.statuses = [
				{
					id: 0,
					name: 'Habilitado'
				},
				{
					id: 1,
					name: 'Deshabilitado'
				}
			];

			userService.findAll = function() {
				return $http({
					url: Api.getRoute('user/findAll/'),
					method: 'GET'
				});
			};

			userService.statusToString = function(statusId) {
				var statuses = userService.statuses;

				for (var i in statuses) {
					var status = statuses[i];

					if (status.id === statusId) {
						return status.name;
					}
				}

				return '';
			};

			userService.findById = function(userId) {
				return $http({
					url: Api.getRoute('user/findById/' + userId),
					method: 'GET'
				});
			};

			userService.findByUsername = function(username) {
				return $http({
					url: Api.getRoute('user/findByUsername/' + username),
					method: 'GET'
				});
			};

			userService.edit = function(userInfo) {
				return $http({
					url: Api.getRoute('user/edit/'),
					method: 'POST',
					data: userInfo
				});
			};

			userService.create = function(user) {
				return $http({
					url: Api.getRoute('user/create/'),
					method: 'POST',
					data: user
				});
			};

			userService.changePassword = function(userId, newPassword) {
				return $http({
					url: Api.getRoute('user/changePassword/'),
					method: 'POST',
					data: {
						userId: userId,
						newPassword: newPassword
					}
				});
			};

			userService.delete = function(userId) {
				return $http({
					url: Api.getRoute('user/delete/'),
					method: 'POST',
					data: {
						userId: userId
					}
				});
			};

			userService.getProfile = function() {
				return $http({
					url: Api.getRoute('user/profile/'),
					method: 'GET'
				});
			};

			userService.current.changePassword = function(newPassword) {
				return $http({
					url: Api.getRoute('user/current/changePassword/'),
					method: 'POST',
					data: {
						newPassword: newPassword
					}
				});
			};

			userService.current.changeEmail = function(newEmail) {
				return $http({
					url: Api.getRoute('user/current/changeEmail/'),
					method: 'POST',
					data: {
						newEmail: newEmail
					}
				});
			};

			userService.current.changeFullName = function(firstName, lastName) {
				return $http({
					url: Api.getRoute('user/current/changeFullName/'),
					method: 'POST',
					data: {
						firstName: firstName,
						lastName: lastName
					}
				});
			};

			userService.current.getPrivileges = function() {
				return $http({
					url: Api.getRoute('user/current/privileges/'),
					method: 'GET'
				});
			};

			return userService;
		}
	]);
})(angular);
