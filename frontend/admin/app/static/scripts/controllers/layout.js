;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('MainLayoutCtrl', ['$scope','$location', 'Auth', 'privileges',
		function($scope, $location, Auth, privileges) {
			$scope.privileges = privileges;
			$scope.window = {};
			$scope.isCollapsed = true;

			$scope.signOut = function() {
				Auth.logout().success(function() {
					$location.path('/login');
				});
			};

			$scope.activeOption = null;
			$scope.topMenu = [];
			$scope.menu = [
				[
					{
						label: 'Inicio',
						icon: 'fa-home',
						link: '/main/home',
						responsiveOnly: false,
						requiresAdmin: false,
						requiresTeacher: false
					},
					{
						label: 'Asignaturas',
						icon: 'fa-book',
						link: '',
						responsiveOnly: true,
						requiresAdmin: false,
						requiresTeacher: true
					},
					{
						label: 'Alumnos',
						icon: 'fa-child',
						link: '/main/students/all',
						responsiveOnly: true,
						requiresAdmin: false,
						requiresTeacher: true
					},
					{
						label: 'Docentes',
						icon: 'fa-male',
						link: '',
						responsiveOnly: true,
						requiresAdmin: false,
						requiresTeacher: true
					},
					{
						label: 'Ingreso de Notas',
						icon: 'fa-file-text-o',
						link: '',
						responsiveOnly: true,
						requiresAdmin: false,
						requiresTeacher: true
					},
					{
						label: 'Boletines',
						icon: 'fa-bullhorn',
						link: '',
						responsiveOnly: true,
						requiresAdmin: false,
						requiresTeacher: false
					}
				],
				[
					{
						label: 'Usuarios',
						icon: 'fa-users',
						link: '/main/users/all',
						responsiveOnly: true,
						requiresAdmin: true,
						requiresTeacher: false
					},
					{
						label: 'Reporte #1',
						icon: 'fa-calendar',
						link: '',
						responsiveOnly: true,
						requiresAdmin: true,
						requiresTeacher: false
					},
					{
						label: 'Reporte #2',
						icon: 'fa-edit',
						link: '',
						responsiveOnly: true,
						requiresAdmin: true,
						requiresTeacher: false
					},
					{
						label: 'Reporte #3',
						icon: 'fa-area-chart',
						link: '',
						responsiveOnly: true,
						requiresAdmin: true,
						requiresTeacher: false
					},
					{
						label: 'Mi Perfil',
						icon: 'fa-user',
						link: '/main/profile',
						responsiveOnly: false,
						requiresAdmin: false,
						requiresTeacher: false
					},
					{
						label: 'Configuración',
						icon: 'fa-cogs',
						link: '/main/configuration',
						responsiveOnly: false,
						requiresAdmin: true,
						requiresTeacher: false
					},
					{
						label: 'Cerrar Sesión',
						icon: 'fa-sign-out',
						link: '',
						onClick: $scope.signOut,
						responsiveOnly: false,
						requiresAdmin: false,
						requiresTeacher: false
					}
				]
			];

			function hideResponsiveMenu() {
				if ($scope.isResponsiveMode()) {
					$scope.isCollapsed = !$scope.isCollapsed;
				}
			}

			function setActiveOption(option) {
				$scope.activeOption = option;
			}

			$scope.optionOnClick = function(option) {
				hideResponsiveMenu();

				if (typeof option.onClick !== 'undefined') {
					option.onClick();
					return;
				}

				setActiveOption(option);
			};

			$scope.isResponsiveMode = function() {
				return $scope.window.width <= 767;
			};

			$scope.showOption = function(option) {
				if (!option.requiresAdmin && !option.requiresTeacher) {
					return true;
				} else if (option.requiresTeacher && $scope.privileges.isTeacher) {
					return true;
				} else if (option.requiresAdmin && $scope.privileges.isAdmin) {
					return true;
				} else {
					return false;
				}
			};

			$scope.showResponsiveOption = function(option) {
				var showResponsive = !option.responsiveOnly || (option.responsiveOnly && $scope.isResponsiveMode());

				return showResponsive && $scope.showOption(option);
			};

			function determineActiveOption() {
				var currentPath = $location.path();

				for (var i in $scope.menu) {
					var options = $scope.menu[i];

					for (var x in options) {
						var option = options[x];

						if (option.link === currentPath) {
							return option;
						}
					}
				}

				return null;
			}

			function generateTopMenu() {
				$scope.topMenu = $scope.menu[0].concat($scope.menu[1]);
			}

			function onLoad() {
				var option = determineActiveOption();
				setActiveOption(option);

				generateTopMenu();
			}

			onLoad();
		}
	]);
})(angular);
