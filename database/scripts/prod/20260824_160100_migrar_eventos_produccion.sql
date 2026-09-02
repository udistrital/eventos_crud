-- Migracion del esquema eventos legacy al modelo vigente - PRODUCCION.

-- Mantiene bloqueos, tablas temporales y cambios dentro de una unica transaccion.
BEGIN;

-- 1. Bloquea el inventario legacy y descarta los procesos que no pertenecen a admisiones.
LOCK TABLE
    eventos.calendario,
    eventos.calendario_evento,
    eventos.calendario_evento_tipo_publico,
    eventos.encargado_evento,
    eventos.participante_sesion,
    eventos.relacion_sesiones,
    eventos.rol_encargado_evento,
    eventos.rol_participante_sesion,
    eventos.sesion,
    eventos.sesion_patron_recurrencia,
    eventos.tipo_evento,
    eventos.tipo_publico,
    eventos.tipo_recurrencia,
    eventos.tipo_sesion
IN ACCESS EXCLUSIVE MODE;

CREATE TEMP TABLE tmp_tipo_evento_descartado (
    tipo_evento_id INTEGER PRIMARY KEY
) ON COMMIT DROP;

INSERT INTO tmp_tipo_evento_descartado
SELECT id
FROM eventos.tipo_evento
WHERE LOWER(TRIM(nombre)) NOT IN (
    'proceso de admisiones',
    'proceso de admisiones - reingreso'
);

CREATE TEMP TABLE tmp_calendario_evento_descartado (
    calendario_evento_id INTEGER PRIMARY KEY
) ON COMMIT DROP;

WITH RECURSIVE eventos_descartados AS (
    SELECT ce.id
    FROM eventos.calendario_evento ce
    JOIN tmp_tipo_evento_descartado d ON d.tipo_evento_id = ce.tipo_evento_id
    UNION
    SELECT hijo.id
    FROM eventos.calendario_evento hijo
    JOIN eventos_descartados padre ON padre.id = hijo.evento_padre_id
)
INSERT INTO tmp_calendario_evento_descartado
SELECT id FROM eventos_descartados;

DELETE FROM eventos.calendario_evento_tipo_publico r
USING tmp_calendario_evento_descartado d
WHERE r.calendario_evento_id = d.calendario_evento_id;

DELETE FROM eventos.encargado_evento e
USING tmp_calendario_evento_descartado d
WHERE e.calendario_evento_id = d.calendario_evento_id;

DELETE FROM eventos.calendario_evento ce
USING tmp_calendario_evento_descartado d
WHERE ce.id = d.calendario_evento_id;

DELETE FROM eventos.tipo_evento te
USING tmp_tipo_evento_descartado d
WHERE te.id = d.tipo_evento_id;

-- 2. Verifica que los objetos y datos legacy restantes coincidan con el inventario aprobado.
DO $$
DECLARE
    tabla TEXT;
    tablas_legacy CONSTANT TEXT[] := ARRAY[
        'calendario',
        'calendario_evento',
        'calendario_evento_tipo_publico',
        'encargado_evento',
        'participante_sesion',
        'relacion_sesiones',
        'rol_encargado_evento',
        'rol_participante_sesion',
        'sesion',
        'sesion_patron_recurrencia',
        'tipo_evento',
        'tipo_publico',
        'tipo_recurrencia',
        'tipo_sesion'
    ];
