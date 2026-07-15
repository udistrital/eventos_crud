package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:AuditoriaEventosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:AuditoriaEventosController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoExtensionProgramaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:CalendarioEventoTipoPublicoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:EventoCatalogoRolGestionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoCatalogoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:ProcesoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RelacionSesionesController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolEncargadoEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolGestionEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:RolParticipanteSesionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:SesionPatronRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoPublicoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoRecurrenciaController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method:           "GetAll",
			Router:           "/",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method:           "GetOne",
			Router:           "/:id",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TipoSesionController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "Post",
			Router:           "/",
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "Put",
			Router:           "/:id",
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "Delete",
			Router:           "/:id",
			AllowHTTPMethods: []string{"delete"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/eventos_crud/controllers:TrEventoController"],
		beego.ControllerComments{
			Method:           "GetAllByPersona",
			Router:           "/:persona",
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
