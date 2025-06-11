package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"trabalhoTcc.com/mod/repository"
)

type OrderController struct {
	repository repository.OrderRepository 
}

func NewOrderController(repository repository.OrderRepository) OrderController {
	return OrderController {
		repository: repository,
	}
}

func(or *OrderController) AuthMeHandler(c *gin.Context) {
	userID := c.GetInt("user_id")
	email := c.GetString("email")
	userRole := c.GetString("user_role")

	c.JSON(http.StatusOK, gin.H{
		"user_id":   userID,
		"email":     email,
		"user_role": userRole,
	})
}

func (or *OrderController) SetUserOrder(ctx *gin.Context) {
    var orderRequest struct {
        Products []int `json:"products"`  // Assuming the frontend sends an array of product IDs
    }

    userID := ctx.GetInt("user_id")
    
    if err := ctx.ShouldBindJSON(&orderRequest); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Invalid request body",
            "error":   err.Error(),
        })
        return
    }

    if len(orderRequest.Products) == 0 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "No products provided in the request",
        })
        return
    }

    err := or.repository.SetUserOrder(userID, orderRequest.Products)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "message": "Failed to process the order",
            "error":   err.Error(),
        })
        return
    }

    // If the order is created successfully, return a success response
    ctx.JSON(http.StatusOK, gin.H{
        "message": "Order created successfully",
    })
}

func (or *OrderController) ReturnOrder(ctx *gin.Context) {
	var user_id int = ctx.GetInt("user_id")

	order, err := or.repository.ReturnOrder(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"message": "cannot return user's orders",
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK,order)
}

func (or *OrderController) ReturnAllOrders(ctx *gin.Context){
    orders, err  := or.repository.ReturnAllOrders();
    if err != nil {
        ctx.JSON(http.StatusBadGateway, gin.H{
            "message": "failed to fetch all orders",
            "error": err.Error(),
        })
    }
    ctx.JSON(http.StatusOK, orders);
}

// func (or *OrderController) AlterOrder(ctx *gin.Context) {

//     var requestData struct {
//         Status string `json:"status"`
//         OrderID int `json:"order_id"`
//     }
    
//     if err := ctx.ShouldBindJSON(&requestData); err != nil {
//         ctx.JSON(http.StatusBadRequest, gin.H{
//             "message": "Invalid request body",
//             "error": err.Error(),
//         })
//         return
//     }


//     err := or.respository.AlterOrder(requestData.status,requestData.OrderID);
//     if err != nil {
//         ctx.JSON(http.StatusInternalServerError, gin.H{
//             "message": "failed to alter status",
//             "error": err.Error(),
//         })
//     }
//     ctx.JSON(http.StatusOK, "Order altered with success")
// }