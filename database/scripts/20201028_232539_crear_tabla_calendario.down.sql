ALTER TABLE eventos.calendario DROP CONSTRAINT IF EXISTS fk_calendario_padre CASCADE;
DROP TABLE IF EXISTS eventos.calendario;