-- PostgreSQL schema for Stocky Reward System

CREATE TABLE users (
    id UUID PRIMARY KEY,
    name VARCHAR(100)
);

CREATE TABLE rewards (
    id UUID PRIMARY KEY,
    user_id UUID REFERENCES users(id),
    stock_symbol VARCHAR(20),
    quantity NUMERIC(18,6),
    timestamp TIMESTAMP
);

CREATE TABLE ledger_entries (
    id UUID PRIMARY KEY,
    reward_id UUID REFERENCES rewards(id),
    entry_type VARCHAR(20),
    stock_symbol VARCHAR(20),
    quantity NUMERIC(18,6),
    amount_inr NUMERIC(18,4),
    description VARCHAR(100),
    timestamp TIMESTAMP
);

CREATE TABLE stock_prices (
    stock_symbol VARCHAR(20) PRIMARY KEY,
    price_inr NUMERIC(18,4),
    timestamp TIMESTAMP
);
