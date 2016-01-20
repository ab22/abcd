;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('CreateUserCtrl', ['$scope', '$stateParams', '$location', 'ngToast', 'User',
		function($scope, $stateParams, $location, ngToast, User) {
			$scope.statuses = User.statuses;
			$scope.passwordRepeat = '';
			$scope.user = {
				username: '',
				password: '',
				firstName: '',
				lastName: '',
				email: '',
				isAdmin: false,
				isTeacher: false
			};

			function passwordsMatch() {
				return $scope.user.password === $scope.passwordRepeat;
			}

			function setPasswordFieldsValidities(match) {
				$scope.userForm.password.$setValidity('invalidPassword', match);
				$scope.userForm.repeatPassword.$setValidity('invalidPassword', match);
			}

			$scope.onUsernameChange = function() {
				if ($scope.user.username === '') {
					$scope.userForm.username.$setValidity('available', true);
					return;
				}

				User.findByUsername($scope.user.username).success(function() {
					$scope.userForm.username.$setValidity('available', false);
				}).error(function(response, status) {
					$scope.userForm.username.$setValidity('available', status === 404);
				});
			};

			$scope.onPasswordChange = function() {
				var match = passwordsMatch();
				setPasswordFieldsValidities(match);
			};

			$scope.createUser = function() {
				var match = passwordsMatch();
				setPasswordFieldsValidities(match);

				if (!match) {
					$scope.user.password = '';
					$scope.passwordRepeat = '';
					return;
				}

				User.create($scope.user).success(function(response) {
					if (!response.success) {
						ngToast.create({
							className: 'danger',
							content: response.errorMessage,
							dismissButton: true
						});
						return;
					}

					ngToast.create('Se creó el usuario!');
					$scope.userForm.$setPristine();
					$scope.userForm.$setUntouched();
					$scope.passwordRepeat = '';
					$scope.user = {
						username: '',
						password: '',
						firstName: '',
						lastName: '',
						email: '',
						isAdmin: false,
						isTeacher: false
					};
				});
			};
		}
	]);
})(angular);
