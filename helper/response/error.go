package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GeneralMessage struct {
	Message string `json:"message"`
}

type GeneralMessageWithData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	TotalPage   int `json:"total_page"`
	TotalItems  int `json:"total_items"`
	NextPage    int `json:"next_page"`
	PrevPage    int `json:"prev_page"`
}

type PaginationRes struct {
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

func GetCurrentUser(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func SendStatusOkResponse(c *gin.Context, message string) {
	c.JSON(http.StatusOK, GeneralMessage{
		Message: message,
	})
}

func SendStatusCreatedResponse(c *gin.Context, message string) {
	c.JSON(http.StatusCreated, GeneralMessage{
		Message: message,
	})
}

func SendStatusCreatedWithDataResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, GeneralMessageWithData{
		Message: message,
		Data:    data,
	})
}

func SendStatusOkWithDataResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, GeneralMessageWithData{
		Message: message,
		Data:    data,
	})
}

func SendStatusBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, GeneralMessage{
		Message: message,
	})
}

func SendStatusNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, GeneralMessage{
		Message: message,
	})
}

func SendStatusInternalServerError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, GeneralMessage{
		Message: message,
	})
}

func SendStatusUnauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, GeneralMessage{
		Message: message,
	})
}

func SendStatusForbidden(c *gin.Context, message string) {
	c.JSON(http.StatusForbidden, GeneralMessage{
		Message: message,
	})
}

func SendPaginationResponse(c *gin.Context, data interface{}, currentPage, totalPages, totalItems, nextPage, prevPage int, message string) {
	pagination := PaginationMeta{
		CurrentPage: currentPage,
		TotalPage:   totalPages,
		TotalItems:  totalItems,
		NextPage:    nextPage,
		PrevPage:    prevPage,
	}
	c.JSON(http.StatusOK, PaginationRes{
		Message: message,
		Data:    data,
		Meta:    pagination,
	})
}
