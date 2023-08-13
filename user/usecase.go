package user

import (
	"SynergizeTest/config"
	"SynergizeTest/models"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

func Register(ctx *gin.Context, form models.User) (err error) {
	form.RegisterAt = time.Now()
	//add user
	if errx := config.DB.Create(&form); errx.Error != nil {
		err = errx.Error
		return
	}
	// automatically create wallet when user register
	err = CreateWallet(ctx, models.Wallet{
		Saldo:     0,
		UserID:    form.UserID,
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return
	}
	return
}

func CreateWallet(ctx *gin.Context, Wallet models.Wallet) (err error) {
	//add wallet
	errx := config.DB.Create(&Wallet)
	if errx.Error != nil {
		err = errx.Error
		return
	}
	return
}

func Login(ctx *gin.Context, form models.LoginForm) (err error) {
	user := models.User{}
	//check if username and password correct
	if errx := config.DB.Where("user_name = ? AND password = ?", form.Username, form.Password).Find(&user); errx.Error != nil {
		err = errx.Error
		return
	}
	if user.UserID == 0 {
		err = fmt.Errorf("user doesn't exist")
		return
	}

	userID := strconv.Itoa(user.UserID)
	//set user login session in redis
	config.Redis.Set(userID, "true", 86400*time.Second)

	return
}

func Logout(ctx *gin.Context, form models.LogoutForm) {
	userID := strconv.Itoa(form.UserID)
	//remove user login
	config.Redis.Set(userID, "false", 10*time.Second)
}

func RegisterBank(ctx *gin.Context, form models.Bank) (err error) {
	//check if user login
	userID := strconv.Itoa(form.UserID)
	res, err := config.Redis.Get(userID).Result()
	if err != nil && err != redis.Nil {
		return
	}
	if res != "true" || err == redis.Nil {
		err = fmt.Errorf("user not login")
		return
	}

	form.UpdatedAt = time.Now()
	//add bank
	rows := config.DB.Create(&form)
	if rows.Error != nil {
		err = rows.Error
		return
	}
	return
}

func TopUp(ctx *gin.Context, form models.TopUpForm) (err error) {
	//check if user login
	userID := strconv.Itoa(form.UserID)
	res, err := config.Redis.Get(userID).Result()
	if err != nil && err != redis.Nil {
		return
	}
	if res != "true" || err == redis.Nil {
		err = fmt.Errorf("user not login")
		return
	}

	var wallet models.Wallet
	//get user wallet
	if errx := config.DB.Where("user_id=?", form.UserID).Find(&wallet); errx.Error != nil {
		err = errx.Error
		fmt.Println(errx)
		return
	}

	//add amount to current saldo in db and save
	wallet.Saldo = wallet.Saldo + form.Amount
	if errx := config.DB.Save(&wallet); errx.Error != nil {
		err = errx.Error
		fmt.Println(errx)
		return
	}
	return
}

func GetUserDetail(ctx *gin.Context, form models.LogoutForm) (res models.UserDetail, err error) {
	//check user login
	userID := strconv.Itoa(form.UserID)
	redisResult, err := config.Redis.Get(userID).Result()
	if err != nil && err != redis.Nil {
		return
	}
	if redisResult != "true" || err == redis.Nil {
		err = fmt.Errorf("user not login")
		return
	}
	//get all user data
	if errx := config.DB.Where("user_id = ?", form.UserID).Find(&res.User); errx.Error != nil {
		err = errx.Error
		fmt.Println(errx)
		return
	}
	if errx := config.DB.Where("user_id=?", form.UserID).Find(&res.Wallet); errx.Error != nil {
		err = errx.Error
		fmt.Println(errx)
		return
	}
	if errx := config.DB.Where("user_id=?", form.UserID).Find(&res.Bank); errx.Error != nil {
		err = errx.Error
		fmt.Println(errx)
		return
	}
	return
}

func GetUserList(ctx *gin.Context, form models.Pagination) (res []models.UserDetail, err error) {
	var user []models.User
	var rows *gorm.DB
	filter := ""
	//check filter for user table
	IsExist := IsExists(form.Keys, []string{"user_name", "email", "phone_number"})
	if IsExist {
		filter = fmt.Sprintf("%v = ?", form.Keys)
		rows = config.DB.Debug().Scopes(config.Paginate(form)).Where(filter, form.Value).Find(&user)
		if rows.Error != nil {
			err = rows.Error
			return
		}
	} else {
		rows = config.DB.Scopes(config.Paginate(form)).Find(&user)
		if rows.Error != nil {
			err = rows.Error
			return
		}
	}

	for i := 0; i < len(user); i++ {
		res = append(res, models.UserDetail{
			User: user[i],
		})
		//check filter for wallet column
		if form.Keys == "saldo" {
			filter := "user_id = ? and saldo = ?"
			rows = config.DB.Debug().Where(filter, user[i].UserID, form.Value).Find(&res[i].Wallet)
			if rows.Error != nil {
				err = rows.Error
				return
			}
		} else {
			if errx := config.DB.Where("user_id = ?", user[i].UserID).Find(&res[i].Wallet); errx.Error != nil {
				err = errx.Error
				fmt.Println(errx)
				return
			}
		}
		//check filter for bank column
		IsExist2 := IsExists(form.Keys, []string{"bank_code", "account_number", "account_name"})
		if IsExist2 {
			filter = fmt.Sprintf("user_id = ? and %v = ?", form.Keys)
			rows = config.DB.Debug().Where(filter, user[i].UserID, form.Value).Find(&res[i].Bank)
			if rows.Error != nil {
				err = rows.Error
				return
			}
		} else {
			if errx := config.DB.Where("user_id=?", user[i].UserID).Find(&res[i].Bank); errx.Error != nil {
				err = errx.Error
				fmt.Println(errx)
				return
			}
		}
	}
	// add all valid data
	var temp []models.UserDetail
	for i := 0; i < len(res); i++ {
		IsExist = IsExists(form.Keys, []string{"bank_code", "account_number", "account_name"})
		if IsExist && res[i].Bank.ID != 0 {
			temp = append(temp, res[i])
		} else if form.Keys == "saldo" && res[i].Wallet.WalletID != 0 {
			temp = append(temp, res[i])
		} else if IsExists(form.Keys, []string{"user_name", "email", "phone_number", ""}) {
			temp = append(temp, res[i])
		}

	}
	return temp, nil
}

func IsExists(value string, array []string) (exists bool) {
	exists = false

	for _, i := range array {
		fmt.Println(i)
		if value == i {
			exists = true
		}
	}
	return
}
