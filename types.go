package mongovault

import (
	"github.com/hashicorp/vault/api"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB defines all the function exposed for interactions
type DB interface {
	// Get loads the first entry, that matches the filters into
	// the given interface
	Get([]Filter, interface{}) error
	// GetAll loads all entries, that match the filters, into
	// the given interface
	GetAll([]Filter, interface{}) error
	// Insert is used to insert the given into the Database
	Insert(interface{}) error
	// Delete removes the first element, that matches the given
	// filters, from the Database
	Delete([]Filter) error
	// DeleteMany removes all the elements, that match the given
	// filters, from the Database
	DeleteMany([]Filter) error
	// Update updates the all elements, that match the given
	// filters, by updating the values according to the given
	// UpdateValue
	Update([]Filter, UpdateValue, ...*options.UpdateOptions) error
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

	VaultSession  *api.Client
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
	Session   *api.Client
	CredsName string
}

// Filter is a single Filter that can be used to identify an entry
type Filter struct {
	Key   string
	Value interface{}
}

// UpdateValue is used as a single update Entry
type UpdateValue map[string]interface{}