BEGIN
    FOREACH tabla IN ARRAY tablas_legacy LOOP
        IF to_regclass(format('eventos.%I', tabla)) IS NULL THEN
            RAISE EXCEPTION 'Falta la tabla legacy eventos.%', tabla;
        END IF;
    END LOOP;

    IF (SELECT COUNT(*) FROM information_schema.tables
        WHERE table_schema = 'eventos' AND table_type = 'BASE TABLE') <> 14 THEN
        RAISE EXCEPTION 'El esquema eventos no tiene exactamente 14 tablas legacy';
    END IF;

    IF EXISTS (
        SELECT 1 FROM information_schema.columns
        WHERE table_schema = 'eventos'
          AND table_name = 'calendario'
          AND column_name = 'multiple_periodo_id'
    ) THEN
        RAISE EXCEPTION 'Aparecio calendario.multiple_periodo_id; requiere una regla explicita';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.calendario) <> 14
       OR (SELECT COUNT(*) FROM eventos.tipo_evento) <> 15
       OR (SELECT COUNT(*) FROM eventos.calendario_evento) <> 29
       OR (SELECT COUNT(*) FROM eventos.calendario_evento_tipo_publico) <> 31 THEN
        RAISE EXCEPTION 'Los conteos legacy cambiaron desde el inventario aprobado';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.sesion) <> 0
       OR (SELECT COUNT(*) FROM eventos.sesion_patron_recurrencia) <> 0
       OR (SELECT COUNT(*) FROM eventos.participante_sesion) <> 0
       OR (SELECT COUNT(*) FROM eventos.relacion_sesiones) <> 0 THEN
        RAISE EXCEPTION 'Aparecieron datos de sesiones no contemplados en el inventario';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.encargado_evento) <> 2 THEN
        RAISE EXCEPTION 'Cambio el numero de encargados legacy';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.calendario_evento
        WHERE LOWER(TRIM(nombre)) = 'prueba') <> 1 THEN
        RAISE EXCEPTION 'Se esperaba exactamente un evento llamado prueba';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.calendario
        WHERE LOWER(TRIM(nombre)) = 'a') <> 1 THEN
        RAISE EXCEPTION 'Se esperaba exactamente un calendario llamado a';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM eventos.tipo_evento te
        JOIN eventos.calendario c ON c.id = te.calendario_id
        WHERE LOWER(TRIM(c.nombre)) = 'a'
    ) THEN
        RAISE EXCEPTION 'El calendario a tiene procesos asociados y no puede descartarse directamente';
    END IF;

    IF EXISTS (SELECT 1 FROM eventos.tipo_evento te JOIN tmp_tipo_evento_descartado d ON d.tipo_evento_id = te.id)
       OR EXISTS (SELECT 1 FROM eventos.calendario_evento ce JOIN tmp_calendario_evento_descartado d ON d.calendario_evento_id = ce.id) THEN
        RAISE EXCEPTION 'No fue posible descartar todos los procesos no contemplados';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM eventos.tipo_evento te
        LEFT JOIN eventos.calendario c ON c.id = te.calendario_id
        LEFT JOIN eventos.tipo_recurrencia tr ON tr.id = te.tipo_recurrencia_id
        WHERE c.id IS NULL OR tr.id IS NULL
    ) OR EXISTS (
        SELECT 1
        FROM eventos.calendario_evento ce
        LEFT JOIN eventos.tipo_evento te ON te.id = ce.tipo_evento_id
        WHERE te.id IS NULL
    ) OR EXISTS (
        SELECT 1
        FROM eventos.calendario_evento_tipo_publico r
        LEFT JOIN eventos.calendario_evento ce ON ce.id = r.calendario_evento_id
        LEFT JOIN eventos.tipo_publico tp ON tp.id = r.tipo_publico_id
        WHERE ce.id IS NULL OR tp.id IS NULL
    ) THEN
        RAISE EXCEPTION 'Existen referencias legacy huerfanas';
    END IF;

    IF EXISTS (
        SELECT periodo_id, nivel
        FROM eventos.calendario
        WHERE activo IS TRUE
        GROUP BY periodo_id, nivel
        HAVING COUNT(*) > 1
    ) THEN
        RAISE EXCEPTION 'Hay calendarios activos duplicados por periodo y nivel';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.calendario
        WHERE dependencia_id::JSONB = '{}'::JSONB) <> 1
       OR EXISTS (
        SELECT 1 FROM eventos.calendario
        WHERE dependencia_id::JSONB <> '{}'::JSONB
          AND (
              JSONB_TYPEOF(dependencia_id::JSONB) <> 'object'
              OR NOT dependencia_id::JSONB ? 'proyectos'
              OR JSONB_TYPEOF(dependencia_id::JSONB -> 'proyectos') <> 'array'
          )
    ) OR EXISTS (
        SELECT 1 FROM eventos.calendario_evento
        WHERE dependencia_id IS NOT NULL
          AND NOT (
              JSONB_TYPEOF(dependencia_id::JSONB) = 'object'
              AND dependencia_id::JSONB ? 'proyectos'
              AND JSONB_TYPEOF(dependencia_id::JSONB -> 'proyectos') = 'array'
              AND (
                  NOT dependencia_id::JSONB ? 'fechas'
                  OR JSONB_TYPEOF(dependencia_id::JSONB -> 'fechas') = 'array'
              )
          )
    ) THEN
        RAISE EXCEPTION 'Hay estructuras de dependencia incompatibles';
    END IF;

    IF EXISTS (
        SELECT 1 FROM eventos.calendario_evento
        WHERE fecha_fin IS NOT NULL AND fecha_fin < fecha_inicio
    ) THEN
        RAISE EXCEPTION 'Hay eventos con rango de fechas invalido';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.calendario_evento WHERE ubicacion_id IS NULL) <> 1
       OR (SELECT COUNT(*) FROM eventos.calendario_evento WHERE poster_url IS NULL) <> 1
       OR EXISTS (SELECT 1 FROM eventos.calendario_evento WHERE fecha_modificacion IS NULL) THEN
        RAISE EXCEPTION 'Cambió el conjunto de normalizaciones esperado';
    END IF;

    IF EXISTS (
        SELECT 1 FROM eventos.calendario
        WHERE COALESCE(aplica_extension, FALSE) IS TRUE
           OR COALESCE(documento_extension_id, 0) <> 0
           OR COALESCE(dependencia_particular_id::JSONB, '{}'::JSONB) <> '{}'::JSONB
    ) THEN
        RAISE EXCEPTION 'Existen extensiones legacy reales que no pueden descartarse';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pg_depend d
        JOIN pg_rewrite r ON r.oid = d.objid
        JOIN pg_class vista ON vista.oid = r.ev_class
        JOIN pg_namespace nv ON nv.oid = vista.relnamespace
        JOIN pg_class origen ON origen.oid = d.refobjid
        JOIN pg_namespace no ON no.oid = origen.relnamespace
        WHERE no.nspname = 'eventos' AND nv.nspname <> 'eventos'
    ) OR EXISTS (
        SELECT 1
        FROM pg_constraint con
        JOIN pg_class origen ON origen.oid = con.confrelid
        JOIN pg_namespace no ON no.oid = origen.relnamespace
        JOIN pg_class destino ON destino.oid = con.conrelid
        JOIN pg_namespace nd ON nd.oid = destino.relnamespace
        WHERE con.contype = 'f'
          AND no.nspname = 'eventos'
          AND nd.nspname <> 'eventos'
    ) THEN
        RAISE EXCEPTION 'Hay dependencias externas sobre el esquema eventos';
    END IF;

END $$;

-- 3. Resuelve los perfiles externos de produccion usados por el publico dirigido.
CREATE TEMP TABLE tmp_perfil_publico (
    tipo_publico_id INTEGER PRIMARY KEY,
    perfil_id INTEGER NOT NULL UNIQUE,
    nombre TEXT NOT NULL
) ON COMMIT DROP;

-- Perfiles verificados en produccion: SGA_MF=35, ASPIRANTE=157 y ESTUDIANTE=187 (2026-08-26).
INSERT INTO tmp_perfil_publico (tipo_publico_id, perfil_id, nombre)
VALUES
    (1, 157, 'ASPIRANTE'),
    (2, 187, 'ESTUDIANTE');

DO $$
BEGIN
    IF (SELECT COUNT(*) FROM tmp_perfil_publico) <> 2 THEN
        RAISE EXCEPTION 'No se cargaron los perfiles de produccion ASPIRANTE=157 y ESTUDIANTE=187';
    END IF;
END $$;

-- 4. Construye los mapas temporales para consolidar tipos legacy en procesos y actividades.
CREATE TEMP TABLE tmp_tipo_evento_catalogo (
    tipo_evento_id INTEGER PRIMARY KEY,
    proceso_catalogo_id INTEGER NOT NULL
) ON COMMIT DROP;

INSERT INTO tmp_tipo_evento_catalogo
SELECT id, 1
FROM eventos.tipo_evento
WHERE LOWER(TRIM(nombre)) IN (
    'proceso de admisiones',
    'proceso de admisiones - reingreso'
);

DO $$
BEGIN
    IF (SELECT COUNT(*) FROM tmp_tipo_evento_catalogo) <> 15 THEN
        RAISE EXCEPTION 'Se esperaban 15 tipos legacy asociados a PROC_ADMISIONES';
    END IF;
END $$;

CREATE TEMP TABLE tmp_proceso ON COMMIT DROP AS
SELECT MIN(te.id)::INTEGER AS id,
       BOOL_OR(te.activo) AS activo,
       MIN(te.fecha_creacion) AS fecha_creacion,
       MAX(te.fecha_modificacion) AS fecha_modificacion,
       MIN(te.tipo_recurrencia_id)::INTEGER AS tipo_recurrencia_id,
       te.calendario_id::INTEGER AS calendario_id,
       m.proceso_catalogo_id
FROM eventos.tipo_evento te
JOIN tmp_tipo_evento_catalogo m ON m.tipo_evento_id = te.id
GROUP BY te.calendario_id, m.proceso_catalogo_id;

