;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('CreateStudentsCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.statuses = Student.statuses;

			$scope.gender = {
				 male : true,
				 female: false
			};

			$scope.algo;

			console.log($scope.gender.male);
			$scope.student = {
				idNumber: '',
				firstName: '',
				lastName: '',
				email: '',
				status: 1,
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


			$scope.handleOfGender = function (){
				console.log($scope.gender.male);
				console.log($scope.gender.female);
				if($scope.gender.male)
					$scope.student.gender = true;
				else
					$scope.student.gender = false;
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

			$scope.createStudent = function() {
				console.log($scope.student);
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

					$scope.studentForm.$setPristine();
					$scope.studentForm.$setUntouched();
					$scope.student = {
							idNumber: '',
							firstName: '',
							lastName: '',
							email: '',
							status: 1,
							placeOfBirth: '',
							address: '',
							birthdate: new Date(),
							gender: false,
							nationality: '',
							phoneNumber: ''
						};
				});
			};
		}
	]);
})(angular);
