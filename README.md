# eventos_crud
API para la gestión de eventos

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `eventos_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/eventos_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `API_EVENTOS_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `EVENTOS_CRUD__PGUSER`: Usuario de la base de datos
 - `EVENTOS_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `EVENTOS_CRUD__PGURLS`: Host de conexión
 - `EVENTOS_CRUD__PGDB`: Nombre de la base de datos
 - `EVENTOS_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
EVENTOS_HTTP_PORT=8081 EVENTOS_CRUD__PGUSER=user EVENTOS_CRUD__PGPASS=password EVENTOS_CRUD__PGURLS=localhost EVENTOS_CRUD__PGDB=bd EVENTOS_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/udistrital/eventos_crud/blob/develop/modelo_eventos_crud.png).
