--
-- PostgreSQL database dump
--

-- Dumped from database version 15.3
-- Dumped by pg_dump version 15.3

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.accounts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id text,
    account_type text,
    account_number text,
    balance text,
    currency text,
    status text,
    interest_rate text
);


--
-- Name: audit_logs; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.audit_logs (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid,
    action text,
    entity_type text,
    entity_id text,
    details text,
    "timestamp" timestamp with time zone
);


--
-- Name: cards; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.cards (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    account_id text,
    card_number text,
    card_type text,
    expiration_date timestamp with time zone,
    cvv text,
    status text
);


--
-- Name: investments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.investments (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id text,
    type text,
    amount text,
    purchase_date timestamp with time zone,
    current_value text,
    status text
);


--
-- Name: loan_payments; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.loan_payments (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    loan_id text,
    amount text,
    due_date timestamp with time zone,
    paid_at timestamp with time zone,
    status text
);


--
-- Name: loans; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.loans (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id text,
    amount text,
    interest_rate text,
    term text,
    status text,
    purpose text,
    approved_at timestamp with time zone
);


--
-- Name: notifications; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.notifications (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id uuid,
    type text,
    message text,
    read boolean
);


--
-- Name: stocks; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.stocks (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    symbol text,
    name text,
    current_price text,
    last_updated timestamp with time zone
);


--
-- Name: support_tickets; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.support_tickets (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id text,
    subject text,
    description text,
    status text,
    priority text
);


--
-- Name: transactions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.transactions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    account_id text,
    user_id text,
    type text,
    amount text,
    currency text,
    description text,
    category text,
    status text,
    "timestamp" timestamp with time zone
);


--
-- Name: user_profiles; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.user_profiles (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    user_id text,
    occupation text,
    income_range text,
    risk_tolerance text,
    preferences text
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    first_name text,
    last_name text,
    email text,
    phone_number text,
    date_of_birth timestamp with time zone,
    address text,
    password text,
    role text DEFAULT 'customer'::text,
    access_token text,
    refresh_token text,
    is_verified boolean DEFAULT false,
    is_admin boolean DEFAULT false,
    is_deleted boolean DEFAULT false
);


--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.accounts (id, created_at, updated_at, deleted_at, user_id, account_type, account_number, balance, currency, status, interest_rate) FROM stdin;
fcc28330-4d7e-4378-a574-b2dcf4d7bfd2	2024-08-31 05:32:35.006339+01	2024-08-31 05:32:35.006339+01	\N	f2ec6b03-5379-47dd-86d9-584fe90c8c17	current	2062194596	0	USD	active	0.01
\.


--
-- Data for Name: audit_logs; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.audit_logs (id, created_at, updated_at, deleted_at, user_id, action, entity_type, entity_id, details, "timestamp") FROM stdin;
\.


--
-- Data for Name: cards; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.cards (id, created_at, updated_at, deleted_at, account_id, card_number, card_type, expiration_date, cvv, status) FROM stdin;
\.


--
-- Data for Name: investments; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.investments (id, created_at, updated_at, deleted_at, user_id, type, amount, purchase_date, current_value, status) FROM stdin;
\.


--
-- Data for Name: loan_payments; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.loan_payments (id, created_at, updated_at, deleted_at, loan_id, amount, due_date, paid_at, status) FROM stdin;
\.


--
-- Data for Name: loans; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.loans (id, created_at, updated_at, deleted_at, user_id, amount, interest_rate, term, status, purpose, approved_at) FROM stdin;
\.


--
-- Data for Name: notifications; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.notifications (id, created_at, updated_at, deleted_at, user_id, type, message, read) FROM stdin;
\.


--
-- Data for Name: stocks; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.stocks (id, created_at, updated_at, deleted_at, symbol, name, current_price, last_updated) FROM stdin;
\.


--
-- Data for Name: support_tickets; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.support_tickets (id, created_at, updated_at, deleted_at, user_id, subject, description, status, priority) FROM stdin;
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.transactions (id, created_at, updated_at, deleted_at, account_id, user_id, type, amount, currency, description, category, status, "timestamp") FROM stdin;
\.


