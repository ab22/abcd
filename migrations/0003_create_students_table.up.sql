CREATE TABLE students
(
  id serial NOT NULL,
  email character varying(60),
  first_name character varying(60),
  last_name character varying(60),
  status integer,
  created_at timestamp with time zone,
  updated_at timestamp with time zone,
  deleted_at timestamp with time zone,
  CONSTRAINT students_pkey PRIMARY KEY (id)
)
WITH (
  OIDS=FALSE
);
