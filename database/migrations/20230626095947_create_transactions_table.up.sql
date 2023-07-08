CREATE TYPE types AS ENUM ('income', 'expenses');

CREATE TYPE currencies AS ENUM ('idr', 'usd');

CREATE TABLE IF NOT EXISTS transactions(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    type types NOT NULL,
    currency currencies NOT NULL,
    note TEXT NOT NULL,
    amount INT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ,

    FOREIGN KEY(category_id) REFERENCES categories(id),
    FOREIGN KEY(user_id) REFERENCES users(id)
);