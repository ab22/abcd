CREATE TABLE students
(
	id serial NOT NULL,
	id_number character varying(40) NOT NULL,
	email character varying(60),
	first_name character varying(60),
	last_name character varying(60),
	status integer,
	place_of_birth character varying(60),
	address character varying(100),
	birthdate date,
	gender character varying(3),
	nationality character varying(30),
	phone_number character varying(40),
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	deleted_at timestamp with time zone,
	CONSTRAINT students_pkey PRIMARY KEY (id)
)
WITH (
	OIDS=FALSE
);

CREATE UNIQUE INDEX id_number_unique_idx
	ON students
	USING btree
	(id_number COLLATE pg_catalog."default");
