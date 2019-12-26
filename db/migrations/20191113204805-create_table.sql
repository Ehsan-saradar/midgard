
-- +migrate Up

CREATE TABLE genesis
(
    genesis_time TIMESTAMPTZ not null,
    primary key (genesis_time)
);

CREATE TABLE events (
    time        TIMESTAMPTZ       not null,
    id bigint not null,
    height bigint not null,
    type varchar not null,
    status varchar not null,
    primary key (id, time)
);

CREATE TABLE stakes (
    time        TIMESTAMPTZ       NOT NULL,
    id SERIAL,
    event_id bigint not null,
    from_address varchar not null,
    pool varchar not null,
    runeAmt bigint,
    assetAmt bigint,
    units bigint,
    primary key (id, time)
);

CREATE TABLE swaps (
    time        TIMESTAMPTZ       NOT NULL,
    id SERIAL,
    event_id bigint not null,
    from_address varchar not null,
    to_address varchar not null,
    pool varchar not null,
    price_target bigint,
    trade_slip real,
    liquidity_fee bigint,
    runeAmt bigint,
    assetAmt bigint,
    primary key (id, time)
);

CREATE TYPE tx_direction as enum('in', 'out');
CREATE TABLE txs (
    time        TIMESTAMPTZ       NOT NULL,
    id SERIAL,
    tx_hash varchar not null,
    event_id bigint not null,
    direction tx_direction not null,
    chain varchar,
    from_address varchar,
    to_address varchar,
    memo varchar,
    primary key (id, time, event_id)
);

CREATE TABLE coins (
    time        TIMESTAMPTZ       NOT NULL,
    id SERIAL,
    tx_hash varchar not null,
    event_id bigint not null,
    chain varchar not null,
    symbol varchar not null,
    ticker varchar not null,
    amount bigint not null,
    primary key (id, time, event_id)
);

CREATE TABLE gas (
    time        TIMESTAMPTZ       NOT NULL,
    event_id bigint not null,
    chain varchar not null,
    symbol varchar not null,
    ticker varchar not null,
    amount bigint not null,
    primary key (time, event_id, chain, symbol)
);

CREATE EXTENSION IF NOT EXISTS timescaledb CASCADE;
SELECT create_hypertable('events', 'time');
SELECT create_hypertable('stakes', 'time');
SELECT create_hypertable('swaps', 'time');
SELECT create_hypertable('txs', 'time');
SELECT create_hypertable('coins', 'time');
SELECT create_hypertable('gas', 'time');

-- +migrate Down

DROP TABLE genesis;
DROP TABLE events;
DROP TABLE stakes;
DROP TABLE swaps;
DROP TABLE txs;
DROP TABLE coins;
DROP TABLE gas;

DROP TYPE tx_direction;
