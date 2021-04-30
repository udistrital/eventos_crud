ALTER TABLE eventos.tipo_evento ADD calendario_id integer;

ALTER TABLE eventos.tipo_evento ADD CONSTRAINT fk_calendario FOREIGN KEY (calendario_id)
REFERENCES eventos.calendario (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;