CREATE TEMP TABLE tmp_tipo_evento_proceso ON COMMIT DROP AS
SELECT te.id AS tipo_evento_id, p.id AS proceso_id
FROM eventos.tipo_evento te
JOIN tmp_tipo_evento_catalogo m ON m.tipo_evento_id = te.id
JOIN tmp_proceso p
  ON p.calendario_id = te.calendario_id
 AND p.proceso_catalogo_id = m.proceso_catalogo_id;

CREATE TEMP TABLE tmp_evento_catalogo (
    calendario_evento_id INTEGER PRIMARY KEY,
    evento_catalogo_id INTEGER NOT NULL
) ON COMMIT DROP;

INSERT INTO tmp_evento_catalogo
SELECT id,
       CASE LOWER(TRIM(nombre))
           WHEN 'inscripción de aspirantes reingreso' THEN 1
           WHEN 'inscripción de aspirantes' THEN 2
           WHEN 'fechas de pago para la inscripción de reingreso' THEN 3
           WHEN 'fechas de pago para la inscripción de aspirante' THEN 4
       END
FROM eventos.calendario_evento
WHERE LOWER(TRIM(nombre)) IN (
    'inscripción de aspirantes reingreso',
    'inscripción de aspirantes',
    'fechas de pago para la inscripción de reingreso',
    'fechas de pago para la inscripción de aspirante'
);

DO $$
BEGIN
    IF (SELECT COUNT(*) FROM tmp_proceso) <> 11
       OR (SELECT COUNT(*) FROM tmp_tipo_evento_proceso) <> 15
       OR (SELECT COUNT(*) FROM tmp_evento_catalogo WHERE evento_catalogo_id IS NOT NULL) <> 26 THEN
        RAISE EXCEPTION 'Los mapas de procesos o eventos no tienen los conteos esperados';
    END IF;
END $$;

-- 5. Transforma los calendarios conservando sus identificadores externos de produccion.
-- Los periodo_id 12, 13, 26, 29, 31, 33, 34, 35, 37, 38 y 47 fueron validados como PA (2026-08-26).
SET LOCAL search_path TO eventos, public;

ALTER TABLE eventos.calendario DROP CONSTRAINT IF EXISTS fk_calendario_padre;
DELETE FROM eventos.calendario WHERE LOWER(TRIM(nombre)) = 'a';
ALTER TABLE eventos.calendario
    DROP COLUMN calendario_padre_id,
    DROP COLUMN documento_extension_id,
    DROP COLUMN aplica_extension,
    DROP COLUMN dependencia_particular_id;
UPDATE eventos.calendario
SET dependencia_id = '{"proyectos":[]}'::JSON
WHERE dependencia_id::JSONB = '{}'::JSONB;
ALTER TABLE eventos.calendario ADD CONSTRAINT ck_calendario_dependencia_estructura
CHECK (
    JSONB_TYPEOF(dependencia_id::JSONB) = 'object'
    AND dependencia_id::JSONB ? 'proyectos'
    AND JSONB_TYPEOF(dependencia_id::JSONB -> 'proyectos') = 'array'
);
CREATE UNIQUE INDEX uq_calendario_periodo_nivel_activo
ON eventos.calendario (periodo_id, nivel) WHERE activo IS TRUE;

-- 6. Normaliza las fechas particulares al formato visible usado por el MID, sin zona horaria.
CREATE FUNCTION pg_temp.normalizar_fecha_visible(fecha TEXT)
RETURNS TEXT
LANGUAGE SQL
IMMUTABLE
AS $$
    SELECT CASE
        WHEN fecha ~ '^[0-9]{4}-[0-9]{2}-[0-9]{2}[T ][0-9]{2}:[0-9]{2}:[0-9]{2}' THEN
            REPLACE(SUBSTRING(fecha FROM 1 FOR 19), ' ', 'T')
        WHEN fecha ~ '^[0-9]{2}-[0-9]{2}-[0-9]{4}$' THEN
            SUBSTRING(fecha FROM 7 FOR 4) || '-'
            || SUBSTRING(fecha FROM 4 FOR 2) || '-'
            || SUBSTRING(fecha FROM 1 FOR 2) || 'T00:00:00'
        ELSE fecha
    END
$$;

WITH dependencias_normalizadas AS (
    SELECT ce.id,
           JSONB_SET(
               ce.dependencia_id::JSONB,
               '{fechas}',
               (
                   SELECT JSONB_AGG(
                       fecha.valor || JSONB_BUILD_OBJECT(
                           'Inicio', pg_temp.normalizar_fecha_visible(fecha.valor ->> 'Inicio'),
                           'Fin', pg_temp.normalizar_fecha_visible(fecha.valor ->> 'Fin'),
                           'Modificacion', pg_temp.normalizar_fecha_visible(fecha.valor ->> 'Modificacion')
                       )
                       ORDER BY fecha.orden
                   )
                   FROM JSONB_ARRAY_ELEMENTS(ce.dependencia_id::JSONB -> 'fechas')
                       WITH ORDINALITY AS fecha(valor, orden)
               ),
               FALSE
           )::JSON AS dependencia_id
    FROM eventos.calendario_evento ce
    WHERE ce.dependencia_id IS NOT NULL
      AND JSONB_TYPEOF(ce.dependencia_id::JSONB -> 'fechas') = 'array'
)
UPDATE eventos.calendario_evento ce
SET dependencia_id = normalizada.dependencia_id
FROM dependencias_normalizadas normalizada
WHERE normalizada.id = ce.id;

DO $$
BEGIN
    IF EXISTS (
        SELECT 1
        FROM eventos.calendario_evento ce
        CROSS JOIN LATERAL JSONB_ARRAY_ELEMENTS(ce.dependencia_id::JSONB -> 'fechas') fecha
        WHERE ce.dependencia_id IS NOT NULL
          AND JSONB_TYPEOF(ce.dependencia_id::JSONB -> 'fechas') = 'array'
          AND (
              fecha ->> 'Inicio' !~ '^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}$'
              OR fecha ->> 'Fin' !~ '^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}$'
              OR fecha ->> 'Modificacion' !~ '^[0-9]{4}-[0-9]{2}-[0-9]{2}T[0-9]{2}:[0-9]{2}:[0-9]{2}$'
          )
    ) THEN
        RAISE EXCEPTION 'No fue posible normalizar todas las fechas particulares';
    END IF;
END $$;

-- 7. Completa restricciones y parametria de sesiones que permanecen en el modelo vigente.
UPDATE eventos.sesion
SET fecha_modificacion = fecha_creacion
WHERE fecha_modificacion IS NULL;
ALTER TABLE eventos.sesion ALTER COLUMN fecha_modificacion SET NOT NULL;
ALTER TABLE eventos.sesion ADD CONSTRAINT ck_sesion_rango_fechas
CHECK (fecha_fin IS NULL OR fecha_fin >= fecha_inicio);
ALTER TABLE eventos.relacion_sesiones ADD CONSTRAINT ck_relacion_sesiones_distintas
CHECK (sesion_padre <> sesion_hijo);

