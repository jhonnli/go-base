package filter

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jhonnli/go-base/controller/common"
	"github.com/jhonnli/golibs"
	"net/http"
	"strings"
)

// JWT基本信息
type JwtInfo struct {
	AppID    int    `json:"appId"`    //应用编号
	Exp      int64  `json:"exp"`      //过期时间
	Iat      int64  `json:"iat"`      //颁发时间
	Token    string `json:"token"`    //token
	UserId   int    `json:"userId"`   //用户编号
	UserName string `json:"userName"` //用户姓名
	IP       string `json:"ip"`       //IP地址
	RoleIds  string `json:"roleIds"`  //用户角色编码列表
}

// 中间件，检查用户token
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		//region 获取jwt信息
		tokenVal := c.Request.Header.Get("Authorization")
		if tokenVal == "" {
			c.AbortWithStatusJSON(http.StatusOK, common.Response{
				Code:    "middleware.jwt.error",
				Message: "请求未携带token，无权限访问",
			})
			return
		}
		if !strings.HasPrefix(tokenVal, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusOK, common.Response{
				Code:    "middleware.jwt.error",
				Message: "请求未携带token，无权限访问",
			})
			return
		}
		if golibs.Length(tokenVal) < 128 {
			c.AbortWithStatusJSON(http.StatusOK, common.Response{
				Code:    "middleware.jwt.error",
				Message: "请求未携带超过128位的token参数，无权限访问",
			})
			return
		}

		tokenVal = golibs.SubString(tokenVal, golibs.Length("Bearer "), golibs.Length(tokenVal)-golibs.Length("Bearer "))
		//endregion

		//region 解析jwt信息
		var jwtInfo JwtInfo
		tokenInfo, err := jwt.Parse(tokenVal, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return []byte(""), errors.New("签名方法不正确")
			}
			arrs := strings.Split(token.Raw, ".")
			if len(arrs) == 3 {
				b64 := golibs.HexStringToBytes(arrs[1])
				err := json.Unmarshal(b64, &jwtInfo)
				if err != nil {
					return []byte(""), err
				}
				if jwtInfo.Exp < golibs.Unix() {
					return []byte(""), errors.New("jwt信息已过期")
				}
				if jwtInfo.IP != c.ClientIP() {
					return []byte(""), errors.New("IP地址错误:" + c.ClientIP())
				}
				if jwtInfo.AppID < 100 {
					return []byte(""), errors.New("无效的应用标识")
				}
			}
			return []byte(""), errors.New("没有找到应用密钥信息")
		})

		if err != nil {
			if !strings.Contains(err.Error(), "没有找到应用密钥信息") {
				c.AbortWithStatusJSON(http.StatusOK, common.Response{
					Code:    "publish.center.middleware.jwt.error",
					Message: "解析token失败:" + err.Error(),
				})
				return
			}
		}

		_, ok := tokenInfo.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusOK, common.Response{
				Code:    "publish.center.middleware.jwt.error",
				Message: "解析token失败了",
			})
			return
		}
		//endregion
	}
}

// 从SSO获取新的jwt信息响应结果
type GetReturnValue struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Token string `json:"token"`
	} `json:"data"`
}
