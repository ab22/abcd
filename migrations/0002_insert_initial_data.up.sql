-- Create the admin user
INSERT INTO users(
	username,
	password,
	email,
	first_name,
	last_name,
	status,
	is_admin,
	is_teacher,
	created_at,
	updated_at
) VALUES (
	'admin',
	'$2a$10$vWQ6Tsu8odmvJ74.dlSZT.XonFRKuqxd/bZzxHP041FOWTzhKj552',
	'admin@abcd.com',
	'Administrador',
	'Administrador',
	0,
	true,
	true,
	timezone('UTC', now()),
	timezone('UTC', now())
)
