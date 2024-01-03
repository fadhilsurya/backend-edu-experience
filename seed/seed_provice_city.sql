INSERT INTO `province` (name, created_at, updated_at)
VALUES
    ('DKI Jakarta', current_date, current_date),
    ('Jawa Barat',current_date, current_date);

INSERT INTO `city` (name, province_id, created_at, updated_at)
VALUES
    ('Jakarta Selatan', 1, current_date, current_date),
    ('Jakarta Pusat', 1, current_date, current_date),
    ('Bekasi', 2, current_date, current_date);
