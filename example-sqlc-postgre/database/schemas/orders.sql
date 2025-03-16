CREATE TYPE order_status AS ENUM ('init', 'pending', 'executing', 'failed', 'done');
CREATE TYPE order_type AS ENUM ('limit', 'market');

CREATE TABLE orders
(
    id         string                NOT NULL primary key,
    account_id uuid                  NOT NULL,
    symbol     character varying(20) NOT NULL,
    quantity   numeric(48, 30),
    fees       jsonb,
    status     order_status          not null default 'init',
    type       order_type            not null,
    version    int                   NOT NULL,
    created_at timestamp without time zone    NOT NULL DEfAULT current_timestamp,
    updated_at timestamp without time zone    NOT NULL DEfAULT current_timestamp
)