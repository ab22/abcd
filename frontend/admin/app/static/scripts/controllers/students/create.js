;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('CreateStudentsCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.statuses = Student.statuses;
			$scope.checkStatus = false;
			$scope.student = {
				firstName: '',
				lastName: '',
				status: 0
			};

			$scope.setStatus = function(){
				if($scope.checkStatus)
					$scope.status = 1;
				$scope.status = 0;
			console.log($scope.status);
			}

			$scope.onStudentIdChange = function() {


				//missing request to backend

			};

			$scope.createStudent = function() {
					console.log("entro");

				Student.create($scope.student).success(function(response) {
					if (!response.success) {
						ngToast.create({
							className: 'danger',
							content: response.errorMessage,
							dismissButton: true
						});
						return;
					}
					ngToast.create('El estudiante se ha creado!');

					$scope.student = {
						firstName: '',
						lastName: '',
						status: 0
					};
				});
			};
		}
	]);
})(angular);
