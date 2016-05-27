;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('EditStudentCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.studentNotFound = false;

			$scope.student = {
				id: 0,
				idNumber: $stateParams.studentIdNumber,
				firstName: '',
				lastName:'',
				placeOfBirth: '',
				address: '',
				birthdate: new Date(),
				gender: false,
				nationality: '',
				phoneNumber: ''
			};

			$scope.datetimePickers = {
				birthDate: {
					opened: false
				},

				open: function(datetimePicker) {
					datetimePicker.opened = true;
				},

				format: 'dd/MM/yyyy'
			};

			Student.findByIdNumber($scope.student.idNumber).success(function(response) {
				$scope.student = response;
			}).error(function(response) {
				ngToast.create({
					className: 'danger',
					content: 'No se encontró el estudiante',
					dismissButton: true
				});
				$scope.studentNotFound  = response.status === 404;
			});

			$scope.onStudentIdChange = function() {
				if ($scope.student.idNumber === '') {
					$scope.studentForm.idNumber.$setValidity('available', true);
					return;
				}

				Student.findByIdNumber($scope.student.idNumber).success(function() {
					$scope.studentForm.idNumber.$setValidity('available', false);
					ngToast.create({
							className: 'danger',
							content: 'Numero de identidad ya existente',
							dismissButton: true
						});
				}).error(function(response, status) {
					$scope.studentForm.idNumber.$setValidity('available', status === 404);
				});
			};

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
					$location.path('/main/students/all');
				});
			};

		}
	]);
})(angular);
