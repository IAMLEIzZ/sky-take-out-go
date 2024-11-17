package admin

import (
	"context"
	"log"
	"sky-take-out-go/db"
	"sky-take-out-go/internal/api/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

var Key string = "SHOP_STATUS"
var ctx = context.Background()
var rdb, _ = db.InitRedis();

// Set Shop Status
// PATH: /admin/shop/:status
func SetShopStatus(c *gin.Context) {
	log.Println("INFO: " + "Set Shop Status")
	status := c.Param("status")

	// set status to redis
	err := rdb.Set(ctx, Key, status, 0).Err()
	if err != nil {
		response.Response_Error(c)
		return 
	}

	response.Response_Success(c, nil)
}


// Get Shop Status
// PATH: /admin/shop/status
func GetShopStatus(c *gin.Context) {
	log.Println("INFO: " + "Get Shop Status")

	status_str, err := rdb.Get(ctx, Key).Result()
	if err != nil {
		response.Response_Error(c)
		return 
	}
	status, _ := strconv.Atoi(status_str)
	response.Response_Success(c, status)
}