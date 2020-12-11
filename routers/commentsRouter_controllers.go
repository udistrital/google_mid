package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/google_mid/controllers:DriveController"] = append(beego.GlobalControllerRouter["github.com/udistrital/google_mid/controllers:DriveController"],
        beego.ControllerComments{
            Method: "PostFileDrive",
            Router: "/:produccion_id/:metadato_id/:meta_produccion_id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/google_mid/controllers:NotificacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/google_mid/controllers:NotificacionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
