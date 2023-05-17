package http

import (
	"net/http"

	"github.com/dev-mockingbird/bird"
	"github.com/dev-mockingbird/errors"
	"github.com/dev-mockingbird/logf"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"udious.com/mockingbird/channel/http/docs"
	"udious.com/mockingbird/channel/service"
)

type Handler struct {
	Service *service.Service
}

// @BasePath /api/v1

// HelloWorld hello world
// @Summary hello world
// @Schemes
// @Description hello world
// @Tags helloworld
// @Accept json
// @Produce json
// @Param name query string true "hello world"
// @Param description query string false "channel description"
// @Param group query string true "channel group"
// @Param creator_id query string true "creator id"
// @Param subscribe_user_id query []string false "subscribe user id"
// @Success 200 {string} model.id
// @Router /users/:userId/channels [post]
func (h Handler) HelloWorld(actor bird.Actor) {
	h.msg(actor, "ok", "hello world")
}

func (h Handler) error(actor bird.Actor, err error, codes ...string) {
	actor.Logger().Logf(logf.Error, err.Error())
	code := func() string {
		if len(codes) > 0 {
			return codes[0]
		}
		return ""
	}()
	tags, msg := errors.Parse(err)
	if len(tags) > 0 {
		code = tags[0]
	}
	if code == "" {
		actor.Write(http.StatusInternalServerError, bird.UnknownError(errors.New("internal error occurred, please try again later")))
		return
	}
	actor.Write(h.statusCode(code), bird.ResponseBody{Code: code, Data: bird.Msg(msg)})
}

func (h Handler) ok(actor bird.Actor, data any) {
	actor.Write(http.StatusOK, bird.OK(data))
}

func (h Handler) msg(actor bird.Actor, code, msg string) {
	actor.Write(h.statusCode(code), bird.ResponseBody{Code: code, Data: bird.Msg(msg)})
}

func (h Handler) statusCode(code string) int {
	return map[string]int{
		service.InvalidArguments: http.StatusBadRequest,
	}[code]
}

func RegisterGin(e *gin.Engine, logger logf.Logger, s *service.Service) {
	h := Handler{Service: s}
	docs.SwaggerInfo.BasePath = "/api/v1"
	r := bird.GinRouter(e, logger)
	v1 := r.Group("/api/v1")
	{
		v1.ON("/hello-world", h.HelloWorld).Prepare(http.MethodGet)
	}
	e.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
