package main

import (
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"sync"
	"time"
)


func main() {

	hostStr := flag.String("host", "127.0.0.1", "MySQL Addr")
	portStr := flag.String("port", "3306", "MySQL Port")
	connectionStr := flag.String("c", "1000", "连接数")
	connectionStrInt, _ := strconv.Atoi(*connectionStr)
	databaseStr := flag.String("database", "test", "MySQL Databases")
	userStr := flag.String("user", "shilei", "MySQL Username")
	passStr := flag.String("pass", "shilei", "MySQL Password")
	helpStr := flag.String("help", "", `
只写入单列数据，所以需要创建特殊的表结构:
	CREATE DATABASES test;
	CREATE TABLE test.tt (id int PRIMARY KEY auto_increment,data varchar(255));
插入的数据格式为:
	INSERT INTO tt(data) VALUES(N+1);
如果数据插入格式错误，请自行排查!`)
	logStr := flag.Bool("log", false, "是否开启 log")
	flag.Parse()

	db, _ := sqlx.Open("mysql", *userStr + ":" + *passStr + "@tcp(" + *hostStr + ":" + *portStr + ")/" + *databaseStr)
	defer db.Close()
	db.SetMaxOpenConns(0)
	db.SetMaxIdleConns(2000)

	var sm sync.WaitGroup

	for j := 1; j < connectionStrInt; j++ {
		ticker := time.NewTicker(1 * time.Second)
		sm.Add(1)
		go func() {

			for j := 1; j < 30000000; j++ {
				result, err := db.Exec("insert into tt(data) values(?);", strconv.Itoa(j))
				if err != nil{
					fmt.Println("Error!", err, helpStr)
				}

				if *logStr {
					fmt.Println(&result)
				}
			}

			sm.Done()
		}()
		ticker.Stop()
	}

	sm.Wait()
}


