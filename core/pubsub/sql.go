package pubsub

import (
	stdSQL "database/sql"
	"fmt"

	mysqldriver "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-sql/v3/pkg/sql"
	"github.com/ThreeDotsLabs/watermill/message"
)

type SQL struct {
	logger watermill.LoggerAdapter
	dbType string
	client *stdSQL.DB
}

func NewSQL(
	dbType string,
	dbClient *stdSQL.DB,
) Container {
	logger := watermill.NewStdLogger(false, false)

	return &SQL{
		logger: logger,
		dbType: dbType,
		client: dbClient,
	}
}

func (pubSub SQL) schemaAdapter() sql.SchemaAdapter {
	if pubSub.dbType == "postgres" {
		return sql.DefaultPostgreSQLSchema{}
	}
	return sql.DefaultMySQLSchema{}
}

func (pubSub SQL) offsetsAdapter() sql.OffsetsAdapter {
	if pubSub.dbType == "postgres" {
		return sql.DefaultPostgreSQLOffsetsAdapter{}
	}
	return sql.DefaultMySQLOffsetsAdapter{}
}

func (pubSub SQL) NewSubscriber() message.Subscriber {

	subscriber, err := sql.NewSubscriber(
		pubSub.client,
		sql.SubscriberConfig{
			SchemaAdapter:    pubSub.schemaAdapter(),
			OffsetsAdapter:   pubSub.offsetsAdapter(),
			InitializeSchema: true,
		},
		pubSub.logger,
	)
	if err != nil {
		panic(err)
	}

	return subscriber
}

func (pubSub SQL) NewPublisher() message.Publisher {
	publisher, err := sql.NewPublisher(
		pubSub.client,
		sql.PublisherConfig{
			SchemaAdapter: pubSub.schemaAdapter(),
		},
		pubSub.logger,
	)
	if err != nil {
		panic(err)
	}

	return publisher
}

func dsn(sqlDBType string, sqlUser string, sqlPass string, sqlAddr string, sqlPort string, sqlDBName string) string {
	switch sqlDBType {
	case "mysql":
		conf := mysqldriver.NewConfig()
		conf.Net = "tcp"
		conf.User = sqlUser
		conf.Passwd = sqlPass
		conf.Addr = sqlAddr + ":" + sqlPort
		conf.DBName = sqlDBName
		return conf.FormatDSN()
	case "postgres":
		return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
			sqlUser, sqlPass, sqlDBName, sqlAddr, sqlPort)
	default:
		return ""
	}
}

func createDB(
	sqlDBType string,
	sqlUser string,
	sqlPass string,
	sqlAddr string,
	sqlPort string,
	sqlDBName string,
) *stdSQL.DB {
	db, err := stdSQL.Open(sqlDBType,
		dsn(sqlDBType, sqlUser, sqlPass, sqlAddr, sqlPort, sqlDBName),
	)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
