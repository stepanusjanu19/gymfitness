package controllers

import (
	"backend/domain/requests"
	"backend/domain/responses"
	"backend/domain/services"
	"backend/lib/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindItem(c *gin.Context) {
	items, err := services.FindItemAll()
	if err != nil {
        response.ERROR(c.Writer, http.StatusInternalServerError, err)
        return
    }

	var itemResponse []responses.ItemResponse

	for _, item := range items {
		itemResponse = append(itemResponse, responses.ItemResponse{
			ItemID: item.ItemID,
			Name: item.Name,
			Description: item.Description,
		})
	}

	if len(itemResponse) < 1 {
		itemResponse = []responses.ItemResponse{}
	}

	response.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "Get all items successfully",
		"item": itemResponse,
	})
}

func FindItemByID(c *gin.Context) {
	itemID := c.Param("item_id")

    item, err := services.FindByItemID(itemID)
    if err!= nil {
        response.ERROR(c.Writer, http.StatusNotFound, err)
        return
    }

    response.JSON(c.Writer, http.StatusOK, gin.H{
        "message": "Get item by ID successfully",
        "item": responses.ItemResponse{
            ItemID: item.ItemID,
            Name: item.Name,
            Description: item.Description,
        },
    })
}

func CreateItem(c *gin.Context) {
	var itemRequest requests.SaveItem

    if err := c.ShouldBindJSON(&itemRequest); err!= nil {
        response.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error": err.Error(),
		})
        return
    }

    item, err := services.CreateItem(itemRequest)
    if err!= nil {
        response.ERROR(c.Writer, http.StatusInternalServerError, err)
        return
    }

    response.JSON(c.Writer, http.StatusCreated, gin.H{
        "message": "Create item successfully",
        "item": responses.ItemResponse{
            ItemID: item.ItemID,
            Name: item.Name,
            Description: item.Description,
        },
    })
}

func UpdateItem(c *gin.Context) {
	var itemRequest requests.UpdateItem
	itemID := c.Param("item_id")

	if err := c.ShouldBindJSON(&itemRequest); err!= nil {
        response.JSON(c.Writer, http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"error": err.Error(),
		})
        return
    }

	item, err := services.UpdateItem(itemRequest, itemID)
	if err!= nil {
        response.ERROR(c.Writer, http.StatusInternalServerError, err)
        return
    }

	response.JSON(c.Writer, http.StatusOK, gin.H{
		"message": "Update item successfully",
        "item": responses.ItemResponse{
            ItemID: item.ItemID,
            Name: item.Name,
            Description: item.Description,
        },
    })
}

func DeleteItem(c *gin.Context) {
	itemID := c.Param("item_id")

    itemResponse, err := services.DeleteItem(itemID)
    if err!= nil {
        response.ERROR(c.Writer, http.StatusInternalServerError, err)
        return
    }

    response.JSON(c.Writer, http.StatusOK, gin.H{
        "message": "Delete item successfully",
		"item": responses.ItemResponse{
			ItemID: itemResponse.ItemID,
            Name: itemResponse.Name,
            Description: itemResponse.Description,
		},
    })
}