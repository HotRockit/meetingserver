package model

import (
	"reflect"
)

//Meeting is a struct to describe meeting
type Meeting struct {
	MeetingID       uint `json:"meeting_id" gorm:"primary_key;not null;auto_increment"`
	MeetingRoomName string `json:"meeting_room_name" gorm:"not null;type:varchar(100)"`
	User            string `json:"user" gorm:"not null;type:varchar(100)"`
	StartTime       int `json:"start_time" gorm:"type:int;not null"`    //这里为了简化操作，所有的时间类型全部都转换成了string类型
	EndTime         int `json:"end_time" gorm:"type:int;not null"`   //这里为了和微信小程序对应，数据库里面其实没有存时间，存的是方格的索引值
	Reason          string `json:"reason" gorm:"not null;type:varchar(100)"`
	BorrowTime      string `json:"borrow_time" gorm:"type:varchar(20);not null"`  //格式为 2020-06-03go
}

//User is a struct to describe user
type User struct {
	//UserID uint `json:"user_id" gorm:"primary_key;not null;auto_increment"`
	Username string `json:"username" gorm:"primary_key;type:varchar(100)"`
	Password string `json:"password" gorm:"not null;type:varchar(100)"`
	Phone    string `json:"phone" gorm:"not null;type:varchar(15)"`
	State	int `json:"state" gorm:"not null;type:int;default:0"`    //0代表状态正常,1代表账号被冻结了
}

//MeetingRoom is a struct to describe MeetingRoom
type MeetingRoom struct {
	RoomID uint `json:"room_id" gorm:"primary_key;not null;auto_increment"`
	RoomName string `json:"room_name" gorm:"not null;type:varchar(100)"`
}

//判断结构体是否为空
func (user User) UserIsEmpty() bool {
	return reflect.DeepEqual(user,User{})
}

func (meeting Meeting) MeetingIsEmpty() bool {
	return reflect.DeepEqual(meeting,Meeting{})
}

func (room MeetingRoom) MeetingRoomIsEmpty() bool {
	return reflect.DeepEqual(room,MeetingRoom{})
}
