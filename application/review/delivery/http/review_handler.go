package http

import (
	"github.com/go-park-mail-ru/2020_1_k-on/application/models"
	"github.com/go-park-mail-ru/2020_1_k-on/application/review"
	"github.com/go-park-mail-ru/2020_1_k-on/application/server/middleware"
	"github.com/go-park-mail-ru/2020_1_k-on/pkg/constants"
	"github.com/labstack/echo"
	"github.com/mailru/easyjson"
	"github.com/microcosm-cc/bluemonday"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type ReviewHandler struct {
	film      review.UseCase
	series    review.UseCase
	logger    *zap.Logger
	sanitizer *bluemonday.Policy
}

func NewReviewHandler(e *echo.Echo, review review.UseCase, series review.UseCase,
	auth middleware.Auth, logger *zap.Logger, sanitizer *bluemonday.Policy) {
	handler := ReviewHandler{
		film:      review,
		series:    series,
		logger:    logger,
		sanitizer: sanitizer,
	}

	e.POST("/films/:id/reviews", handler.AddFilmReview, auth.GetSession, middleware.ParseErrors, middleware.CSRF)
	e.POST("/series/:id/reviews", handler.AddSeriesReview, auth.GetSession, middleware.ParseErrors, middleware.CSRF)
	e.GET("/films/:id/reviews", handler.GetByFilm, middleware.ParseErrors)
	e.GET("/series/:id/reviews", handler.GetBySeries, middleware.ParseErrors)
	e.GET("/films/:id/reviews/user", handler.GetByFilmAndUser, auth.GetSession, middleware.ParseErrors)
	e.GET("/series/:id/reviews/user", handler.GetBySeriesAndUser, auth.GetSession, middleware.ParseErrors)
}

func (r *ReviewHandler) AddFilmReview(ctx echo.Context) error {
	id, err := r.getProductId(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	rev, err := r.parseRequestBody(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "request parser error")
	}

	rev.ProductId = id
	if err := r.film.Add(rev); err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, rev)
}

func (r *ReviewHandler) AddSeriesReview(ctx echo.Context) error {
	id, err := r.getProductId(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	rev, err := r.parseRequestBody(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "request parser error")
	}

	rev.ProductId = id
	if err := r.series.Add(rev); err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, rev)
}

func (r *ReviewHandler) GetByFilm(ctx echo.Context) error {
	id, err := r.getProductId(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	reviews, err := r.film.GetByProductId(uint(id))
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, reviews)
}

func (r *ReviewHandler) GetBySeries(ctx echo.Context) error {
	id, err := r.getProductId(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}

	reviews, err := r.series.GetByProductId(uint(id))
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, reviews)
}

func (r *ReviewHandler) GetByFilmAndUser(ctx echo.Context) error {
	productId, err := r.getProductId(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	rev, err := r.film.GetReview(productId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, rev)
}

func (r *ReviewHandler) GetBySeriesAndUser(ctx echo.Context) error {
	productId, err := r.getProductId(ctx)
	if err != nil {
		return middleware.WriteErrResponse(ctx, http.StatusBadRequest, "wrong parameter")
	}
	userId := ctx.Get(constants.UserIdKey).(uint)

	rev, err := r.series.GetReview(productId, userId)
	if err != nil {
		return err
	}

	return middleware.WriteOkResponse(ctx, rev)
}

func (r *ReviewHandler) parseRequestBody(ctx echo.Context) (*models.Review, error) {
	rev := new(models.Review)
	if err := easyjson.UnmarshalFromReader(ctx.Request().Body, rev); err != nil {
		r.logger.Error("request parser error")
		return nil, err
	}

	rev.Body = r.sanitizer.Sanitize(rev.Body)
	rev.UserId = ctx.Get(constants.UserIdKey).(uint)

	return rev, nil
}

func (r *ReviewHandler) getProductId(ctx echo.Context) (uint, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil || id < 0 {
		r.logger.Error("wrong parameter")
		return 0, err
	}

	return uint(id), nil
}
