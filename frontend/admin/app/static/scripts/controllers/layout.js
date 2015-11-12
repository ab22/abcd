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
			$scope.menu = [
				[
					{
						label: 'Inicio',
						icon: 'fa-home',
						link: '/main/home'
					},
					{
						label: 'Asignaturas',
						icon: 'fa-book',
						link: ''
					},
					{
						label: 'Alumnos',
						icon: 'fa-child',
						link: ''
					},
					{
						label: 'Docentes',
						icon: 'fa-male',
						link: ''
					},
					{
						label: 'Ingreso de Notas',
						icon: 'fa-file-text-o',
						link: ''
					}
				],
				[
					{
						label: 'Boletines',
						icon: 'fa-bullhorn',
						link: ''
					},
					{
						label: 'Reporte #1',
						icon: 'fa-calendar',
						link: ''
					},
					{
						label: 'Reporte #2',
						icon: 'fa-edit',
						link: ''
					},
					{
						label: 'Reporte #3',
						icon: 'fa-area-chart',
						link: ''
					}
				],
				[
					{
						label: 'Usuarios',
						icon: 'fa-users',
						link: '/main/users/all'
					},
					{
						label: 'Roles',
						icon: 'fa-legal',
						link: ''
					},
					{
						label: 'Mi Perfil',
						icon: 'fa-user',
						link: ''
					},
					{
						label: 'Configuración',
						icon: 'fa-cogs',
						link: ''
					},
					{
						label: 'Cerrar Sesión',
						icon: 'fa-sign-out',
						link: '',
						onClick: $scope.signOut
					}
				]
			];

			$scope.topMenu = [
				$scope.menu[0][0],
				$scope.menu[2][2],
				$scope.menu[2][3],
				$scope.menu[2][4]
			];

			$scope.optionOnClick = function(option) {
				if (typeof option.onClick !== 'undefined') {
					option.onClick();
					return;
				}

				setActiveOption(option);
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

			function onLoad() {
				var option = determineActiveOption();
				setActiveOption(option);
			}

			onLoad();
		}
	]);
})(angular);
