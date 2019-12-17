
CREATE TABLE eventos.tipo_sesion(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_sesion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_sesion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE eventos.tipo_sesion IS 'Tabla que permite registrar los diferentes tipos de sesion que tienen las diferentes organizaciones';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.id IS 'Identificador de la tabla tipo sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.nombre IS 'Nombre del tipo de sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.descripcion IS 'Descripcion del tipo de sesion que se referencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.codigo_abreviacion IS 'Codigo de abreviación del tipo de sesión que se maneja';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.activo IS 'Booleano que define si el tipo de sesión se encuentra actualmente activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.numero_orden IS 'Número de orden para la tabla tipo_sesion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_sesion ON eventos.tipo_sesion  IS 'Contrait para pk de la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_sesion ON eventos.tipo_sesion  IS 'UK para garantizar unicidad del nombre del tipo de la sesión';
-- ddl-end --
ALTER TABLE eventos.tipo_sesion OWNER TO postgres;
-- ddl-end --

-- object: eventos.sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.sesion CASCADE;
CREATE TABLE eventos.sesion(
	id serial NOT NULL,
	descripcion varchar(250),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	periodo integer,
	recurrente boolean NOT NULL,
	numero_recurrencias integer,
	tipo_sesion integer NOT NULL,
	CONSTRAINT pk_sesion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE eventos.sesion IS 'Tabla que almacena cada uno de las sesiones y sus datos';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.id IS 'Identificador de la sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.descripcion IS 'Descripción de la sesión qeu se registra';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_creacion IS 'Fecha de creacion de una sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_modificacion IS 'Fecha de modifiación del evento para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_inicio IS 'Fecha de inicio de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_fin IS 'Fecha de finalización de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.periodo IS 'Se almacena el periodo al cual pertenece el evento';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.recurrente IS 'Booleano qeu permite identificar si la sesion es recurrente o no';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.numero_recurrencias IS 'Numero de veces que va a ser necesarioqeu se repita la sesión';
-- ddl-end --
COMMENT ON CONSTRAINT pk_sesion ON eventos.sesion  IS 'Restriccion pk de la tabla sesión';
-- ddl-end --
ALTER TABLE eventos.sesion OWNER TO postgres;
-- ddl-end --

-- object: fk_sesion_tipo_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.sesion DROP CONSTRAINT IF EXISTS fk_sesion_tipo_sesion CASCADE;
ALTER TABLE eventos.sesion ADD CONSTRAINT fk_sesion_tipo_sesion FOREIGN KEY (tipo_sesion)
REFERENCES eventos.tipo_sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: eventos.participante_sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.participante_sesion CASCADE;
CREATE TABLE eventos.participante_sesion(
	id serial NOT NULL,
	sesion integer NOT NULL,
	rol_participante_sesion integer NOT NULL,
	ente integer NOT NULL,
	CONSTRAINT pk_participante_sesion PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE eventos.participante_sesion IS 'Tabla de rompimiento entre los difenrentes participantes y la sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.participante_sesion.id IS 'Identificador de la tabla responsable_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.participante_sesion.ente IS 'Columna que referencia el Id de la tabla donde se encuentra el ente.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_participante_sesion ON eventos.participante_sesion  IS 'Identificador de la tabla responsable sesión';
-- ddl-end --
ALTER TABLE eventos.participante_sesion OWNER TO postgres;
-- ddl-end --

-- object: fk_participante_sesion_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS fk_participante_sesion_sesion CASCADE;
ALTER TABLE eventos.participante_sesion ADD CONSTRAINT fk_participante_sesion_sesion FOREIGN KEY (sesion)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: eventos.rol_participante_sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.rol_participante_sesion CASCADE;
CREATE TABLE eventos.rol_participante_sesion(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_rol_participante_sesion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_rol_participante_sesion UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE eventos.rol_participante_sesion IS 'Permite almacenar los diferentes roles que puede tener un participante en una sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.id IS 'Identificador de la tabla rol_participante_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.nombre IS 'Nombre del rol que puede tener un participante de una sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.descripcion IS 'Descripcion que permite especificar el rol y las funciones en la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.codigo_abreviacion IS 'Codigo de abreviacion del nombre del rol que cumple un participante en la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.activo IS 'Booleano que permite identificar si el rol se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.numero_orden IS 'Numero de orden para el rol_participante_sesion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_rol_participante_sesion ON eventos.rol_participante_sesion  IS 'UK para garantizar unicidad del nombre';
-- ddl-end --
ALTER TABLE eventos.rol_participante_sesion OWNER TO postgres;
-- ddl-end --

-- object: fk_participante_sesion_rol_participante_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS fk_participante_sesion_rol_participante_sesion CASCADE;
ALTER TABLE eventos.participante_sesion ADD CONSTRAINT fk_participante_sesion_rol_participante_sesion FOREIGN KEY (rol_participante_sesion)
REFERENCES eventos.rol_participante_sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: eventos.tipo_recurrencia | type: TABLE --
-- DROP TABLE IF EXISTS eventos.tipo_recurrencia CASCADE;
CREATE TABLE eventos.tipo_recurrencia(
	id serial NOT NULL,
	nombre varchar(50) NOT NULL,
	descripcion varchar(250),
	codigo_abreviacion varchar(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_recurrencia PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_recurrencia UNIQUE (nombre)

);
-- ddl-end --
COMMENT ON TABLE eventos.tipo_recurrencia IS 'Tabla que almacena los tipos de recurrencia en tiempo que se tienen (por minuto, hora, diario...)';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.id IS 'Identificador de la tabla tipo_recurrencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.nombre IS 'Nombre del tipo de recurrencia que se registra en la tabla';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.descripcion IS 'Breve descripcion del tipo de recurrencia que se registra';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.codigo_abreviacion IS 'Codigo de abreviacion del tipo de recurrencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.activo IS 'Booleano que permite identificar si el tipo de recurrencia se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.numero_orden IS 'Número de orden para la tabla tipo de recurrencia';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_recurrencia ON eventos.tipo_recurrencia  IS 'UK para garantizar unicidad del nombre del tipo de recurrencia';
-- ddl-end --
ALTER TABLE eventos.tipo_recurrencia OWNER TO postgres;
-- ddl-end --

-- object: eventos.sesion_patron_recurrencia | type: TABLE --
-- DROP TABLE IF EXISTS eventos.sesion_patron_recurrencia CASCADE;
CREATE TABLE eventos.sesion_patron_recurrencia(
	id serial NOT NULL,
	tipo_recurrencia integer NOT NULL,
	sesion integer NOT NULL,
	valor varchar(100) NOT NULL,
	CONSTRAINT pk_sesion_patron_recurrencia PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE eventos.sesion_patron_recurrencia IS 'Permite almacenar los valores del patron de recurrencia de la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion_patron_recurrencia.id IS 'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion_patron_recurrencia.valor IS 'Valor que se establece para las recurrencias ';
-- ddl-end --
ALTER TABLE eventos.sesion_patron_recurrencia OWNER TO postgres;
-- ddl-end --

-- object: fk_sesion_patron_recurrencia_tipo_recurrencia | type: CONSTRAINT --
-- ALTER TABLE eventos.sesion_patron_recurrencia DROP CONSTRAINT IF EXISTS fk_sesion_patron_recurrencia_tipo_recurrencia CASCADE;
ALTER TABLE eventos.sesion_patron_recurrencia ADD CONSTRAINT fk_sesion_patron_recurrencia_tipo_recurrencia FOREIGN KEY (tipo_recurrencia)
REFERENCES eventos.tipo_recurrencia (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_sesion_patron_recurrencia_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.sesion_patron_recurrencia DROP CONSTRAINT IF EXISTS fk_sesion_patron_recurrencia_sesion CASCADE;
ALTER TABLE eventos.sesion_patron_recurrencia ADD CONSTRAINT fk_sesion_patron_recurrencia_sesion FOREIGN KEY (sesion)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: eventos.relacion_sesiones | type: TABLE --
-- DROP TABLE IF EXISTS eventos.relacion_sesiones CASCADE;
CREATE TABLE eventos.relacion_sesiones(
	id serial NOT NULL,
	sesion_padre integer NOT NULL,
	sesion_hijo integer NOT NULL,
	CONSTRAINT pk_relacion_sesiones PRIMARY KEY (id)

);
-- ddl-end --
COMMENT ON TABLE eventos.relacion_sesiones IS 'Tabla que almacena la relación entre las sesión padre y la sesión hijo.';
-- ddl-end --
COMMENT ON COLUMN eventos.relacion_sesiones.id IS 'Identificador de la tabla relacion_sesiones';
-- ddl-end --
COMMENT ON COLUMN eventos.relacion_sesiones.sesion_hijo IS 'Hijo de la relacion de sesiones';
-- ddl-end --
ALTER TABLE eventos.relacion_sesiones OWNER TO postgres;
-- ddl-end --

-- object: fk_relacion_sesiones_sesion_padre | type: CONSTRAINT --
-- ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS fk_relacion_sesiones_sesion_padre CASCADE;
ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT fk_relacion_sesiones_sesion_padre FOREIGN KEY (sesion_padre)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: uq_participante_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS uq_participante_sesion CASCADE;
ALTER TABLE eventos.participante_sesion ADD CONSTRAINT uq_participante_sesion UNIQUE (sesion,rol_participante_sesion,ente);
-- ddl-end --

-- object: uq_sesion_padre_sesion_hijo | type: CONSTRAINT --
-- ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS uq_sesion_padre_sesion_hijo CASCADE;
ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT uq_sesion_padre_sesion_hijo UNIQUE (sesion_padre,sesion_hijo);
-- ddl-end --

-- object: uq_sesion_patron_recurrencia | type: CONSTRAINT --
-- ALTER TABLE eventos.sesion_patron_recurrencia DROP CONSTRAINT IF EXISTS uq_sesion_patron_recurrencia CASCADE;
ALTER TABLE eventos.sesion_patron_recurrencia ADD CONSTRAINT uq_sesion_patron_recurrencia UNIQUE (tipo_recurrencia,sesion);
-- ddl-end --

-- object: fk_relacion_sesiones_sesion_hijo | type: CONSTRAINT --
-- ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS fk_relacion_sesiones_sesion_hijo CASCADE;
ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT fk_relacion_sesiones_sesion_hijo FOREIGN KEY (sesion_hijo)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --