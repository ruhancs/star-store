package db

import (
	"fmt"

	"github.com/gocql/gocql"
)

func ConnetToCassandraCluster() (*gocql.Session,error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "store"
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return session,nil
}