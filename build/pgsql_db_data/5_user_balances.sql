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
-- Name: user_balances; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.user_balances (
                                      user_id INT REFERENCES public.users(id) ON DELETE CASCADE,  -- Связь с таблицей users
                                      balance NUMERIC(10, 2) NOT NULL DEFAULT 0,                   -- Баланс пользователя
                                      PRIMARY KEY (user_id)                                        -- user_id - уникальный ключ
);

ALTER TABLE public.user_balances OWNER TO myuser;

--
-- PostgreSQL database dump complete
--

