;(function(angular) {
	'use strict';

	angular.module('app.services').factory('User', ['$http', 'Api',
		function($http, Api) {
			var userService = {
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
					url: Api.getRoute('user/findById/'),
					method: 'POST',
					data: {
						userId: userId
					}
				});
			};

			userService.findByUsername = function(username) {
				return $http({
					url: Api.getRoute('user/findByUsername/'),
					method: 'POST',
					data: {
						username: username
					}
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

			userService.delete = function(userId, newPassword) {
				return $http({
					url: Api.getRoute('user/delete/'),
					method: 'POST',
					data: {
						userId: userId,
						newPassword: newPassword
					}
				});
			};

			userService.getProfile = function() {
				return $http({
					url: Api.getRoute('user/profile/'),
					method: 'POST'
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

			return userService;
		}
	]);
})(angular);
