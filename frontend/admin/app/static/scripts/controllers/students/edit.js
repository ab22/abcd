;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('EditStudentCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.studentNotFound = false;
			$scope.student = {
				id: parseInt($stateParams.studentId) || 0,
				email:'',
				firstname: '',
				lastname: '',
				status: 0
			};

			Student.findById($scope.student.id).success(function(response) {
				$scope.checkStatus = $scope.student.status;
				$scope.student = response;
				$scope.checkStatus = $scope.student.status;
			}).error(function(response) {
				ngToast.create({
					className: 'danger',
					content: 'No se encontró el estudiante',
					dismissButton: true
				});
				$scope.studentNotFound  = response.status === 404;
			});

			$scope.editStudent = function() {
				Student.edit($scope.student).success(function(response) {
					if (!response.success) {
						ngToast.create({
							className: 'danger',
							content: response.errorMessage,
							dismissButton: true
						});

						return;
					}

					ngToast.create('Se actualizó el usuario!');
					$location.path('/main/student/all');
				});
			};

		}
	]);
})(angular);
