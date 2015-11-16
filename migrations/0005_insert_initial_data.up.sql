-- Create initial privileges
INSERT INTO privileges(
	key,
	description
) VALUES (
	'USERS_MODULE',
	'Acceso a editar, crear y listar usuarios.'
);

INSERT INTO privileges(
	key,
	description
) VALUES (
	'ROLES_MODULE',
	'Acceso a editar, crear y listar usuarios.'
);

-- Create initial roles and insert all
-- privileges into admin role.
WITH admin_role AS (
	INSERT INTO roles(
		name
	) VALUES (
		'Administrador'
	) RETURNING id
)
INSERT INTO roles_privileges(role_id, privilege_key)
	SELECT id as role_id, key AS privilege_key
	FROM privileges
	CROSS JOIN
	admin_role;

INSERT INTO roles(
	name
) VALUES (
	'Docente'
);

-- Create the admin user
INSERT INTO users(
	username,
	password,
	email,
	first_name,
	last_name,
	status,
	role_id,
	created_at,
	updated_at
) VALUES (
	'admin',
	'$2a$10$vWQ6Tsu8odmvJ74.dlSZT.XonFRKuqxd/bZzxHP041FOWTzhKj552',
	'admin@abcd.com',
	'Administrador',
	'Administrador',
	0,
	1,
	timezone('UTC', now()),
	timezone('UTC', now())
)
