package db

import (
	"crm"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"gopkg.in/gorp.v2"
)

var log *logrus.Entry
var pdb *sql.DB
var mapper *gorp.DbMap

type typeConverter struct{}

func (c typeConverter) ToDb(val interface{}) (interface{}, error) {
	switch t := val.(type) {
	case crm.Contact:
		b, err := json.Marshal(t)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}

	return val, nil
}

func (c typeConverter) FromDb(target interface{}) (gorp.CustomScanner, bool) {
	switch target.(type) {
	case *crm.Contact:
		binder := func(holder, target interface{}) error {
			s, ok := holder.(*string)
			if !ok {
				return errors.New("FromDb: Unable to convert Inventory to *string")
			}
			b := []byte(*s)
			return json.Unmarshal(b, target)
		}
		return gorp.CustomScanner{new(string), target, binder}, true
	}

	return gorp.CustomScanner{}, false
}

func init() {
	log = logrus.WithField("component", "db")
}

func InitPostgreSQL(host string) {
	pdb, _ = sql.Open("postgres", fmt.Sprintf("postgres://postgres@%s/postgres?sslmode=disable", host))
	mapper = &gorp.DbMap{Db: pdb, Dialect: gorp.PostgresDialect{}}
	mapper.TypeConverter = typeConverter{}
	mapper.AddTableWithName(crm.Contact{}, "contact").SetKeys(true, "ID")

	log.Infof("Using PostgreSQL @ %s", host)
}
