SELECT 
    wr.Id, 
    wr.money_balance, 
    wr.ro_balance, 
    wr.ro_money_balance,
    wr.created_at,
    wr.updated_at,
    wr.approved,
    u.id as user_id,
    u.fullname,
    u.phone_number,
    ba.id as bank_acc_id,
    ba.bank_name,
    ba.bank_number,
    ba.name_on_bank
FROM withdraw_requests wr
JOIN users u ON u.id = wr.user_id
JOIN bank_accounts ba ON ba.id = wr.bank_acc_id;