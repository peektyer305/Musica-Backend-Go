package postgre

import (
	valueobject "Musica-Backend/internal/domain/value_object"
	"Musica-Backend/internal/infrastructure/postgre/model"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertInitialData(db *gorm.DB) {
	// User データの作成
	var postCout int64
	db.Table("posts").Count(&postCout)
	if postCout != 0 {
		return
	}

	userId, err := valueobject.NewUserId("0191af5d-abb0-7cbe-b096-1e2907188ac6")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	userIconUrl := "https://example.com/userIcon.jpg"
	passwordHash, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	user := model.User{
		Id:          userId,
		Username:    "Achro",
		Email:       "pektyer305@gmail.com",
		Password:    string(passwordHash),
		UserIconUrl: &userIconUrl,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	// User をデータベースに挿入
	if err := db.Table("users").Create(&user).Error; err != nil {
		fmt.Println(err)
		panic(err)
	}

	// Post データの作成
	postId, err := valueobject.NewPostId("0191af68-a755-7a17-9d33-f6d58aa85197")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	content := "test content"
	imageUrl := "https://example.com/image.jpg"
	musicUrl := "https://example.com/music.mp3"

	post := model.Post{
		Id:        postId,
		UserId:    userId,
		Title:     "test title",
		Content:   &content,
		MusicUrl:  musicUrl,
		ImageUrl:  &imageUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Post をデータベースに挿入
	if err := db.Table("posts").Create(&post).Error; err != nil {
		fmt.Println(err)
		panic(err)
	}
}
