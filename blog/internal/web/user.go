package web

import (
	"blog/internal/domain"
	"blog/internal/service"
	"net/http"
	"time"

	// "regexp"
	regexp "github.com/dlclark/regexp2"
	// "github.com/golang-jwt/jwt/v5"
	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/gin-gonic/gin"
)

const (
	emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
)

var JWTKey = []byte("k6CswdUm77WKcbM68UQUuxVsHSpTCwgK")

type UserHandler struct {
	emailRexExp *regexp.Regexp
	svc         *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{
		emailRexExp: regexp.MustCompile(emailRegexPattern, regexp.None),
		svc:         svc,
	}
}

func (h *UserHandler) RegisterRoutes(server *gin.Engine) {
	ug := server.Group("/users")
	//注册
	ug.POST("/signup", h.SignUp)
	//登录
	ug.POST("/login", h.Login)
}

func (h *UserHandler) SignUp(ctx *gin.Context) {
	type SignUpReq struct {
		UserName        string `json:"username"`
		Email           string `json:"email"`
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	if err := ctx.Bind(&req); err != nil {
		return
	}

	isEmail, err := h.emailRexExp.MatchString(req.Email)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, Response{
			ErrorCode: "system_error",
			Message:   "系统错误",
		})
		return
	}
	if !isEmail {
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "email_error",
			Message:   "邮箱错误",
		})
		return
	}
	if req.Password != req.ConfirmPassword {
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "password_error",
			Message:   "密码错误",
		})
		return
	}

	err = h.svc.Signup(ctx, domain.User{
		UserName: req.UserName,
		Email:    req.Email,
		Password: req.Password,
	})
	switch err {
	case nil:
		ctx.JSON(http.StatusOK, "注册成功")
	case service.ErrDuplicateEmail:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "email_Duplicate",
			Message:   "邮箱冲突",
		})
	default:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "system_error",
			Message:   "系统错误",
		})
	}
}

func (h *UserHandler) Login(ctx *gin.Context) {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var req Req
	if err := ctx.Bind(&req); err != nil {
		return
	}
	u, err := h.svc.Login(ctx, req.Email, req.Password)
	switch err {
	case nil:
		uc := UserClaims{
			Uid:       u.Id,
			UserAgent: ctx.GetHeader("User-Agent"),
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, uc)
		tokenStr, err := token.SignedString(JWTKey)
		if err != nil {
			ctx.JSON(http.StatusOK, SystemError)
		}
		ctx.Header("x-jwt-token", tokenStr)
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "",
			Message:   "登录成功",
			Data: jwtStr{
				Token: "Bearer " + tokenStr,
			},
		})
	case service.ErrInvalidUserOrPassword:
		ctx.JSON(http.StatusOK, Response{
			ErrorCode: "UserInfoError",
			Message:   "用户名或者密码错误",
		})
	default:
		ctx.JSON(http.StatusOK, SystemError)
	}
}

type jwtStr struct {
	Token string
}

type UserClaims struct {
	jwt.RegisteredClaims
	Uid       int64
	UserAgent string
}
