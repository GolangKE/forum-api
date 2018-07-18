--
-- Create users table
--

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() NOT NULL,
    email character varying DEFAULT ''::character varying NOT NULL,
    username character varying,
    full_name character varying,
    password_digest character varying DEFAULT ''::character varying NOT NULL,
    reset_password_token character varying,
    reset_password_sent_at timestamp without time zone,
    sign_in_count integer DEFAULT 0 NOT NULL,
    current_sign_in_at timestamp without time zone,
    last_sign_in_at timestamp without time zone,
    current_sign_in_ip inet,
    last_sign_in_ip inet,
    confirmation_token character varying,
    confirmed_at timestamp without time zone,
    confirmation_sent_at timestamp without time zone,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);

--
-- Set id primary key constraint
--

ALTER TABLE ONLY users ADD CONSTRAINT users_pkey PRIMARY KEY (id);

--
-- Add unique index to email
--

CREATE UNIQUE INDEX users_email_index ON users USING btree (email);

--
-- Add unique index to username
--

CREATE UNIQUE INDEX users_username_index ON users USING btree (username);

