CREATE TABLE privileges
(
  key character varying(255) NOT NULL,
  description character varying(200),
  CONSTRAINT privileges_pkey PRIMARY KEY (key)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE privileges
  OWNER TO postgres;
