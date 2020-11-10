CREATE TABLE eventos.calendario_evento_tipo_publico 
(id serial NOT NULL,
fecha_creacion timestamp NOT NULL,
fecha_modificacion timestamp NOT NULL,
activo boolean NOT NULL,
calendario_evento_id integer,
tipo_publico_id integer,
CONSTRAINT pk_calendario_evento_tipo_publico PRIMARY KEY (id));


ALTER TABLE eventos.calendario_evento_tipo_publico OWNER TO desarrollooas;
COMMENT ON TABLE eventos.calendario_evento_tipo_publico IS 'Tabla de rompimiento entre los diferentes eventos de calendarios y el tipos de publico';
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.id IS 'Identificador de la tabla responsable_sesion';
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.fecha_creacion IS 'Fecha de creacion de un participante del evento';
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.fecha_modificacion IS 'Fecha de modifiación de un participante del evento';
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.activo IS 'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
COMMENT ON CONSTRAINT pk_calendario_evento_tipo_publico ON eventos.calendario_evento_tipo_publico  IS 'Identificador de la tabla responsable evento';


ALTER TABLE eventos.calendario_evento_tipo_publico ADD CONSTRAINT fk_calendario_evento_tipo_publico_calendario_evento FOREIGN KEY (calendario_evento_id) REFERENCES eventos.calendario_evento (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;
ALTER TABLE eventos.calendario_evento_tipo_publico ADD CONSTRAINT fk_calendario_evento_tipo_publico_tipo_publico FOREIGN KEY (tipo_publico_id) REFERENCES eventos.tipo_publico (id) MATCH FULL ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE eventos.tipo_publico DROP CONSTRAINT IF EXISTS fk_tipo_publico_calendario_evento CASCADE;
ALTER TABLE eventos.tipo_publico DROP COLUMN calendario_evento_id;	