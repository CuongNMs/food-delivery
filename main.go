package main

import (
	"food-delivery/component/appctx"
	"food-delivery/component/uploadprovider"
	"food-delivery/middleware"
	ginrestaurant "food-delivery/module/restaurant/transport/gin"
	"food-delivery/module/upload/uploadtransport/ginupload"
	"food-delivery/module/user/usertransport/ginuser"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"math"
	"os"
)

type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id;"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}
type RestaurantUpdate struct {
	Name *string `json:"name" gorm:"column:name"`
	Addr *string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}

func main() {
	dns := os.Getenv("MYSQL_CONN_STRING")
	s3BucketName := os.Getenv("S3BucketName")
	s3Region := os.Getenv("S3Region")
	s3ApiKey := os.Getenv("S3ApiKey")
	s3Secret := os.Getenv("S3Secret")
	s3Domain := os.Getenv("S3Domain")
	var db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	s3Provider := uploadprovider.NewS3Provider(s3BucketName, s3Region, s3ApiKey, s3Secret, s3Domain)
	appCtx := appctx.NewAppContext(db, s3Provider)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(db, err)
	db = db.Debug()

	r := gin.Default()
	r.Use(middleware.Recover(appCtx))
	r.GET("/ping", func(c *gin.Context) {
		c.JSONP(200, gin.H{
			"message": "pong",
		})
	})
	r.Static("static", "./static")
	v1 := r.Group("/v1")
	v1.POST("/uploadprovider", ginupload.UploadImage(appCtx))
	v1.POST("/restaurants", ginrestaurant.CreateRestaurant(appCtx))
	v1.DELETE("/:id", ginrestaurant.DeleteRestaurant(appCtx))
	v1.GET("", ginrestaurant.ListCreateRestaurant(appCtx))

	v1.POST("/register", ginuser.Register(appCtx))
	r.Run()

	//fmt.Println(countBits(2))
	//fmt.Println(generate(3))
	//test := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
	//fmt.Println(minCostClimbingStairs(test))
}

func countBits(n int) []int {
	result := make([]int, n+1)
	if n == 0 {
		result = []int{1}
		return result
	}

	result[0] = 0
	result[1] = 1
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			result[i] = result[i/2]
		} else {
			result[i] = result[i/2] + 1
		}
	}
	return result

}

func maxProfit(prices []int) int {
	math.Max(1, 2)
	minIndex := 0
	maxIndex := 1
	result := -999999999
	current := 0
	for maxIndex < len(prices) {
		current = prices[maxIndex] - prices[minIndex]
		if current < 0 {
			minIndex = maxIndex
		} else {
			if current > result {
				result = current
			}
		}
		maxIndex++
	}
	if result > 0 {
		return result
	} else {
		return 0
	}
}

func minCostClimbingStairs(cost []int) int {
	result := make([]int, len(cost))

	result[0] = cost[0]
	result[1] = cost[1]
	// totalCost := 0
	i := 2
	for i < len(cost) {
		if result[i-1] <= result[i-2] {
			result[i] = cost[i] + result[i-1]
		} else {
			result[i] = cost[i] + result[i-2]
		}
		// totalCost = result[i]
		i++
	}

	if result[len(result)-1] < result[len(result)-2] {
		return result[len(result)-1]
	} else {
		return result[len(result)-2]
	}
}
