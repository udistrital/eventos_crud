// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/planesticud/eventos_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/calendario_evento",
			beego.NSInclude(
				&controllers.CalendarioEventoController{},
			),
		),

		beego.NSNamespace("/tipo_recurrencia",
			beego.NSInclude(
				&controllers.TipoRecurrenciaController{},
			),
		),

		beego.NSNamespace("/tipo_evento",
			beego.NSInclude(
				&controllers.TipoEventoController{},
			),
		),

		beego.NSNamespace("/rol_encargado_evento",
			beego.NSInclude(
				&controllers.RolEncargadoEventoController{},
			),
		),

		beego.NSNamespace("/encargado_evento",
			beego.NSInclude(
				&controllers.EncargadoEventoController{},
			),
		),

		beego.NSNamespace("/tipo_publico",
			beego.NSInclude(
				&controllers.TipoPublicoController{},
			),
		),

		beego.NSNamespace("/tr_evento",
			beego.NSInclude(
				&controllers.TrEventoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
