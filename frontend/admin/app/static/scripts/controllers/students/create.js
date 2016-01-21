;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('CreateStudentsCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'Student',
		function($scope, $stateParams, $location, ngToast, Student) {
			$scope.statuses = Student.statuses;

			$scope.student = {
				id :0,
				idNumber: '',
				firstName: '',
				lastName: '',
				email: '',
				status: '',
				placeOfBirth: '',
				address: '',
				birthdate: '',
				gender: false,
				nationality: '',
				phoneNumber: ''
			};
			sweetAlert("Oops...", "Something went wrong!", "error");
			$scope.onStudentIdChange = function() {
				//missing request to backend
			};

			$scope.createStudent = function() {
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
							Id: 0,
							firstName: '',
							lastName: '',
							email: '',
							status: '',
							placeOfBirth: '',
							address: '',
							birthdate: '',
							gender: false,
							nationality: '',
							phoneNumber: ''
						};
				});
			};
		}
	]);
})(angular);
