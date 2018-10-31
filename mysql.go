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
func NewMysqlOperate(address string, user string, password string, db string) *DbWorker {
	once.Do(func() {
		ins = &DbWorker{
			Dsn: user + `:` + password + `@tcp(` + address + `)/` + db + `?charset=utf8`,
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

// 设置mysql的链接信息
func (dbw *DbWorker) SetConnInfo(address string, user string, password string, db string) *DbWorker {
	dbw.Dsn = user + `:` + password + `@tcp(` + address + `)/` + db + `?charset=utf8`
	return dbw
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

// update delete insert 都统一使用这函数进行操作
func (dbw *DbWorker) UpdateData(sqlString string, args ...interface{}) error {
	fmt.Println(dbw.Dsn)
	stmt, testerr := dbw.Db.Prepare(sqlString)

	if testerr != nil {
		return testerr
	}
	defer stmt.Close()

	ret, err := stmt.Exec(args...)
	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		return err
	}
	if LastInsertId, err1 := ret.LastInsertId(); nil == err {
		fmt.Println("LastInsertId:", LastInsertId)
	} else {
		return err1
	}
	if RowsAffected, err2 := ret.RowsAffected(); nil == err {
		fmt.Println("RowsAffected:", RowsAffected)
	} else {
		return err2
	}
	return nil
}

func (dbw *DbWorker) QueryDataPre(ResStruct interface{}) {
	dbw.ResStruct = ResStruct
}

func (dbw *DbWorker) DeleteData(deleteString string) error {
	stmt, testerr := dbw.Db.Prepare(deleteString)
	if testerr != nil {
		fmt.Println(testerr)
		return nil
	}

	defer stmt.Close()

	ret, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ret.LastInsertId)
	return nil
}

func (dbw *DbWorker) QueryData(queryString string, args ...interface{}) ([]map[string]string, error) {
	// stmt, _ := dbw.Db.Prepare(`SELECT * From user where age >= ? AND age < ?`)
	stmt, testerr := dbw.Db.Prepare(queryString)
	if testerr != nil {
		return make([]map[string]string, 0), testerr
	}
	defer stmt.Close()

	// dbw.QueryDataPre(ResStruct)

	rows, err := stmt.Query(args...)
	defer rows.Close()
	if err != nil {
		return make([]map[string]string, 0), err
	}

	retunSlice := make([]map[string]string, 0)

	for rows.Next() {
		// fmt.Println(1)
		// rows.Scan(&test.IndexName, &test.RoleName, &test.Roles)
		// fmt.Println(test.IndexName)
		columns, cerror := rows.Columns()
		if cerror != nil {
			return make([]map[string]string, 0), cerror
		}

		scanArgs := make([]interface{}, len(columns))
		values := make([]interface{}, len(columns))

		for i := range values {
			scanArgs[i] = &values[i]
		}

		//将数据保存到 record 字典
		serr := rows.Scan(scanArgs...)
		if serr != nil {
			return make([]map[string]string, 0), cerror
		}
		record := make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
		retunSlice = append(retunSlice, record)
	}

	// fmt.Println(rows)
	// for rows.Next() {
	// 	rows.Scan(&ResStruct.IndexName, &ResStruct.RoleName, &ResStruct.Roles)
	// 	// rows.Scan(args)
	// 	if err != nil {
	// 		fmt.Printf(err.Error())
	// 		continue
	// 	}
	// 	// if !dbw.UserInfo.IndexName.Valid {
	// 	// 	dbw.UserInfo.IndexName.String = ""
	// 	// }
	// 	// if !dbw.UserInfo.RoleName.Valid {
	// 	// 	dbw.UserInfo.RoleName.String = ""
	// 	// }
	// 	fmt.Println(ResStruct.IndexName)
	// 	// fmt.Println("get data, id: ", dbw.UserInfo.IndexName, " name: ", dbw.UserInfo.Name.String, " age: ", int(dbw.UserInfo.Age.Int64))
	// }

	rerr := rows.Err()
	if rerr != nil {
		return make([]map[string]string, 0), rerr
	}
	return retunSlice, nil
}
