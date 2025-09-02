-- +migrate Down
DELETE FROM roles where id IN (1,2);