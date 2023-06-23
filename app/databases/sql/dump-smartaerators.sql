--
-- PostgreSQL database dump
--

-- Dumped from database version 13.2
-- Dumped by pg_dump version 14.6

-- Started on 2023-06-21 09:32:28 WIB

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
-- TOC entry 3359 (class 1262 OID 17471)
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
-- TOC entry 3360 (class 0 OID 0)
-- Dependencies: 4
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: postgres
--

COMMENT ON SCHEMA public IS 'standard public schema';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 254 (class 1259 OID 17556)
-- Name: brand; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.brand (
    brand_id integer NOT NULL,
    brand_name character varying,
    date_created timestamp without time zone
);


ALTER TABLE public.brand OWNER TO postgres;

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
    brand_id integer,
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
    date_created timestamp without time zone
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
-- TOC entry 3353 (class 0 OID 17556)
-- Dependencies: 254
-- Data for Name: brand; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.brand VALUES (1, 'Amara AA-222', '2023-06-21 00:20:39.811275');
INSERT INTO public.brand VALUES (2, 'Amara AA-333', '2023-06-21 00:20:39.811275');
INSERT INTO public.brand VALUES (3, 'Amara AA-666', '2023-06-21 00:20:39.811275');


--
-- TOC entry 3346 (class 0 OID 17474)
-- Dependencies: 247
-- Data for Name: device_history; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device_history VALUES (1084, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5625, 21.97006, 48.53256, '2023-06-17 19:53:38.426783');
INSERT INTO public.device_history VALUES (1085, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5625, 21.97006, 48.53256, '2023-06-17 19:53:45.22244');
INSERT INTO public.device_history VALUES (1086, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5, 21.97006, 48.47006, '2023-06-17 19:53:52.262134');
INSERT INTO public.device_history VALUES (1087, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5625, 21.97006, 48.53256, '2023-06-17 19:53:59.017916');
INSERT INTO public.device_history VALUES (1088, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.625, 21.97006, 48.59506, '2023-06-17 19:54:05.796881');
INSERT INTO public.device_history VALUES (1089, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5625, 21.97006, 48.53256, '2023-06-17 19:54:12.876442');
INSERT INTO public.device_history VALUES (1090, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5, 21.97006, 48.47006, '2023-06-17 19:54:19.922162');
INSERT INTO public.device_history VALUES (1091, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.5, 21.97006, 48.47006, '2023-06-17 19:54:26.680255');
INSERT INTO public.device_history VALUES (1092, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.5, 21.97006, 50.47006, '2023-06-17 19:54:33.698727');
INSERT INTO public.device_history VALUES (1093, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.5625, 21.97006, 51.53256, '2023-06-17 19:54:40.696168');
INSERT INTO public.device_history VALUES (1094, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.5625, 21.97006, 51.53256, '2023-06-17 19:54:47.683926');
INSERT INTO public.device_history VALUES (1095, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.5, 21.97006, 51.47006, '2023-06-17 19:54:54.475277');
INSERT INTO public.device_history VALUES (1096, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.375, 21.97006, 51.34506, '2023-06-17 19:55:01.391421');
INSERT INTO public.device_history VALUES (1097, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.25, 21.97006, 51.22006, '2023-06-17 19:55:08.456898');
INSERT INTO public.device_history VALUES (1098, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.125, 21.97006, 51.09506, '2023-06-17 19:55:15.516857');
INSERT INTO public.device_history VALUES (1099, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.0625, 21.97006, 51.03256, '2023-06-17 19:55:22.576988');
INSERT INTO public.device_history VALUES (1100, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.9375, 21.97006, 50.90756, '2023-06-17 19:55:29.757458');
INSERT INTO public.device_history VALUES (1296, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 0, 0, 0, '2023-06-21 02:37:50.516757');
INSERT INTO public.device_history VALUES (1297, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 0, 0, 0, '2023-06-21 02:37:52.146281');
INSERT INTO public.device_history VALUES (1103, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.8125, 21.97006, 50.78256, '2023-06-17 19:55:36.69623');
INSERT INTO public.device_history VALUES (1104, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.75, 21.97006, 50.72006, '2023-06-17 19:55:43.794486');
INSERT INTO public.device_history VALUES (1105, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.5, 21.97006, 50.47006, '2023-06-17 19:55:50.737666');
INSERT INTO public.device_history VALUES (1106, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.3125, 21.97006, 50.28256, '2023-06-17 19:55:57.796194');
INSERT INTO public.device_history VALUES (1207, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 0, 0, 0, '2023-06-18 20:17:14.20823');
INSERT INTO public.device_history VALUES (1208, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 0, 0, 0, '2023-06-18 20:17:15.856452');
INSERT INTO public.device_history VALUES (1251, 11, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 00:44:47.95779');
INSERT INTO public.device_history VALUES (1252, 11, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 00:44:49.669993');
INSERT INTO public.device_history VALUES (1257, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:14:24.569027');
INSERT INTO public.device_history VALUES (1258, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:14:26.174588');
INSERT INTO public.device_history VALUES (1259, 11, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:14:33.837779');
INSERT INTO public.device_history VALUES (1260, 11, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:14:35.864816');
INSERT INTO public.device_history VALUES (1261, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:16:23.028161');
INSERT INTO public.device_history VALUES (1262, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:16:24.99622');
INSERT INTO public.device_history VALUES (1263, 10, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:16:35.162415');
INSERT INTO public.device_history VALUES (1264, 10, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:16:37.177912');
INSERT INTO public.device_history VALUES (1269, 10, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 03:26:14.881818');
INSERT INTO public.device_history VALUES (1270, 10, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 03:26:16.466481');
INSERT INTO public.device_history VALUES (1271, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 03:29:31.398673');
INSERT INTO public.device_history VALUES (1272, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 03:29:33.37157');
INSERT INTO public.device_history VALUES (1273, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 31.49, 27.8, 45.39, '2023-06-19 03:36:15.631781');
INSERT INTO public.device_history VALUES (1107, 10, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.125, 21.97006, 50.09506, '2023-06-17 19:56:04.66359');
INSERT INTO public.device_history VALUES (1298, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 30.79, 27.56, 44.57, '2023-06-21 09:28:34.389833');
INSERT INTO public.device_history VALUES (1299, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 28.32, 34, 45.32, '2023-06-21 09:28:37.120337');
INSERT INTO public.device_history VALUES (1300, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 28.19, 24.84, 40.61, '2023-06-21 09:28:52.204506');
INSERT INTO public.device_history VALUES (1111, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.6875, 21.97006, 49.65756, '2023-06-17 19:56:37.840654');
INSERT INTO public.device_history VALUES (1301, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 25.06, 24.92, 37.52, '2023-06-21 09:28:54.945786');
INSERT INTO public.device_history VALUES (1302, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 32.74, 32.6, 49.04, '2023-06-21 09:28:57.68373');
INSERT INTO public.device_history VALUES (1303, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 31.33, 32.09, 47.375, '2023-06-21 09:29:00.432456');
INSERT INTO public.device_history VALUES (1304, 11, 1, '9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 26.43, 27.04, 39.95, '2023-06-21 09:29:03.174155');
INSERT INTO public.device_history VALUES (1116, 11, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.625, 21.97006, 49.59506, '2023-06-17 19:56:46.563093');
INSERT INTO public.device_history VALUES (1118, 11, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.6875, 21.97006, 49.65756, '2023-06-17 19:56:51.021126');
INSERT INTO public.device_history VALUES (1122, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.5625, 21.97006, 49.53256, '2023-06-17 19:56:58.03961');
INSERT INTO public.device_history VALUES (1123, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.625, 21.97006, 50.59506, '2023-06-17 19:57:05.079392');
INSERT INTO public.device_history VALUES (1124, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 30.5625, 21.97006, 52.53256, '2023-06-17 19:57:12.00176');
INSERT INTO public.device_history VALUES (1125, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 31.8125, 21.97006, 53.78256, '2023-06-17 19:57:18.882709');
INSERT INTO public.device_history VALUES (1126, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 32.875, 21.97006, 54.84506, '2023-06-17 19:57:25.498558');
INSERT INTO public.device_history VALUES (1127, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 33.375, 21.97006, 55.34506, '2023-06-17 19:57:32.518448');
INSERT INTO public.device_history VALUES (1128, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 33.6875, 21.97006, 55.65756, '2023-06-17 19:57:39.457526');
INSERT INTO public.device_history VALUES (1129, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 33.9375, 21.97006, 55.90756, '2023-06-17 19:57:46.316423');
INSERT INTO public.device_history VALUES (1130, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.0625, 21.97006, 56.03256, '2023-06-17 19:57:53.4486');
INSERT INTO public.device_history VALUES (1131, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.1875, 21.97006, 56.15756, '2023-06-17 19:58:00.639911');
INSERT INTO public.device_history VALUES (1132, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.3125, 21.97006, 56.28256, '2023-06-17 19:58:07.540374');
INSERT INTO public.device_history VALUES (1133, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.4375, 21.97006, 56.40756, '2023-06-17 19:58:14.417867');
INSERT INTO public.device_history VALUES (1134, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 31.1875, 21.97006, 53.15756, '2023-06-17 19:58:21.077675');
INSERT INTO public.device_history VALUES (1135, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.0625, 21.97006, 49.03256, '2023-06-17 19:58:27.917238');
INSERT INTO public.device_history VALUES (1136, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 25.75, 21.97006, 47.72006, '2023-06-17 19:58:34.804273');
INSERT INTO public.device_history VALUES (1137, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.9375, 21.97006, 48.90756, '2023-06-17 19:58:41.640598');
INSERT INTO public.device_history VALUES (1138, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 30.5, 21.97006, 52.47006, '2023-06-17 19:58:48.424855');
INSERT INTO public.device_history VALUES (1139, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 31.9375, 21.97006, 53.90756, '2023-06-17 19:58:55.219808');
INSERT INTO public.device_history VALUES (1140, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 32.6875, 21.97006, 54.65756, '2023-06-17 19:59:02.320314');
INSERT INTO public.device_history VALUES (1141, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 31.875, 21.97006, 53.84506, '2023-06-17 19:59:09.31912');
INSERT INTO public.device_history VALUES (1142, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.6875, 21.97006, 51.65756, '2023-06-17 19:59:16.440741');
INSERT INTO public.device_history VALUES (1143, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.875, 21.97006, 49.84506, '2023-06-17 19:59:21.500127');
INSERT INTO public.device_history VALUES (1144, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 26.1875, 21.97006, 48.15756, '2023-06-17 19:59:28.27337');
INSERT INTO public.device_history VALUES (1145, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 25, 21.97006, 46.97006, '2023-06-17 19:59:35.080996');
INSERT INTO public.device_history VALUES (1146, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.875, 21.97006, 45.84506, '2023-06-17 19:59:41.922029');
INSERT INTO public.device_history VALUES (1147, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.125, 21.97006, 44.09506, '2023-06-17 19:59:48.882321');
INSERT INTO public.device_history VALUES (1148, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 20.3125, 21.97006, 42.28256, '2023-06-17 19:59:55.901707');
INSERT INTO public.device_history VALUES (1149, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 19.3125, 21.97006, 41.28256, '2023-06-17 20:00:02.601685');
INSERT INTO public.device_history VALUES (1150, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 18.5625, 21.97006, 40.53256, '2023-06-17 20:00:09.507851');
INSERT INTO public.device_history VALUES (1151, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 17.9375, 21.97006, 39.90756, '2023-06-17 20:00:18.400211');
INSERT INTO public.device_history VALUES (1152, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 17.3125, 21.97006, 39.28256, '2023-06-17 20:00:25.140502');
INSERT INTO public.device_history VALUES (1153, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 16.875, 21.97006, 38.84506, '2023-06-17 20:00:32.568369');
INSERT INTO public.device_history VALUES (1154, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 17.5, 21.97006, 39.47006, '2023-06-17 20:00:39.487196');
INSERT INTO public.device_history VALUES (1155, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 20.625, 21.97006, 42.59506, '2023-06-17 20:00:46.180391');
INSERT INTO public.device_history VALUES (1156, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 21.1875, 21.97006, 43.15756, '2023-06-17 20:00:53.079023');
INSERT INTO public.device_history VALUES (1157, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 21.5625, 21.97006, 43.53256, '2023-06-17 20:00:59.899255');
INSERT INTO public.device_history VALUES (1158, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 21.75, 21.97006, 43.72006, '2023-06-17 20:01:06.627176');
INSERT INTO public.device_history VALUES (1159, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22, 21.97006, 43.97006, '2023-06-17 20:01:13.239144');
INSERT INTO public.device_history VALUES (1160, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.1875, 21.97006, 44.15756, '2023-06-17 20:01:20.262285');
INSERT INTO public.device_history VALUES (1161, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.3125, 21.97006, 44.28256, '2023-06-17 20:01:27.028306');
INSERT INTO public.device_history VALUES (1162, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.5, 21.97006, 44.47006, '2023-06-17 20:01:34.064206');
INSERT INTO public.device_history VALUES (1163, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.625, 21.97006, 44.59506, '2023-06-17 20:01:40.701716');
INSERT INTO public.device_history VALUES (1164, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.75, 21.97006, 44.72006, '2023-06-17 20:01:47.738672');
INSERT INTO public.device_history VALUES (1165, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 22.9375, 21.97006, 44.90756, '2023-06-17 20:01:54.541157');
INSERT INTO public.device_history VALUES (1166, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.0625, 21.97006, 45.03256, '2023-06-17 20:02:01.362191');
INSERT INTO public.device_history VALUES (1211, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.94, 32.98, 41.43, '2023-06-18 20:24:49.716108');
INSERT INTO public.device_history VALUES (1212, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 35.08, 33.46, 51.81, '2023-06-18 20:24:52.623272');
INSERT INTO public.device_history VALUES (1213, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.94, 31.75, 50.815, '2023-06-18 20:24:55.529164');
INSERT INTO public.device_history VALUES (1214, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.32, 31.58, 45.11, '2023-06-18 20:24:58.43784');
INSERT INTO public.device_history VALUES (1215, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.8, 26.62, 42.11, '2023-06-18 20:25:01.345149');
INSERT INTO public.device_history VALUES (1216, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 25.8, 33.65, 42.625, '2023-06-18 20:25:04.397085');
INSERT INTO public.device_history VALUES (1217, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.9, 27.84, 42.82, '2023-06-18 20:25:07.309568');
INSERT INTO public.device_history VALUES (1218, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 32.53, 34.29, 49.675, '2023-06-18 20:25:10.223134');
INSERT INTO public.device_history VALUES (1219, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 29.96, 34.56, 47.24, '2023-06-18 20:25:13.130819');
INSERT INTO public.device_history VALUES (1220, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.26, 28.09, 42.305, '2023-06-18 20:25:16.050489');
INSERT INTO public.device_history VALUES (1221, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.36, 27.25, 47.985, '2023-06-18 20:25:18.9613');
INSERT INTO public.device_history VALUES (1222, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 25.52, 34.12, 42.58, '2023-06-18 20:25:21.865811');
INSERT INTO public.device_history VALUES (1223, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 28.84, 25.9, 41.79, '2023-06-18 20:25:24.781335');
INSERT INTO public.device_history VALUES (1224, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 27.09, 24.83, 39.505, '2023-06-18 20:25:27.675731');
INSERT INTO public.device_history VALUES (1225, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.3, 31.43, 50.015, '2023-06-18 20:25:30.590075');
INSERT INTO public.device_history VALUES (1226, 11, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 34.82, 27.75, 48.695, '2023-06-18 20:25:33.492957');
INSERT INTO public.device_history VALUES (1227, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 34.23, 28.44, 48.45, '2023-06-18 20:26:25.695883');
INSERT INTO public.device_history VALUES (1228, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.7, 31.98, 46.69, '2023-06-18 20:26:28.609523');
INSERT INTO public.device_history VALUES (1229, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 29.44, 26.05, 42.465, '2023-06-18 20:26:31.526057');
INSERT INTO public.device_history VALUES (1230, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 32.83, 31.02, 48.34, '2023-06-18 20:26:34.450788');
INSERT INTO public.device_history VALUES (1231, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 27.12, 32.09, 43.165, '2023-06-18 20:26:37.371515');
INSERT INTO public.device_history VALUES (1232, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 26.16, 34.85, 43.585, '2023-06-18 20:26:40.289535');
INSERT INTO public.device_history VALUES (1233, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.09, 35.18, 47.68, '2023-06-18 20:26:43.208746');
INSERT INTO public.device_history VALUES (1234, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.46, 33.29, 47.105, '2023-06-18 20:26:46.119962');
INSERT INTO public.device_history VALUES (1235, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 27.57, 32.47, 43.805, '2023-06-18 20:26:49.038986');
INSERT INTO public.device_history VALUES (1236, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 28.89, 28.37, 43.075, '2023-06-18 20:26:51.954548');
INSERT INTO public.device_history VALUES (1237, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 33.89, 31.91, 49.845, '2023-06-18 20:26:54.863313');
INSERT INTO public.device_history VALUES (1238, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 24.57, 31.87, 40.505, '2023-06-18 20:26:57.96046');
INSERT INTO public.device_history VALUES (1239, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.88, 24.56, 43.16, '2023-06-18 20:27:00.879904');
INSERT INTO public.device_history VALUES (1240, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.35, 29.53, 45.115, '2023-06-18 20:27:03.787624');
INSERT INTO public.device_history VALUES (1241, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 31.98, 25.03, 44.495, '2023-06-18 20:27:06.697242');
INSERT INTO public.device_history VALUES (1242, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 33.11, 29.39, 47.805, '2023-06-18 20:27:09.609365');
INSERT INTO public.device_history VALUES (1243, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 26.43, 28.18, 40.52, '2023-06-18 20:27:12.52374');
INSERT INTO public.device_history VALUES (1244, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 32.62, 31.95, 48.595, '2023-06-18 20:27:15.434798');
INSERT INTO public.device_history VALUES (1245, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 35.08, 34.94, 52.55, '2023-06-18 20:27:18.3451');
INSERT INTO public.device_history VALUES (1246, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.15, 35.12, 47.71, '2023-06-18 20:27:21.256692');
INSERT INTO public.device_history VALUES (1253, 11, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 01:15:34.647601');
INSERT INTO public.device_history VALUES (1254, 11, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 01:15:36.23818');
INSERT INTO public.device_history VALUES (1265, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:57:59.585704');
INSERT INTO public.device_history VALUES (1266, 10, 2, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:58:01.176148');
INSERT INTO public.device_history VALUES (1274, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 29.29, 27.87, 43.225, '2023-06-19 03:36:45.587799');
INSERT INTO public.device_history VALUES (1275, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 29.41, 26.95, 42.885, '2023-06-19 03:36:48.187416');
INSERT INTO public.device_history VALUES (1276, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 28.59, 31.39, 44.285, '2023-06-19 03:36:51.152117');
INSERT INTO public.device_history VALUES (1277, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 28.85, 34.44, 46.07, '2023-06-19 03:36:54.109984');
INSERT INTO public.device_history VALUES (1278, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 29.25, 25.39, 41.945, '2023-06-19 03:36:57.070801');
INSERT INTO public.device_history VALUES (1167, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.1875, 21.97006, 45.15756, '2023-06-17 20:02:08.099058');
INSERT INTO public.device_history VALUES (1168, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.25, 21.97006, 45.22006, '2023-06-17 20:02:14.96768');
INSERT INTO public.device_history VALUES (1169, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.4375, 21.97006, 45.40756, '2023-06-17 20:02:22.07019');
INSERT INTO public.device_history VALUES (1170, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.5, 21.97006, 45.47006, '2023-06-17 20:02:28.860056');
INSERT INTO public.device_history VALUES (1171, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.625, 21.97006, 45.59506, '2023-06-17 20:02:35.955888');
INSERT INTO public.device_history VALUES (1172, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.6875, 21.97006, 45.65756, '2023-06-17 20:02:42.601128');
INSERT INTO public.device_history VALUES (1173, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.75, 21.97006, 45.72006, '2023-06-17 20:02:49.539521');
INSERT INTO public.device_history VALUES (1174, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.875, 21.97006, 45.84506, '2023-06-17 20:02:56.360134');
INSERT INTO public.device_history VALUES (1175, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 23.9375, 21.97006, 45.90756, '2023-06-17 20:03:03.099652');
INSERT INTO public.device_history VALUES (1176, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24, 21.97006, 45.97006, '2023-06-17 20:03:09.800398');
INSERT INTO public.device_history VALUES (1177, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.0625, 21.97006, 46.03256, '2023-06-17 20:03:16.583152');
INSERT INTO public.device_history VALUES (1178, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.125, 21.97006, 46.09506, '2023-06-17 20:03:23.239389');
INSERT INTO public.device_history VALUES (1179, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.1875, 21.97006, 46.15756, '2023-06-17 20:03:30.068099');
INSERT INTO public.device_history VALUES (1180, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.25, 21.97006, 46.22006, '2023-06-17 20:03:36.81955');
INSERT INTO public.device_history VALUES (1181, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.3125, 21.97006, 46.28256, '2023-06-17 20:03:43.870188');
INSERT INTO public.device_history VALUES (1182, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.375, 21.97006, 46.34506, '2023-06-17 20:03:50.648927');
INSERT INTO public.device_history VALUES (1183, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.375, 21.97006, 46.34506, '2023-06-17 20:03:57.540119');
INSERT INTO public.device_history VALUES (1184, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.4375, 21.97006, 46.40756, '2023-06-17 20:04:04.319208');
INSERT INTO public.device_history VALUES (1185, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.5, 21.97006, 46.47006, '2023-06-17 20:04:11.149017');
INSERT INTO public.device_history VALUES (1186, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.5625, 21.97006, 46.53256, '2023-06-17 20:04:17.960029');
INSERT INTO public.device_history VALUES (1187, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.625, 21.97006, 46.59506, '2023-06-17 20:04:24.779793');
INSERT INTO public.device_history VALUES (1188, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.625, 21.97006, 46.59506, '2023-06-17 20:04:31.585824');
INSERT INTO public.device_history VALUES (1189, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.5625, 21.97006, 46.53256, '2023-06-17 20:04:38.388916');
INSERT INTO public.device_history VALUES (1190, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.6875, 21.97006, 46.65756, '2023-06-17 20:04:45.367889');
INSERT INTO public.device_history VALUES (1191, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.6875, 21.97006, 46.65756, '2023-06-17 20:04:52.28323');
INSERT INTO public.device_history VALUES (1192, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.6875, 21.97006, 46.65756, '2023-06-17 20:04:59.121471');
INSERT INTO public.device_history VALUES (1193, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.75, 21.97006, 46.72006, '2023-06-17 20:05:05.700958');
INSERT INTO public.device_history VALUES (1194, 10, 1, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 24.75, 21.97006, 46.72006, '2023-06-17 20:05:12.936787');
INSERT INTO public.device_history VALUES (1247, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 31.2, 29.07, 45.735, '2023-06-19 00:07:37.273051');
INSERT INTO public.device_history VALUES (1248, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 25.41, 29.37, 40.095, '2023-06-19 00:07:39.76294');
INSERT INTO public.device_history VALUES (1249, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 31.78, 26.48, 45.02, '2023-06-19 00:08:04.071664');
INSERT INTO public.device_history VALUES (1250, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 31.18, 29.44, 45.9, '2023-06-19 00:08:07.259794');
INSERT INTO public.device_history VALUES (1255, 10, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:08:34.73315');
INSERT INTO public.device_history VALUES (1256, 10, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 0, 0, 0, '2023-06-19 02:08:36.312331');
INSERT INTO public.device_history VALUES (1267, 11, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 0, 0, 0, '2023-06-19 03:13:45.087822');
INSERT INTO public.device_history VALUES (1268, 11, 2, 'e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 0, 0, 0, '2023-06-19 03:13:46.521685');
INSERT INTO public.device_history VALUES (1279, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 30.8, 30.85, 46.225, '2023-06-19 03:40:54.808924');
INSERT INTO public.device_history VALUES (1280, 11, 1, 'cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 25.22, 33.15, 41.795, '2023-06-19 03:40:57.410426');


--
-- TOC entry 3348 (class 0 OID 17479)
-- Dependencies: 249
-- Data for Name: device_mode; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device_mode VALUES (1, 'Otomatis', '2023-06-01 03:07:18.071393');
INSERT INTO public.device_mode VALUES (2, 'Manual', '2023-06-01 03:07:27.461798');


--
-- TOC entry 3349 (class 0 OID 17485)
-- Dependencies: 250
-- Data for Name: device_status; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device_status VALUES (11, 'Aktif', '2023-03-30 06:14:05.677865');
INSERT INTO public.device_status VALUES (10, 'Tidak Aktif', '2023-03-30 06:14:21.438698');


--
-- TOC entry 3350 (class 0 OID 17491)
-- Dependencies: 251
-- Data for Name: devices; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.devices VALUES ('cc5aa4c7-112b-47d1-a07e-1df6d3c9a2db', 'UYBOPZUzuQIOAL2Z', 'Dummy Perangkat 2', 'Depan Rumah Kanan Atas, Depan Rumah Kanan Atas Depan Rumah Kanan Atas Depan Rumah Kanan Atas Depan Rumah Kanan Atas', 11, 3, 'a962321c-6b3a-4b92-8a70-9729a1f15b75', '-7.9742665249933244', '110.22720336500485', '2023-06-16 23:22:30.825766', '2023-06-19 09:44:22.203231', 2);
INSERT INTO public.devices VALUES ('0258dfd4-c21f-43dc-adc2-ef620d39098a', '00000000', 'aaa', 'disini', 10, 1, 'a962321c-6b3a-4b92-8a70-9729a1f15b75', '-7.974246499999996', '110.22721861999996', '2023-06-21 00:02:23.690128', '2023-06-21 02:34:22.396452', 2);
INSERT INTO public.devices VALUES ('e5d415f7-a96b-4dc2-84b8-64a1830b4c01', 'ps9t5UiX15TVLxYB', 'Aerator Utama', 'Aquarium Kiri Hitam', 10, 3, 'a962321c-6b3a-4b92-8a70-9729a1f15b75', '-6.974875218957293', '107.65222215777261', '2023-03-30 06:15:57.118381', '2023-06-21 02:37:52.148962', 2);
INSERT INTO public.devices VALUES ('9474e5c1-82eb-4576-95e9-4b3c4c2b981f', 'fWxWMJbgMGTOd0a9', 'Perangkat Public', 'Tempat umum 1', 11, 2, '03eb0947-8d83-416e-8ac5-4a595685cf79', '-7.974250074965296', '110.22722624992589', '2023-06-21 04:33:56.132465', '2023-06-21 09:29:03.175942', 1);


--
-- TOC entry 3351 (class 0 OID 17497)
-- Dependencies: 252
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.roles VALUES (1, 'Owner', '2023-06-21 00:17:16.504666');


--
-- TOC entry 3352 (class 0 OID 17503)
-- Dependencies: 253
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.users VALUES ('a962321c-6b3a-4b92-8a70-9729a1f15b75', 1, 'GuppyTech Admin', 'admin@guppytech.id', '$2a$12$xOeTPV2cIAcadxrGrkuxYemySlKNYoVjvtcAxvL1IEqY5Jk.XETb6
', 'Bandung', NULL, '2023-03-30 05:46:05.620484', '2023-03-30 05:46:05.620484');
INSERT INTO public.users VALUES ('03eb0947-8d83-416e-8ac5-4a595685cf79', 1, 'Public Demo', 'public@guppy.tech', '$2a$12$mA2cnP4HO.3sbEuaXXSU6uq9K0epA0ad6H3/QDXsGZaZGL/dsMrd.
', 'Bandung', NULL, '2023-06-21 04:00:23.518107', '2023-06-21 04:00:23.518107');


--
-- TOC entry 3361 (class 0 OID 0)
-- Dependencies: 248
-- Name: device_historya_history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.device_historya_history_id_seq', 1304, true);


--
-- TOC entry 3197 (class 2606 OID 17576)
-- Name: brand brand_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.brand
    ADD CONSTRAINT brand_pk PRIMARY KEY (brand_id);


--
-- TOC entry 3185 (class 2606 OID 17510)
-- Name: device_history device_historya_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT device_historya_pkey PRIMARY KEY (history_id);


--
-- TOC entry 3187 (class 2606 OID 17512)
-- Name: device_mode device_mode_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_mode
    ADD CONSTRAINT device_mode_pk PRIMARY KEY (mode_id);


--
-- TOC entry 3189 (class 2606 OID 17514)
-- Name: device_status device_status_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_status
    ADD CONSTRAINT device_status_pk PRIMARY KEY (status_id);


--
-- TOC entry 3191 (class 2606 OID 17516)
-- Name: devices devices_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_pk PRIMARY KEY (device_id);


--
-- TOC entry 3195 (class 2606 OID 17518)
-- Name: users newtable_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT newtable_pk PRIMARY KEY (user_id);


--
-- TOC entry 3193 (class 2606 OID 17520)
-- Name: roles role_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT role_pk PRIMARY KEY (role_id);


--
-- TOC entry 3204 (class 2606 OID 17596)
-- Name: devices brand_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT brand_fk FOREIGN KEY (brand_id) REFERENCES public.brand(brand_id);


--
-- TOC entry 3201 (class 2606 OID 17521)
-- Name: devices devices_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT devices_fk FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- TOC entry 3198 (class 2606 OID 17526)
-- Name: device_history devices_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT devices_fk FOREIGN KEY (device_id) REFERENCES public.devices(device_id);


--
-- TOC entry 3202 (class 2606 OID 17531)
-- Name: devices mode_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT mode_fk FOREIGN KEY (mode_id) REFERENCES public.device_mode(mode_id);


--
-- TOC entry 3199 (class 2606 OID 17536)
-- Name: device_history mode_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT mode_fk FOREIGN KEY (mode_id) REFERENCES public.device_mode(mode_id);


--
-- TOC entry 3203 (class 2606 OID 17541)
-- Name: devices status_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.devices
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.device_status(status_id);


--
-- TOC entry 3200 (class 2606 OID 17546)
-- Name: device_history status_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.device_history
    ADD CONSTRAINT status_fk FOREIGN KEY (status_id) REFERENCES public.device_status(status_id);


--
-- TOC entry 3205 (class 2606 OID 17551)
-- Name: users users_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_fk FOREIGN KEY (role_id) REFERENCES public.roles(role_id);


-- Completed on 2023-06-21 09:32:55 WIB

--
-- PostgreSQL database dump complete
--
