;(function(angular) {
	'use strict';

	angular.module('app.controllers').controller('MainLayoutCtrl', ['$scope','$location', 'Auth',
		function($scope, $location, Auth) {
			$scope.window = {};

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
						responsiveOnly: false
					},
					{
						label: 'Asignaturas',
						icon: 'fa-book',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Alumnos',
						icon: 'fa-child',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Docentes',
						icon: 'fa-male',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Ingreso de Notas',
						icon: 'fa-file-text-o',
						link: '',
						responsiveOnly: true
					}
				],
				[
					{
						label: 'Boletines',
						icon: 'fa-bullhorn',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Reporte #1',
						icon: 'fa-calendar',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Reporte #2',
						icon: 'fa-edit',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Reporte #3',
						icon: 'fa-area-chart',
						link: '',
						responsiveOnly: true
					}
				],
				[
					{
						label: 'Usuarios',
						icon: 'fa-users',
						link: '/main/users/all',
						responsiveOnly: true
					},
					{
						label: 'Roles',
						icon: 'fa-legal',
						link: '',
						responsiveOnly: true
					},
					{
						label: 'Mi Perfil',
						icon: 'fa-user',
						link: '',
						responsiveOnly: false
					},
					{
						label: 'Configuración',
						icon: 'fa-cogs',
						link: '',
						responsiveOnly: false
					},
					{
						label: 'Cerrar Sesión',
						icon: 'fa-sign-out',
						link: '',
						onClick: $scope.signOut,
						responsiveOnly: false
					}
				]
			];

			$scope.optionOnClick = function(option) {
				if (typeof option.onClick !== 'undefined') {
					option.onClick();
					return;
				}

				setActiveOption(option);
			};

			$scope.isResponsiveMode = function() {
				return $scope.window.width <= 767;
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

			function setActiveOption(option) {
				$scope.activeOption = option;
			}

			function generateTopMenu() {
				$scope.topMenu = $scope.menu[0].concat($scope.menu[1].concat($scope.menu[2]));
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
