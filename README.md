# sesiones_crud
Api para administración de sesiones o periodos.
API de core, Integración con CI
sesiones_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/sesiones_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `SESIONES_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `SESIONES_CRUD__PGUSER`: Usuario de la base de datos
 - `SESIONES_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `SESIONES_CRUD__PGURLS`: Host de conexión
 - `SESIONES_CRUD__PGDB`: Nombre de la base de datos
 - `SESIONES_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
SESIONES_HTTP_PORT=8095 SESIONES_CRUD__PGUSER=postgres SESIONES_CRUD__PGPASS=password SESIONES_CRUD__PGURLS=localhost SESIONES_CRUD__PGDB=local SESIONES_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/sesiones_crud/blob/develop/modelo_sesiones_crud.png).
