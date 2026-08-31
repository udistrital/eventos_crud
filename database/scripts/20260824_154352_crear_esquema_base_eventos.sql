-- Script para crear el esquema base de eventos.

CREATE SCHEMA eventos;

CREATE TABLE eventos.tipo_recurrencia (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    descripcion VARCHAR(250),
    codigo_abreviacion VARCHAR(20),
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE TABLE eventos.tipo_sesion (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    descripcion VARCHAR(250),
    codigo_abreviacion VARCHAR(20),
    activo BOOLEAN NOT NULL,
    numero_orden NUMERIC(5, 2),
    CONSTRAINT uq_nombre_tipo_sesion UNIQUE (nombre)
);

CREATE TABLE eventos.rol_participante_sesion (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    descripcion VARCHAR(250),
    codigo_abreviacion VARCHAR(20),
    activo BOOLEAN NOT NULL,
    numero_orden NUMERIC(5, 2),
    CONSTRAINT uq_nombre_rol_participante_sesion UNIQUE (nombre)
);

CREATE TABLE eventos.calendario (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(50) NOT NULL,
    descripcion VARCHAR(250),
    dependencia_id JSON NOT NULL,
    documento_id INTEGER NOT NULL,
    periodo_id INTEGER NOT NULL,
    aplicacion_id INTEGER NOT NULL,
    nivel INTEGER NOT NULL,
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    CONSTRAINT ck_calendario_dependencia_estructura
        CHECK (
            JSONB_TYPEOF(dependencia_id::JSONB) = 'object'
            AND dependencia_id::JSONB ? 'proyectos'
            AND JSONB_TYPEOF(dependencia_id::JSONB -> 'proyectos') = 'array'
        )
);

CREATE TABLE eventos.proceso_catalogo (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(50) NOT NULL,
    nombre VARCHAR(150) NOT NULL,
    descripcion VARCHAR(250),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_proceso_catalogo_codigo UNIQUE (codigo_abreviacion)
);

CREATE TABLE eventos.evento_catalogo (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(50) NOT NULL,
    nombre VARCHAR(150) NOT NULL,
    descripcion VARCHAR(250),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_evento_catalogo_codigo UNIQUE (codigo_abreviacion)
);

CREATE TABLE eventos.proceso (
    id SERIAL PRIMARY KEY,
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    tipo_recurrencia_id INTEGER NOT NULL,
    calendario_id INTEGER NOT NULL,
    proceso_catalogo_id INTEGER NOT NULL,
    CONSTRAINT fk_proceso_tipo_recurrencia
        FOREIGN KEY (tipo_recurrencia_id)
        REFERENCES eventos.tipo_recurrencia(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_proceso_calendario
        FOREIGN KEY (calendario_id)
        REFERENCES eventos.calendario(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_proceso_proceso_catalogo
        FOREIGN KEY (proceso_catalogo_id)
        REFERENCES eventos.proceso_catalogo(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT
);

CREATE TABLE eventos.evento_catalogo_proceso_catalogo (
    id SERIAL PRIMARY KEY,
    evento_catalogo_id INTEGER NOT NULL,
    proceso_catalogo_id INTEGER NOT NULL,
    repetible BOOLEAN NOT NULL DEFAULT FALSE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_evento_catalogo_proceso_catalogo_evento
        FOREIGN KEY (evento_catalogo_id)
        REFERENCES eventos.evento_catalogo(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_evento_catalogo_proceso_catalogo_proceso
        FOREIGN KEY (proceso_catalogo_id)
        REFERENCES eventos.proceso_catalogo(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_evento_catalogo_proceso_catalogo
        UNIQUE (evento_catalogo_id, proceso_catalogo_id)
);

CREATE TABLE eventos.evento_catalogo_rol_gestion (
    id SERIAL PRIMARY KEY,
    evento_catalogo_id INTEGER NOT NULL,
    perfil_id INTEGER NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_evento_catalogo_rol_gestion_evento
        FOREIGN KEY (evento_catalogo_id)
        REFERENCES eventos.evento_catalogo(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_evento_catalogo_rol_gestion_perfil
        UNIQUE (evento_catalogo_id, perfil_id)
);

CREATE TABLE eventos.calendario_evento (
    id SERIAL PRIMARY KEY,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_inicio TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_fin TIMESTAMP WITHOUT TIME ZONE,
    activo BOOLEAN NOT NULL,
    dependencia_id JSON,
    proceso_id INTEGER NOT NULL,
    evento_catalogo_id INTEGER NOT NULL,
    numero_ocurrencia INTEGER NOT NULL DEFAULT 1,
    ubicacion_id INTEGER NOT NULL DEFAULT 0,
    poster_url VARCHAR(250) NOT NULL DEFAULT '',
    CONSTRAINT fk_calendario_evento_proceso
        FOREIGN KEY (proceso_id)
        REFERENCES eventos.proceso(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_calendario_evento_evento_catalogo
        FOREIGN KEY (evento_catalogo_id)
        REFERENCES eventos.evento_catalogo(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_calendario_evento_dependencia_estructura
        CHECK (
            dependencia_id IS NULL
            OR (
                JSONB_TYPEOF(dependencia_id::JSONB) = 'object'
                AND dependencia_id::JSONB ? 'proyectos'
                AND JSONB_TYPEOF(dependencia_id::JSONB -> 'proyectos') = 'array'
                AND (
                    NOT dependencia_id::JSONB ? 'fechas'
                    OR JSONB_TYPEOF(dependencia_id::JSONB -> 'fechas') = 'array'
                )
            )
        ),
    CONSTRAINT ck_calendario_evento_rango_fechas
        CHECK (
            fecha_fin IS NULL
            OR fecha_fin >= fecha_inicio
        ),
    CONSTRAINT ck_calendario_evento_numero_ocurrencia
        CHECK (
            numero_ocurrencia > 0
        )
);

CREATE TABLE eventos.calendario_evento_tipo_publico (
    id SERIAL PRIMARY KEY,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    calendario_evento_id INTEGER NOT NULL,
    perfil_id INTEGER NOT NULL,
    CONSTRAINT fk_calendario_evento_tipo_publico_evento
        FOREIGN KEY (calendario_evento_id)
        REFERENCES eventos.calendario_evento(id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT uq_calendario_evento_tipo_publico
        UNIQUE (calendario_evento_id, perfil_id)
);

CREATE TABLE eventos.calendario_evento_extension (
    id SERIAL PRIMARY KEY,
    calendario_evento_id INTEGER NOT NULL,
    fecha_fin TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    documento_id INTEGER,
    descripcion TEXT,
    numero_extension INTEGER NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_calendario_evento_extension_evento
        FOREIGN KEY (calendario_evento_id)
        REFERENCES eventos.calendario_evento(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_calendario_evento_extension_numero
        UNIQUE (calendario_evento_id, numero_extension),
    CONSTRAINT ck_calendario_evento_extension_numero
        CHECK (numero_extension > 0)
);

CREATE TABLE eventos.calendario_evento_extension_programa (
    id SERIAL PRIMARY KEY,
    calendario_evento_extension_id INTEGER NOT NULL,
    extension_padre_id INTEGER,
    dependencia_id INTEGER NOT NULL,
    vigente BOOLEAN NOT NULL DEFAULT TRUE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_calendario_evento_extension_programa_extension
        FOREIGN KEY (calendario_evento_extension_id)
        REFERENCES eventos.calendario_evento_extension(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_calendario_evento_extension_programa_padre
        FOREIGN KEY (extension_padre_id)
        REFERENCES eventos.calendario_evento_extension(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_extension_programa_vigencia
        CHECK (NOT vigente OR activo),
    CONSTRAINT ck_extension_programa_padre_distinto
        CHECK (
            extension_padre_id IS NULL
            OR extension_padre_id <> calendario_evento_extension_id
        )
);

CREATE TABLE eventos.sesion (
    id SERIAL PRIMARY KEY,
    descripcion VARCHAR(250),
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_inicio TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    fecha_fin TIMESTAMP WITHOUT TIME ZONE,
    periodo INTEGER,
    recurrente BOOLEAN NOT NULL,
    numero_recurrencias INTEGER,
    tipo_sesion INTEGER NOT NULL,
    CONSTRAINT fk_sesion_tipo_sesion
        FOREIGN KEY (tipo_sesion)
        REFERENCES eventos.tipo_sesion(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT ck_sesion_rango_fechas
        CHECK (
            fecha_fin IS NULL
            OR fecha_fin >= fecha_inicio
        )
);

CREATE TABLE eventos.sesion_patron_recurrencia (
    id SERIAL PRIMARY KEY,
    tipo_recurrencia INTEGER NOT NULL,
    sesion INTEGER NOT NULL,
    valor VARCHAR(100) NOT NULL,
    CONSTRAINT fk_sesion_patron_tipo_recurrencia
        FOREIGN KEY (tipo_recurrencia)
        REFERENCES eventos.tipo_recurrencia(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_sesion_patron_sesion
        FOREIGN KEY (sesion)
        REFERENCES eventos.sesion(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_sesion_patron_recurrencia
        UNIQUE (tipo_recurrencia, sesion)
);

CREATE TABLE eventos.participante_sesion (
    id SERIAL PRIMARY KEY,
    sesion INTEGER NOT NULL,
    rol_participante_sesion INTEGER NOT NULL,
    ente INTEGER NOT NULL,
    CONSTRAINT fk_participante_sesion_sesion
        FOREIGN KEY (sesion)
        REFERENCES eventos.sesion(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_participante_sesion_rol
        FOREIGN KEY (rol_participante_sesion)
        REFERENCES eventos.rol_participante_sesion(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_participante_sesion
        UNIQUE (sesion, rol_participante_sesion, ente)
);

CREATE TABLE eventos.relacion_sesiones (
    id SERIAL PRIMARY KEY,
    sesion_padre INTEGER NOT NULL,
    sesion_hijo INTEGER NOT NULL,
    CONSTRAINT fk_relacion_sesiones_padre
        FOREIGN KEY (sesion_padre)
        REFERENCES eventos.sesion(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT fk_relacion_sesiones_hijo
        FOREIGN KEY (sesion_hijo)
        REFERENCES eventos.sesion(id)
        ON UPDATE CASCADE
        ON DELETE RESTRICT,
    CONSTRAINT uq_sesion_padre_sesion_hijo
        UNIQUE (sesion_padre, sesion_hijo),
    CONSTRAINT ck_relacion_sesiones_distintas
        CHECK (sesion_padre <> sesion_hijo)
);

CREATE TABLE eventos.calendario_evento_auditoria (
    id BIGSERIAL PRIMARY KEY,
    calendario_evento_id INTEGER NOT NULL,
    operacion TEXT NOT NULL,
    tercero_id INTEGER NOT NULL,
    fecha_creacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    estado_anterior JSONB,
    estado_nuevo JSONB NOT NULL,
    cambios JSONB NOT NULL DEFAULT '{}'::JSONB,
    CONSTRAINT fk_auditoria_calendario_evento
        FOREIGN KEY (calendario_evento_id)
        REFERENCES eventos.calendario_evento(id)
        ON UPDATE RESTRICT
        ON DELETE RESTRICT,
    CONSTRAINT ck_auditoria_estado_anterior
        CHECK (
            (operacion = 'INSERT' AND estado_anterior IS NULL)
            OR (operacion <> 'INSERT' AND estado_anterior IS NOT NULL)
        ),
    CONSTRAINT ck_auditoria_fechas_append_only
        CHECK (fecha_modificacion = fecha_creacion),
    CONSTRAINT ck_auditoria_activo
        CHECK (activo IS TRUE),
    CONSTRAINT ck_auditoria_tercero
        CHECK (tercero_id > 0)
);

CREATE UNIQUE INDEX uq_calendario_periodo_nivel_activo
    ON eventos.calendario (periodo_id, nivel)
    WHERE activo IS TRUE;

CREATE UNIQUE INDEX uq_proceso_calendario_catalogo_activo
    ON eventos.proceso (calendario_id, proceso_catalogo_id)
    WHERE activo IS TRUE;

CREATE UNIQUE INDEX uq_calendario_evento_ocurrencia_activa
    ON eventos.calendario_evento (
        proceso_id,
        evento_catalogo_id,
        numero_ocurrencia
    )
    WHERE activo IS TRUE;

CREATE UNIQUE INDEX uq_calendario_evento_rango_activo_cerrado
    ON eventos.calendario_evento (
        proceso_id,
        evento_catalogo_id,
        fecha_inicio,
        fecha_fin
    )
    WHERE activo IS TRUE
      AND fecha_fin IS NOT NULL;

CREATE UNIQUE INDEX uq_calendario_evento_rango_activo_abierto
    ON eventos.calendario_evento (
        proceso_id,
        evento_catalogo_id,
        fecha_inicio
    )
    WHERE activo IS TRUE
      AND fecha_fin IS NULL;

CREATE UNIQUE INDEX uq_extension_programa_dependencia_activa
    ON eventos.calendario_evento_extension_programa (
        calendario_evento_extension_id,
        dependencia_id
    )
    WHERE activo IS TRUE;

CREATE INDEX idx_proceso_calendario
    ON eventos.proceso (calendario_id);

CREATE INDEX idx_proceso_tipo_recurrencia
    ON eventos.proceso (tipo_recurrencia_id);

CREATE INDEX idx_proceso_catalogo
    ON eventos.proceso (proceso_catalogo_id);

CREATE INDEX idx_evento_catalogo_proceso_proceso
    ON eventos.evento_catalogo_proceso_catalogo (proceso_catalogo_id);

CREATE INDEX idx_evento_catalogo_rol_perfil
    ON eventos.evento_catalogo_rol_gestion (perfil_id);

CREATE INDEX idx_calendario_evento_proceso
    ON eventos.calendario_evento (proceso_id);

CREATE INDEX idx_calendario_evento_catalogo
    ON eventos.calendario_evento (evento_catalogo_id);

CREATE INDEX idx_calendario_evento_ocurrencia
    ON eventos.calendario_evento (
        proceso_id,
        evento_catalogo_id,
        numero_ocurrencia DESC
    );

CREATE INDEX idx_calendario_evento_publico_perfil
    ON eventos.calendario_evento_tipo_publico (perfil_id);

CREATE INDEX idx_extension_programa_extension
    ON eventos.calendario_evento_extension_programa (
        calendario_evento_extension_id
    );

CREATE INDEX idx_extension_programa_padre
    ON eventos.calendario_evento_extension_programa (extension_padre_id);

CREATE INDEX idx_extension_programa_dependencia
    ON eventos.calendario_evento_extension_programa (dependencia_id);

CREATE INDEX idx_sesion_tipo
    ON eventos.sesion (tipo_sesion);

CREATE INDEX idx_sesion_patron_sesion
    ON eventos.sesion_patron_recurrencia (sesion);

CREATE INDEX idx_participante_sesion_rol
    ON eventos.participante_sesion (rol_participante_sesion);

CREATE INDEX idx_relacion_sesiones_hijo
    ON eventos.relacion_sesiones (sesion_hijo);

CREATE INDEX idx_auditoria_evento_fecha
    ON eventos.calendario_evento_auditoria (
        calendario_evento_id,
        fecha_creacion DESC
    );

CREATE INDEX idx_auditoria_tercero_fecha
    ON eventos.calendario_evento_auditoria (
        tercero_id,
        fecha_creacion DESC
    );

INSERT INTO eventos.tipo_recurrencia (
    id,
    nombre,
    descripcion,
    codigo_abreviacion,
    activo,
    fecha_creacion,
    fecha_modificacion
) VALUES
    (1, 'Semestral', 'Semestral', 'SEM', TRUE, NOW(), NOW()),
    (2, 'Anual', 'Anual', 'ANUA', TRUE, NOW(), NOW());

INSERT INTO eventos.proceso_catalogo (
    id,
    codigo_abreviacion,
    nombre,
    descripcion,
    activo
) VALUES (
    1,
    'PROC_ADMISIONES',
    'Admisiones',
    'Proceso de admision, reingreso, transferencia, aspirantes, resultados y oficializacion.',
    TRUE
);

INSERT INTO eventos.evento_catalogo (
    id,
    codigo_abreviacion,
    nombre,
    descripcion,
    activo
) VALUES
    (
        1,
        'REIN',
        'INSCRIPCION REINGRESOS',
        'Periodo de inscripcion para aspirantes por reingreso.',
        TRUE
    ),
    (
        2,
        'INSCR',
        'INSCRIPCION DE ASPIRANTES',
        'Periodo de inscripcion para aspirantes a programas academicos.',
        TRUE
    ),
    (
        3,
        'PAGO_REIN',
        'FECHAS DE PAGO PARA REINGRESOS',
        'Rango de fechas para recibo de pago de reingresos',
        TRUE
    ),
    (
        4,
        'PAGO_INSC',
        'FECHAS DE PAGO PARA INSCRIPCION',
        'Rango de fechas para recibo de pago de inscripciones',
        TRUE
    );

INSERT INTO eventos.evento_catalogo_proceso_catalogo (
    id,
    evento_catalogo_id,
    proceso_catalogo_id,
    repetible,
    activo
) VALUES
    (1, 1, 1, FALSE, TRUE),
    (2, 2, 1, FALSE, TRUE),
    (3, 3, 1, FALSE, TRUE),
    (4, 4, 1, FALSE, TRUE);

INSERT INTO eventos.rol_participante_sesion (
    id,
    nombre,
    descripcion,
    codigo_abreviacion,
    activo,
    numero_orden
) VALUES
    (1, 'Creador', 'Creador del evento', 'CE', TRUE, 1.00),
    (2, 'Participante', 'Participante del evento', 'PE', TRUE, 2.00);

INSERT INTO eventos.tipo_sesion (
    id,
    nombre,
    descripcion,
    codigo_abreviacion,
    activo,
    numero_orden
) VALUES
    (
        1,
        'Sesiones de modalidad de materias de posgrado',
        'Sesiones relacionadas con la modalidad de grado de materias de posgrado',
        'EMP',
        TRUE,
        1.00
    ),
    (
        2,
        'Publicacion de materias modalidad posgrado',
        'Publicacion de espacios academicos por parte de los proyectos curriculares',
        'PMP',
        TRUE,
        2.00
    ),
    (
        3,
        'Solicitud de materias de posgrado',
        'Solicitud de espacios academicos para cursar materias de posgrado',
        'SMP',
        TRUE,
        3.00
    ),
    (
        4,
        'Primera seleccion de admitidos a materias',
        'Primer corte de seleccion de estudiantes admitidos',
        'PFSAMP',
        TRUE,
        4.00
    ),
    (
        5,
        'Primera formalizacion de solicitudes',
        'Primera fecha para formalizar solicitudes de materias de posgrado',
        'PFFMP',
        TRUE,
        5.00
    ),
    (
        6,
        'Segunda seleccion de admitidos a materias',
        'Segundo corte de seleccion de estudiantes admitidos',
        'SCPPE',
        TRUE,
        6.00
    ),
    (
        7,
        'Segunda formalizacion de solicitudes',
        'Segunda fecha para formalizar solicitudes de materias de posgrado',
        'SFFMP',
        TRUE,
        7.00
    ),
    (
        8,
        'Aprobacion de solicitudes iniciales de materias',
        'Aprobacion de solicitudes iniciales de materias de posgrado',
        'ASIMP',
        TRUE,
        8.00
    );

SELECT setval(
    pg_get_serial_sequence('eventos.tipo_recurrencia', 'id'),
    (SELECT MAX(id) FROM eventos.tipo_recurrencia),
    TRUE
);

SELECT setval(
    pg_get_serial_sequence('eventos.proceso_catalogo', 'id'),
    (SELECT MAX(id) FROM eventos.proceso_catalogo),
    TRUE
);

SELECT setval(
    pg_get_serial_sequence('eventos.evento_catalogo', 'id'),
    (SELECT MAX(id) FROM eventos.evento_catalogo),
    TRUE
);

SELECT setval(
    pg_get_serial_sequence('eventos.evento_catalogo_proceso_catalogo', 'id'),
    (SELECT MAX(id) FROM eventos.evento_catalogo_proceso_catalogo),
    TRUE
);

SELECT setval(
    pg_get_serial_sequence('eventos.rol_participante_sesion', 'id'),
    (SELECT MAX(id) FROM eventos.rol_participante_sesion),
    TRUE
);

SELECT setval(
    pg_get_serial_sequence('eventos.tipo_sesion', 'id'),
    (SELECT MAX(id) FROM eventos.tipo_sesion),
    TRUE
);

COMMENT ON SCHEMA eventos IS
    'Modelo vigente de calendarios, eventos academicos, extensiones y sesiones.';

COMMENT ON TABLE eventos.proceso_catalogo IS
    'Catalogo institucional de procesos academicos.';

COMMENT ON TABLE eventos.evento_catalogo IS
    'Catalogo institucional de actividades de calendario.';

COMMENT ON TABLE eventos.evento_catalogo_proceso_catalogo IS
    'Relacion N:M entre actividades y procesos donde pueden utilizarse.';

COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.repetible IS
    'Indica si la actividad admite varias ocurrencias activas con rangos diferentes.';

COMMENT ON COLUMN eventos.calendario_evento.numero_ocurrencia IS
    'Consecutivo administrado por la aplicacion dentro del proceso y catalogo de actividad.';

COMMENT ON TABLE eventos.evento_catalogo_rol_gestion IS
    'Perfiles externos autorizados para gestionar actividades de catalogo.';

COMMENT ON TABLE eventos.calendario_evento_tipo_publico IS
    'Perfiles externos que conforman el publico dirigido de una actividad.';

COMMENT ON TABLE eventos.calendario_evento_extension IS
    'Extensiones de fecha autorizadas para una actividad de calendario.';

COMMENT ON TABLE eventos.calendario_evento_extension_programa IS
    'Programas cubiertos por una extension y su encadenamiento previo.';

COMMENT ON TABLE eventos.calendario_evento_auditoria IS
    'Registro append-only de cambios transaccionales sobre calendario_evento.';

COMMENT ON COLUMN eventos.calendario_evento_auditoria.operacion IS
    'Operacion informada por la aplicacion; el campo admite valores abiertos.';

COMMENT ON COLUMN eventos.calendario_evento_auditoria.estado_anterior IS
    'Instantanea previa; debe ser nula unicamente para INSERT.';

COMMENT ON COLUMN eventos.calendario_evento_auditoria.estado_nuevo IS
    'Instantanea posterior a la operacion.';

COMMENT ON COLUMN eventos.calendario_evento_auditoria.cambios IS
    'Detalle JSON de los valores globales y particulares modificados.';