INSERT INTO eventos.rol_participante_sesion
    (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden)
VALUES
    (1, 'Creador', 'Creador del evento', 'CE', TRUE, 1.00),
    (2, 'Participante', 'Participante del evento', 'PE', TRUE, 2.00);

INSERT INTO eventos.tipo_sesion
    (id, nombre, descripcion, codigo_abreviacion, activo, numero_orden)
VALUES
    (1, 'Sesiones de modalidad de materias de posgrado', 'Sesiones relacionadas con la modalidad de grado de materias de posgrado', 'EMP', TRUE, 1.00),
    (2, 'Publicacion de materias modalidad posgrado', 'Publicacion de espacios academicos por parte de los proyectos curriculares', 'PMP', TRUE, 2.00),
    (3, 'Solicitud de materias de posgrado', 'Solicitud de espacios academicos para cursar materias de posgrado', 'SMP', TRUE, 3.00),
    (4, 'Primera seleccion de admitidos a materias', 'Primer corte de seleccion de estudiantes admitidos', 'PFSAMP', TRUE, 4.00),
    (5, 'Primera formalizacion de solicitudes', 'Primera fecha para formalizar solicitudes de materias de posgrado', 'PFFMP', TRUE, 5.00),
    (6, 'Segunda seleccion de admitidos a materias', 'Segundo corte de seleccion de estudiantes admitidos', 'SCPPE', TRUE, 6.00),
    (7, 'Segunda formalizacion de solicitudes', 'Segunda fecha para formalizar solicitudes de materias de posgrado', 'SFFMP', TRUE, 7.00),
    (8, 'Aprobacion de solicitudes iniciales de materias', 'Aprobacion de solicitudes iniciales de materias de posgrado', 'ASIMP', TRUE, 8.00);

