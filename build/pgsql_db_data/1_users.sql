--
-- PostgreSQL database dump
--

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: users; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.users (
                              id SERIAL PRIMARY KEY,
                              name character varying(255),
                              role character varying(255),
                              email character varying(255),
                              password character varying(255),
                              notification_url character varying(255),
                              created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                              updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE public.users OWNER TO myuser;

--
-- PostgreSQL database dump complete
--

