CREATE TABLE purchase (
           id SERIAL PRIMARY KEY,
           procurement_id INT NULL,
           date DATE NOT NULL,
           initiator VARCHAR(100) NOT NULL,
           object VARCHAR(100) NOT NULL,
           status VARCHAR(50) NOT NULL,
           budget INT NOT NULL,
           category VARCHAR(100) NOT NULL,  -- Новое поле для категории
           expense_item VARCHAR(100) NOT NULL,  -- Новое поле для статьи расходов
           required_agreement_date DATE NOT NULL,  -- Уже было в запросе
           required_delivery_date DATE NOT NULL,  -- Уже было в запросе
           economy_or_overrun INT,  -- Экономия/Перевышение
           days_in_work INT NOT NULL,  -- Количество дней в работе
           purchase_count INT NOT NULL,  -- Количество закупок
           CONSTRAINT fk_procurement FOREIGN KEY (procurement_id) REFERENCES procurement(id) ON DELETE CASCADE
);
