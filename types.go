package mongovault

import (
	"github.com/Lol3rrr/cvault"
	"go.mongodb.org/mongo-driver/mongo"
)

// DB defines all the function exposed for interactions
type DB interface {
	Connect() error
	Get([]Filter, interface{}) error
	Insert(interface{}) error
	Delete([]Filter) error
	GetAll([]Filter, interface{}) error
	Update([]Filter, UpdateValue) error
}

// Session is the actual struct that implements the DB interface
type Session struct {
	URL             string
	Port            string
	Database        string
	Collection      string
	Username        string
	Password        string
	ApplicationName string

	MongoClient     *mongo.Client
	MongoCollection *mongo.Collection

	VaultSession  cvault.Session
	CredsEndpoint string
}

// DBOptions is used to provide the Base Settings for a Session
type DBOptions struct {
	URL             string
	Port            string
	Database        string
	Collection      string
	ApplicationName string
}

// VaultSettings is used to provide the Base Settings for a Session
type VaultSettings struct {
	Session   cvault.Session
	CredsName string
}

// Filter is a single Filter that can be used to identify an entry
type Filter struct {
	Key   string
	Value interface{}
}

// UpdateValue is used as a single update Entry
type UpdateValue map[string]interface{}
