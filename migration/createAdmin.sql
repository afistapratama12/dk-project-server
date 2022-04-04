INSERT INTO users(id, role, fullname, phone_number, username, password, sas_balance, ro_balance, money_balance, parent_id, position) 
VALUES (1, 'admin', 'admin', '082132132132', 'DK-1-admin', '1234', 100, 100, 0, NULL, NULL),
(2, 'user', 'pandu', '082232232232', 'DK-2-pandu', '1234', 20, 20, 1000000, 1, "left");