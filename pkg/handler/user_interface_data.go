package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get General User Interface Data
// @Security ApiKeyAuth
// @Tags userInterfaceData
// @Description get user interface data
// @ID get-user-interface-data
// @Accept json
// @Produce json
// @Success 200 {object} UserInterfaceData
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user-interface-data [get]
func (h *Handler) getGeneralUserInterfaceData(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}

	userInterfaceData, err := h.services.UserInterfaceData.GetAll(false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userInterfaceData)
}

// @Summary Get General User Interface Data by Battle ID
// @Security ApiKeyAuth
// @Tags userInterfaceData
// @Description get user interface data by battle id
// @ID get-user-interface-data
// @Accept json
// @Produce json
// @Param id path int true "battle id"
// @Success 200 {object} Battle
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/user-interface-data/{id} [get]
func (h *Handler) getGeneralUserInterfaceDataByBattleId(c *gin.Context) {
	_, err := getUserId(c)
	if err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	userInterfaceData, err := h.services.UserInterfaceData.GetById(id, false)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, userInterfaceData)
}
