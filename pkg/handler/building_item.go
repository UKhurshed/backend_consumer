package handler

import (
	"backend_consumer/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary CreateBuildingItem
// @Description create a new building item
// @Accept json
// @Produce json
// @Param input body domain.Building true "building"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/ [post]
func (h *Handler) CreateBuildingItem(c *gin.Context) {
	var input domain.Building

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateBuildingItem(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

// @Summary GetAllBuildings
// @Description Get all buildings
// @Accept json
// @Produce json
// @Success 200 {array} domain.Building
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/ [get]
func (h *Handler) GetAllBuildings(c *gin.Context) {

	items, err := h.services.GetAll()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, DataResponse{items})
}

// @Summary UpdateBuildingItem
// @Description update chosen item
// @Accept json
// @Produce json
// @Param input body domain.BuildingUpdateInput true "update building"
// @Success 200 {string} Status "ok"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/ [put]
func (h *Handler) UpdateBuildingItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	var input domain.BuildingUpdateInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.services.Update(id, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})

}

// @Summary DeleteBuildingItem
// @Description delete item with id
// @Accept json
// @Produce json
// @Param id path string true "building id"
// @Success 200 {string} Status "ok"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/ [delete]
func (h *Handler) DeleteBuildingItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	err = h.services.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
