package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"time"
)

type MysqlClient struct {
	Conn        *sql.DB
	ConnCluster *sql.DB
}

func NewMysqlDatabase(env *Env) MysqlClient {
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dbHost := env.DBMysqlHost
	dbPort := env.DBMysqlPort
	dbUser := env.DBMysqlUser
	dbPass := env.DBMysqlPass
	dbName := env.DBMysqlName

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	dbClusterHost := env.DBMysqlClusterHost
	dbClusterPort := env.DBMysqlClusterPort
	dbClusterUser := env.DBMysqlClusterUser
	dbClusterPass := env.DBMysqlClusterPass
	dbClusterName := env.DBMysqlClusterName
	connectionCluster := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbClusterUser, dbClusterPass, dbClusterHost, dbClusterPort, dbClusterName)
	client, err := NewClient(connection, connectionCluster)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Conn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func NewClient(connection string, connectionCluster string) (MysqlClient, error) {
	time.Local = time.UTC
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	val.Add("multiStatements", "true")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())

	dbConn, err := sql.Open(`mysql`, dsn)
	err = dbConn.Ping()
	if err != nil {
		log.Fatal(err)
	}

	dsnCluster := fmt.Sprintf("%s?%s", connectionCluster, val.Encode())
	dbConnCluster, err := sql.Open(`mysql`, dsnCluster)
	err = dbConnCluster.Ping()
	if err != nil {
		fmt.Println("failed connection to dsnCluster")
		log.Fatal(err)
	}

	//
	//
	//path := fmt.Sprintf("mysql://%s", dsn)
	//m, err := migrate.New("file://./db/migrations", path)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//if err := m.Up(); err != nil {
	//	if err == migrate.ErrNoChange {
	//		log.Println("No migrations to apply")
	//	} else {
	//		log.Fatal(err)
	//	}
	//}
	//
	//// cluster
	//dsnCluster := fmt.Sprintf("%s?%s", connectionCluster, val.Encode())
	//dbConnCluster, err := sql.Open(`mysql`, dsnCluster)
	//err = dbConnCluster.Ping()
	//if err != nil {
	//	fmt.Println("failed connection to dsnCluster")
	//	log.Fatal(err)
	//}
	//
	//pathCluster := fmt.Sprintf("mysql://%s", dsnCluster)
	//mCluster, errCLuster := migrate.New("file://./db/migrations", pathCluster)
	//if errCLuster != nil {
	//	fmt.Println("migrationError")
	//	log.Fatal(err)
	//}
	//
	//if err := mCluster.Up(); err != nil {
	//	if err == migrate.ErrNoChange {
	//		log.Println("No migrations to apply cluster")
	//	} else {
	//		log.Fatal(err)
	//	}
	//}

	return MysqlClient{Conn: dbConn, ConnCluster: dbConnCluster}, err
}
