package models

/*model层用于定义结构体和对数据库进行操作*/

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Agency struct { //创建一个strcut用于与表的元素做映射
	Ano     string `xorm:"not null pk CHAR(8)"`
	Aname   string `xorm:"not null CHAR(8)"`
	Asex    string `xorm:"not null CHAR(1)"`
	Aphone  string `xorm:"CHAR(12)"`
	Aremark string `xorm:"comment('备注') VARCHAR(50)"`
}

var engine *xorm.Engine   // 定义一个engine对象用于操作数据库
var session *xorm.Session // 定义一个session对象用于事务处理

func init() {
	var errdb error
	engine, errdb = xorm.NewEngine("mysql", "root:root@/medidb?charset=utf8")
	// 第二个参数的含义为"数据库名称:数据库连接密码@(数据库地址:3306)/数据库实例名称?charset=utf8"
	if errdb != nil {
		fmt.Println(errdb.Error())
		return
	}
	session = engine.NewSession()
	defer session.Close()
}

/*xorm通过find方法获取多条数据
可以传入Slice用于查询，如：
allagency := make([]Agency, 0)
也可以传入map用于查询，如：
allagency := make(map[int64]Agency)
当然也可传入各种条件用于搜索，如：
err := engine.Where("cno =?",3).Find(&allagency)*/

func GetAllAgency() (*[]Agency, error) {
	allagency := make([]Agency, 0)
	err := engine.Find(&allagency)
	if err != nil {
		return nil, err
	}
	return &allagency, nil
}

/*xorm通过get方法获取单条数据
返回两个参数，一个has是bool类型，表明数据是否存在；一个err表明是否发生错误。
不管err是否为nil，has都有可能为true或者false。
可通过主键字段值作为查询条件，如：
has, err := engine.Id(1).Get(agency)
可通过where查询条件，如：
engine.Where("cno=?", 1).Get(agency)
也可通过结构体非空数据查询，如：
	agency := &Agency{Ano: id}
	has, err := engine.Get(agency)*/

/*但需要判断某个记录是否存在可以使用Exist方法, 相比Get，Exist性能更好
多表联合查询时使用Join方法，如：
	engine.Table("client").Join("INNER", "agency", "agency.ano=client.ano").
		Join("INNER", "medicine", "medicine.mno=client.mno").Find(&clients)*/

func GetAgencybyId(id string) (*Agency, error) {
	agency := &Agency{Ano: id}
	has, err := engine.Get(agency)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	} else if has == true {
		return agency, nil
	}
	return nil, err
}

/*xorm通过Insert方法进行插入操作
Insert方法将返回两个参数，第一个为插入的记录数，第二个参数为错误。
既可支持单条插入，传入的是结构体指针,如：
affected, err := engine.Insert(&agency)
也可传入结构体指针切片，用于多条插入，如：
users := make([]*agency, 2)
affected, err := engine.Insert(&agency) */

func AddAgency(agency *Agency) (bool, error) {
	errsession := session.Begin() // session的Begin方法用于表示事务开始
	if errsession != nil {
		fmt.Println(errsession.Error())
		return true, errsession
	}
	_, err := session.Insert(agency)
	if err != nil {
		fmt.Println(err.Error())
		session.Rollback() // 当发生错误时便通过Rollback方法进行数据回滚
		return false, err
	}
	session.Commit() // 最后需要提交事务处理
	return true, nil

}

/*xorm通过Update方法进行更新操作
Update方法将返回两个参数，第一个为更新的记录数，第二个参数为错误。
可以传入结构体指针更新，如：
_, err := session.ID(agency.Ano).Update(agency)
也可以传入一个Map[string]interface{}类型更新，但此时需要指定更新的是哪个表如：
_, err := session.Table(agency).Id(id).Update(map[string]interface{}{"Ano":2}*/

func UpdateAgency(agency *Agency) (bool, error) {
	errsession := session.Begin()
	if errsession != nil {
		fmt.Println(errsession.Error())
		return true, errsession
	}
	_, err := session.ID(agency.Ano).Update(agency)
	if err != nil {
		fmt.Println(err.Error())
		session.Rollback()
		return false, err
	}
	session.Commit()
	return true, nil

}

/*xorm通过Delete方法进行删除操作
Delete的返回值第一个参数为删除的记录数，第二个参数为错误。
传入结构体指针用于删除，如：
affected, err := engine.Id(id).Delete(agency)
需要注意的是，在struck中使用deleted标记的话可以进行软删除，对应的字段必须为time.Time类型，如：
type User struct {
    Id int64
    Name string
    DeletedAt time.Time `xorm:"deleted"`
}*/

func DeleteAgency(id string) (bool, error) {
	errsession := session.Begin()
	if errsession != nil {
		fmt.Println(errsession.Error())
		return true, errsession
	}
	agency := new(Agency)
	_, err := session.ID(id).Delete(agency)
	if err != nil {
		fmt.Println(err.Error())
		session.Rollback()
		return false, err
	}
	session.Commit()
	return true, nil
}
