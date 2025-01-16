CREATE TABLE procurement (
                             id SERIAL PRIMARY KEY,
                             date DATE NOT NULL,
                             initiator VARCHAR(100) NOT NULL,
                             object VARCHAR(100) NOT NULL,
                             status VARCHAR(50) NOT NULL,
                             budget INT NOT NULL,
                             required_agreement_date DATE NOT NULL,
                             required_delivery_date DATE NOT NULL,
                             economy_or_overrun VARCHAR(50),
                             days_in_work INT NOT NULL,
                             purchase_count INT NOT NULL
);
