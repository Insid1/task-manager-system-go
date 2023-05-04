package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	todo "go-task-manager-system"
	"net/http"
)

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	var input todo.TodoList

	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	todoListId, err := h.services.TodoList.Create(input, userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"listId": todoListId,
	})
}

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}

	lists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": lists,
	})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	list, err := h.services.TodoList.GetById(userId, listId)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNoContent, "")
		return
	} else if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	var input todo.UpdateTodoListInput
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	list, err := h.services.TodoList.Update(userId, listId, input)

	if err == sql.ErrNoRows {
		newErrorResponse(c, http.StatusBadRequest, err.Error())

	} else if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserIdFromContext(c)
	if err != nil {
		return
	}
	listId, err := getIdFromParam(c)
	if err != nil {
		return
	}

	err = h.services.TodoList.Delete(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]uint64{
		"id": listId,
	})
}
