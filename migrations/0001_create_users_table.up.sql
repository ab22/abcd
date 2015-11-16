CREATE TABLE users
(
  id serial NOT NULL,
  password character varying(255),
  email character varying(60),
  first_name character varying(60),
  last_name character varying(60),
  status integer,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  username character varying(30),
  role_id integer,
  CONSTRAINT users_pkey PRIMARY KEY (id),
  CONSTRAINT users_username_lowercase_ck CHECK (username::text = lower(username::text))
)
WITH (
  OIDS=FALSE
);
ALTER TABLE users
  OWNER TO postgres;

CREATE UNIQUE INDEX uix_users_email
  ON users
  USING btree
  (email COLLATE pg_catalog."default");

