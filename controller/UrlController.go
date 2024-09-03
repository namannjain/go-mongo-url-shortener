package controller

import (
	"fmt"
	"net/http"
	"time"
	"urlShortenerMongo/constant"
	"urlShortenerMongo/database"
	"urlShortenerMongo/helper"
	"urlShortenerMongo/types"

	"github.com/gin-gonic/gin"
)

func ShortTheUrl(c *gin.Context) {
	var shortUrlBody types.ShortUrlBody

	err := c.BindJSON(&shortUrlBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": constant.BindError,
		})
		return
	}

	code := helper.GenerateRandomString(6)

	dbRecord, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	if dbRecord.UrlCode != "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "this code is already in use",
		})
		return
	}

	var url types.UrlDao
	url.CreatedAt = time.Now().Unix()
	url.ExpiredAt = time.Now().Unix()
	url.UrlCode = code
	url.LongUrl = shortUrlBody.LongUrl
	url.ShortUrl = constant.BaseUrl + code

	resp, err := database.Mgr.Insert(url, constant.UrlCollection)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":     false,
		"data":      resp,
		"short_url": url.ShortUrl,
	})
}

func RedirectUrl(c *gin.Context) {
	code := c.Param("code")

	dbRecord, _ := database.Mgr.GetUrlFromCode(code, constant.UrlCollection)
	if dbRecord.UrlCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "there is no url found",
		})
		return
	}
	fmt.Println(dbRecord.LongUrl)

	c.Redirect(http.StatusPermanentRedirect, dbRecord.LongUrl)
}
