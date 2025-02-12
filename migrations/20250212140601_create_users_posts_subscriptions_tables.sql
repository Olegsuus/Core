-- +goose Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
    );

CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id UUID NOT NULL,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id)
    );

CREATE TABLE IF NOT EXISTS subscriptions (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    subscriber_id UUID NOT NULL,
    subscribed_to_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_subscriber FOREIGN KEY (subscriber_id) REFERENCES users(id),
    CONSTRAINT fk_subscribed_to FOREIGN KEY (subscribed_to_id) REFERENCES users(id)
    );

CREATE INDEX IF NOT EXISTS idx_subscriber_id ON subscriptions(subscriber_id);
CREATE INDEX IF NOT EXISTS idx_subscribed_to_id ON subscriptions(subscribed_to_id);

-- +goose Down
DROP TABLE IF EXISTS subscriptions;
DROP TABLE IF EXISTS posts;
DROP TABLE IF EXISTS users;