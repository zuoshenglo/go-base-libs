package go_base_libs

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var ins *DbWorker
var once sync.Once

// 用sync.Once 使mysql的连接为单例模式
func NewMysqlOperate(dsn string) *DbWorker {
	once.Do(func() {
		ins = &DbWorker{
			Dsn: dsn,
		}
	})

	var err error
	ins.Db, err = sql.Open("mysql", ins.Dsn)
	if err != nil {
		// panic(err)
		fmt.Println("打开数据库失败：", err)
		return ins
	}
	// defer ins.Db.Close()
	return ins
}

// 手动关闭
func (dbw *DbWorker) CloseCnn() {
	dbw.Db.Close()
}

type DbWorker struct {
	Dsn       string
	Db        *sql.DB
	ResStruct interface{}
}

type userTB struct {
	IndexName sql.NullString
	RoleName  sql.NullString
	Roles     sql.NullString
}

// func main() {
// 	var err error
// 	dbw := DbWorker{
// 		Dsn: "root:123456@tcp(localhost:3306)/sqlx_db?charset=utf8mb4",
// 	}
// 	dbw.Db, err = sql.Open("mysql", dbw.Dsn)
// 	if err != nil {
// 		panic(err)
// 		return
// 	}
// 	defer dbw.Db.Close()

// 	dbw.insertData()
// 	dbw.queryData()
// }

func (dbw *DbWorker) InsertData(insertString string, args ...interface{}) {
	stmt, testerr := dbw.Db.Prepare(insertString)
	if testerr != nil {
		fmt.Println("初始化失败：", testerr)
	}
	// stmt, err := db.Prepare("Insert userinfo set username=?,departname=?,created=?")
	defer stmt.Close()

	fmt.Println(insertString)
	ret, err := stmt.Exec()
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return
	}
	if LastInsertId, err := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	}
	if RowsAffected, err := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	}
}

func (dbw *DbWorker) QueryDataPre(ResStruct interface{}) {
	dbw.ResStruct = ResStruct
}

func (dbw *DbWorker) QueryData(ResStruct TestStruct, args ...*string) {
	// stmt, _ := dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`)
	stmt, testerr := dbw.Db.Prepare(`SELECT * FROM alter_roles`)
	if testerr != nil {
		fmt.Println(testerr)
	}
	defer stmt.Close()

	dbw.QueryDataPre(ResStruct)

	rows, err := stmt.Query()
	defer rows.Close()
	if err != nil {
		fmt.Printf("select data error: %v\n", err)
		return
	}
	fmt.Println(args)
	for rows.Next() {
		rows.Scan(args)
		// rows.Scan(args)
		if err != nil {
			fmt.Printf(err.Error())
			continue
		}
		// if !dbw.UserInfo.IndexName.Valid {
		// 	dbw.UserInfo.IndexName.String = ""
		// }
		// if !dbw.UserInfo.RoleName.Valid {
		// 	dbw.UserInfo.RoleName.String = ""
		// }
		fmt.Println(ResStruct.IndexName)
		// fmt.Println("get data, id: ", dbw.UserInfo.IndexName, " name: ", dbw.UserInfo.Name.String, " age: ", int(dbw.UserInfo.Age.Int64))
	}

	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
	}
}

type TestStruct struct {
	IndexName sql.NullString
	RoleName  sql.NullString
	Roles     sql.NullString
}
