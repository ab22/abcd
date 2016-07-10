;(function(angular) {
	'use strict';

	angular.module('app.directives').directive('resizableWindow', ['$window',
		function($window) {
			return {
				restrict: 'A',
				scope: {
					window: '='
				},
				link: function(scope) {
					var window = angular.element($window);
					scope.window = {
						height: 0,
						width: 0
					};

					scope.getWindowDimensions = function() {
						return {
							height: $window.innerHeight,
							width: $window.innerWidth
						};
					};

					scope.$watch(scope.getWindowDimensions, function(newValue) {
						scope.window.height = newValue.height;
						scope.window.width = newValue.width;
					}, true);

					window.bind('resize', function() {
						scope.$apply();
					});
				}
			};
		}
	]);
})(angular);
