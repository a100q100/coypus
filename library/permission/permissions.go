package permission

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
	"github.com/gogf/gf/g/net/ghttp"
	"github.com/hequan2017/coypus/library/inject"
	jwtGet "github.com/hequan2017/coypus/library/jwt"
	"strings"
)

func CasbinMiddleware(r *ghttp.Request)  {

	Authorization := r.Header.Get("Authorization")
	token := strings.Split(Authorization, " ")
		t, _ := jwt.Parse(token[1], func(*jwt.Token) (interface{}, error) {
			return jwtGet.JwtSecret(), nil
		})
		glog.Info("-----权限验证-----",jwtGet.GetIdFromClaims("username", t.Claims), r.Request.URL.Path, r.Request.Method)

		if b, err := inject.Obj.Enforcer.EnforceSafe(jwtGet.GetIdFromClaims("username", t.Claims), r.Request.URL.Path, r.Request.Method); err != nil {
			_ = r.Response.WriteJson(g.Map{
				"err":  2,
				"msg":  "错误,未成功获取用户名称",
				"data": nil,
			})
			r.ExitAll()
		} else if !b {
			_ = r.Response.WriteJson(g.Map{
				"err":  2,
				"msg":  "登录用户 没有权限",
				"data": nil,
			})
			r.ExitAll()
		}
}
