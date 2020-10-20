package controller

import (
	"github.com/gin-gonic/gin"
	"meetingserver/service"
	"meetingserver/util"
	"strconv"
)

//MeetingController is a controller to handle
type AppController struct {

}

var appService = new(service.AppService)

func (meetingController *AppController) UserController(engine *gin.Engine) {
	userGroup := engine.Group("/user")
	{
		userGroup.GET("/login",LoginHandler )
		userGroup.GET("/register",RegisterHandler)
		userGroup.GET("/updatePassword",UpdatePasswordHandler)
		userGroup.GET("/deleteUser",DeleteUserHandler)
	}
}

func DeleteUserHandler(c *gin.Context) {
	username := c.Query("username")
	util.JsonResult(c,appService.DeleteUserService(username))
}

func UpdatePasswordHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	util.JsonResult(c,appService.UpdatePasswordService(username,password))
}

func LoginHandler(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")
	util.JsonResult(c,appService.LoginService(username,password))
}

func RegisterHandler(c *gin.Context){
	//这里本来准备使用post请求方式，但是由于微信小程序端发起post请求比较麻烦，所以就采用了get请求方式
	username := c.Query("username")
	password := c.Query("password")
	phone := c.Query("phone")
	result := appService.RegisterService(username,password,phone)
	util.JsonResult(c,result)
}

func (meetingController *AppController) MeetingController(engine *gin.Engine) {
	meetingGroup := engine.Group("/meeting")
	{
		meetingGroup.GET("/addMeeting",meetingController.AddMeetingHandler)
		meetingGroup.GET("/queryMeetingRoom",meetingController.QueryMeetingRoomHandler)
		meetingGroup.GET("/queryMeeting",meetingController.QueryMeetingHandler)
		meetingGroup.GET("deleteMeeting",meetingController.DeleteMeetingHandler)
	}
}

func (meetingController *AppController) AddMeetingHandler(c *gin.Context){

	meetingRoomName := c.Query("meetingRoomName")
	user := c.Query("user")
	start,_ := strconv.Atoi(c.Query("start"))
	end, _ := strconv.Atoi(c.Query("end"))
	reason := c.Query("reason")
	borrow := c.Query("borrow")
	result := appService.AddMeetingService(meetingRoomName,user,start,end,reason,borrow)  //0代表插入失败，1代表插入成功
	util.JsonResult(c,result)
}

func (meetingController *AppController) QueryMeetingRoomHandler(c *gin.Context) {
	rooms := appService.QueryMeetingRoomService()
	util.JsonResult(c,rooms)
}

func (meetingController *AppController) QueryMeetingHandler(c *gin.Context) {
	//这里分两种情况，一种是查询某个用户的会议记录，一个是查询某一天某一个房间的所有会议
	user := c.Query("user")
	day := c.Query("day")   //格式为 2020-02-03
	room := c.Query("room")
	meetings := appService.QueryMeetingService(user,day,room)
	util.JsonResult(c,meetings)
}

func (meetingController *AppController) DeleteMeetingHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("meeting_id"))
	util.JsonResult(c,appService.DeleteMeetingService(id))
}
