package routes

import (
	"project/group3/factory"
	_middleware "project/group3/features/middlewares"
	_validatorEvents "project/group3/validator/events"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New(presenter factory.Presenter) *echo.Echo {
	// presenter := factory.InitFactory()
	e := echo.New()

	e.HTTPErrorHandler = _validatorEvents.ErroHandlerEvent

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE, echo.PATCH, echo.OPTIONS},
	}))
	e.Pre(middleware.RemoveTrailingSlash())

	// users
	e.POST("/users", presenter.UserPresenter.PostUser)
	e.POST("/login", presenter.UserPresenter.LoginAuth)
	e.PUT("/users", presenter.UserPresenter.PutUser, _middleware.JWTMiddleware())
	e.GET("/myprofile", presenter.UserPresenter.GetByMe, _middleware.JWTMiddleware())
	e.DELETE("/users", presenter.UserPresenter.DeleteByID, _middleware.JWTMiddleware())

	// events
	e.POST("/events", presenter.EventPresenter.CreateEvent, _middleware.JWTMiddleware())
	e.GET("/events/:id", presenter.EventPresenter.DetailEvent)
	e.PUT("/events/:id", presenter.EventPresenter.EditEvent, _middleware.JWTMiddleware())
	e.DELETE("/events/:id", presenter.EventPresenter.DestroyEvent, _middleware.JWTMiddleware())
	e.POST("/join-event/:id", presenter.EventPresenter.JoinEvent, _middleware.JWTMiddleware())
	e.GET("/events", presenter.EventPresenter.AllEvent)
	e.GET("/my-events", presenter.EventPresenter.MyEvent, _middleware.JWTMiddleware())
	// attendee
	e.GET("/attendees/:id", presenter.EventPresenter.AttendeeEvent)

	// comments //
	e.POST("/comments/:id", presenter.CommentPresenter.PostComment, _middleware.JWTMiddleware())
	e.GET("/comments/:id", presenter.CommentPresenter.GetComment)
	e.DELETE("/comments/:id", presenter.CommentPresenter.DeleteComment, _middleware.JWTMiddleware())
	return e

}
