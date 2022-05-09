-- Database generated with pgModeler (PostgreSQL Database Modeler).
-- pgModeler version: 0.9.4
-- PostgreSQL version: 13.0
-- Project Site: pgmodeler.io
-- Model Author: ---
-- object: desarrollooas | type: ROLE --
-- DROP ROLE IF EXISTS desarrollooas;
CREATE ROLE desarrollooas WITH 
	SUPERUSER
	CREATEDB
	CREATEROLE
	INHERIT
	LOGIN
	ENCRYPTED PASSWORD '********'
	VALID UNTIL '2020-10-29 21:19:22';
-- ddl-end --


-- Database creation must be performed outside a multi lined SQL file. 
-- These commands were put in this file only as a convenience.
-- 
-- object: eventos | type: DATABASE --
-- DROP DATABASE IF EXISTS eventos;
CREATE DATABASE eventos
	ENCODING = 'UTF8'
	LC_COLLATE = 'en_US.UTF-8'
	LC_CTYPE = 'en_US.UTF-8'
	TABLESPACE = pg_default
	OWNER = postgres;
-- ddl-end --


-- object: eventos | type: SCHEMA --
-- DROP SCHEMA IF EXISTS eventos CASCADE;
CREATE SCHEMA eventos;
-- ddl-end --
ALTER SCHEMA eventos OWNER TO desarrollooas;
-- ddl-end --

SET search_path TO pg_catalog,public,eventos;
-- ddl-end --

-- object: eventos.migrations_status | type: TYPE --
-- DROP TYPE IF EXISTS eventos.migrations_status CASCADE;
CREATE TYPE eventos.migrations_status AS
ENUM ('update','rollback');
-- ddl-end --
ALTER TYPE eventos.migrations_status OWNER TO postgres;
-- ddl-end --

-- object: eventos.migrations_id_migration_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.migrations_id_migration_seq CASCADE;
CREATE SEQUENCE eventos.migrations_id_migration_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.migrations_id_migration_seq OWNER TO postgres;
-- ddl-end --

-- object: eventos.migrations | type: TABLE --
-- DROP TABLE IF EXISTS eventos.migrations CASCADE;
CREATE TABLE eventos.migrations (
	id_migration integer NOT NULL DEFAULT nextval('eventos.migrations_id_migration_seq'::regclass),
	name character varying(255) DEFAULT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	statements text,
	rollback_statements text,
	status eventos.migrations_status,
	CONSTRAINT migrations_pkey PRIMARY KEY (id_migration)
);
-- ddl-end --
ALTER TABLE eventos.migrations OWNER TO postgres;
-- ddl-end --

