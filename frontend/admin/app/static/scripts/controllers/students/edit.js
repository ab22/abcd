;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('EditStudentCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.studentNotFound = false;

			$scope.gender = '';

			$scope.student = {
				id: 0,
				idNumber: $stateParams.studentIdNumber,
				firstName: '',
				lastName:'',
				status: '',
				placeOfBirth: '',
				address: '',
				birthdate: new Date(),
				gender: false,
				nationality: '',
				phoneNumber: ''
			};

			$scope.datetimePickers = {
				birthDate: {
					opened: false,
					date: new Date(),
				},
				open: function(datetimePicker) {
					datetimePicker.opened = true;
				},

				format: 'dd/MM/yyyy'
			};

			Student.findByIdNumber($scope.student.idNumber).success(function(response) {
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



			$scope.handleOfGender = function (){
				if ($scope.gender === "1") {
					$scope.student.gender = true;
				} else {
					$scope.student.gender = false;
				}
			};

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
					$location.path('/main/student/all');
				});
			};

		}
	]);
})(angular);
