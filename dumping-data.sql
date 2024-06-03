--
-- PostgreSQL database dump
--

-- Dumped from database version 15.2
-- Dumped by pg_dump version 15.2

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
-- Data for Name: category; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.category (category_id, name) FROM stdin;
\.


--
-- Data for Name: user; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public."user" (user_id, email, username, whatsapp_number, password, created_at, updated_at) FROM stdin;
12	agung@gmail.com	agus123	098765431	password	2024-05-15 21:15:47.640992	2024-05-15 21:15:47.640992
14	acha@gmail.com	acha_d4itb	081938833350	acha_d4itb	2024-05-31 10:52:51.681931	2024-05-31 10:52:51.681931
16	tsaqif@gmail.com	tsaqif_d4itb	082133473078	tsaqif_d4itb	2024-05-31 10:53:54.38404	2024-05-31 10:53:54.38404
17	zaki@gmail.com	zaki_d4itb	085888429506	zaki_d4itb	2024-05-31 10:54:37.310462	2024-05-31 10:54:37.310462
18	nandha@gmail.com	nandha_d4itb	081283230218	nandha_d4itb	2024-05-31 10:55:22.398308	2024-05-31 10:55:22.398308
20	desy@gmail.com	desy_d4itb	081210831115	desy_d4itb	2024-05-31 10:56:02.156545	2024-05-31 10:56:02.156545
21	alif@gmail.com	alif_d4itb	081212548640	alif_d4itb	2024-05-31 10:56:32.411366	2024-05-31 10:56:32.411366
\.


--
-- Data for Name: store; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.store (store_id, user_id, name, whatsapp_number, photo_profile, closed_at) FROM stdin;
16	16	Toko Seblak Mas Tsaqif	082133473078	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	\N
17	18	Toko Komputer Pak Nandha	081283230218	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	\N
15	21	Toko Kopi Bang Alif	081212548640	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	\N
18	17	Es Buah Abah Zaki	085888429506	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	\N
19	20	Baju Muslimah Bu Desy	081210831115	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	\N
20	14	Toko Buku Mba Acha	081938833350	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	\N
\.


--
-- Data for Name: product; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.product (product_id, store_id, name, category_id, photo, description, deleted_at) FROM stdin;
36	18	Sup Buah	\N	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	Sueger cik	\N
38	18	Es Jeruk	\N	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	Asem & seger	\N
37	18	Jus Alpukat	\N	https://img.freepik.com/free-vector/businessman-character-avatar-isolated_24877-60111.jpg	Dingin🥶	\N
48	18	Es Buah Segar	\N	https://img.example.com/es-buah-segar.jpg	Deskripsi produk Es Buah Segar	\N
49	18	Es Campur	\N	https://img.example.com/es-campur.jpg	Deskripsi produk Es Campur	\N
50	18	Es Kelapa Muda	\N	https://img.example.com/es-kelapa-muda.jpg	Deskripsi produk Es Kelapa Muda	\N
51	19	Baju Gamis Motif Bunga	\N	https://img.example.com/baju-gamis-motif-bunga.jpg	Deskripsi produk Baju Gamis Motif Bunga	\N
52	19	Khimar Polos	\N	https://img.example.com/khimar-polos.jpg	Deskripsi produk Khimar Polos	\N
53	19	Mukena Katun	\N	https://img.example.com/mukena-katun.jpg	Deskripsi produk Mukena Katun	\N
54	20	Buku Anak-Anak Ceria	\N	https://img.example.com/buku-anak-ceria.jpg	Deskripsi produk Buku Anak-Anak Ceria	\N
39	15	Kopi Arabika	\N	https://img.example.com/kopi-arabika.jpg	Bikin melek seharian	\N
41	15	Kopi Ekselsa	\N	https://img.example.com/kopi-ekselsa.jpg	Bikin semangat tahajud semalaman	\N
42	16	Seblak Pedas Level 5	\N	https://img.example.com/seblak-pedas-level-5.jpg	Seblak bikin doer	\N
43	16	Seblak Kuah	\N	https://img.example.com/seblak-kuah.jpg	Cocok untuk the nuruls	\N
44	16	Seblak Instan	\N	https://img.example.com/seblak-instan.jpg	Seblak wong kere	\N
45	17	Laptop ASUS ROG	\N	https://img.example.com/laptop-asus-rog.jpg	Laptop anti karat	\N
46	17	Monitor 27 Inch	\N	https://img.example.com/monitor-27-inch.jpg	Monitor tahan banting	\N
47	17	Keyboard Gaming	\N	https://img.example.com/keyboard-gaming.jpg	Keyboard anti maling	\N
55	20	Novel Romantis	\N	https://img.example.com/novel-romantis.jpg	Cocok buat balita	\N
56	20	Buku Pelajaran SD	\N	https://img.example.com/buku-pelajaran-sd.jpg	Anak tk dilarang baca	\N
40	15	Kopi Robusta	\N	https://img.example.com/kopi-robusta.jpg	Bikin semangat kerja	\N
\.


