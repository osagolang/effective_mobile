package handler

import (
	"github.com/jackc/pgx/v5"
	"net/http"
	"strconv"
	"time"

	"effective_mobile/internal/model"
	"effective_mobile/internal/repository"
	"effective_mobile/internal/service"
	"effective_mobile/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Handler struct {
	service *service.Service
	repo    *repository.Repo
}

func NewHandler(service *service.Service, repo *repository.Repo) *Handler {
	return &Handler{service: service, repo: repo}
}

// CreateSubscriptionHandler godoc
// @Summary      Создать подписку
// @Description  Создает новую подписку
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        input body model.CreateSubscriptionRequest true "Данные подписки"
// @Success      201 {object} model.Subscription
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/subscriptions [post]
func (h *Handler) CreateSubscriptionHandler(c *gin.Context) {

	var request *model.CreateSubscriptionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		logger.Warn("incorrect body request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect body request", "error": err})
		return
	}

	sub, err := h.service.CreateSubscription(c.Request.Context(), request)
	if err != nil {
		logger.Error("fail create subscription", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"type": "incorrect request", "error": err})
		return
	}

	c.JSON(http.StatusCreated, sub)
}

// GetSubscriptionHandler godoc
// @Summary      Получить подписку
// @Description  Получает подписку по id
// @Tags         subscriptions
// @Produce      json
// @Param        id path int true "ID подписки"
// @Success      200 {object} model.Subscription
// @Failure      400 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/subscriptions/{id} [get]
func (h *Handler) GetSubscriptionHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect subscription id", "error": err})
		logger.Warn("incorrect subscription id", zap.Error(err))
		return
	}

	sub, err := h.repo.GetSubscriptionByID(c.Request.Context(), int64(id))
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"type": "id not found", "error": err})
			logger.Warn("id not found", zap.Error(err))
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"type": "incorrect request", "error": err})
		logger.Error("fail get subscription", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, sub)
}

// UpdateSubscriptionHandler godoc
// @Summary      Обновить подписку
// @Description  Обновляет данные подписки по id
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id path int true "ID подписки"
// @Param        input body model.UpdateSubscriptionRequest true "Данные для обновления"
// @Success      200 {object} model.Subscription
// @Failure      400 {object} map[string]interface{}
// @Failure      404 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/subscriptions/{id} [put]
func (h *Handler) UpdateSubscriptionHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect subscription id", "error": err})
		logger.Warn("incorrect subscription id", zap.Error(err))
		return
	}

	var request *model.UpdateSubscriptionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect body request", "error": err})
		logger.Warn("incorrect body request", zap.Error(err))
		return
	}

	sub, err := h.service.UpdateSubscription(c.Request.Context(), int64(id), request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"type": "incorrect request", "error": err})
		logger.Error("fail update subscription", zap.Error(err))
		return
	}
	if sub == nil {
		c.JSON(http.StatusNotFound, gin.H{"type": "subscription not found", "error": err})
		logger.Warn("subscription not found", zap.Error(err))
	}

	c.JSON(http.StatusOK, sub)

}

// DeleteSubscriptionHandler godoc
// @Summary      Удалить подписку
// @Description  Удаляет подписку по id
// @Tags         subscriptions
// @Param        id path int true "ID подписки"
// @Success      204
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/subscriptions/{id} [delete]
func (h *Handler) DeleteSubscriptionHandler(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect subscription id", "error": err})
		logger.Warn("incorrect subscription id", zap.Error(err))
		return
	}

	if err := h.repo.DeleteSubscription(c.Request.Context(), int64(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"type": "incorrect request", "error": err})
		logger.Error("fail delete subscription", zap.Error(err))
		return
	}

	c.Status(http.StatusNoContent)
}

// ListSubscriptionHandler godoc
// @Summary      Получить список подписок
// @Description  Получает список подписок с фильтрами
// @Tags         subscriptions
// @Produce      json
// @Param        user_id query string false "ID пользователя (UUID)"
// @Param        service_name query string false "Название сервиса"
// @Param        price query int false "Стоимость подписки"
// @Param        start_date query string false "Дата начала (01-2006)"
// @Param        end_date query string false "Дата окончания (01-2006)"
// @Param        limit query int false "Лимит"
// @Param        offset query int false "Смещение"
// @Success      200 {array} model.Subscription
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/subscriptions [get]
func (h *Handler) ListSubscriptionHandler(c *gin.Context) {

	filter := &model.Filter{}

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect filter userID", "error": err})
			logger.Warn("incorrect filter userID", zap.Error(err))
			return
		}
		filter.UserID = &userID
	}

	if serviceName := c.Query("service_name"); serviceName != "" {
		filter.ServiceName = &serviceName
	}

	if priceStr := c.Query("price"); priceStr != "" {
		price, err := strconv.Atoi(priceStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect filter price", "error": err})
			logger.Warn("incorrect filter price", zap.Error(err))
			return
		}
		filter.Price = &price
	}

	if startDateStr := c.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse("01-2006", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect start date format", "error": err})
			logger.Warn("incorrect start date format", zap.Error(err))
			return
		}
		filter.StartDate = &startDate
	}

	if endDateStr := c.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse("01-2006", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect end date format", "error": err})
			logger.Warn("incorrect end date format", zap.Error(err))
			return
		}
		filter.EndDate = &endDate
	}

	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect filter limit", "error": err})
			logger.Warn("incorrect filter limit", zap.Error(err))
			return
		}
		filter.Limit = &limit
	}

	if offsetStr := c.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect filter offset", "error": err})
			logger.Warn("incorrect filter offset", zap.Error(err))
			return
		}
		filter.Offset = &offset
	}

	filter.Normalize()

	subs, err := h.repo.ListSubscription(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"type": "internal error", "error": err})
		logger.Error("fail list subscription", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, subs)
}

// GetTotalCostHandler godoc
// @Summary      Получить суммарную стоимость подписок
// @Description  Считает суммарную стоимость подписок с фильтрами
// @Tags         subscriptions
// @Produce      json
// @Param        user_id query string false "ID пользователя (UUID)"
// @Param        service_name query string false "Название сервиса"
// @Param        start_date query string false "Дата начала (01-2006)"
// @Param        end_date query string false "Дата окончания (01-2006)"
// @Success      200 {object} map[string]int
// @Failure      400 {object} map[string]interface{}
// @Failure      500 {object} map[string]interface{}
// @Router       /api/subscriptions/totalcost [get]
func (h *Handler) GetTotalCostHandler(c *gin.Context) {

	filter := &model.Filter{}

	if userIDStr := c.Query("user_id"); userIDStr != "" {
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect filter userID", "error": err})
			logger.Warn("incorrect filter userID", zap.Error(err))
			return
		}
		filter.UserID = &userID
	}

	if serviceName := c.Query("service_name"); serviceName != "" {
		filter.ServiceName = &serviceName
	}

	if startDateStr := c.Query("start_date"); startDateStr != "" {
		startDate, err := time.Parse("01-2006", startDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect start date format", "error": err})
			logger.Warn("incorrect start date format", zap.Error(err))
			return
		}
		filter.StartDate = &startDate
	}

	if endDateStr := c.Query("end_date"); endDateStr != "" {
		endDate, err := time.Parse("01-2006", endDateStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"type": "incorrect end date format", "error": err})
			logger.Warn("incorrect end date format", zap.Error(err))
			return
		}
		filter.EndDate = &endDate
	}

	totalCost, err := h.repo.TotalCost(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"type": "filed get total cost", "error": err})
		logger.Error("fail get total cost", zap.Error(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{"Total_price": totalCost})
}
