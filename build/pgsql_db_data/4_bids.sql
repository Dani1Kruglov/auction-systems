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
-- Name: bids; Type: TABLE; Schema: public; Owner: myuser
--

CREATE TABLE public.bids (
                             id SERIAL PRIMARY KEY,
                             user_id integer REFERENCES public.users(id) ON DELETE CASCADE, -- Связь с пользователем (покупателем)
                             auction_id integer REFERENCES public.auctions(id) ON DELETE CASCADE, -- Связь с аукционом
                             bid_amount numeric(12, 2) NOT NULL, -- Сумма ставки
                             created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP,
                             updated_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE public.bids OWNER TO myuser;

--
-- PostgreSQL database dump complete
--

