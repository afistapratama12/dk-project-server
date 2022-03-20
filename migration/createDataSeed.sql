INSERT INTO users(role, fullname, phone_number, username, password) 
VALUES ('admin', 'admin', '082132132132', 'DK-1-admin', '1234');

INSERT INTO users(role, fullname, phone_number, username, password, parent_id, position) 
VALUES ('user', 'user2', '082132132132', 'DK-2-user2', '1234', 1, 'left'),
    ('user', 'user3', '082132132132', 'DK-3-user3', '1234', 1, 'center'),
    ('user', 'user4', '082132132132', 'DK-4-user4', '1234', 2, 'left');

INSERT INTO bank_accounts(bank_name, bank_number, name_on_bank, user_id) 
VALUES ('BCA', '123123123', 'user user user', 2),
('BCA', '321321321', 'user user user', 3);

INSERT INTO withdraw_requests(user_id, bank_acc_id, money_balance, ro_balance, ro_money_balance, created_at, updated_at, approved) 
VALUES (2, 1, 100000, 1, 2000, NOW(), NOW(), false),
(3, 2, 0, 2, 4000, NOW(), NOW(), false);

-- SCENARIO untuk untuk withdraw request si RO
INSERT INTO users(role, fullname, phone_number, username, password, parent_id, position) 
VALUES ('user', 'user5', '082132131234', 'DK-5-user', '1234', 4, 'left'),
    ('user', 'user6', '082132131234', 'DK-6-user', '1234', 5, 'left'),
    ('user', 'user7', '082132131234', 'DK-7-user', '1234', 6, 'left');

INSERT INTO bank_accounts(bank_name, bank_number, name_on_bank, user_id) 
VALUES ('BCA', '123123123', 'user user user', 4),
('BCA', '321321321', 'user user user', 5),
('BCA', '321321321', 'user user user', 6);

-- user5 create ro 2
-- user6 create ro 1
-- user7 create ro 4

