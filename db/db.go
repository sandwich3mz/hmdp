package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
	"hmdp/config"
	"hmdp/ent"
	"strconv"
)

func InitializeDB() *ent.Client {
	dbConfig := config.Conf.Database
	dsn := dbConfig.UserName + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + strconv.Itoa(dbConfig.Port) +
		")/" + dbConfig.Database + "?charset=" + dbConfig.Charset + "&parseTime=True&loc=Local"
	logrus.Debugf("dsn: %v", dsn)
	client, err := ent.Open(dbConfig.Driver, dsn)
	if err != nil {
		logrus.Errorf("failed at creating ent client with db %s, err: %v", dsn, err)
		return nil
	}
	//client.SetMaxIdleConns(dbConfig.MaxIdleConns)
	//client.SetMaxOpenConns(dbConfig.MaxOpenConns)
	//drv := entsql.OpenDB(dbConfig.Driver, db)

	return client
}