-- object: eventos.tipo_evento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.tipo_evento_id_seq CASCADE;
CREATE SEQUENCE eventos.tipo_evento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.tipo_evento_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_evento | type: TABLE --
-- DROP TABLE IF EXISTS eventos.tipo_evento CASCADE;
CREATE TABLE eventos.tipo_evento (
	id integer NOT NULL DEFAULT nextval('eventos.tipo_evento_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	tipo_recurrencia_id integer,
	calendario_id integer,
	CONSTRAINT pk_tipo_evento PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.tipo_evento IS E'Tabla que permite registrar los diferentes tipos de sesion que tienen las diferentes organizaciones';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.id IS E'Identificador de la tabla tipo sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.nombre IS E'Nombre del tipo de sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.descripcion IS E'Descripcion del tipo de sesion que se referencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.codigo_abreviacion IS E'Codigo de abreviación del tipo de sesión que se maneja';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.activo IS E'Booleano que define si el tipo de sesión se encuentra actualmente activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.fecha_creacion IS E'Fecha de creacion de un tipo de evento';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_evento.fecha_modificacion IS E'Fecha de modifiación de un tipo de evento';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_evento ON eventos.tipo_evento  IS E'Contrait para pk de la tabla';
-- ddl-end --
ALTER TABLE eventos.tipo_evento OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.calendario_evento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.calendario_evento_id_seq CASCADE;
CREATE SEQUENCE eventos.calendario_evento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.calendario_evento_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.calendario_evento | type: TABLE --
-- DROP TABLE IF EXISTS eventos.calendario_evento CASCADE;
CREATE TABLE eventos.calendario_evento (
	id integer NOT NULL DEFAULT nextval('eventos.calendario_evento_id_seq'::regclass),
	descripcion character varying(250),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp,
	fecha_inicio timestamp NOT NULL,
	fecha_fin timestamp,
	activo boolean NOT NULL,
	evento_padre_id integer,
	tipo_evento_id integer,
	nombre character varying(50) NOT NULL,
	dependencia_id json,
	CONSTRAINT pk_calendario_evento PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.calendario_evento IS E'Tabla que almacena cada uno de las sesiones y sus datos';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.id IS E'Identificador de la sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.descripcion IS E'Descripción de la sesión qeu se registra';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.fecha_creacion IS E'Fecha de creacion de una sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.fecha_modificacion IS E'Fecha de modifiación del evento para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.fecha_inicio IS E'Fecha de inicio de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.fecha_fin IS E'Fecha de finalización de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.activo IS E'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.nombre IS E'Nombre del tipo de sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento.dependencia_id IS E'se referencia al esquema de dependencia del sistema OKIOS';
-- ddl-end --
COMMENT ON CONSTRAINT pk_calendario_evento ON eventos.calendario_evento  IS E'Restriccion pk de la tabla calendario evento';
-- ddl-end --
ALTER TABLE eventos.calendario_evento OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.encargado_evento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.encargado_evento_id_seq CASCADE;
CREATE SEQUENCE eventos.encargado_evento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.encargado_evento_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.encargado_evento | type: TABLE --
-- DROP TABLE IF EXISTS eventos.encargado_evento CASCADE;
CREATE TABLE eventos.encargado_evento (
	id integer NOT NULL DEFAULT nextval('eventos.encargado_evento_id_seq'::regclass),
	encargado_id integer NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	rol_encargado_evento_id integer,
	calendario_evento_id integer,
	CONSTRAINT pk_encargado_evento PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.encargado_evento IS E'Tabla de rompimiento entre los diferentes encargados y el evento';
-- ddl-end --
COMMENT ON COLUMN eventos.encargado_evento.id IS E'Identificador de la tabla responsable_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.encargado_evento.encargado_id IS E'Se hace referencia al esquema de persona';
-- ddl-end --
COMMENT ON COLUMN eventos.encargado_evento.fecha_creacion IS E'Fecha de creacion de un participante del evento';
-- ddl-end --
COMMENT ON COLUMN eventos.encargado_evento.fecha_modificacion IS E'Fecha de modifiación de un participante del evento';
-- ddl-end --
COMMENT ON COLUMN eventos.encargado_evento.activo IS E'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_encargado_evento ON eventos.encargado_evento  IS E'Identificador de la tabla responsable evento';
-- ddl-end --
ALTER TABLE eventos.encargado_evento OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.rol_encargado_evento_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.rol_encargado_evento_id_seq CASCADE;
CREATE SEQUENCE eventos.rol_encargado_evento_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.rol_encargado_evento_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.rol_encargado_evento | type: TABLE --
-- DROP TABLE IF EXISTS eventos.rol_encargado_evento CASCADE;
CREATE TABLE eventos.rol_encargado_evento (
	id integer NOT NULL DEFAULT nextval('eventos.rol_encargado_evento_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_de_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_rol_encargado_evento PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.rol_encargado_evento IS E'Permite almacenar los diferentes roles que puede tener un encargado en un evento';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.id IS E'Identificador de la tabla rol_participante_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.nombre IS E'Nombre del rol que puede tener un participante de una sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.descripcion IS E'Descripcion que permite especificar el rol y las funciones en la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.codigo_abreviacion IS E'Codigo de abreviacion del nombre del rol que cumple un participante en la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.activo IS E'Booleano que permite identificar si el rol se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.fecha_creacion IS E'Fecha de creacion de un rol participante del evento';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_encargado_evento.fecha_modificacion IS E'Fecha de modifiación de un rol participante del evento';
-- ddl-end --
ALTER TABLE eventos.rol_encargado_evento OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_recurrencia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.tipo_recurrencia_id_seq CASCADE;
CREATE SEQUENCE eventos.tipo_recurrencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.tipo_recurrencia_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_recurrencia | type: TABLE --
-- DROP TABLE IF EXISTS eventos.tipo_recurrencia CASCADE;
CREATE TABLE eventos.tipo_recurrencia (
	id integer NOT NULL DEFAULT nextval('eventos.tipo_recurrencia_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_recurrencia PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.tipo_recurrencia IS E'Tabla que almacena los tipos de recurrencia en tiempo que se tienen (por minuto, hora, diario...)';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.id IS E'Identificador de la tabla tipo_recurrencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.nombre IS E'Nombre del tipo de recurrencia que se registra en la tabla';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.descripcion IS E'Breve descripcion del tipo de recurrencia que se registra';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.codigo_abreviacion IS E'Codigo de abreviacion del tipo de recurrencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.activo IS E'Booleano que permite identificar si el tipo de recurrencia se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.fecha_creacion IS E'Fecha de creacion de un tipo de recurencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_recurrencia.fecha_modificacion IS E'Fecha de modifiación de un tipo de recurrencia';
-- ddl-end --
ALTER TABLE eventos.tipo_recurrencia OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_publico_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.tipo_publico_id_seq CASCADE;
CREATE SEQUENCE eventos.tipo_publico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.tipo_publico_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_publico | type: TABLE --
-- DROP TABLE IF EXISTS eventos.tipo_publico CASCADE;
CREATE TABLE eventos.tipo_publico (
	id integer NOT NULL DEFAULT nextval('eventos.tipo_publico_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	codigo_abreviacion character varying(250),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	CONSTRAINT pk_tipo_publico PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.tipo_publico IS E'Tabla para almacenar el tipo de publico que puede participar en un evento';
-- ddl-end --
ALTER TABLE eventos.tipo_publico OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_sesion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.tipo_sesion_id_seq CASCADE;
CREATE SEQUENCE eventos.tipo_sesion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.tipo_sesion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.tipo_sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.tipo_sesion CASCADE;
CREATE TABLE eventos.tipo_sesion (
	id integer NOT NULL DEFAULT nextval('eventos.tipo_sesion_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_tipo_sesion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_tipo_sesion UNIQUE (nombre)
);
-- ddl-end --
COMMENT ON TABLE eventos.tipo_sesion IS E'Tabla que permite registrar los diferentes tipos de sesion que tienen las diferentes organizaciones';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.id IS E'Identificador de la tabla tipo sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.nombre IS E'Nombre del tipo de sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.descripcion IS E'Descripcion del tipo de sesion que se referencia';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.codigo_abreviacion IS E'Codigo de abreviación del tipo de sesión que se maneja';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.activo IS E'Booleano que define si el tipo de sesión se encuentra actualmente activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.tipo_sesion.numero_orden IS E'Número de orden para la tabla tipo_sesion';
-- ddl-end --
COMMENT ON CONSTRAINT pk_tipo_sesion ON eventos.tipo_sesion  IS E'Contrait para pk de la tabla';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_tipo_sesion ON eventos.tipo_sesion  IS E'UK para garantizar unicidad del nombre del tipo de la sesión';
-- ddl-end --
ALTER TABLE eventos.tipo_sesion OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.sesion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.sesion_id_seq CASCADE;
CREATE SEQUENCE eventos.sesion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.sesion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.sesion CASCADE;
CREATE TABLE eventos.sesion (
	id integer NOT NULL DEFAULT nextval('eventos.sesion_id_seq'::regclass),
	descripcion character varying(250),
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
COMMENT ON TABLE eventos.sesion IS E'Tabla que almacena cada uno de las sesiones y sus datos';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.id IS E'Identificador de la sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.descripcion IS E'Descripción de la sesión qeu se registra';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_creacion IS E'Fecha de creacion de una sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_modificacion IS E'Fecha de modifiación del evento para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_inicio IS E'Fecha de inicio de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.fecha_fin IS E'Fecha de finalización de la sesión para la organización';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.periodo IS E'Se almacena el periodo al cual pertenece el evento';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.recurrente IS E'Booleano qeu permite identificar si la sesion es recurrente o no';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion.numero_recurrencias IS E'Numero de veces que va a ser necesarioqeu se repita la sesión';
-- ddl-end --
COMMENT ON CONSTRAINT pk_sesion ON eventos.sesion  IS E'Restriccion pk de la tabla sesión';
-- ddl-end --
ALTER TABLE eventos.sesion OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.participante_sesion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.participante_sesion_id_seq CASCADE;
CREATE SEQUENCE eventos.participante_sesion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.participante_sesion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.participante_sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.participante_sesion CASCADE;
CREATE TABLE eventos.participante_sesion (
	id integer NOT NULL DEFAULT nextval('eventos.participante_sesion_id_seq'::regclass),
	sesion integer NOT NULL,
	rol_participante_sesion integer NOT NULL,
	ente integer NOT NULL,
	CONSTRAINT pk_participante_sesion PRIMARY KEY (id),
	CONSTRAINT uq_participante_sesion UNIQUE (sesion,rol_participante_sesion,ente)
);
-- ddl-end --
COMMENT ON TABLE eventos.participante_sesion IS E'Tabla de rompimiento entre los difenrentes participantes y la sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.participante_sesion.id IS E'Identificador de la tabla responsable_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.participante_sesion.ente IS E'Columna que referencia el Id de la tabla donde se encuentra el ente.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_participante_sesion ON eventos.participante_sesion  IS E'Identificador de la tabla responsable sesión';
-- ddl-end --
ALTER TABLE eventos.participante_sesion OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.rol_participante_sesion_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.rol_participante_sesion_id_seq CASCADE;
CREATE SEQUENCE eventos.rol_participante_sesion_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.rol_participante_sesion_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.rol_participante_sesion | type: TABLE --
-- DROP TABLE IF EXISTS eventos.rol_participante_sesion CASCADE;
CREATE TABLE eventos.rol_participante_sesion (
	id integer NOT NULL DEFAULT nextval('eventos.rol_participante_sesion_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	codigo_abreviacion character varying(20),
	activo boolean NOT NULL,
	numero_orden numeric(5,2),
	CONSTRAINT pk_rol_participante_sesion PRIMARY KEY (id),
	CONSTRAINT uq_nombre_rol_participante_sesion UNIQUE (nombre)
);
-- ddl-end --
COMMENT ON TABLE eventos.rol_participante_sesion IS E'Permite almacenar los diferentes roles que puede tener un participante en una sesión';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.id IS E'Identificador de la tabla rol_participante_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.nombre IS E'Nombre del rol que puede tener un participante de una sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.descripcion IS E'Descripcion que permite especificar el rol y las funciones en la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.codigo_abreviacion IS E'Codigo de abreviacion del nombre del rol que cumple un participante en la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.activo IS E'Booleano que permite identificar si el rol se encuentra activo o no';
-- ddl-end --
COMMENT ON COLUMN eventos.rol_participante_sesion.numero_orden IS E'Numero de orden para el rol_participante_sesion';
-- ddl-end --
COMMENT ON CONSTRAINT uq_nombre_rol_participante_sesion ON eventos.rol_participante_sesion  IS E'UK para garantizar unicidad del nombre';
-- ddl-end --
ALTER TABLE eventos.rol_participante_sesion OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.sesion_patron_recurrencia_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.sesion_patron_recurrencia_id_seq CASCADE;
CREATE SEQUENCE eventos.sesion_patron_recurrencia_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.sesion_patron_recurrencia_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.sesion_patron_recurrencia | type: TABLE --
-- DROP TABLE IF EXISTS eventos.sesion_patron_recurrencia CASCADE;
CREATE TABLE eventos.sesion_patron_recurrencia (
	id integer NOT NULL DEFAULT nextval('eventos.sesion_patron_recurrencia_id_seq'::regclass),
	tipo_recurrencia integer NOT NULL,
	sesion integer NOT NULL,
	valor character varying(100) NOT NULL,
	CONSTRAINT pk_sesion_patron_recurrencia PRIMARY KEY (id),
	CONSTRAINT uq_sesion_patron_recurrencia UNIQUE (tipo_recurrencia,sesion)
);
-- ddl-end --
COMMENT ON TABLE eventos.sesion_patron_recurrencia IS E'Permite almacenar los valores del patron de recurrencia de la sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion_patron_recurrencia.id IS E'Identificador de la tabla';
-- ddl-end --
COMMENT ON COLUMN eventos.sesion_patron_recurrencia.valor IS E'Valor que se establece para las recurrencias';
-- ddl-end --
ALTER TABLE eventos.sesion_patron_recurrencia OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.relacion_sesiones_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.relacion_sesiones_id_seq CASCADE;
CREATE SEQUENCE eventos.relacion_sesiones_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.relacion_sesiones_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.relacion_sesiones | type: TABLE --
-- DROP TABLE IF EXISTS eventos.relacion_sesiones CASCADE;
CREATE TABLE eventos.relacion_sesiones (
	id integer NOT NULL DEFAULT nextval('eventos.relacion_sesiones_id_seq'::regclass),
	sesion_padre integer NOT NULL,
	sesion_hijo integer NOT NULL,
	CONSTRAINT pk_relacion_sesiones PRIMARY KEY (id),
	CONSTRAINT uq_sesion_padre_sesion_hijo UNIQUE (sesion_padre,sesion_hijo)
);
-- ddl-end --
COMMENT ON TABLE eventos.relacion_sesiones IS E'Tabla que almacena la relación entre las sesión padre y la sesión hijo.';
-- ddl-end --
COMMENT ON COLUMN eventos.relacion_sesiones.id IS E'Identificador de la tabla relacion_sesiones';
-- ddl-end --
COMMENT ON COLUMN eventos.relacion_sesiones.sesion_hijo IS E'Hijo de la relacion de sesiones';
-- ddl-end --
ALTER TABLE eventos.relacion_sesiones OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.calendario_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.calendario_id_seq CASCADE;
CREATE SEQUENCE eventos.calendario_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.calendario_id_seq OWNER TO postgres;
-- ddl-end --

-- object: eventos.calendario | type: TABLE --
-- DROP TABLE IF EXISTS eventos.calendario CASCADE;
CREATE TABLE eventos.calendario (
	id integer NOT NULL DEFAULT nextval('eventos.calendario_id_seq'::regclass),
	nombre character varying(50) NOT NULL,
	descripcion character varying(250),
	dependencia_id json NOT NULL,
	documento_id integer NOT NULL,
	periodo_id integer NOT NULL,
	aplicacion_id integer NOT NULL,
	nivel integer NOT NULL,
	activo boolean NOT NULL,
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	calendario_padre_id integer,
	documento_extension_id integer,
	aplica_extension boolean,
	aplica_edicion_actividades boolean,
	dependencia_particular_id json,
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
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.documento_extension_id IS E'Campo id para documento soporte si hay extensión';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.aplica_extension IS E'Campo para verificar si es calendario por extensión';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.aplica_edicion_actividades IS E'Campo para verficar si al calendario aplica edicion de fechas actividades';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario.dependencia_particular_id IS E'Campo json para listar dependencias con ajustes particulares de fechas';
-- ddl-end --
ALTER TABLE eventos.calendario OWNER TO postgres;
-- ddl-end --

-- object: eventos.calendario_evento_tipo_publico_id_seq | type: SEQUENCE --
-- DROP SEQUENCE IF EXISTS eventos.calendario_evento_tipo_publico_id_seq CASCADE;
CREATE SEQUENCE eventos.calendario_evento_tipo_publico_id_seq
	INCREMENT BY 1
	MINVALUE 1
	MAXVALUE 9223372036854775807
	START WITH 1
	CACHE 1
	NO CYCLE
	OWNED BY NONE;

-- ddl-end --
ALTER SEQUENCE eventos.calendario_evento_tipo_publico_id_seq OWNER TO desarrollooas;
-- ddl-end --

-- object: eventos.calendario_evento_tipo_publico | type: TABLE --
-- DROP TABLE IF EXISTS eventos.calendario_evento_tipo_publico CASCADE;
CREATE TABLE eventos.calendario_evento_tipo_publico (
	id integer NOT NULL DEFAULT nextval('eventos.calendario_evento_tipo_publico_id_seq'::regclass),
	fecha_creacion timestamp NOT NULL,
	fecha_modificacion timestamp NOT NULL,
	activo boolean NOT NULL,
	calendario_evento_id integer,
	tipo_publico_id integer,
	CONSTRAINT pk_calendario_evento_tipo_publico PRIMARY KEY (id)
);
-- ddl-end --
COMMENT ON TABLE eventos.calendario_evento_tipo_publico IS E'Tabla de rompimiento entre los diferentes eventos de calendarios y el tipos de publico';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.id IS E'Identificador de la tabla responsable_sesion';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.fecha_creacion IS E'Fecha de creacion de un participante del evento';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.fecha_modificacion IS E'Fecha de modifiación de un participante del evento';
-- ddl-end --
COMMENT ON COLUMN eventos.calendario_evento_tipo_publico.activo IS E'Campo para registrar si el  registro se encuentra activo o no, solo a nivel de registro.';
-- ddl-end --
COMMENT ON CONSTRAINT pk_calendario_evento_tipo_publico ON eventos.calendario_evento_tipo_publico  IS E'Identificador de la tabla responsable evento';
-- ddl-end --
ALTER TABLE eventos.calendario_evento_tipo_publico OWNER TO desarrollooas;
-- ddl-end --

-- object: fk_tipo_evento_tipo_recurrencia | type: CONSTRAINT --
-- ALTER TABLE eventos.tipo_evento DROP CONSTRAINT IF EXISTS fk_tipo_evento_tipo_recurrencia CASCADE;
ALTER TABLE eventos.tipo_evento ADD CONSTRAINT fk_tipo_evento_tipo_recurrencia FOREIGN KEY (tipo_recurrencia_id)
REFERENCES eventos.tipo_recurrencia (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_calendario | type: CONSTRAINT --
-- ALTER TABLE eventos.tipo_evento DROP CONSTRAINT IF EXISTS fk_calendario CASCADE;
ALTER TABLE eventos.tipo_evento ADD CONSTRAINT fk_calendario FOREIGN KEY (calendario_id)
REFERENCES eventos.calendario (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_calendario_evento_tipo_evento | type: CONSTRAINT --
-- ALTER TABLE eventos.calendario_evento DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_evento CASCADE;
ALTER TABLE eventos.calendario_evento ADD CONSTRAINT fk_calendario_evento_tipo_evento FOREIGN KEY (tipo_evento_id)
REFERENCES eventos.tipo_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_evento_padre | type: CONSTRAINT --
-- ALTER TABLE eventos.calendario_evento DROP CONSTRAINT IF EXISTS fk_evento_padre CASCADE;
ALTER TABLE eventos.calendario_evento ADD CONSTRAINT fk_evento_padre FOREIGN KEY (evento_padre_id)
REFERENCES eventos.calendario_evento (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --

-- object: fk_encargado_evento_rol_encargado_evento | type: CONSTRAINT --
-- ALTER TABLE eventos.encargado_evento DROP CONSTRAINT IF EXISTS fk_encargado_evento_rol_encargado_evento CASCADE;
ALTER TABLE eventos.encargado_evento ADD CONSTRAINT fk_encargado_evento_rol_encargado_evento FOREIGN KEY (rol_encargado_evento_id)
REFERENCES eventos.rol_encargado_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_encargado_evento_calendario_evento | type: CONSTRAINT --
-- ALTER TABLE eventos.encargado_evento DROP CONSTRAINT IF EXISTS fk_encargado_evento_calendario_evento CASCADE;
ALTER TABLE eventos.encargado_evento ADD CONSTRAINT fk_encargado_evento_calendario_evento FOREIGN KEY (calendario_evento_id)
REFERENCES eventos.calendario_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_sesion_tipo_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.sesion DROP CONSTRAINT IF EXISTS fk_sesion_tipo_sesion CASCADE;
ALTER TABLE eventos.sesion ADD CONSTRAINT fk_sesion_tipo_sesion FOREIGN KEY (tipo_sesion)
REFERENCES eventos.tipo_sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_participante_sesion_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS fk_participante_sesion_sesion CASCADE;
ALTER TABLE eventos.participante_sesion ADD CONSTRAINT fk_participante_sesion_sesion FOREIGN KEY (sesion)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_participante_sesion_rol_participante_sesion | type: CONSTRAINT --
-- ALTER TABLE eventos.participante_sesion DROP CONSTRAINT IF EXISTS fk_participante_sesion_rol_participante_sesion CASCADE;
ALTER TABLE eventos.participante_sesion ADD CONSTRAINT fk_participante_sesion_rol_participante_sesion FOREIGN KEY (rol_participante_sesion)
REFERENCES eventos.rol_participante_sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
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

-- object: fk_relacion_sesiones_sesion_padre | type: CONSTRAINT --
-- ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS fk_relacion_sesiones_sesion_padre CASCADE;
ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT fk_relacion_sesiones_sesion_padre FOREIGN KEY (sesion_padre)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_relacion_sesiones_sesion_hijo | type: CONSTRAINT --
-- ALTER TABLE eventos.relacion_sesiones DROP CONSTRAINT IF EXISTS fk_relacion_sesiones_sesion_hijo CASCADE;
ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT fk_relacion_sesiones_sesion_hijo FOREIGN KEY (sesion_hijo)
REFERENCES eventos.sesion (id) MATCH FULL
ON DELETE RESTRICT ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_calendario_padre | type: CONSTRAINT --
-- ALTER TABLE eventos.calendario DROP CONSTRAINT IF EXISTS fk_calendario_padre CASCADE;
ALTER TABLE eventos.calendario ADD CONSTRAINT fk_calendario_padre FOREIGN KEY (calendario_padre_id)
REFERENCES eventos.calendario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;
-- ddl-end --
COMMENT ON CONSTRAINT fk_calendario_padre ON eventos.calendario  IS E'FK del calendario';
-- ddl-end --


-- object: fk_calendario_evento_tipo_publico_calendario_evento | type: CONSTRAINT --
-- ALTER TABLE eventos.calendario_evento_tipo_publico DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_publico_calendario_evento CASCADE;
ALTER TABLE eventos.calendario_evento_tipo_publico ADD CONSTRAINT fk_calendario_evento_tipo_publico_calendario_evento FOREIGN KEY (calendario_evento_id)
REFERENCES eventos.calendario_evento (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --

-- object: fk_calendario_evento_tipo_publico_tipo_publico | type: CONSTRAINT --
-- ALTER TABLE eventos.calendario_evento_tipo_publico DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_publico_tipo_publico CASCADE;
ALTER TABLE eventos.calendario_evento_tipo_publico ADD CONSTRAINT fk_calendario_evento_tipo_publico_tipo_publico FOREIGN KEY (tipo_publico_id)
REFERENCES eventos.tipo_publico (id) MATCH FULL
ON DELETE SET NULL ON UPDATE CASCADE;
-- ddl-end --


