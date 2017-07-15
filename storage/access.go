package storage

import (
	"database/sql"
	"log"
	"qqpaygc/models"

	"goutils"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db = &sql.DB{}
var err error

func init() {
	//db, err = sql.Open("mysql", "root:y0701003@tcp(localhost:3306)/slt") //公司
	db, err = sql.Open("mysql", "allenslt:y0701003@tcp(allen.com:3306)/qqpay") //>>公司的 VPN 要關掉才能連
	log.Println("Hello!!!")
}

func DB_BOLoginVerify(membercode string, password string) bool {
	//Begin函数内部会去获取连接
	dbQueryStr := `
	SELECT COUNT(*) as count
	FROM bomember WHERE membercode=? and password=?
	`
	log.Println("DB_BOLoginVerify membercode: " + membercode + " password: " + password)
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(membercode, password)
	goutils.CheckErr(err)
	defer rows.Close()
	var count int8
	rows.Next()
	err = rows.Scan(&count)
	log.Println("DB_BOLoginVerify count: " + fmt.Sprint(count))
	return count == 1
}

func DB_AddQQPay(qq, idGroup, startdate, extenddays, amount string) {
	//增加 payrecord
	dbQueryStr := `
	INSERT INTO payrecord (accountQQ, idgroup, paydate, amount, extenddays) VALUES (?, ?, NOW(), ? , ?);
	`
	log.Println("DB_AddQQPay qq:" + qq + " idGroup:" + idGroup + " startdate:" + startdate + " extenddays:" + extenddays + " amount:" + amount)
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(qq, idGroup, amount, extenddays)
	goutils.CheckErr(err)
	defer rows.Close()

	//查該QQ是否已存在此群
	dbQueryStr = `
	SELECT COUNT(*) as count
	FROM accountgroup WHERE accountQQ=? and idGroup=?
	`
	log.Println("DB_AddQQPay qq: " + qq + " idGroup: " + idGroup)
	stm, err = db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err = stm.Query(qq, idGroup)
	goutils.CheckErr(err)
	defer rows.Close()
	var count int8
	rows.Next()
	err = rows.Scan(&count)
	log.Println("DB_AddQQPay count: " + fmt.Sprint(count))
	if count == 1 {
		//若該QQ已存在此群 則 更新  accountgroup
		dbQueryStr = `
		UPDATE accountgroup SET 
		startdate=?,
		expiredate=DATE_ADD(startdate,INTERVAL ? DAY)
		WHERE idgroup=? AND accountQQ=?
		`
		stm, err = db.Prepare(dbQueryStr)
		defer stm.Close()
		goutils.CheckErr(err)
		rows, err = stm.Query(startdate, extenddays, idGroup, qq)
		goutils.CheckErr(err)
		defer rows.Close()

	} else {
		//若該QQ不存在此群 則 增加  accountgroup
		dbQueryStr = `
		INSERT INTO accountgroup (accountQQ, idGroup, startdate, expiredate) VALUES (?, ?, ?, DATE_ADD(startdate,INTERVAL ? DAY));
		`
		stm, err = db.Prepare(dbQueryStr)
		defer stm.Close()
		goutils.CheckErr(err)
		rows, err = stm.Query(qq, idGroup, startdate, extenddays)
		goutils.CheckErr(err)
		defer rows.Close()
	}

	//若該QQ不存在此群 則 增加  accountgroup

}

func DB_AddQQGroup(groupQQ string, groupname string) {
	//Begin函数内部会去获取连接
	dbQueryStr := `
	INSERT INTO qqgroup (groupQQ, groupname) VALUES (?, ?)
	`
	log.Println("DB_AddQQGroup groupQQ: " + groupQQ + " groupname: " + groupname)
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(groupQQ, groupname)
	goutils.CheckErr(err)
	defer rows.Close()
}

func DB_GetQQExpireData(accountQQ string) []*models.QQExpireInfo {
	if accountQQ == "" {
		return nil
	}
	//Begin函数内部会去获取连接
	dbQueryStr := `
	SELECT idaccountgroup, accountQQ, (select groupname from qqgroup where idgroup=ag.idgroup) groupname, startdate, expiredate
	FROM accountgroup ag
	WHERE accountQQ = ?
	`
	log.Println("DB_GetQQExpireData accountQQ:" + accountQQ)
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(accountQQ)
	goutils.CheckErr(err)
	defer rows.Close()

	var qqExpireInfoMap []*models.QQExpireInfo
	for rows.Next() { //有下一筆就會一直true下去
		qqExpireInfo := new(models.QQExpireInfo)
		err = rows.Scan(&qqExpireInfo.IdAccountGroup,
			&qqExpireInfo.QQ,
			&qqExpireInfo.GroupName,
			&qqExpireInfo.StartDate,
			&qqExpireInfo.ExpireDate)
		goutils.CheckErr(err)
		qqExpireInfoMap = append(qqExpireInfoMap, qqExpireInfo)
	}
	return qqExpireInfoMap

}

func DB_GetAllQQGroup() []*models.QQGroupInfo {
	//Begin函数内部会去获取连接
	dbQueryStr := `
	SELECT *
	FROM qqgroup
	`
	log.Println("DB_GetAllQQGroup")
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query()
	goutils.CheckErr(err)
	defer rows.Close()

	var groupInfoMap []*models.QQGroupInfo

	for rows.Next() { //有下一筆就會一直true下去
		qqGroupInfo := new(models.QQGroupInfo)
		err = rows.Scan(&qqGroupInfo.IdGroup,
			&qqGroupInfo.QQ,
			&qqGroupInfo.GroupName)
		goutils.CheckErr(err)
		groupInfoMap = append(groupInfoMap, qqGroupInfo)
	}
	return groupInfoMap

}

func DB_DelQQGroup(IdGroup string) {
	//Begin函数内部会去获取连接
	dbQueryStr := `
	DELETE FROM qqgroup WHERE idgroup = ?
	`
	log.Println("DB_DelQQGroup IdGroup: " + IdGroup)
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(IdGroup)
	goutils.CheckErr(err)
	defer rows.Close()
}

func DB_EditQQGroup(idGroup string, groupQQ string, groupname string) {
	//Begin函数内部会去获取连接
	dbQueryStr := `
	UPDATE qqgroup SET 
	groupQQ=?,
	groupname=?
	WHERE idgroup=?
	`

	log.Println("DB_EditQQGroup groupQQ: " + groupQQ + " groupname: " + groupname + " idGroup: " + idGroup)
	stm, err := db.Prepare(dbQueryStr)
	defer stm.Close()
	goutils.CheckErr(err)
	rows, err := stm.Query(groupQQ, groupname, idGroup)
	goutils.CheckErr(err)
	defer rows.Close()
}
