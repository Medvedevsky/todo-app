package handler

import (
	"net/http"

	"github.com/Medvedevsky/todo-app"
	"github.com/gin-gonic/gin"
)

func (h *Handler) createList(c *gin.Context) {
	var input todo.TodoList

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	list_id, err := h.services.TodoList.Create(0, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": list_id,
	})

}

func (h *Handler) getAllLists(c *gin.Context) {

}

func (h *Handler) getListById(c *gin.Context) {

}

func (h *Handler) updateList(c *gin.Context) {

}

func (h *Handler) deleteList(c *gin.Context) {

}
