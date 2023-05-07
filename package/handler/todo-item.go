package handler

import (
	"github.com/gin-gonic/gin"
	todo "go-task-manager-system"
	"net/http"
)

func (h *Handler) createItem(c *gin.Context) {
	var input todo.TodoItem

	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}

	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	itemId, err := h.services.TodoItem.Create(&input, userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"itemId": itemId,
	})
}

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}

	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	items, err := h.services.TodoItem.GetAll(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]*[]todo.TodoItem{
		"data": items,
	})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}

	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	itemId, err := getItemIdFromParam(c)
	if err != nil {
		return
	}

	item, err := h.services.TodoItem.GetById(userId, listId, itemId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	var input todo.UpdateTodoItemInput

	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}

	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	itemId, err := getItemIdFromParam(c)
	if err != nil {
		return
	}

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoItem.Update(userId, listId, itemId, &input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"itemId": itemId,
	})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}

	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	itemId, err := getItemIdFromParam(c)
	if err != nil {
		return
	}

	err = h.services.TodoItem.Delete(userId, listId, itemId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"itemId": itemId,
	})
}
