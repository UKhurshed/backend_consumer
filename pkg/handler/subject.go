package handler

import (
	"backend_consumer/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary Get All Subjects
// @Security ApiKeyAuth
// @Tags Субъект
// @Description get all subjects
// @ID get-all-subjects
// @Accept  json
// @Produce  json
// @Success 200 {array} dataResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/subject [get]
func (h *Handler) GetAllSubject(c *gin.Context) {
	subjects, err := h.services.GetAllSubjects()

	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, dataResponse{subjects})
}


// @Summary Create subject
// @Security ApiKeyAuth
// @Tags Субъект
// @Description create subject
// @ID create-subject
// @Accept  json
// @Produce  json
// @Param input body domain.Subject true "subject"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/subject [post]
func (h *Handler) CreateSubject(c *gin.Context) {
	var input domain.Subject

	if err := c.BindJSON(&input); err != nil {
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.CreateSubject(input)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Update subject
// @Security ApiKeyAuth
// @Tags Субъект
// @Description Update subject
// @ID Update-subject
// @Accept  json
// @Produce  json
// @Param input body domain.SubjectInput true "update subject" id path string true "subject id"
// @Success 200 {string} status "ok"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/subject/:id [put]
func (h *Handler) UpdateSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	var input domain.SubjectInput
	if err := c.BindJSON(&input); err != nil{
		newResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	if err := h.services.UpdateSubject(id, input); err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		"ok",
	})
}

// @Summary Delete Subject
// @Security ApiKeyAuth
// @Description delete item with id
// @Tags Субъект
// @Accept json
// @Produce json
// @Param id path string true "subject id"
// @Success 200 {string} Status "ok"
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /api/subject/:id [delete]
func (h *Handler) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		newResponse(c, http.StatusBadRequest, "invalid id params")
		return
	}

	err = h.services.DeleteSubject(id)
	if err != nil {
		newResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{"ok"})
}
