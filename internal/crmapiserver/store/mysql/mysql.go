package mysql

import (
	"fmt"
	"github.com/marmotedu/errors"
	"github.com/marmotedu/iam/internal/crmapiserver/store"
	"github.com/marmotedu/iam/internal/pkg/logger"
	"github.com/marmotedu/iam/internal/pkg/options"
	"github.com/marmotedu/iam/pkg/db"
	"gorm.io/gorm"
	"sync"
)

type datastore struct {
	db *gorm.DB
}

func (ds *datastore) CrmApplet() store.CrmAppletStore {
	return newCrmApplet(ds)
}

func (ds *datastore) Admin() store.AdminStore {
	return newAdmins(ds)
}

func (ds *datastore) Close() error {
	db, err := ds.db.DB()
	if err != nil {
		return errors.Wrap(err, "get gorm db instance failed")
	}
	return db.Close()
}

var (
	once    sync.Once
	factory store.Factory
)

func GetMySQLFactory(opts *options.MySQLOptions) (store.Factory, error) {
	if opts == nil && factory == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory")
	}

	var err error
	var dbIns *gorm.DB
	once.Do(func() {
		options := &db.Options{
			Host:                  opts.Host,
			Username:              opts.Username,
			Password:              opts.Password,
			Database:              opts.Database,
			MaxIdleConnections:    opts.MaxIdleConnections,
			MaxOpenConnections:    opts.MaxOpenConnections,
			MaxConnectionLifeTime: opts.MaxConnectionLifeTime,
			LogLevel:              opts.LogLevel,
			Logger:                logger.New(opts.LogLevel),
		}
		dbIns, err = db.New(options)

		// uncomment the following line if you need auto migration the given models
		// not suggested in production environment.
		// migrateDatabase(dbIns)

		factory = &datastore{dbIns}
	})

	if factory == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", factory, err)
	}

	return factory, nil
}
