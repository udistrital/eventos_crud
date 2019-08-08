INSERT INTO EVENTO.ROL_ENCARGADO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion, numero_de_orden) VALUES (1,'Creador','Creador', 'CRE', true, now(), now(), 1);
INSERT INTO EVENTO.ROL_ENCARGADO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion, numero_de_orden) VALUES (2,'Actualizador','Actualizador', 'act', true, now(), now(), 2);

INSERT INTO EVENTO.TIPO_RECURRENCIA (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (1,'Anual','Anual', 'ANO', true, now(), now());
INSERT INTO EVENTO.TIPO_RECURRENCIA (id, nombre, descripcion, codigo_abreviacion, activo, fecha_creacion, fecha_modificacion) VALUES (2,'Semestral','Semestral', 'SEM', true, now(), now());

INSERT INTO EVENTO.TIPO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, dependencia_id, fecha_creacion, fecha_modificacion, tipo_recurrencia_id) VALUES (1,'Inscripcion','Inscripcion de espacios', 'INS', true, 122, now(), now(), 1);
INSERT INTO EVENTO.TIPO_EVENTO (id, nombre, descripcion, codigo_abreviacion, activo, dependencia_id, fecha_creacion, fecha_modificacion, tipo_recurrencia_id) VALUES (2,'Grado','Grado', 'GRD', true, 122, now(), now(), 2);