--
-- Data for Name: store_pickup_place; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.store_pickup_place (store_pickup_place_id, store_id, name, deleted_at) FROM stdin;
17	17	Pos Satpam D4	\N
11	18	Pos Satpam	\N
21	18	Lobby Pasca Sarjana	\N
14	18	Kantin PENS	\N
19	17	Parkiran Maba	\N
22	15	Pos Satpam D4	\N
23	15	Depan Masjid An-Nahl	\N
24	15	Depan Perpus D3	\N
25	16	Kantin PENS	\N
26	16	Depan Gedung D4	\N
27	16	Sebelah Timur Lapangan Merah	\N
29	17	Jalan setapak PENS-PPNS	\N
30	17	Lobby Pasca Sarjana	\N
31	19	Kantin PENS	\N
32	19	Depan Perpus D4	\N
33	19	Depan BAK	\N
34	20	Kantin PENS	\N
35	20	Perpustakaan D4	\N
36	20	Parkiran PENS	\N
\.


--
-- Data for Name: batches; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.batches (batch_id, product_id, store_pickup_place_id, stock, close_order_time, pickup_time, price) FROM stdin;
10	37	11	5	2024-06-15 17:00:00+07	2024-06-20 09:00:00+07	10000
11	37	11	5	2024-06-15 17:00:00+07	2024-06-20 09:00:00+07	10000
12	37	11	3	2024-06-25 17:00:00+07	2024-06-30 09:00:00+07	11000
13	39	22	10	2024-07-15 17:00:00+07	2024-07-20 09:00:00+07	\N
14	39	23	15	2024-07-16 17:00:00+07	2024-07-21 09:00:00+07	\N
15	39	24	20	2024-07-17 17:00:00+07	2024-07-22 09:00:00+07	\N
16	40	22	12	2024-07-18 17:00:00+07	2024-07-23 09:00:00+07	\N
17	40	23	18	2024-07-19 17:00:00+07	2024-07-24 09:00:00+07	\N
18	40	24	22	2024-07-20 17:00:00+07	2024-07-25 09:00:00+07	\N
19	41	22	14	2024-07-21 17:00:00+07	2024-07-26 09:00:00+07	\N
20	41	23	16	2024-07-22 17:00:00+07	2024-07-27 09:00:00+07	\N
21	41	24	19	2024-07-23 17:00:00+07	2024-07-28 09:00:00+07	\N
22	42	25	10	2024-07-15 17:00:00+07	2024-07-20 09:00:00+07	\N
23	42	26	12	2024-07-16 17:00:00+07	2024-07-21 09:00:00+07	\N
24	42	27	15	2024-07-17 17:00:00+07	2024-07-22 09:00:00+07	\N
25	43	25	8	2024-07-18 17:00:00+07	2024-07-23 09:00:00+07	\N
26	43	26	14	2024-07-19 17:00:00+07	2024-07-24 09:00:00+07	\N
27	43	27	20	2024-07-20 17:00:00+07	2024-07-25 09:00:00+07	\N
28	44	25	10	2024-07-21 17:00:00+07	2024-07-26 09:00:00+07	\N
29	44	26	11	2024-07-22 17:00:00+07	2024-07-27 09:00:00+07	\N
30	44	27	13	2024-07-23 17:00:00+07	2024-07-28 09:00:00+07	\N
31	45	17	5	2024-07-15 17:00:00+07	2024-07-20 09:00:00+07	\N
32	45	19	7	2024-07-16 17:00:00+07	2024-07-21 09:00:00+07	\N
33	45	29	9	2024-07-17 17:00:00+07	2024-07-22 09:00:00+07	\N
34	46	17	4	2024-07-18 17:00:00+07	2024-07-23 09:00:00+07	\N
35	46	19	6	2024-07-19 17:00:00+07	2024-07-24 09:00:00+07	\N
36	46	29	8	2024-07-20 17:00:00+07	2024-07-25 09:00:00+07	\N
37	47	17	3	2024-07-21 17:00:00+07	2024-07-26 09:00:00+07	\N
38	47	19	5	2024-07-22 17:00:00+07	2024-07-27 09:00:00+07	\N
39	47	29	7	2024-07-23 17:00:00+07	2024-07-28 09:00:00+07	\N
40	36	11	30	2024-07-15 17:00:00+07	2024-07-20 09:00:00+07	\N
41	36	14	25	2024-07-16 17:00:00+07	2024-07-21 09:00:00+07	\N
42	36	21	35	2024-07-17 17:00:00+07	2024-07-22 09:00:00+07	\N
43	37	11	40	2024-07-18 17:00:00+07	2024-07-23 09:00:00+07	\N
44	37	14	45	2024-07-19 17:00:00+07	2024-07-24 09:00:00+07	\N
45	37	21	50	2024-07-20 17:00:00+07	2024-07-25 09:00:00+07	\N
46	38	11	20	2024-07-21 17:00:00+07	2024-07-26 09:00:00+07	\N
47	38	14	30	2024-07-22 17:00:00+07	2024-07-27 09:00:00+07	\N
48	38	21	40	2024-07-23 17:00:00+07	2024-07-28 09:00:00+07	\N
49	51	31	15	2024-07-15 17:00:00+07	2024-07-20 09:00:00+07	\N
50	51	32	20	2024-07-16 17:00:00+07	2024-07-21 09:00:00+07	\N
51	51	33	25	2024-07-17 17:00:00+07	2024-07-22 09:00:00+07	\N
52	52	31	10	2024-07-18 17:00:00+07	2024-07-23 09:00:00+07	\N
53	52	32	15	2024-07-19 17:00:00+07	2024-07-24 09:00:00+07	\N
54	52	33	20	2024-07-20 17:00:00+07	2024-07-25 09:00:00+07	\N
55	53	31	30	2024-07-21 17:00:00+07	2024-07-26 09:00:00+07	\N
56	53	32	35	2024-07-22 17:00:00+07	2024-07-27 09:00:00+07	\N
57	53	33	40	2024-07-23 17:00:00+07	2024-07-28 09:00:00+07	\N
58	54	34	50	2024-07-15 17:00:00+07	2024-07-20 09:00:00+07	\N
59	54	35	60	2024-07-16 17:00:00+07	2024-07-21 09:00:00+07	\N
60	54	36	70	2024-07-17 17:00:00+07	2024-07-22 09:00:00+07	\N
61	55	34	30	2024-07-18 17:00:00+07	2024-07-23 09:00:00+07	\N
62	55	35	40	2024-07-19 17:00:00+07	2024-07-24 09:00:00+07	\N
63	55	36	50	2024-07-20 17:00:00+07	2024-07-25 09:00:00+07	\N
64	56	34	20	2024-07-21 17:00:00+07	2024-07-26 09:00:00+07	\N
65	56	35	25	2024-07-22 17:00:00+07	2024-07-27 09:00:00+07	\N
66	56	36	30	2024-07-23 17:00:00+07	2024-07-28 09:00:00+07	\N
\.


