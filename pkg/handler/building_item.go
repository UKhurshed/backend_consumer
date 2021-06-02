package handler

import (
	"backend_consumer/pkg/domain"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary CreateBuildingItem
// @Security ApiKeyAuth
// @Tags Объект
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
		newResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.CreateBuildingItem(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

// @Summary GetAllBuildings
// @Security ApiKeyAuth
// @Tags Объект
// @Description Get all buildings
// @Accept json
// @Produce json
// @Success 200 {array} dataResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/ [get]
func (h *Handler) GetAllBuildings(c *gin.Context) {

	//params
	nameBuilding := c.Request.URL.Query().Get("name_building")
	typeOfObject := c.Request.URL.Query().Get("type_object")
	networkTrading := c.Request.URL.Query().Get("network_trading")
	region := c.Request.URL.Query().Get("name_region")
	microDistrict := c.Request.URL.Query().Get("micro_district")
	streetName := c.Request.URL.Query().Get("street_name")
	openIn := c.Request.URL.Query().Get("open_in")


	if len(typeOfObject) == 0 {
		typeOfObject = strconv.Itoa(0)
	}

	if len(networkTrading) == 0 {
		networkTrading = strconv.Itoa(0)
	}

	if len(region) == 0 {
		region = strconv.Itoa(0)
	}

	fmt.Println("Network Trading: ", networkTrading)
	fmt.Println("TypeOfObject: ", typeOfObject)
	fmt.Println("Region: ", region)

		items, err := h.services.GetAll(nameBuilding, typeOfObject, networkTrading, region, microDistrict, streetName, openIn)

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dataResponse{items})
}

// @Summary UpdateBuildingItem
// @Security ApiKeyAuth
// @Tags Объект
// @Description update chosen item
// @Accept json
// @Produce json
// @Param input body domain.BuildingUpdateInput true "update building"
// @Param id path string true "building id"
// @Success 200 {string} Status "ok"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/:id [put]
func (h *Handler) UpdateBuildingItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	var input domain.BuildingUpdateInput
	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.services.Update(id, input); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})

}

// @Summary DeleteBuildingItem
// @Security ApiKeyAuth
// @Tags Объект
// @Description delete item with id
// @Accept json
// @Produce json
// @Param id path string true "building id"
// @Success 200 {string} Status "ok"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/:id [delete]
func (h *Handler) DeleteBuildingItem(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	err = h.services.Delete(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
