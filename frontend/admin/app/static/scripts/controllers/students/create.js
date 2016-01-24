;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('CreateStudentsCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.statuses = Student.statuses;

			$scope.gender = {
				male : false,
				female: false
			};

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


			$scope.handleOfGender = function (gender){
				console.log("entre");
				if($scope.gender.male)
					$scope.student.gender = true;
				if($scope.gender.female)
					$scope.student.gender = false;
			};

			$scope.onStudentIdChange = function() {
				if ($scope.student.idNumber === '') {
					$scope.studentForm.idNumber.$setValidity('available', true);
					return;
				}

				Student.findByIdNumber($scope.student.idNumber).success(function() {
					$scope.studentForm.idNumber.$setValidity('available', false);
					console.log("estudiante ya existe");
				}).error(function(response, status) {
					$scope.studentForm.idNumber.$setValidity('available', status === 404);
					console.log("estudiante no existe");
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
							status: 0,
							placeOfBirth: '',
							address: '',
							birthdate: new Data(),
							gender: null,
							nationality: '',
							phoneNumber: ''
						};
				});
			};
		}
	]);
})(angular);
