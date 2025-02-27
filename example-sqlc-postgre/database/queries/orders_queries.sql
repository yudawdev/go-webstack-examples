-- name: UpsertOrders :exec
INSERT INTO "orders"(id,
                     account_id,
                     symbol,
                     quantity,
                     fees,
                     status,
                     version)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7) ON CONFLICT (id) DO
UPDATE
    SET quantity = EXCLUDED.quantity,
    fees = EXCLUDED.fees,
    status = EXCLUDED.status,
    version = EXCLUDED.version
WHERE orders.version <= EXCLUDED.version;