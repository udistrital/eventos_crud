// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/sesiones_crud/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/tipo_sesion",
			beego.NSInclude(
				&controllers.TipoSesionController{},
			),
		),

		beego.NSNamespace("/rol_participante_sesion",
			beego.NSInclude(
				&controllers.RolParticipanteSesionController{},
			),
		),

		beego.NSNamespace("/participante_sesion",
			beego.NSInclude(
				&controllers.ParticipanteSesionController{},
			),
		),

		beego.NSNamespace("/sesion",
			beego.NSInclude(
				&controllers.SesionController{},
			),
		),

		beego.NSNamespace("/relacion_sesiones",
			beego.NSInclude(
				&controllers.RelacionSesionesController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
