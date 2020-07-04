package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type (
	databaseContext struct {
		cxt    context.Context
		cancel context.CancelFunc
	}
	Database struct {
		user   string
		pass   string
		addr   string
		port   uint16
		cxt    *databaseContext
		client *mongo.Client
	}
)

func (d *Database) createConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%d", d.user, d.pass, d.addr, d.port)
}

func (d *Database) Connect() error {
	connStr := d.createConnectionString()
	log.Printf("Connecting to %s", connStr)

	cxt, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	d.cxt = &databaseContext{cxt, cancel}

	client, err := mongo.Connect(d.cxt.cxt, options.Client().ApplyURI(connStr))
	if err != nil {
		return fmt.Errorf("failed to connect to database")
	}
	d.client = client

	log.Printf("Testing connection with ping")
	if err = d.client.Ping(d.cxt.cxt, readpref.Primary()); err != nil {
		return fmt.Errorf("ping failed")
	}

	return nil
}

func (d *Database) Close() {
	log.Print("Closing database connection")
	d.cxt.cancel()
	if err := d.client.Disconnect(d.cxt.cxt); err != nil {
		log.Fatalf("failed to disconnect from database")
	}
}

func New(user, pass, addr string, port uint16) *Database {
	return &Database{user: user, pass: pass, addr: addr, port: port}
}
