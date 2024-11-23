CREATE TABLE customer (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    xid VARCHAR(255) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE credential (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id uuid NOT NULL REFERENCES customer(id),
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE wallet (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    customer_id uuid NOT NULL UNIQUE REFERENCES customer(id),
    balance BIGINT DEFAULT 0 NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('enabled', 'disabled')),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

CREATE TABLE transaction (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    wallet_id uuid NOT NULL REFERENCES wallet(id),
    amount BIGINT NOT NULL,
    status VARCHAR(50) NOT NULL CHECK (status IN ('success', 'failed')),
    type VARCHAR(50) NOT NULL CHECK (type IN ('deposit', 'withdrawal')),
    reference_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);
