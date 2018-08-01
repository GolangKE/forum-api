--
-- Name: categories; Type: TABLE; Schema: public
--

CREATE TABLE categories (
    id uuid DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    name character varying,
    code character varying,
    description text,
    is_active boolean DEFAULT true NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    deleted_at timestamp without time zone
);

CREATE UNIQUE INDEX index_categories_on_name ON categories USING btree (name);
CREATE UNIQUE INDEX index_categories_on_code ON categories USING btree (code);