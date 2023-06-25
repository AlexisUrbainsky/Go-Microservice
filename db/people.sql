-- Table: public.people

-- DROP TABLE IF EXISTS public.people;

CREATE TABLE IF NOT EXISTS public.people
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    name text COLLATE pg_catalog."default" NOT NULL,
    phone bigint NOT NULL,
    birthday date NOT NULL,
    CONSTRAINT people_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.people
    OWNER to postgres;



INSERT INTO public.people (name, phone, birthday) OVERRIDING SYSTEM VALUE VALUES ( 'Alex', 123456781, '1987-02-18');
INSERT INTO public.people (name, phone, birthday) OVERRIDING SYSTEM VALUE VALUES ( 'Micaela', 657849231, '1993-06-13');
INSERT INTO public.people (name, phone, birthday) OVERRIDING SYSTEM VALUE VALUES ( 'Omar', 548741236, '1981-09-22');
