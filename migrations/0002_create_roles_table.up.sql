CREATE TABLE roles
(
  id serial NOT NULL,
  name character varying(30),
  CONSTRAINT roles_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE roles
  OWNER TO postgres;
