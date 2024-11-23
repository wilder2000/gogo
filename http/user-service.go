package http

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"wilder.cn/gogo/comm"
	"wilder.cn/gogo/config"
	"wilder.cn/gogo/database"
	"wilder.cn/gogo/log"
	"wilder.cn/gogo/sys/db"
)

type (
	Service interface {
		RegisterUser(input RegisterUserInput, rname string) (db.SUser, *MVCError)
		//AutoRegisterUser(input AutoUserInput, rname string) (db.SUser, *MVCError)
		Login(input LoginInput) (db.SUser, *MVCError)
		IsEmailAvailable(input CheckEmailInput) (bool, error)
		SaveAvatar(ID string, fileLocation string) (db.SUser, error)
		GetUserByID(ID string) (db.SUser, error)
	}
)
type service struct {
	repository Repository
}

var (
	UserProxy *service
)

func init() {
	if config.AConfig.UserService {
		database.LoadDatabaseConfig()
		UserProxy = NewService()
	}

}
func NewService() *service {
	userRepository := NewRepository(database.DBHander)
	return &service{userRepository}
}

//	func (s *service) AutoRegisterUser(input AutoUserInput, rname string) (db.SUser, *MVCError) {
//		user := db.SUser{}
//		user.Name = rname
//		user.Sex = 2               //未知
//		user.State = UserStateAuto //自动注册
//		user.Icon = config.AConfig.DefaultAvatar
//		user.Createtime = comm.LocalTime()
//		user.Modtime = user.Createtime
//		uid, err := uuid.NewUUID()
//		user.Email = uid.String() + "@youhua.space" //虚邮箱
//		if err != nil {
//			log.Logger.Warn(err.Error())
//		}
//		user.ID = uid.URN()
//		// user.Role = "user"
//		newUser, err := s.repository.Save(user)
//		if err != nil {
//			return user, NewMVCError(RFailed, "server save to db failed.")
//		}
//
//		return newUser, nil
//	}
func (s *service) RegisterUser(input RegisterUserInput, rname string) (db.SUser, *MVCError) {

	user := db.SUser{}
	user.Email = input.Email
	user.Name = rname
	user.Sex = 2 //未知
	user.Icon = config.AConfig.DefaultAvatar
	user.State = db.UserStateNormal

	oldU, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		log.Logger.InfoF("User regist failed:%s", err.Error())
		return user, NewMVCError(RFailed, "server internal error.")
	}
	if len(oldU.ID) > 1 {
		return user, NewMVCError(DataExistFound, "email is already exist.")
	}

	hpwd, err := EPassword(input.Password)
	if err != nil {
		log.Logger.InfoF("User regist failed:%s", err.Error())
		return user, NewMVCError(RFailed, "password encode failed.")
	}

	user.Password = string(hpwd)
	user.Createtime = comm.LocalTime()
	user.Modtime = user.Createtime
	uid, err := uuid.NewUUID()
	if err != nil {
		log.Logger.Warn(err.Error())
	}
	user.ID = uid.URN()
	// user.Role = "user"
	newUser, err := s.repository.Save(user)
	if err != nil {
		return user, NewMVCError(RFailed, "server save to db failed.")
	}

	return newUser, nil
}

// Login 后台管理登录使用，不用于app
func (s *service) Login(input LoginInput) (db.SUser, *MVCError) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, NewMVCError(RFailed, err.Error())
	}
	if len(user.ID) == 0 {
		return user, NewMVCError(UserNotFound, "No user found by that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return user, NewMVCError(PwdWrong, err.Error())
	}
	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if len(user.ID) == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(ID string, fileLocation string) (db.SUser, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	user.Icon = fileLocation
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}
	return updatedUser, nil
}

func (s *service) GetUserByID(ID string) (db.SUser, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}
	if len(user.ID) == 0 {
		return user, errors.New("No user found by that ID")
	}
	return user, nil
}
func (s *service) ChangePwd(pwd string, email string) error {
	hpwd, err := EPassword(pwd)
	if err != nil {
		return err
	}
	ok := s.repository.UpdatePwd(email, string(hpwd))
	if ok {
		return nil
	} else {
		return errors.New("Not update password where email=" + email)
	}
}
