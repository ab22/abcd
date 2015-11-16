CREATE TABLE roles_privileges
(
  role_id integer NOT NULL,
  privilege_key character varying(255) NOT NULL,
  CONSTRAINT roles_privileges_pkey PRIMARY KEY (role_id, privilege_key)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE roles_privileges
  OWNER TO postgres;
