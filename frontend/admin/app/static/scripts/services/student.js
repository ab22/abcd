;(function(angular) {
	'use strict';

	angular.module('app.services').factory('Student', ['$http', 'Api',
		function($http, Api) {
			var studentService = {};

			studentService.findAll = function() {
				return $http({
					url: Api.getRoute('student/findAll/'),
					method: 'GET'
				});
			};

			studentService.FindByUser =	function(id){
				return $htpp({
					url: Api.getRoute('student/findById/'+id),
					method: 'POST'
				});
			};

			return studentService;
		}
	]);
})(angular);
