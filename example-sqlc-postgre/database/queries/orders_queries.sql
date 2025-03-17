-- name: UpsertOrders :exec
INSERT INTO "orders"(id,
                     account_id,
                     symbol,
                     quantity,
                     fees,
                     status,
                     type,
                     version)
VALUES ($1,
        $2,
        $3,
        $4,
        $5,
        $6,
        $7,
        $8) ON CONFLICT (id) DO
UPDATE
    SET quantity = EXCLUDED.quantity,
    fees = EXCLUDED.fees,
    status = EXCLUDED.status,
    type = EXCLUDED.type,
    version = EXCLUDED.version
WHERE orders.version <= EXCLUDED.version;

-- name: GetOrdersByStatuses :many
SELECT id, account_id, symbol, quantity, fees, status, type, version, created_at, updated_at
FROM orders
WHERE status::text = ANY($1::text[])
ORDER BY created_at DESC;