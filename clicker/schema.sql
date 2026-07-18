--
-- PostgreSQL database dump
--

\restrict dOIUgkDMC2fekLymJbGrcHy3veRyWdgoGc9PITAfZ2DGtvERk0AXoknIMJpuOw0

-- Dumped from database version 17.10 (Debian 17.10-1.pgdg13+1)
-- Dumped by pg_dump version 17.10 (Debian 17.10-1.pgdg13+1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: player; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.player (
    id integer NOT NULL,
    login character varying NOT NULL,
    money bigint DEFAULT 0 NOT NULL,
    email text
);


ALTER TABLE public.player OWNER TO "user";

--
-- Name: player_id_seq; Type: SEQUENCE; Schema: public; Owner: user
--

CREATE SEQUENCE public.player_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.player_id_seq OWNER TO "user";

--
-- Name: player_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: user
--

ALTER SEQUENCE public.player_id_seq OWNED BY public.player.id;


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: user
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO "user";

--
-- Name: player id; Type: DEFAULT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.player ALTER COLUMN id SET DEFAULT nextval('public.player_id_seq'::regclass);


--
-- Name: player player_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.player
    ADD CONSTRAINT player_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: user
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--

\unrestrict dOIUgkDMC2fekLymJbGrcHy3veRyWdgoGc9PITAfZ2DGtvERk0AXoknIMJpuOw0

