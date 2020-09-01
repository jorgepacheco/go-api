CREATE TABLE accounts.account (
	id uuid NOT NULL,
	iban varchar(140) NULL,
    balance numeric(22,3) NULL,
	CONSTRAINT account_pkey PRIMARY KEY (id)
);