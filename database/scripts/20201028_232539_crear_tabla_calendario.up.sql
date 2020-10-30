-- object: eventos.calendario | type: TABLE --
-- DROP TABLE IF EXISTS eventos.calendario;
CREATE TABLE eventos.calendario (
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	dependencia_id json NOT NULL,
	documento_id integer NOT NULL,
	periodo_id integer NOT NULL,
	aplicacion_id integer NOT NULL,
	nivel integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	calendario_padre_id integer,
	CONSTRAINT pk_calendario PRIMARY KEY (id)

);

-- ddl-end --
COMMENT ON TABLE eventos.calendario IS E'tabla calendario';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.id IS E'id tabla calendario';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.nombre IS E'nombre parametria';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.descripcion IS E'descripcion parametrica';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.dependencia_id IS E'se referencia al esquema de dependencia del sistema OKIOS';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.documento_id IS E'Documento';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.periodo_id IS E'Se almacena el periodo al cual pertenece el eventos';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.aplicacion_id IS E'Campo obligatorio para realcionar el eventos creado a una aplicacion. Es un id relacionado a la tabla de aplicacion en el esquema configuracion';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.nivel IS E'nivel';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.activo IS E'estado parametrica';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.fecha_creacion IS E'fecha creacion del registro';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.fecha_modificacion IS E'fecha modificacion del registro';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.calendario_padre_id IS E'Calendario padre id';


ALTER TABLE eventos.calendario ADD CONSTRAINT fk_calendario_padre FOREIGN KEY (calendario_padre_id)
REFERENCES eventos.calendario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
COMMENT ON CONSTRAINT fk_calendario_padre ON eventos.calendario  IS E'FK del calendario';
-- ddl-end --