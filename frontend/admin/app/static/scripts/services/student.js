;(function(angular) {
	'use strict';

	angular.module('app.services').factory('Student', ['$http', 'Api',
		function($http, Api) {
			var studentService = {};

			studentService.statuses = [
				{
					id: 0,
					name: 'Habilitado'
				},
				{
					id: 1,
					name: 'Deshabilitado'
				}
			];

			studentService.findAll = function() {
				return $http({
					url: Api.getRoute('student/findAll/'),
					method: 'GET'
				});
			};

			studentService.findById = function(studentId) {
				return $http({
					url: Api.getRoute('student/findById/' + studentId),
					method: 'GET'
				});

			};

			studentService.findByIdNumber = function(studentIdNumber) {
				return $http({
					url: Api.getRoute('student/findByIdNumber/' + studentIdNumber),
					method: 'GET'
				});
			};

			studentService.statusToString = function(statusId) {
				var statuses = studentService.statuses;

				for (var i in statuses) {
					var status = statuses[i];

					if (status.id === statusId) {
						return status.name;
					}
				}

				return '';
			};

			studentService.delete = function(studentId) {
				return $http({
					url: Api.getRoute('student/delete/'),
					method: 'POST',
					data: {
						studentId: studentId
					}
				});
			};

			studentService.create = function(student) {
				return $http({
					url: Api.getRoute('student/create/'),
					method: 'POST',
					data: student
				});
			};

			studentService.edit = function(studentInfo) {
				return $http({
					url: Api.getRoute('student/edit/'),
					method: 'POST',
					data: studentInfo
				});
			};

			return studentService;
		}
	]);
})(angular);
