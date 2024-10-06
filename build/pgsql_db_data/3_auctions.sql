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
-- Name: auctions; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.auctions (
                                 id SERIAL PRIMARY KEY,
                                 lot_id integer UNIQUE REFERENCES public.lots(id) ON DELETE CASCADE, -- Связь с лотом (один лот - один аукцион)
                                 winner_id integer REFERENCES public.users(id) ON DELETE SET NULL, -- Победитель аукциона
                                 status character varying(50) NOT NULL, -- Статус аукциона (например, 'open', 'closed', 'wait')
                                 min_step numeric(12, 2) NOT NULL,
                                 start_time timestamp without time zone NOT NULL, -- Время начала аукциона
                                 end_time timestamp without time zone NOT NULL, -- Время окончания аукциона
                                 created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                                 updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE public.auctions OWNER TO myuser;

--
-- PostgreSQL database dump complete
--

