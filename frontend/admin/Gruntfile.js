'use strict';

module.exports = function (grunt) {
	var config = {};

	grunt.loadNpmTasks('grunt-contrib-clean');
	grunt.loadNpmTasks('grunt-contrib-concat');
	grunt.loadNpmTasks('grunt-contrib-copy');
	grunt.loadNpmTasks('grunt-contrib-cssmin');
	grunt.loadNpmTasks('grunt-contrib-htmlmin');
	grunt.loadNpmTasks('grunt-contrib-jshint');
	grunt.loadNpmTasks('grunt-contrib-uglify');
	grunt.loadNpmTasks('grunt-filerev');
	grunt.loadNpmTasks('grunt-jscs');
	grunt.loadNpmTasks('grunt-text-replace');
	grunt.loadNpmTasks('grunt-usemin');
	grunt.loadNpmTasks('grunt-wiredep');

	config.globals = {
		appPath: 'app/',
		staticFilesPath: 'app/static/',
		dist: 'dist'
	};

	config.clean = {
		dist: {
			files: [{
				dot: true,
				src: [
					'.tmp',
					'<%= globals.dist %>'
				]
			}]
		}
	};

	config.copy = {
		index: {
			expand: true,
			cwd: '<%= globals.appPath %>/',
			dest: '<%= globals.dist %>',
			src: [
				'index.html'
			]
		},
		views: {
			expand: true,
			cwd: '<%= globals.staticFilesPath %>/views/',
			dest: '<%= globals.dist %>/static/views/',
			src: '**/*.html'
		},
		images: {
			expand: true,
			cwd: '<%= globals.staticFilesPath %>/images/',
			dest: '<%= globals.dist %>/static/images/',
			src: '**'
		},
		fonts: {
			expand: true,
			cwd: '<%= globals.staticFilesPath %>/fonts/',
			dest: '<%= globals.dist %>/static/fonts/',
			src: '**'
		}
	};

	config.jshint = {
		options: {
			jshintrc: '.jshintrc'
		},
		all: [
			'<%= globals.staticFilesPath %>/scripts/**/*.js'
		]
	};

	config.jscs = {
		src: '<%= globals.staticFilesPath %>/scripts/**/*.js',
		options: {
			config: '.jscsrc'
		}
	};

	config.filerev = {
		dist: {
			src: [
		 		'<%= globals.dist %>/static/scripts/**/*.js',
				'<%= globals.dist %>/static/styles/{,*/}*.css',
				'<%= globals.dist %>/static/images/{,*/}*.{png,jpg,jpeg,gif,webp,svg}',
				'<%= globals.dist %>/static/styles/fonts/*'
			]
		}
	};

	config.wiredep = {
		task: {
			src: [
				'<%= globals.appPath %>/index.html'
			],
			ignorePath: /\.\.\//
		}
	};

	config.useminPrepare = {
		html: '<%= globals.appPath %>/index.html',
		options: {
			dest: '<%= globals.dist %>/',
			flow: {
				html: {
					steps: {
						js: ['concat', 'uglify'],
						css: ['concat', 'cssmin']
					},
					post: {}
				}
			}
		}
	};

	config.usemin = {
		html: ['<%= globals.dist %>/index.html']
	};

	config.htmlmin = {
		dist: {
			options: {
				collapseWhitespace: true,
				conservativeCollapse: true,
				collapseBooleanAttributes: true,
				removeCommentsFromCDATA: true,
				removeOptionalTags: false
			},
			files: [{
				expand: true,
				cwd: '<%= globals.dist %>',
				src: [
					'*.html',
					'static/views/**/*.html'
				],
				dest: '<%= globals.dist %>'
			}]
		}
	};

	grunt.initConfig(config);

	grunt.registerTask('test', [
		'jshint',
		'jscs',
	]);

	grunt.registerTask('build', [
		'test',
		'clean:dist',
		'wiredep',
		'useminPrepare',
		'concat',
		'cssmin',
		'uglify',
		'filerev',
		'copy',
		'usemin',
		'htmlmin'
	]);

	grunt.registerTask('default', [
		'test'
	]);
};