--
-- Data for Name: user_profiles; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.user_profiles (id, created_at, updated_at, deleted_at, user_id, occupation, income_range, risk_tolerance, preferences) FROM stdin;
938dd5ba-eff6-448c-a279-943fecfea70b	2024-08-31 05:32:33.786312+01	2024-08-31 05:32:33.786312+01	\N	f2ec6b03-5379-47dd-86d9-584fe90c8c17	Software Engineer	medium	moderate	Technology, Investing, Travel
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: -
--

COPY public.users (id, created_at, updated_at, deleted_at, first_name, last_name, email, phone_number, date_of_birth, address, password, role, access_token, refresh_token, is_verified, is_admin, is_deleted) FROM stdin;
f2ec6b03-5379-47dd-86d9-584fe90c8c17	2024-08-31 05:32:32.658567+01	2024-08-31 07:41:45.022012+01	\N	John	Doe	johndoe@example.com	+1 (555) 123-4567	1990-05-15 01:00:00+01	123 Main St, Anytown, AN 12345	$2a$10$wG7Anf5X0nC54LzikTtCuuDhYIwkFh3SvUEx1UufcMPT2AXsHwIRy	customer	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImpvaG5kb2VAZXhhbXBsZS5jb20iLCJleHAiOjE3MjUxNzI5MDUsInJvbGUiOiJjdXN0b21lciIsInVzZXJfaWQiOiJmMmVjNmIwMy01Mzc5LTQ3ZGQtODZkOS01ODRmZTkwYzhjMTcifQ.Yur8-H-7nzhOV8buXmxPzhUyt3TgR0A__y3Pn5tbs4A	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjU2OTEzMDUsInVzZXJfaWQiOiJmMmVjNmIwMy01Mzc5LTQ3ZGQtODZkOS01ODRmZTkwYzhjMTcifQ.zayE_8MIlN7cmdA8G5kr6ADe1utozO8Xhs7ZrrfHGjk	f	f	f
\.


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: audit_logs audit_logs_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT audit_logs_pkey PRIMARY KEY (id);


--
-- Name: cards cards_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.cards
    ADD CONSTRAINT cards_pkey PRIMARY KEY (id);


--
-- Name: investments investments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.investments
    ADD CONSTRAINT investments_pkey PRIMARY KEY (id);


--
-- Name: loan_payments loan_payments_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.loan_payments
    ADD CONSTRAINT loan_payments_pkey PRIMARY KEY (id);


--
-- Name: loans loans_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.loans
    ADD CONSTRAINT loans_pkey PRIMARY KEY (id);


--
-- Name: notifications notifications_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT notifications_pkey PRIMARY KEY (id);


--
-- Name: stocks stocks_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.stocks
    ADD CONSTRAINT stocks_pkey PRIMARY KEY (id);


--
-- Name: support_tickets support_tickets_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.support_tickets
    ADD CONSTRAINT support_tickets_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: user_profiles user_profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.user_profiles
    ADD CONSTRAINT user_profiles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_accounts_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_accounts_deleted_at ON public.accounts USING btree (deleted_at);


--
-- Name: idx_audit_logs_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_audit_logs_deleted_at ON public.audit_logs USING btree (deleted_at);


--
-- Name: idx_cards_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_cards_deleted_at ON public.cards USING btree (deleted_at);


--
-- Name: idx_investments_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_investments_deleted_at ON public.investments USING btree (deleted_at);


--
-- Name: idx_loan_payments_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_loan_payments_deleted_at ON public.loan_payments USING btree (deleted_at);


--
-- Name: idx_loans_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_loans_deleted_at ON public.loans USING btree (deleted_at);


--
-- Name: idx_notifications_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_notifications_deleted_at ON public.notifications USING btree (deleted_at);


--
-- Name: idx_stocks_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_stocks_deleted_at ON public.stocks USING btree (deleted_at);


--
-- Name: idx_support_tickets_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_support_tickets_deleted_at ON public.support_tickets USING btree (deleted_at);


--
-- Name: idx_transactions_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_transactions_deleted_at ON public.transactions USING btree (deleted_at);


--
-- Name: idx_user_profiles_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_user_profiles_deleted_at ON public.user_profiles USING btree (deleted_at);


--
-- Name: idx_user_profiles_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_user_profiles_user_id ON public.user_profiles USING btree (user_id);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: audit_logs fk_audit_logs_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.audit_logs
    ADD CONSTRAINT fk_audit_logs_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: notifications fk_notifications_user; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.notifications
    ADD CONSTRAINT fk_notifications_user FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

