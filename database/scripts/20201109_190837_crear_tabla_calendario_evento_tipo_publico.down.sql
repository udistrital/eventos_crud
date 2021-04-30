DROP TABLE eventos.calendario_evento_tipo_publico CASCADE;

ALTER TABLE eventos.tipo_publico DROP COLUMN IF EXISTS calendario_evento_id;

ALTER TABLE eventos.tipo_publico ADD COLUMN calendario_evento_id integer;

ALTER TABLE eventos.tipo_publico DROP CONSTRAINT IF EXISTS fk_tipo_publico_calendario_evento CASCADE;

ALTER TABLE eventos.tipo_publico ADD CONSTRAINT fk_tipo_publico_calendario_evento FOREIGN KEY (calendario_evento_id) REFERENCES eventos.calendario_evento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;

COMMENT ON COLUMN eventos.tipo_publico.calendario_evento_id IS 'FK de tabla calendario_evento';