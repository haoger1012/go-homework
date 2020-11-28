package main

import (
	"net/http"
	"strconv"

	"golang.org/x/net/context"

	"github.com/gin-gonic/gin"
)

func main() {
	context.Background()
	router := gin.Default()
	http.HandleFunc("ss", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	router.GET("/role", Get)

	router.GET("/role/:id", GetOne)

	router.POST("/role", Post)

	router.PUT("/role/:id", Put)

	router.DELETE("/role/:id", Delete)

	router.Run(":8080")
}

// 取得全部資料
func Get(c *gin.Context) {
	c.JSON(http.StatusOK, Data)
}

// 取得單一筆資料
func GetOne(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(Data); i++ {
		if Data[i].ID == uint(id) {
			c.JSON(http.StatusOK, Data[i])
			return
		}
	}
	c.Status(http.StatusNotFound)
}

// 新增資料
func Post(c *gin.Context) {
	var r Role
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	Data = append(Data, r)
	c.JSON(http.StatusOK, r)
}

type RoleVM struct {
	ID      uint   `json:"id"`      // Key
	Name    string `json:"name"`    // 角色名稱
	Summary string `json:"summary"` // 介紹
}

// 更新資料, 更新角色名稱與介紹
func Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	var r Role
	if err := c.ShouldBind(&r); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i := 0; i < len(Data); i++ {
		if Data[i].ID == uint(id) {
			Data[i].Name = r.Name
			Data[i].Summary = r.Summary
			c.JSON(http.StatusOK, Data[i])
			return
		}
	}
}

// 刪除資料
func Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	for i, role := range Data {
		if role.ID == uint(id) {
			Data = append(Data[0:i], Data[i+1:]...)
			break
		}
	}

	c.Status(http.StatusNoContent)
}
