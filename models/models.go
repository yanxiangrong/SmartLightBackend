package models

import (
	"SmartLightBackend/pkg/logging"
	"errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strconv"
	"time"
)

var db *gorm.DB

type Device struct {
	Addr     int       `gorm:"type:int;not null"`
	Current  *int      `gorm:"type:int;default:0;not null"`
	Switch   *bool     `gorm:"type:boolean;default:0;not null"`
	Fault    *bool     `gorm:"type:boolean;default:0;not null"`
	Remark   *string   `gorm:"type:varchar(20);default:null;"`
	UpdateAt time.Time `gorm:"type:datetime;default:0;not null"`
}

type User struct {
	Id     string `gorm:"type:varchar(20);not null"`
	Passwd string `gorm:"type:varchar(20);not null"`
}

type JsonDevice struct {
	Addr     int    `json:"addr"`
	Current  int    `json:"current"`
	Switch   bool   `json:"switch"`
	Fault    bool   `json:"fault"`
	UpdateAt string `json:"update_at"`
	Remark   string `json:"remark"`
}

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("runtime/sqlite.db"), &gorm.Config{})
	if err != nil {
		logging.Fatal("Open Sqlite Error ", err.Error())
	}
	// Migrate the schema
	err = db.AutoMigrate(&Device{})
	err = db.AutoMigrate(&User{})
	if err != nil {
		logging.Fatal("Sqlite AutoMigrate Error ", err.Error())
	}
}

func CloseDB() {
	logging.Debug("Close Sqlite")
	//err := db.Close()
	//if err != nil {
	//	logging.Error("Close Sqlite Error ", err.Error())
	//}
}

func AddDevice(device Device) {
	device.UpdateAt = time.Now()
	db.Create(&device)
}

func UpdateDevice(device Device) {
	device.UpdateAt = time.Now()
	db.Where("addr=?", device.Addr).Updates(&device)
}

func UpdateDeviceWithoutTime(device Device) {
	db.Where("addr=?", device.Addr).Updates(&device)
}

func DeleteDevice(device Device) {
	db.Where("addr=?", device.Addr).Delete(Device{})
}

func IsDeviceExist(device Device) bool {
	deviceQuery := Device{}
	result := db.Where("addr=?", device.Addr).First(&deviceQuery)
	// check error ErrRecordNotFound
	return !errors.Is(result.Error, gorm.ErrRecordNotFound) && deviceQuery.Addr == device.Addr
}

func toJsonDevice(device Device) JsonDevice {
	var remark = ""
	if device.Remark != nil {
		remark = *device.Remark
	}
	return JsonDevice{
		Addr:     device.Addr,
		Current:  *device.Current,
		Switch:   *device.Switch,
		Fault:    *device.Fault,
		Remark:   remark,
		UpdateAt: strconv.FormatInt(device.UpdateAt.Unix(), 10),
	}
}

func GetAllDevices() []JsonDevice {
	var devices []Device
	var jsonDevices []JsonDevice
	db.Find(&devices)
	for i := 0; i < len(devices); i++ {
		jsonDevices = append(jsonDevices, toJsonDevice(devices[i]))
	}
	return jsonDevices
}

func AddUser(user User) {
	db.Create(&user)
}

func Verify(user User) bool {
	var retUser User
	db.Find(&User{Id: user.Id}).First(&retUser)

	if retUser.Passwd == user.Passwd {
		return true
	}
	return false
}

func UpdateUser(user User) {
	db.Where("id=?", user.Id).Updates(&user)
}

func DeleteUser(user User) {
	db.Where("id=?", user.Id).Delete(&user)
}
