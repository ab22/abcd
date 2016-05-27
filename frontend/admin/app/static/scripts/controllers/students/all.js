;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('AllStudentsCtrl', ['$scope', '$modal', 'ngToast', 'Student',
		function($scope, $modal, ngToast, Student) {
			$scope.students = null;
			$scope.rowCollection = [];

			$scope.statusToString = Student.statusToString;

			function requestStudents() {
				Student.findAll().success(function(response) {
					$scope.students = response;
					$scope.rowCollection = [];

					$scope.students.forEach(function(student) {
						$scope.rowCollection.push(student);
					});
				});
			}

			function removeStudent(studentId) {
				var students = $scope.students;

				for (var i in students) {
					var student = students[i];

					if (student.id === studentId) {
						students.splice(i, 1);
					}
				}
			}

			$scope.deleteStudent = function(studentId, studentname) {
				var modalData = {
					title: 'Eliminar estudiante: ' + studentname,
					body: '¿Estás seguro de eliminar el studiante <strong>' + studentname + '</strong>?'
				};

				var modalInstance = $modal.open({
					animation: true,
					controller: 'confirmationModalCtrl',
					templateUrl: 'static/views/common_modals/confirmation_modal.html',
					resolve: {
						data: function() {
							return modalData;
						}
					}
				});

				modalInstance.result.then(function() {
					Student.delete(studentId).success(function() {
						ngToast.create('Se eliminó el estudiante!');
						removeStudent(studentId);
					});
				});
			};

			$scope.reloadStudents = function() {
				requestStudents();
			};

			requestStudents();
		}
	]);
})(angular);
