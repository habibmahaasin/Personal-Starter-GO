--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 14.6

-- Started on 2023-06-17 19:52:31 WIB

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP DATABASE railway;
--
-- TOC entry 3350 (class 1262 OID 17471)
-- Name: railway; Type: DATABASE; Schema: -; Owner: postgres
--

CREATE DATABASE railway WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'en_US.utf8';


ALTER DATABASE railway OWNER TO postgres;

\connect railway

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- TOC entry 4 (class 2615 OID 2200)
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA public;


ALTER SCHEMA public OWNER TO postgres;

--
-- TOC entry 3351 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 247 (class 1259 OID 17474)
-- Name: device_history; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_history (
    history_id integer NOT NULL,
    status_id integer,
    mode_id integer,
    device_id uuid,
    temperature real,
    ph real,
    dissolved_oxygen real,
    history_date timestamp without time zone
);


ALTER TABLE public.device_history OWNER TO postgres;

--
-- TOC entry 248 (class 1259 OID 17477)
-- Name: device_historya_history_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

ALTER TABLE public.device_history ALTER COLUMN history_id ADD GENERATED BY DEFAULT AS IDENTITY (
    SEQUENCE NAME public.device_historya_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- TOC entry 249 (class 1259 OID 17479)
-- Name: device_mode; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_mode (
    mode_id integer NOT NULL,
    mode_name character varying,
    date_created timestamp without time zone
);


ALTER TABLE public.device_mode OWNER TO postgres;

--
-- TOC entry 250 (class 1259 OID 17485)
-- Name: device_status; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.device_status (
    status_id integer NOT NULL,
    status_name character varying,
    date_created timestamp without time zone
);


ALTER TABLE public.device_status OWNER TO postgres;

--
-- TOC entry 251 (class 1259 OID 17491)
-- Name: devices; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.devices (
    device_id uuid NOT NULL,
    antares_id character varying,
    device_name character varying,
    device_location character varying,
    status_id integer,
    brand character varying,
    user_id uuid,
    latitude character varying,
    longitude character varying,
    date_created timestamp without time zone,
    date_updated timestamp without time zone,
    mode_id integer
);


ALTER TABLE public.devices OWNER TO postgres;

--
-- TOC entry 252 (class 1259 OID 17497)
-- Name: roles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.roles (
    role_id integer NOT NULL,
    role_name character varying,
    date_created character varying
);


ALTER TABLE public.roles OWNER TO postgres;

--
-- TOC entry 253 (class 1259 OID 17503)
-- Name: users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.users (
    user_id uuid NOT NULL,
    role_id integer,
    full_name character varying,
    email character varying,
    password character varying,
    address character varying,
    avatar character varying,
    date_created timestamp without time zone,
    date_updated timestamp without time zone
);


ALTER TABLE public.users OWNER TO postgres;

--
-- TOC entry 3338 (class 0 OID 17474)
-- Dependencies: 247
-- Data for Name: device_history; Type: TABLE DATA; Schema: public; Owner: postgres
--



--
-- TOC entry 3340 (class 0 OID 17479)
-- Dependencies: 249
-- Data for Name: device_mode; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device_mode VALUES (1, 'Otomatis', '2023-06-01 03:07:18.071393');
INSERT INTO public.device_mode VALUES (2, 'Manual', '2023-06-01 03:07:27.461798');


--
-- TOC entry 3341 (class 0 OID 17485)
-- Dependencies: 250
-- Data for Name: device_status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device_status VALUES (11, 'Aktif', '2023-03-30 06:14:05.677865');
INSERT INTO public.device_status VALUES (10, 'Tidak Aktif', '2023-03-30 06:14:21.438698');


--
-- TOC entry 3342 (class 0 OID 17491)
-- Dependencies: 251
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.devices VALUES ('cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 'UYBOPZUzuQIOAL2Z', 'Dummy Perangkat 2', 'Depan Rumah Kanan Atas, Depan Rumah Kanan Atas Depan Rumah Kanan Atas Depan Rumah Kanan Atas Depan Rumah Kanan Atas', 10, 'Amara 2 Lubang', 'a962321c-6b3a-4b92-8a70-9729a1f15b75', '-7.9742665249933244', '110.22720336500485', '2023-06-16 23:22:30.825766', '2023-06-17 00:41:51.046816', 1);
INSERT INTO public.devices VALUES ('cc7464ba-eef6-4883-96d0-af882252b666', '3764736527362u3', 'GuppyCiganitri', 'Ciganitri Land', 11, 'Amara 2 Lubang', 'a962321c-6b3a-4b92-8a70-9729a1f15b75', '-6.974875218957293', '107.65222215777261', '2023-06-17 18:04:34.9146', '2023-06-17 18:05:16.885912', 2);
INSERT INTO public.devices VALUES ('e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 'ps9t5UiX15TVLxYB', 'Aerator Utama', 'Aquarium Kiri Hitam', 10, 'Amara', 'a962321c-6b3a-4b92-8a70-9729a1f15b75', NULL, NULL, '2023-03-30 06:15:57.118381', '2023-06-17 19:49:02.788773', 1);
INSERT INTO public.devices VALUES ('9e88d893-d546-4eb1-8065-83a22d28d50d', '', 'Dummy no Antares', 'Aquarium Kanan Putih', 11, 'Amara', 'a962321c-6b3a-4b92-8a70-9729a1f15b75', NULL, NULL, '2023-06-01 06:10:34.319206', '2023-06-16 23:42:06.85617', 2);


--
-- TOC entry 3343 (class 0 OID 17497)
-- Dependencies: 252
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.roles VALUES (1, 'Super Admin', 'now()');


--
-- TOC entry 3344 (class 0 OID 17503)
-- Dependencies: 253
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('a962321c-6b3a-4b92-8a70-9729a1f15b75', 1, 'GuppyTech Admin', 'admin@guppytech.id', '$2a$12$xOeTPV2cIAcadxrGrkuxYemySlKNYoVjvtcAxvL1IEqY5Jk.XETb6
', 'Bandung', NULL, '2023-03-30 05:46:05.620484', '2023-03-30 05:46:05.620484');


--
-- TOC entry 3352 (class 0 OID 0)
-- Dependencies: 248
-- Name: device_historya_history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.device_historya_history_id_seq', 1083, true);


--
-- TOC entry 3180 (class 2606 OID 17510)
-- Name: device_history device_historya_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT device_historya_pkey PRIMARY KEY (history_id);


--
-- TOC entry 3182 (class 2606 OID 17512)
-- Name: device_mode device_mode_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_mode
    ADD CONSTRAINT device_mode_pk PRIMARY KEY (mode_id);


--
-- TOC entry 3184 (class 2606 OID 17514)
-- Name: device_status device_status_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_status
    ADD CONSTRAINT device_status_pk PRIMARY KEY (status_id);


--
-- TOC entry 3186 (class 2606 OID 17516)
-- Name: devices devices_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pk PRIMARY KEY (device_id);


--
-- TOC entry 3190 (class 2606 OID 17518)
-- Name: users newtable_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT newtable_pk PRIMARY KEY (user_id);


--
-- TOC entry 3188 (class 2606 OID 17520)
-- Name: roles role_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT role_pk PRIMARY KEY (role_id);


--
-- TOC entry 3194 (class 2606 OID 17521)
-- Name: devices devices_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- TOC entry 3191 (class 2606 OID 17526)
-- Name: device_history devices_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT devices_fk FOREIGN KEY (device_id) REFERENCES public.devices(device_id);


--
-- TOC entry 3195 (class 2606 OID 17531)
-- Name: devices mode_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT mode_fk FOREIGN KEY (mode_id) REFERENCES public.device_mode(mode_id);


--
-- TOC entry 3192 (class 2606 OID 17536)
-- Name: device_history mode_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT mode_fk FOREIGN KEY (mode_id) REFERENCES public.device_mode(mode_id);


--
-- TOC entry 3196 (class 2606 OID 17541)
-- Name: devices status_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.device_status(status_id);


--
-- TOC entry 3193 (class 2606 OID 17546)
-- Name: device_history status_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.device_status(status_id);


--
-- TOC entry 3197 (class 2606 OID 17551)
-- Name: users users_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.roles(role_id);


-- Completed on 2023-06-17 19:52:55 WIB

--
-- PostgreSQL database dump complete
--
