INSERT INTO procurement (id, date, initiator, object, status, budget, required_agreement_date, required_delivery_date, economy_or_overrun, days_in_work, purchase_count)
VALUES
    (1, '2025-01-01', 'Alice', 'Office Supplies', 'Completed', 50000, '2024-12-20', '2025-01-05', 2000, 15, 3),
    (2, '2025-01-10', 'Bob', 'IT Equipment', 'In Progress', 150000, '2025-01-08', '2025-01-20', -5000, 7, 5),
    (3, '2025-01-15', 'Charlie', 'Construction Materials', 'Pending', 300000, '2025-01-14', '2025-02-10', 10000, 2, 8),
    (4, '2024-12-25', 'Diana', 'Office Furniture', 'Completed', 75000, '2024-12-20', '2025-01-01', 3000, 10, 4),
    (5, '2025-01-12', 'Eve', 'Marketing Campaign', 'Canceled', 20000, '2025-01-11', '2025-01-25', 0, 5, 1),
    (6, '2025-01-05', 'Frank', 'Software Licenses', 'Completed', 60000, '2025-01-01', '2025-01-07', 1500, 6, 2),
    (7, '2025-01-02', 'Grace', 'Consulting Services', 'In Progress', 120000, '2024-12-30', '2025-01-10', -2000, 8, 3),
    (8, '2025-01-08', 'Henry', 'Training Programs', 'Completed', 45000, '2025-01-05', '2025-01-12', 500, 7, 2),
    (9, '2024-12-30', 'Ivy', 'Event Management', 'Pending', 100000, '2024-12-28', '2025-01-15', -3000, 3, 6),
    (10, '2025-01-14', 'Jack', 'Research Projects', 'Completed', 250000, '2025-01-10', '2025-01-22', 5000, 9, 4);

INSERT INTO purchase (id, procurement_id, date, initiator, object, status, budget, category, expense_item, required_agreement_date, required_delivery_date, economy_or_overrun, days_in_work, purchase_count)
VALUES
-- Procurement 1
(1, 1, '2025-01-01', 'Alice', 'Office Supplies - Paper', 'Completed', 20000, 'Office', 'Stationery', '2024-12-20', '2025-01-05', 500, 10, 1),
(2, 1, '2025-01-01', 'Alice', 'Office Supplies - Pens', 'Completed', 15000, 'Office', 'Stationery', '2024-12-20', '2025-01-05', 300, 5, 1),
(3, 1, '2025-01-01', 'Alice', 'Office Supplies - Notebooks', 'Completed', 15000, 'Office', 'Stationery', '2024-12-20', '2025-01-05', 1200, 15, 1),

-- Procurement 2
(4, 2, '2025-01-10', 'Bob', 'Laptops', 'In Progress', 100000, 'Technology', 'Hardware', '2025-01-08', '2025-01-20', -3000, 5, 2),
(5, 2, '2025-01-10', 'Bob', 'Monitors', 'In Progress', 50000, 'Technology', 'Hardware', '2025-01-08', '2025-01-20', -2000, 3, 1),

-- Procurement 3
(6, 3, '2025-01-15', 'Charlie', 'Bricks', 'Pending', 150000, 'Construction', 'Raw Materials', '2025-01-14', '2025-02-10', 7000, 2, 4),
(7, 3, '2025-01-15', 'Charlie', 'Cement', 'Pending', 100000, 'Construction', 'Raw Materials', '2025-01-14', '2025-02-10', 2000, 2, 2),
(8, 3, '2025-01-15', 'Charlie', 'Steel Beams', 'Pending', 50000, 'Construction', 'Raw Materials', '2025-01-14', '2025-02-10', 1000, 1, 2),

-- Procurement 4
(9, 4, '2024-12-25', 'Diana', 'Office Chairs', 'Completed', 50000, 'Office', 'Furniture', '2024-12-20', '2025-01-01', 2000, 10, 2),
(10, 4, '2024-12-25', 'Diana', 'Office Desks', 'Completed', 25000, 'Office', 'Furniture', '2024-12-20', '2025-01-01', 1000, 8, 2),

-- Procurement 5
(11, 5, '2025-01-12', 'Eve', 'Online Ads', 'Canceled', 20000, 'Marketing', 'Advertising', '2025-01-11', '2025-01-25', 0, 5, 1),

-- Procurement 6
(12, 6, '2025-01-05', 'Frank', 'Antivirus Licenses', 'Completed', 30000, 'Technology', 'Software', '2025-01-01', '2025-01-07', 800, 3, 1),
(13, 6, '2025-01-05', 'Frank', 'Office Software', 'Completed', 30000, 'Technology', 'Software', '2025-01-01', '2025-01-07', 700, 3, 1),

-- Procurement 7
(14, 7, '2025-01-02', 'Grace', 'Consulting Services - Strategy', 'In Progress', 70000, 'Services', 'Consulting', '2024-12-30', '2025-01-10', -1500, 5, 1),
(15, 7, '2025-01-02', 'Grace', 'Consulting Services - HR', 'In Progress', 50000, 'Services', 'Consulting', '2024-12-30', '2025-01-10', -500, 3, 1),

-- Procurement 8
(16, 8, '2025-01-08', 'Henry', 'Training - Leadership', 'Completed', 25000, 'Education', 'Training', '2025-01-05', '2025-01-12', 300, 4, 1),
(17, 8, '2025-01-08', 'Henry', 'Training - Technical Skills', 'Completed', 20000, 'Education', 'Training', '2025-01-05', '2025-01-12', 200, 3, 1),

-- Procurement 9
(18, 9, '2024-12-30', 'Ivy', 'Event Planning - Venue', 'Pending', 60000, 'Events', 'Management', '2024-12-28', '2025-01-15', -2000, 2, 1),
(19, 9, '2024-12-30', 'Ivy', 'Event Planning - Catering', 'Pending', 40000, 'Events', 'Management', '2024-12-28', '2025-01-15', -1000, 1, 1),

-- Procurement 10
(20, 10, '2025-01-14', 'Jack', 'Research Equipment', 'Completed', 200000, 'Research', 'Innovation', '2025-01-10', '2025-01-22', 4000, 6, 2),
(21, 10, '2025-01-14', 'Jack', 'Research Personnel', 'Completed', 50000, 'Research', 'Innovation', '2025-01-10', '2025-01-22', 1000, 3, 1);
