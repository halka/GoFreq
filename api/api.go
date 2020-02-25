// Package api ルーターで受け取ってDBに読み書きすることを書きます
package api

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	"../model"
	"github.com/gin-gonic/gin"

	"golang.org/x/crypto/bcrypt"
)

// Handler データベースのインスタンスを入れておく構造体
type Handler struct {
	Db *gorm.DB
}

// Registration ユーザ登録
func (handler *Handler) Registration(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	// パスワードを暗号化して返す
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err.Error())
	}
	user.Password = string(hash)
	c.JSON(http.StatusOK, gin.H{"user": user.Username})

	handler.Db.Create(&user)
}

// Login ログインする
func (handler *Handler) Login(c *gin.Context) {
	var loginUser model.User
	var fetchUser model.User

	c.BindJSON(&loginUser)
	handler.Db.Where("username = ?", loginUser.Username).First(&fetchUser)
	log.Print(fetchUser)
	loginChallenge := bcrypt.CompareHashAndPassword([]byte(fetchUser.Password), []byte(loginUser.Password))
	if loginChallenge == nil {
		c.JSON(http.StatusOK, gin.H{"user": loginUser.Username, "message": "Welcome"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"user": loginUser.Username, "message": "Unauthorized"})
	}
}

// Logout ログアウトする
func (handler *Handler) Logout(c *gin.Context) {
}

// AddFreq 周波数情報を追加する
func (handler *Handler) AddFreq(c *gin.Context) {
	var freq model.Freq
	c.Bind(&freq)
	handler.Db.Create(&freq)
	c.BindJSON(&freq)
}

// FreqList 周波数情報をすべて取得する
func (handler *Handler) FreqList(c *gin.Context) {
	var freq []model.Freq
	handler.Db.Find(&freq)
	c.BindJSON(&freq)
}

// FreqDetail 個別の周波数情報を取得する
func (handler *Handler) FreqDetail(c *gin.Context) {
	var freq model.Freq
	var id = c.Param("id")
	handler.Db.Where("id = ?", id).First(&freq)
	c.BindJSON(&freq)
}

// FreqDelete 周波数情報を削除する
func (handler *Handler) FreqDelete(c *gin.Context) {
	var freq model.Freq
	var id = c.Param("id")
	handler.Db.Where("id = ?", id).Delete(&freq)
	c.BindJSON(&freq)
}
