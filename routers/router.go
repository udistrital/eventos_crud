// @APIVersion 1.0.0
// @Title beego Eventos Crud
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/eventos_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {

	ns1 := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_sesion",
			beego.NSInclude(
				&controllers.TipoSesionController{},
			),
		),

		beego.NSNamespace("/sesion",
			beego.NSInclude(
				&controllers.SesionController{},
			),
		),

		beego.NSNamespace("/participante_sesion",
			beego.NSInclude(
				&controllers.ParticipanteSesionController{},
			),
		),

		beego.NSNamespace("/rol_participante_sesion",
			beego.NSInclude(
				&controllers.RolParticipanteSesionController{},
			),
		),

		beego.NSNamespace("/sesion_patron_recurrencia",
			beego.NSInclude(
				&controllers.SesionPatronRecurrenciaController{},
			),
		),

		beego.NSNamespace("/relacion_sesiones",
			beego.NSInclude(
				&controllers.RelacionSesionesController{},
			),
		),

		beego.NSNamespace("/tipo_recurrencia",
			beego.NSInclude(
				&controllers.TipoRecurrenciaController{},
			),
		),

		beego.NSNamespace("/tipo_recurrencia",
			beego.NSInclude(
				&controllers.TipoRecurrenciaController{},
			),
		),

		beego.NSNamespace("/proceso",
			beego.NSInclude(
				&controllers.ProcesoController{},
			),
		),

		beego.NSNamespace("/proceso_catalogo",
			beego.NSInclude(
				&controllers.ProcesoCatalogoController{},
			),
		),

		beego.NSNamespace("/evento_catalogo",
			beego.NSInclude(
				&controllers.EventoCatalogoController{},
			),
		),

		beego.NSNamespace("/evento_catalogo_rol_gestion",
			beego.NSInclude(
				&controllers.EventoCatalogoRolGestionController{},
			),
		),

		beego.NSNamespace("/evento_catalogo_proceso_catalogo",
			beego.NSInclude(
				&controllers.EventoCatalogoProcesoCatalogoController{},
			),
		),

		beego.NSNamespace("/calendario_evento",
			beego.NSInclude(
				&controllers.CalendarioEventoController{},
			),
		),

		beego.NSNamespace("/calendario",
			beego.NSInclude(
				&controllers.CalendarioController{},
			),
		),

		beego.NSNamespace("/calendario_evento_tipo_publico",
			beego.NSInclude(
				&controllers.CalendarioEventoTipoPublicoController{},
			),
		),

		beego.NSNamespace("/calendario_evento_extension",
			beego.NSInclude(
				&controllers.CalendarioEventoExtensionController{},
			),
		),

		beego.NSNamespace("/calendario_evento_extension_programa",
			beego.NSInclude(
				&controllers.CalendarioEventoExtensionProgramaController{},
			),
		),
	)

	beego.AddNamespace(ns1)
}
