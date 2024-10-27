package entity

import (
	"time"
)

// Employee struct for GORM
type Employee struct {
	ID         uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username   string    `gorm:"type:varchar(255);not null" json:"username"`
	Name       string    `gorm:"type:varchar(255);not null" json:"name"`
	Password   string    `gorm:"type:varchar(255);not null" json:"password"`
	Phone      string    `gorm:"type:varchar(20);not null" json:"phone"`
	Sex        string    `gorm:"type:varchar(10);not null" json:"sex"`
	IDNumber   string    `gorm:"type:varchar(20);not null" json:"idNumber"`
	Status     int       `gorm:"type:int;not null" json:"status"`
	CreateTime time.Time `gorm:"autoCreateTime" json:"createTime"` // 自动插入创建时间
	UpdateTime time.Time `gorm:"autoUpdateTime" json:"updateTime"` // 自动更新修改时间
	CreateUser uint64    `gorm:"type:bigint;not null" json:"createUser"`
	UpdateUser uint64    `gorm:"type:bigint;not null" json:"updateUser"`
}

func (Employee) TableName() string {
 
	return "employee"  // 指定表名为 employee
}


/*

{
    "code": 0,
    "data": {
        "records": [
            {
                "id": 2,
                "username": "zhangsan",
                "name": "张三",
                "password": "e10adc3949ba59abbe56e057f20f883e",
                "phone": "13556788991",
                "sex": "1",
                "idNumber": "123456789012345678",
                "status": 1,
                "createTime": "2024-10-02T21:02:39+08:00",
                "updateTime": "2024-10-02T21:02:39+08:00",
                "createUser": 10,
                "updateUser": 10
            }
        ],
        "total": 1
    },
    "msg": null
}

*/

/*

{
    "code": 1,
    "msg": null,
    "data": {
        "total": 1,
        "records": [
            {
                "id": 2,
                "username": "zhangsan",
                "name": "张三",
                "password": "e10adc3949ba59abbe56e057f20f883e",
                "phone": "13556788991",
                "sex": "1",
                "idNumber": "123456789012345678",
                "status": 1,
                "createTime": "2024-10-02 21:02",
                "updateTime": "2024-10-02 21:02",
                "createUser": 10,
                "updateUser": 10
            }
        ]
    }
}
*/