-- 8. Crea las tablas estructurales de catalogos, procesos y roles funcionales de gestion.
CREATE TABLE eventos.proceso_catalogo (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(50) NOT NULL UNIQUE,
    nombre VARCHAR(150) NOT NULL,
    descripcion VARCHAR(250),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE eventos.evento_catalogo (
    id SERIAL PRIMARY KEY,
    codigo_abreviacion VARCHAR(50) NOT NULL UNIQUE,
    nombre VARCHAR(150) NOT NULL,
    descripcion VARCHAR(250),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE eventos.proceso (
    id SERIAL PRIMARY KEY,
    activo BOOLEAN NOT NULL,
    fecha_creacion TIMESTAMP NOT NULL,
    fecha_modificacion TIMESTAMP NOT NULL,
    tipo_recurrencia_id INTEGER NOT NULL REFERENCES eventos.tipo_recurrencia(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    calendario_id INTEGER NOT NULL REFERENCES eventos.calendario(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    proceso_catalogo_id INTEGER NOT NULL REFERENCES eventos.proceso_catalogo(id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE eventos.evento_catalogo_proceso_catalogo (
    id SERIAL PRIMARY KEY,
    evento_catalogo_id INTEGER NOT NULL REFERENCES eventos.evento_catalogo(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    proceso_catalogo_id INTEGER NOT NULL REFERENCES eventos.proceso_catalogo(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    repetible BOOLEAN NOT NULL DEFAULT FALSE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_evento_catalogo_proceso_catalogo UNIQUE (evento_catalogo_id, proceso_catalogo_id)
);

CREATE TABLE eventos.evento_catalogo_rol_gestion (
    id SERIAL PRIMARY KEY,
    evento_catalogo_id INTEGER NOT NULL REFERENCES eventos.evento_catalogo(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    perfil_id INTEGER NOT NULL,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_evento_catalogo_rol_gestion_perfil UNIQUE (evento_catalogo_id, perfil_id)
);

-- 9. Migra los datos legacy hacia los nuevos catalogos y procesos consolidados.
INSERT INTO eventos.proceso_catalogo
    (id, codigo_abreviacion, nombre, descripcion, activo, fecha_creacion, fecha_modificacion)
SELECT catalogo.id, catalogo.codigo, catalogo.nombre, catalogo.descripcion,
       BOOL_OR(te.activo), MIN(te.fecha_creacion), MAX(te.fecha_modificacion)
FROM (VALUES
    (1, 'PROC_ADMISIONES', 'Admisiones', 'Proceso de admision, reingreso, transferencia, aspirantes, resultados y oficializacion.')
) AS catalogo(id, codigo, nombre, descripcion)
JOIN tmp_tipo_evento_catalogo m ON m.proceso_catalogo_id = catalogo.id
JOIN eventos.tipo_evento te ON te.id = m.tipo_evento_id
GROUP BY catalogo.id, catalogo.codigo, catalogo.nombre, catalogo.descripcion;

INSERT INTO eventos.evento_catalogo
    (id, codigo_abreviacion, nombre, descripcion, activo, fecha_creacion, fecha_modificacion)
SELECT catalogo.id, catalogo.codigo, catalogo.nombre, catalogo.descripcion,
       BOOL_OR(ce.activo), MIN(ce.fecha_creacion), MAX(ce.fecha_modificacion)
FROM (VALUES
    (1, 'REIN', 'INSCRIPCION REINGRESOS', 'Periodo de inscripcion para aspirantes por reingreso.'),
    (2, 'INSCR', 'INSCRIPCION DE ASPIRANTES', 'Periodo de inscripcion para aspirantes a programas academicos.'),
    (3, 'PAGO_REIN', 'FECHAS DE PAGO PARA REINGRESOS', 'Rango de fechas para recibo de pago de reingresos'),
    (4, 'PAGO_INSC', 'FECHAS DE PAGO PARA INSCRIPCION', 'Rango de fechas para recibo de pago de inscripciones')
) AS catalogo(id, codigo, nombre, descripcion)
JOIN tmp_evento_catalogo m ON m.evento_catalogo_id = catalogo.id
JOIN eventos.calendario_evento ce ON ce.id = m.calendario_evento_id
GROUP BY catalogo.id, catalogo.codigo, catalogo.nombre, catalogo.descripcion;

INSERT INTO eventos.proceso
    (id, activo, fecha_creacion, fecha_modificacion, tipo_recurrencia_id, calendario_id, proceso_catalogo_id)
SELECT id, activo, fecha_creacion, fecha_modificacion, tipo_recurrencia_id, calendario_id, proceso_catalogo_id
FROM tmp_proceso;

INSERT INTO eventos.evento_catalogo_proceso_catalogo
    (id, evento_catalogo_id, proceso_catalogo_id, repetible, activo, fecha_creacion, fecha_modificacion)
SELECT ec.id, ec.id, 1, FALSE, ec.activo, ec.fecha_creacion, ec.fecha_modificacion
FROM eventos.evento_catalogo ec;

CREATE UNIQUE INDEX uq_proceso_calendario_catalogo_activo
ON eventos.proceso (calendario_id, proceso_catalogo_id) WHERE activo IS TRUE;
CREATE INDEX idx_proceso_calendario ON eventos.proceso (calendario_id);
CREATE INDEX idx_proceso_tipo_recurrencia ON eventos.proceso (tipo_recurrencia_id);
CREATE INDEX idx_proceso_catalogo ON eventos.proceso (proceso_catalogo_id);
CREATE INDEX idx_evento_catalogo_proceso_proceso
ON eventos.evento_catalogo_proceso_catalogo (proceso_catalogo_id);
CREATE INDEX idx_evento_catalogo_rol_perfil
ON eventos.evento_catalogo_rol_gestion (perfil_id);

-- 10. Convierte los eventos legacy en actividades catalogadas y migra su publico dirigido.
DELETE FROM eventos.calendario_evento_tipo_publico r
USING eventos.calendario_evento ce
WHERE r.calendario_evento_id = ce.id
  AND (
      NOT EXISTS (
          SELECT 1 FROM tmp_evento_catalogo m
          WHERE m.calendario_evento_id = ce.id
      )
      OR r.tipo_publico_id IN (3, 4)
  );

ALTER TABLE eventos.calendario_evento
    DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_evento,
    DROP CONSTRAINT IF EXISTS fk_evento_padre;
ALTER TABLE eventos.calendario_evento
    ADD COLUMN proceso_id INTEGER,
    ADD COLUMN evento_catalogo_id INTEGER,
    ADD COLUMN numero_ocurrencia INTEGER NOT NULL DEFAULT 1;

UPDATE eventos.calendario_evento ce
SET proceso_id = pm.proceso_id,
    evento_catalogo_id = em.evento_catalogo_id,
    ubicacion_id = COALESCE(ce.ubicacion_id, 0),
    poster_url = COALESCE(ce.poster_url, '')
FROM tmp_tipo_evento_proceso pm, tmp_evento_catalogo em
WHERE pm.tipo_evento_id = ce.tipo_evento_id
  AND em.calendario_evento_id = ce.id;

DELETE FROM eventos.calendario_evento ce
WHERE NOT EXISTS (
    SELECT 1 FROM tmp_evento_catalogo m
    WHERE m.calendario_evento_id = ce.id
);

ALTER TABLE eventos.calendario_evento
    ALTER COLUMN fecha_modificacion SET NOT NULL,
    ALTER COLUMN proceso_id SET NOT NULL,
    ALTER COLUMN evento_catalogo_id SET NOT NULL,
    ALTER COLUMN ubicacion_id SET DEFAULT 0,
    ALTER COLUMN ubicacion_id SET NOT NULL,
    ALTER COLUMN poster_url SET DEFAULT '',
    ALTER COLUMN poster_url SET NOT NULL,
    DROP COLUMN descripcion,
    DROP COLUMN evento_padre_id,
    DROP COLUMN tipo_evento_id,
    DROP COLUMN nombre,
    DROP COLUMN aplica_edicion_actividades;

ALTER TABLE eventos.calendario_evento
    ADD CONSTRAINT fk_calendario_evento_proceso FOREIGN KEY (proceso_id)
        REFERENCES eventos.proceso(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    ADD CONSTRAINT fk_calendario_evento_evento_catalogo FOREIGN KEY (evento_catalogo_id)
        REFERENCES eventos.evento_catalogo(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    ADD CONSTRAINT ck_calendario_evento_dependencia_estructura CHECK (
        dependencia_id IS NULL OR (
            JSONB_TYPEOF(dependencia_id::JSONB) = 'object'
            AND dependencia_id::JSONB ? 'proyectos'
            AND JSONB_TYPEOF(dependencia_id::JSONB -> 'proyectos') = 'array'
            AND (NOT dependencia_id::JSONB ? 'fechas' OR JSONB_TYPEOF(dependencia_id::JSONB -> 'fechas') = 'array')
        )
    ),
    ADD CONSTRAINT ck_calendario_evento_rango_fechas CHECK (fecha_fin IS NULL OR fecha_fin >= fecha_inicio),
    ADD CONSTRAINT ck_calendario_evento_numero_ocurrencia CHECK (numero_ocurrencia > 0);

CREATE UNIQUE INDEX uq_calendario_evento_ocurrencia_activa
ON eventos.calendario_evento (proceso_id, evento_catalogo_id, numero_ocurrencia) WHERE activo IS TRUE;
CREATE UNIQUE INDEX uq_calendario_evento_rango_activo_cerrado
ON eventos.calendario_evento (proceso_id, evento_catalogo_id, fecha_inicio, fecha_fin)
WHERE activo IS TRUE AND fecha_fin IS NOT NULL;
CREATE UNIQUE INDEX uq_calendario_evento_rango_activo_abierto
ON eventos.calendario_evento (proceso_id, evento_catalogo_id, fecha_inicio)
WHERE activo IS TRUE AND fecha_fin IS NULL;
CREATE INDEX idx_calendario_evento_proceso ON eventos.calendario_evento (proceso_id);
CREATE INDEX idx_calendario_evento_catalogo ON eventos.calendario_evento (evento_catalogo_id);
CREATE INDEX idx_calendario_evento_ocurrencia
ON eventos.calendario_evento (proceso_id, evento_catalogo_id, numero_ocurrencia DESC);

ALTER TABLE eventos.calendario_evento_tipo_publico
    DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_publico_calendario_evento,
    DROP CONSTRAINT IF EXISTS fk_calendario_evento_tipo_publico_tipo_publico;
ALTER TABLE eventos.calendario_evento_tipo_publico RENAME COLUMN tipo_publico_id TO perfil_id;
UPDATE eventos.calendario_evento_tipo_publico r
SET perfil_id = m.perfil_id
FROM tmp_perfil_publico m
WHERE r.perfil_id = m.tipo_publico_id;
ALTER TABLE eventos.calendario_evento_tipo_publico
    ALTER COLUMN calendario_evento_id SET NOT NULL,
    ALTER COLUMN perfil_id SET NOT NULL,
    ADD CONSTRAINT fk_calendario_evento_tipo_publico_evento FOREIGN KEY (calendario_evento_id)
        REFERENCES eventos.calendario_evento(id) ON UPDATE CASCADE ON DELETE CASCADE,
    ADD CONSTRAINT uq_calendario_evento_tipo_publico UNIQUE (calendario_evento_id, perfil_id);
CREATE INDEX idx_calendario_evento_publico_perfil
ON eventos.calendario_evento_tipo_publico (perfil_id);

-- 11. Crea el modelo de extensiones por programa y la auditoria append-only de actividades.
CREATE TABLE eventos.calendario_evento_extension (
    id SERIAL PRIMARY KEY,
    calendario_evento_id INTEGER NOT NULL REFERENCES eventos.calendario_evento(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    fecha_fin TIMESTAMP NOT NULL,
    documento_id INTEGER,
    descripcion TEXT,
    numero_extension INTEGER NOT NULL CHECK (numero_extension > 0),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_calendario_evento_extension_numero UNIQUE (calendario_evento_id, numero_extension)
);

CREATE TABLE eventos.calendario_evento_extension_programa (
    id SERIAL PRIMARY KEY,
    calendario_evento_extension_id INTEGER NOT NULL REFERENCES eventos.calendario_evento_extension(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    extension_padre_id INTEGER REFERENCES eventos.calendario_evento_extension(id) ON UPDATE CASCADE ON DELETE RESTRICT,
    dependencia_id INTEGER NOT NULL,
    vigente BOOLEAN NOT NULL DEFAULT TRUE,
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT ck_extension_programa_vigencia CHECK (NOT vigente OR activo),
    CONSTRAINT ck_extension_programa_padre_distinto CHECK (
        extension_padre_id IS NULL OR extension_padre_id <> calendario_evento_extension_id
    )
);

CREATE UNIQUE INDEX uq_extension_programa_dependencia_activa
ON eventos.calendario_evento_extension_programa (calendario_evento_extension_id, dependencia_id)
WHERE activo IS TRUE;
CREATE INDEX idx_extension_programa_extension
ON eventos.calendario_evento_extension_programa (calendario_evento_extension_id);
CREATE INDEX idx_extension_programa_padre
ON eventos.calendario_evento_extension_programa (extension_padre_id);
CREATE INDEX idx_extension_programa_dependencia
ON eventos.calendario_evento_extension_programa (dependencia_id);

CREATE TABLE eventos.calendario_evento_auditoria (
    id BIGSERIAL PRIMARY KEY,
    calendario_evento_id INTEGER NOT NULL REFERENCES eventos.calendario_evento(id) ON UPDATE RESTRICT ON DELETE RESTRICT,
    operacion TEXT NOT NULL,
    tercero_id INTEGER NOT NULL CHECK (tercero_id > 0),
    fecha_creacion TIMESTAMP NOT NULL DEFAULT NOW(),
    fecha_modificacion TIMESTAMP NOT NULL DEFAULT NOW(),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    estado_anterior JSONB,
    estado_nuevo JSONB NOT NULL,
    cambios JSONB NOT NULL DEFAULT '{}'::JSONB,
    CONSTRAINT ck_auditoria_estado_anterior CHECK (
        (operacion = 'INSERT' AND estado_anterior IS NULL)
        OR (operacion <> 'INSERT' AND estado_anterior IS NOT NULL)
    ),
    CONSTRAINT ck_auditoria_fechas_append_only CHECK (fecha_modificacion = fecha_creacion),
    CONSTRAINT ck_auditoria_activo CHECK (activo IS TRUE)
);
CREATE INDEX idx_auditoria_evento_fecha
ON eventos.calendario_evento_auditoria (calendario_evento_id, fecha_creacion DESC);
CREATE INDEX idx_auditoria_tercero_fecha
ON eventos.calendario_evento_auditoria (tercero_id, fecha_creacion DESC);

-- 12. Completa indices auxiliares y sincroniza las secuencias con los identificadores migrados.
CREATE INDEX IF NOT EXISTS idx_sesion_tipo ON eventos.sesion (tipo_sesion);
CREATE INDEX IF NOT EXISTS idx_sesion_patron_sesion ON eventos.sesion_patron_recurrencia (sesion);
CREATE INDEX IF NOT EXISTS idx_participante_sesion_rol ON eventos.participante_sesion (rol_participante_sesion);
CREATE INDEX IF NOT EXISTS idx_relacion_sesiones_hijo ON eventos.relacion_sesiones (sesion_hijo);

DO $$
DECLARE
    tabla TEXT;
    secuencia TEXT;
    maximo BIGINT;
BEGIN
    FOREACH tabla IN ARRAY ARRAY[
        'tipo_recurrencia', 'tipo_sesion', 'rol_participante_sesion', 'calendario',
        'proceso_catalogo', 'evento_catalogo', 'proceso',
        'evento_catalogo_proceso_catalogo', 'evento_catalogo_rol_gestion',
        'calendario_evento', 'calendario_evento_tipo_publico',
        'calendario_evento_extension', 'calendario_evento_extension_programa',
        'sesion', 'sesion_patron_recurrencia', 'participante_sesion',
        'relacion_sesiones', 'calendario_evento_auditoria'
    ] LOOP
        secuencia := pg_get_serial_sequence(format('eventos.%I', tabla), 'id');
        EXECUTE format('SELECT MAX(id) FROM eventos.%I', tabla) INTO maximo;
        PERFORM setval(secuencia::REGCLASS, COALESCE(maximo, 1), maximo IS NOT NULL);
    END LOOP;
END $$;

-- 13. Documenta el esquema, las tablas y las columnas incorporadas por la migracion.
COMMENT ON SCHEMA eventos IS 'Modelo vigente de calendarios, eventos academicos, extensiones y sesiones.';
COMMENT ON TABLE eventos.proceso IS 'Instancias de procesos academicos asociadas a un calendario.';
COMMENT ON TABLE eventos.proceso_catalogo IS 'Catalogo institucional de procesos academicos.';
COMMENT ON TABLE eventos.evento_catalogo IS 'Catalogo institucional de actividades de calendario.';
COMMENT ON TABLE eventos.evento_catalogo_proceso_catalogo IS 'Relacion N:M entre actividades y procesos donde pueden utilizarse.';
COMMENT ON TABLE eventos.evento_catalogo_rol_gestion IS 'Perfiles externos autorizados para gestionar actividades de catalogo.';
COMMENT ON TABLE eventos.calendario_evento_tipo_publico IS 'Perfiles externos que conforman el publico dirigido de una actividad.';
COMMENT ON TABLE eventos.calendario_evento_extension IS 'Extensiones de fecha autorizadas para una actividad de calendario.';
COMMENT ON TABLE eventos.calendario_evento_extension_programa IS 'Programas cubiertos por una extension y su encadenamiento previo.';
COMMENT ON TABLE eventos.calendario_evento_auditoria IS 'Registro append-only de cambios transaccionales sobre calendario_evento.';

COMMENT ON COLUMN eventos.calendario_evento.proceso_id IS 'Proceso del calendario al que pertenece la actividad.';
COMMENT ON COLUMN eventos.calendario_evento.evento_catalogo_id IS 'Actividad institucional representada por el registro.';
COMMENT ON COLUMN eventos.calendario_evento.numero_ocurrencia IS 'Numero consecutivo de la actividad dentro del mismo proceso y catalogo.';

COMMENT ON COLUMN eventos.proceso_catalogo.id IS 'Identificador del catalogo de procesos academicos.';
COMMENT ON COLUMN eventos.proceso_catalogo.codigo_abreviacion IS 'Codigo institucional unico del proceso academico.';
COMMENT ON COLUMN eventos.proceso_catalogo.nombre IS 'Nombre institucional del proceso academico.';
COMMENT ON COLUMN eventos.proceso_catalogo.descripcion IS 'Descripcion funcional del proceso academico.';
COMMENT ON COLUMN eventos.proceso_catalogo.activo IS 'Indica si el proceso catalogado esta disponible para uso.';
COMMENT ON COLUMN eventos.proceso_catalogo.fecha_creacion IS 'Fecha de creacion del registro.';
COMMENT ON COLUMN eventos.proceso_catalogo.fecha_modificacion IS 'Fecha de la ultima modificacion del registro.';

COMMENT ON COLUMN eventos.evento_catalogo.id IS 'Identificador del catalogo de actividades.';
COMMENT ON COLUMN eventos.evento_catalogo.codigo_abreviacion IS 'Codigo institucional unico de la actividad.';
COMMENT ON COLUMN eventos.evento_catalogo.nombre IS 'Nombre institucional de la actividad.';
COMMENT ON COLUMN eventos.evento_catalogo.descripcion IS 'Descripcion funcional de la actividad.';
COMMENT ON COLUMN eventos.evento_catalogo.activo IS 'Indica si la actividad catalogada esta disponible para uso.';
COMMENT ON COLUMN eventos.evento_catalogo.fecha_creacion IS 'Fecha de creacion del registro.';
COMMENT ON COLUMN eventos.evento_catalogo.fecha_modificacion IS 'Fecha de la ultima modificacion del registro.';

COMMENT ON COLUMN eventos.proceso.id IS 'Identificador de la instancia del proceso.';
COMMENT ON COLUMN eventos.proceso.activo IS 'Indica si la instancia del proceso esta activa.';
COMMENT ON COLUMN eventos.proceso.fecha_creacion IS 'Fecha de creacion de la instancia del proceso.';
COMMENT ON COLUMN eventos.proceso.fecha_modificacion IS 'Fecha de la ultima modificacion de la instancia del proceso.';
COMMENT ON COLUMN eventos.proceso.tipo_recurrencia_id IS 'Tipo de recurrencia aplicado al proceso.';
COMMENT ON COLUMN eventos.proceso.calendario_id IS 'Calendario academico al que pertenece el proceso.';
COMMENT ON COLUMN eventos.proceso.proceso_catalogo_id IS 'Definicion institucional del proceso.';

COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.id IS 'Identificador de la relacion entre actividad y proceso.';
COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.evento_catalogo_id IS 'Actividad institucional permitida para el proceso.';
COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.proceso_catalogo_id IS 'Proceso institucional que admite la actividad.';
COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.repetible IS 'Indica si la actividad puede tener varias ocurrencias dentro del proceso.';
COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.activo IS 'Indica si la relacion entre actividad y proceso esta activa.';
COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.fecha_creacion IS 'Fecha de creacion de la relacion.';
COMMENT ON COLUMN eventos.evento_catalogo_proceso_catalogo.fecha_modificacion IS 'Fecha de la ultima modificacion de la relacion.';

COMMENT ON COLUMN eventos.evento_catalogo_rol_gestion.id IS 'Identificador del permiso de gestion de la actividad.';
COMMENT ON COLUMN eventos.evento_catalogo_rol_gestion.evento_catalogo_id IS 'Actividad institucional sobre la que se concede gestion.';
COMMENT ON COLUMN eventos.evento_catalogo_rol_gestion.perfil_id IS 'Perfil externo autorizado para gestionar la actividad.';
COMMENT ON COLUMN eventos.evento_catalogo_rol_gestion.activo IS 'Indica si el permiso de gestion esta activo.';
COMMENT ON COLUMN eventos.evento_catalogo_rol_gestion.fecha_creacion IS 'Fecha de creacion del permiso.';
COMMENT ON COLUMN eventos.evento_catalogo_rol_gestion.fecha_modificacion IS 'Fecha de la ultima modificacion del permiso.';

COMMENT ON COLUMN eventos.calendario_evento_extension.id IS 'Identificador de la extension de fecha.';
COMMENT ON COLUMN eventos.calendario_evento_extension.calendario_evento_id IS 'Actividad de calendario cuya fecha se extiende.';
COMMENT ON COLUMN eventos.calendario_evento_extension.fecha_fin IS 'Nueva fecha limite autorizada para la actividad.';
COMMENT ON COLUMN eventos.calendario_evento_extension.documento_id IS 'Documento que autoriza o soporta la extension.';
COMMENT ON COLUMN eventos.calendario_evento_extension.descripcion IS 'Justificacion o detalle de la extension.';
COMMENT ON COLUMN eventos.calendario_evento_extension.numero_extension IS 'Consecutivo de la extension para la actividad.';
COMMENT ON COLUMN eventos.calendario_evento_extension.activo IS 'Indica si la extension esta activa.';
COMMENT ON COLUMN eventos.calendario_evento_extension.fecha_creacion IS 'Fecha de creacion de la extension.';
COMMENT ON COLUMN eventos.calendario_evento_extension.fecha_modificacion IS 'Fecha de la ultima modificacion de la extension.';

COMMENT ON COLUMN eventos.calendario_evento_extension_programa.id IS 'Identificador de la aplicacion de una extension a un programa.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.calendario_evento_extension_id IS 'Extension de fecha aplicada al programa.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.extension_padre_id IS 'Extension previa de la que deriva este registro.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.dependencia_id IS 'Identificador del programa academico cubierto por la extension.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.vigente IS 'Indica si esta es la extension vigente para el programa.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.activo IS 'Indica si la aplicacion de la extension esta activa.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.fecha_creacion IS 'Fecha de creacion de la aplicacion de la extension.';
COMMENT ON COLUMN eventos.calendario_evento_extension_programa.fecha_modificacion IS 'Fecha de la ultima modificacion de la aplicacion de la extension.';

COMMENT ON COLUMN eventos.calendario_evento_auditoria.id IS 'Identificador del registro de auditoria.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.calendario_evento_id IS 'Actividad de calendario auditada.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.operacion IS 'Operacion realizada sobre la actividad.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.tercero_id IS 'Tercero responsable de la operacion.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.fecha_creacion IS 'Fecha inmutable de registro de la operacion.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.fecha_modificacion IS 'Fecha de modificacion, igual a la fecha de creacion por ser append-only.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.activo IS 'Estado inmutable del registro de auditoria.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.estado_anterior IS 'Estado de la actividad antes de la operacion.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.estado_nuevo IS 'Estado de la actividad despues de la operacion.';
COMMENT ON COLUMN eventos.calendario_evento_auditoria.cambios IS 'Diferencias entre el estado anterior y el estado nuevo.';

-- 14. Verifica estructura, comentarios, conteos y referencias antes de permitir el Commit de DBeaver.
DO $$
DECLARE
    tablas_esperadas CONSTANT TEXT[] := ARRAY[
        'calendario', 'calendario_evento', 'calendario_evento_auditoria',
        'calendario_evento_extension', 'calendario_evento_extension_programa',
        'calendario_evento_tipo_publico', 'evento_catalogo',
        'evento_catalogo_proceso_catalogo', 'evento_catalogo_rol_gestion',
        'participante_sesion', 'proceso', 'proceso_catalogo', 'relacion_sesiones',
        'rol_participante_sesion', 'sesion', 'sesion_patron_recurrencia',
        'tipo_recurrencia', 'tipo_sesion'
    ];
    tabla TEXT;
BEGIN
    FOREACH tabla IN ARRAY tablas_esperadas LOOP
        IF to_regclass(format('eventos.%I', tabla)) IS NULL THEN
            RAISE EXCEPTION 'Falta la tabla objetivo eventos.%', tabla;
        END IF;
    END LOOP;

    IF (SELECT COUNT(*) FROM information_schema.tables
        WHERE table_schema = 'eventos' AND table_type = 'BASE TABLE') <> 22 THEN
        RAISE EXCEPTION 'El esquema no contiene las 18 tablas objetivo y las 4 legacy pendientes de eliminar';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        WHERE n.nspname = 'eventos'
          AND c.relkind = 'r'
          AND NULLIF(BTRIM(COALESCE(OBJ_DESCRIPTION(c.oid, 'pg_class'), '')), '') IS NULL
    ) OR EXISTS (
        SELECT 1
        FROM pg_class c
        JOIN pg_namespace n ON n.oid = c.relnamespace
        JOIN pg_attribute a ON a.attrelid = c.oid
        WHERE n.nspname = 'eventos'
          AND c.relkind = 'r'
          AND a.attnum > 0
          AND NOT a.attisdropped
          AND NULLIF(BTRIM(COALESCE(COL_DESCRIPTION(c.oid, a.attnum), '')), '') IS NULL
    ) THEN
        RAISE EXCEPTION 'Existen tablas o columnas del esquema eventos sin comentario';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.calendario) <> 13
       OR (SELECT COUNT(*) FROM eventos.proceso_catalogo) <> 1
       OR (SELECT COUNT(*) FROM eventos.proceso) <> 11
       OR (SELECT COUNT(*) FROM eventos.evento_catalogo) <> 4
       OR (SELECT COUNT(*) FROM eventos.evento_catalogo_proceso_catalogo) <> 4
       OR (SELECT COUNT(*) FROM eventos.calendario_evento) <> 26
       OR (SELECT COUNT(*) FROM eventos.calendario_evento_tipo_publico) <> 24
       OR (SELECT COUNT(*) FROM eventos.rol_participante_sesion) <> 2
       OR (SELECT COUNT(*) FROM eventos.tipo_sesion) <> 8 THEN
        RAISE EXCEPTION 'Los conteos objetivo no coinciden con el plan aprobado';
    END IF;

    IF (SELECT ARRAY_AGG(codigo_abreviacion ORDER BY id) FROM eventos.proceso_catalogo)
       IS DISTINCT FROM ARRAY['PROC_ADMISIONES']::VARCHAR[] THEN
        RAISE EXCEPTION 'El catalogo de procesos debe contener unicamente PROC_ADMISIONES';
    END IF;

    IF EXISTS (SELECT 1 FROM eventos.calendario_evento WHERE proceso_id IS NULL OR evento_catalogo_id IS NULL)
       OR EXISTS (SELECT 1 FROM eventos.calendario_evento_tipo_publico WHERE perfil_id NOT IN (157, 187)) THEN
        RAISE EXCEPTION 'El postflight encontro referencias o fixtures no permitidos';
    END IF;

    IF (SELECT ARRAY_AGG(codigo_abreviacion ORDER BY id) FROM eventos.tipo_recurrencia)
       IS DISTINCT FROM ARRAY['SEM', 'ANUA']::VARCHAR[] THEN
        RAISE EXCEPTION 'La recurrencia objetivo no coincide con SEM/ANUA';
    END IF;

    IF (SELECT ARRAY_AGG(codigo_abreviacion ORDER BY id) FROM eventos.evento_catalogo WHERE id <= 4)
       IS DISTINCT FROM ARRAY['REIN', 'INSCR', 'PAGO_REIN', 'PAGO_INSC']::VARCHAR[] THEN
        RAISE EXCEPTION 'Los cuatro eventos canonicos no coinciden con el baseline';
    END IF;

    IF (SELECT ARRAY_AGG(codigo_abreviacion ORDER BY id) FROM eventos.rol_participante_sesion)
       IS DISTINCT FROM ARRAY['CE', 'PE']::VARCHAR[]
       OR (SELECT ARRAY_AGG(codigo_abreviacion ORDER BY id) FROM eventos.tipo_sesion)
       IS DISTINCT FROM ARRAY['EMP', 'PMP', 'SMP', 'PFSAMP', 'PFFMP', 'SCPPE', 'SFFMP', 'ASIMP']::VARCHAR[] THEN
        RAISE EXCEPTION 'La parametria de sesiones no coincide con el baseline';
    END IF;

    IF EXISTS (SELECT 1 FROM eventos.calendario_evento_extension)
       OR EXISTS (SELECT 1 FROM eventos.calendario_evento_extension_programa)
       OR EXISTS (SELECT 1 FROM eventos.calendario_evento_auditoria)
       OR EXISTS (SELECT 1 FROM eventos.evento_catalogo_rol_gestion) THEN
        RAISE EXCEPTION 'Las tablas nuevas que deben iniciar vacias contienen datos';
    END IF;

    IF (SELECT COUNT(*) FROM eventos.tipo_evento) <> 15
       OR (SELECT COUNT(*) FROM eventos.tipo_publico) <> 4
       OR (SELECT COUNT(*) FROM eventos.encargado_evento) <> 2 THEN
        RAISE EXCEPTION 'Las tablas descartadas no conservan la evidencia esperada';
    END IF;
END $$;

-- 15. Elimina las tablas legacy cuya informacion ya fue migrada o descartada.
DROP TABLE eventos.encargado_evento;
DROP TABLE eventos.rol_encargado_evento;
DROP TABLE eventos.tipo_publico;
DROP TABLE eventos.tipo_evento;

-- 16. Confirma que solo permanezcan las 18 tablas del modelo vigente.
DO $$
BEGIN
    IF (SELECT COUNT(*) FROM information_schema.tables
        WHERE table_schema = 'eventos' AND table_type = 'BASE TABLE') <> 18 THEN
        RAISE EXCEPTION 'El esquema actualizado no contiene exactamente 18 tablas';
    END IF;
END $$;
