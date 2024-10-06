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
-- Name: lots; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.lots (
                             id SERIAL PRIMARY KEY,
                             seller_id integer REFERENCES public.users(id) ON DELETE CASCADE, -- Связь с пользователем (продавцом)
                             title character varying(255) NOT NULL, -- Название лота
                             description text, -- Описание лота
                             starting_price numeric(12, 2) NOT NULL, -- Начальная цена
                             is_closed boolean DEFAULT FALSE,
                             created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                             updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE public.lots OWNER TO myuser;

--
-- PostgreSQL database dump complete
--

