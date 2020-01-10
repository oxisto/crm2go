CREATE SCHEMA public;

ALTER SCHEMA public OWNER TO postgres;

CREATE TABLE public.contact (
    id SERIAL PRIMARY KEY,
    name character varying NOT NULL,
    company character varying,
    email character varying

);

ALTER TABLE public.contact OWNER TO postgres;
