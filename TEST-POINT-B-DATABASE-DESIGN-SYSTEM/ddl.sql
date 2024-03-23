-- public.users definition

-- Drop table

-- DROP TABLE public.users;

CREATE TABLE public.users (
	id uuid NOT NULL,
	"name" varchar NULL,
	email varchar NULL,
	address text NULL,
	telp varchar NULL,
	"role" varchar NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	"password" text NULL
);

-- public.film definition

-- Drop table

-- DROP TABLE public.film;

CREATE TABLE public.film (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	bioskop_id uuid NULL,
	"name" varchar NULL,
	show_date date NULL,
	show_time varchar NULL,
	price int8 NULL DEFAULT 0,
	status varchar NULL,
	url_trailer text NULL,
	url_image text NULL
);

-- public.bioskop definition

-- Drop table

-- DROP TABLE public.bioskop;

CREATE TABLE public.bioskop (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	"name" varchar NULL,
	address text NULL,
	latitude varchar NULL,
	longitude varchar NULL
);

-- public.studio definition

-- Drop table

-- DROP TABLE public.studio;

CREATE TABLE public.studio (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	bioskop_id uuid NULL,
	"name" varchar NULL,
	no_chair int4 NULL
);

-- public.transactions definition

-- Drop table

-- DROP TABLE public.transactions;

CREATE TABLE public.transactions (
	id uuid NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	user_id uuid NULL,
	film_id uuid NULL,
	studio_id uuid NULL,
	status varchar NULL,
	price int4 NULL,
	total int4 NULL
);