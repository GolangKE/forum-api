--
-- Name: user_tokens; Type: TABLE; Schema: public
--

CREATE TABLE user_tokens (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    token character varying NOT NULL,
    user_id uuid NOT NULL,
    is_valid boolean DEFAULT true NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);

CREATE INDEX index_user_tokens_on_user_id ON user_tokens USING btree (user_id);
CREATE UNIQUE INDEX index_user_tokens_on_token ON user_tokens USING btree (token);