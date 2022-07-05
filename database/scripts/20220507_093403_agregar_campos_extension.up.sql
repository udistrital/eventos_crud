ALTER TABLE eventos.calendario ADD COLUMN documento_extension_id integer;

COMMENT ON COLUMN eventos.calendario.documento_extension_id IS 'Campo id para documento soporte si hay extensión';

ALTER TABLE eventos.calendario ADD COLUMN aplica_extension boolean;

COMMENT ON COLUMN eventos.calendario.aplica_extension IS 'Campo para verificar si es calendario por extensión';

ALTER TABLE eventos.calendario ADD COLUMN aplica_edicion_actividades boolean;

COMMENT ON COLUMN eventos.calendario.aplica_edicion_actividades IS 'Campo para verficar si al calendario aplica edicion de fechas actividades';

ALTER TABLE eventos.calendario ADD COLUMN dependencia_particular_id json;

COMMENT ON COLUMN eventos.calendario.dependencia_particular_id IS 'Campo json para listar dependencias con ajustes particulares de fechas'