--
-- Data for Name: cart; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.cart (cart_id, user_id, batch_id, quantity) FROM stdin;
1	12	10	2
2	12	11	3
3	12	12	1
4	12	13	4
5	12	14	5
6	14	15	2
7	14	16	3
8	14	17	1
9	14	18	4
10	14	19	5
11	16	20	2
12	16	21	3
13	16	22	1
14	16	23	4
15	16	24	5
16	17	25	2
17	17	26	3
18	17	27	1
19	17	28	4
20	17	29	5
21	18	30	2
22	18	31	3
23	18	32	1
24	18	33	4
25	18	34	5
26	20	35	2
27	20	36	3
28	20	37	1
29	20	38	4
30	20	39	5
31	21	40	2
32	21	41	3
33	21	42	1
34	21	43	4
35	21	44	5
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.transactions (transaction_id, user_id, product_id, batch_id, transaction_date, total_payment, quantity, transaction_status, cancelled_transaction_date, cancelled_transaction_reason) FROM stdin;
\.


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.orders (order_id, transaction_id, batch_id, total_payment, stock) FROM stdin;
\.


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.schema_migrations (version, dirty) FROM stdin;
20240603114953	f
\.


--
-- Name: batches_batch_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.batches_batch_id_seq', 66, true);


--
-- Name: cart_cart_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.cart_cart_id_seq', 35, true);


--
-- Name: category_category_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.category_category_id_seq', 4, true);


--
-- Name: orders_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.orders_order_id_seq', 18, true);


--
-- Name: product_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.product_product_id_seq', 56, true);


--
-- Name: store_pickup_place_store_pickup_place_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.store_pickup_place_store_pickup_place_id_seq', 36, true);


--
-- Name: store_store_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.store_store_id_seq', 20, true);


--
-- Name: transactions_transaction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.transactions_transaction_id_seq', 1, false);


--
-- Name: user_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.user_user_id_seq', 21, true);


--
-- PostgreSQL database dump complete
--

