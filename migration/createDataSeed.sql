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

INSERT INTO users(role, fullname, phone_number, username, password, parent_id, position) 
VALUES ('user', 'user5', '082132131234', 'DK-5-user', '1234', 4, 'left'),
    ('user', 'user6', '082132131234', 'DK-6-user', '1234', 5, 'left'),
    ('user', 'user7', '082132131234', 'DK-7-user', '1234', 6, 'left');

INSERT INTO bank_accounts(bank_name, bank_number, name_on_bank, user_id) 
VALUES ('BCA', '123123123', 'user user user', 4),
('BCA', '321321321', 'user user user', 5),
('BCA', '321321321', 'user user user', 6);

INSERT INTO users(role, fullname, phone_number, username, password, parent_id, position) 
VALUES ('user', 'user8', '082132131234', 'DK-8-user', '1234', 4, 'right'),
    ('user', 'user9', '082132131234', 'DK-9-user', '1234', 5, 'center'),
    ('user', 'user10', '082132131234', 'DK-10-user', '1234', 6, 'center'),
    ('user', 'user11', '082132131234', 'DK-11-user', '1234', 7, 'left'),
    ('user', 'user12', '082132131234', 'DK-12-user', '1234', 8, 'left'),
    ('user', 'user13', '082132131234', 'DK-13-user', '1234', 9, 'left'),
    ('user', 'user14', '082132131234', 'DK-14-user', '1234', 10, 'left'),
    ('user', 'user15', '082132131234', 'DK-15-user', '1234', 11, 'left'),
    ('user', 'user16', '082132131234', 'DK-16-user', '1234', 12, 'left'),
    ('user', 'user17', '082132131234', 'DK-17-user', '1234', 13, 'left'),
    ('user', 'user18', '082132131234', 'DK-18-user', '1234', 14, 'left'),
    ('user', 'user19', '082132131234', 'DK-19-user', '1234', 15, 'left'),
    ('user', 'user20', '082132131234', 'DK-20-user', '1234', 16, 'left'),
    ('user', 'user21', '082132131234', 'DK-21-user', '1234', 17, 'left'),
    ('user', 'user22', '082132131234', 'DK-22-user', '1234', 18, 'left'),
    ('user', 'user23', '082132131234', 'DK-23-user', '1234', 19, 'left'),
    ('user', 'user24', '082132131234', 'DK-24-user', '1234', 20, 'left'),
    ('user', 'user25', '082132131234', 'DK-25-user', '1234', 21, 'left'),
    ('user', 'user26', '082132131234', 'DK-26-user', '1234', 22, 'left'),
    ('user', 'user27', '082132131234', 'DK-27-user', '1234', 23, 'left'),
    ('user', 'user28', '082132131234', 'DK-28-user', '1234', 24, 'left');
