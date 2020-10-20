package service

import (
	"log"
	"meetingserver/dao"
	"meetingserver/model"
	"meetingserver/util"
)

type AppService struct {

}


func (appService *AppService)LoginService(username string,password string) model.User {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.LoginDao(username,password)
}

func (appService *AppService)RegisterService(username string,password string,phone string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.RegisterDao(username,password,phone)  //0代表用户存在，1代表注册失败，2代表注册成功
}

func (appService *AppService) QueryMeetingRoomService() []model.MeetingRoom {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	var rooms []model.MeetingRoom
	if err := appDao.Find(&rooms).Error; err != nil {
		log.Print("查询失败")
	}
	return rooms
}

func (appService *AppService) QueryMeetingService(user string,day string,room string) []model.Meeting {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	if user != ""{
		return appDao.QueryMeetingDao1(user,day,room)
	}else{
		return appDao.QueryMeetingDao2(day,room)
	}
}

func (appService *AppService) AddMeetingService(meetingRoomName string, user string,start int,end int, reason string,borrow string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.AddMeetingDao(meetingRoomName,user,start,end,reason,borrow)
}

func (appService *AppService) UpdatePasswordService(username string, password string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.UpdatePasswordDao(username,password)
}

func (appService *AppService) DeleteUserService(username string) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.DeleteUserDao(username)
}

func (appService *AppService) DeleteMeetingService(id int) uint {
	appDao := dao.AppDao{GormEngine: util.GlobalGormEngine}
	return appDao.DeleteMeetingDao(id)
}