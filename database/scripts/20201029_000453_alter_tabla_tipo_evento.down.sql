ALTER TABLE eventos.tipo_evento DROP COLUMN calendario_id;
ALTER TABLE eventos.tipo_evento DROP CONSTRAINT IF EXISTS fk_calendario CASCADE;