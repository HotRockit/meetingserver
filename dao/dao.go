package dao

import (
	"log"
	"meetingserver/model"
	"meetingserver/util"
)

type AppDao struct {
	*util.GormEngine
}

func (appDao *AppDao) LoginDao(username string,password string) model.User {
	var user model.User
	appDao.Where("username = ?",username).Where("password = ?",password).First(&user)

	//可以使用下面这种方式判断结构体是否为空，但这个时候要求结构体的所有字段都是可比较的
	//if user == (model.User{}){
	//	return false
	//}

	//下面这种反射的方式就没有上面方法的局限性，但是反射的性能较低
	if user.UserIsEmpty(){
		appDao.Where("phone = ?",username).Where("password = ?",password).First(&user)
		if user.UserIsEmpty(){
			return user
		}
		return user
	}
	return  user
}

func (appDao *AppDao) RegisterDao(username string, password string, phone string) uint {
	//首先判断用户名是否存在
	var user model.User
	appDao.Where("username = ?",username).First(&user)
	if !user.UserIsEmpty(){
		log.Print("用户存在")
		return 0  //用户存在
	}
	if err := appDao.Create(&model.User{Username: username,Password: password,Phone: phone}).Error;err !=nil{
		log.Print("插入失败")
		return 1
	}
	return  2
}

func (appDao *AppDao) AddMeetingDao(name string, user string,start int,end int, reason string,borrow string) uint {
	if err := appDao.Create(&model.Meeting{MeetingRoomName: name,User: user,StartTime: start,
		EndTime: end,Reason: reason,BorrowTime: borrow}).Error;err !=nil{
		log.Print("插入失败")
		return 0
	}
	return  1
}

func (appDao *AppDao) QueryMeetingDao1(user string,day string,room string) []model.Meeting {
	var meetings []model.Meeting
	if err := appDao.Where("user = ?",user).Where("borrow_time = ?",day).Where("meeting_room_name = ?",room).Find(&meetings).Error; err != nil{
		log.Print("获取失败")
		return nil
	}else {
		return meetings
	}
}

func (appDao *AppDao) QueryMeetingDao2(day string, room string) []model.Meeting{
	var meetings []model.Meeting
	if err := appDao.Where("borrow_time = ?",day).Where("meeting_room_name = ?",room).Find(&meetings).Error; err != nil{
		log.Print("获取失败")
		return nil
	}else {
		return meetings
	}
}

func (appDao *AppDao) UpdatePasswordDao(username string, password string) uint {
	var user model.User
	appDao.Where("username = ?",username).First(&user)
	if err := appDao.Model(&user).Update("password",password).Error;err != nil{
		log.Print("更新失败")
		return 0
	}
	return 1
}

func (appDao *AppDao) DeleteUserDao(username string) uint {
	//删除用户的同时把他对应的会议记录一同删除
	if err := appDao.Where("username = ?",username).Delete(&model.User{}).Error;err != nil{
		log.Print("删除失败")
		return 0
	}
	appDao.Where("user = ?",username).Delete(&model.Meeting{})
	return 1
}

func (appDao *AppDao) DeleteMeetingDao(id int) uint {
	if err := appDao.Where("meeting_id = ?",id).Delete(&model.Meeting{}).Error;err != nil{
		log.Print("删除失败")
		return 0
	}
	return 1
}
