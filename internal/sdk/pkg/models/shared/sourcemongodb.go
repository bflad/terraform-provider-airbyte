// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/airbytehq/terraform-provider-airbyte/internal/sdk/pkg/utils"
)

type SourceMongodbSchemasInstanceTypeInstance string

const (
	SourceMongodbSchemasInstanceTypeInstanceAtlas SourceMongodbSchemasInstanceTypeInstance = "atlas"
)

func (e SourceMongodbSchemasInstanceTypeInstance) ToPointer() *SourceMongodbSchemasInstanceTypeInstance {
	return &e
}

func (e *SourceMongodbSchemasInstanceTypeInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "atlas":
		*e = SourceMongodbSchemasInstanceTypeInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbSchemasInstanceTypeInstance: %v", v)
	}
}

// SourceMongodbMongoDBAtlas - The MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type SourceMongodbMongoDBAtlas struct {
	AdditionalProperties interface{} `additionalProperties:"true" json:"-"`
	// The URL of a cluster to connect to.
	ClusterURL string                                   `json:"cluster_url"`
	instance   SourceMongodbSchemasInstanceTypeInstance `const:"atlas" json:"instance"`
}

func (s SourceMongodbMongoDBAtlas) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbMongoDBAtlas) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbMongoDBAtlas) GetAdditionalProperties() interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *SourceMongodbMongoDBAtlas) GetClusterURL() string {
	if o == nil {
		return ""
	}
	return o.ClusterURL
}

func (o *SourceMongodbMongoDBAtlas) GetInstance() SourceMongodbSchemasInstanceTypeInstance {
	return SourceMongodbSchemasInstanceTypeInstanceAtlas
}

type SourceMongodbSchemasInstance string

const (
	SourceMongodbSchemasInstanceReplica SourceMongodbSchemasInstance = "replica"
)

func (e SourceMongodbSchemasInstance) ToPointer() *SourceMongodbSchemasInstance {
	return &e
}

func (e *SourceMongodbSchemasInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "replica":
		*e = SourceMongodbSchemasInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbSchemasInstance: %v", v)
	}
}

// SourceMongodbReplicaSet - The MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type SourceMongodbReplicaSet struct {
	instance SourceMongodbSchemasInstance `const:"replica" json:"instance"`
	// A replica set in MongoDB is a group of mongod processes that maintain the same data set.
	ReplicaSet *string `json:"replica_set,omitempty"`
	// The members of a replica set. Please specify `host`:`port` of each member separated by comma.
	ServerAddresses string `json:"server_addresses"`
}

func (s SourceMongodbReplicaSet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbReplicaSet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbReplicaSet) GetInstance() SourceMongodbSchemasInstance {
	return SourceMongodbSchemasInstanceReplica
}

func (o *SourceMongodbReplicaSet) GetReplicaSet() *string {
	if o == nil {
		return nil
	}
	return o.ReplicaSet
}

func (o *SourceMongodbReplicaSet) GetServerAddresses() string {
	if o == nil {
		return ""
	}
	return o.ServerAddresses
}

type SourceMongodbInstance string

const (
	SourceMongodbInstanceStandalone SourceMongodbInstance = "standalone"
)

func (e SourceMongodbInstance) ToPointer() *SourceMongodbInstance {
	return &e
}

func (e *SourceMongodbInstance) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "standalone":
		*e = SourceMongodbInstance(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbInstance: %v", v)
	}
}

// SourceMongodbStandaloneMongoDbInstance - The MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
type SourceMongodbStandaloneMongoDbInstance struct {
	// The host name of the Mongo database.
	Host     string                `json:"host"`
	instance SourceMongodbInstance `const:"standalone" json:"instance"`
	// The port of the Mongo database.
	Port *int64 `default:"27017" json:"port"`
}

func (s SourceMongodbStandaloneMongoDbInstance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodbStandaloneMongoDbInstance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodbStandaloneMongoDbInstance) GetHost() string {
	if o == nil {
		return ""
	}
	return o.Host
}

func (o *SourceMongodbStandaloneMongoDbInstance) GetInstance() SourceMongodbInstance {
	return SourceMongodbInstanceStandalone
}

func (o *SourceMongodbStandaloneMongoDbInstance) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type SourceMongodbMongoDbInstanceTypeType string

const (
	SourceMongodbMongoDbInstanceTypeTypeStandaloneMongoDbInstance SourceMongodbMongoDbInstanceTypeType = "StandaloneMongoDbInstance"
	SourceMongodbMongoDbInstanceTypeTypeReplicaSet                SourceMongodbMongoDbInstanceTypeType = "ReplicaSet"
	SourceMongodbMongoDbInstanceTypeTypeMongoDBAtlas              SourceMongodbMongoDbInstanceTypeType = "MongoDBAtlas"
)

type SourceMongodbMongoDbInstanceType struct {
	StandaloneMongoDbInstance *SourceMongodbStandaloneMongoDbInstance
	ReplicaSet                *SourceMongodbReplicaSet
	MongoDBAtlas              *SourceMongodbMongoDBAtlas

	Type SourceMongodbMongoDbInstanceTypeType
}

func CreateSourceMongodbMongoDbInstanceTypeStandaloneMongoDbInstance(standaloneMongoDbInstance SourceMongodbStandaloneMongoDbInstance) SourceMongodbMongoDbInstanceType {
	typ := SourceMongodbMongoDbInstanceTypeTypeStandaloneMongoDbInstance

	return SourceMongodbMongoDbInstanceType{
		StandaloneMongoDbInstance: &standaloneMongoDbInstance,
		Type:                      typ,
	}
}

func CreateSourceMongodbMongoDbInstanceTypeReplicaSet(replicaSet SourceMongodbReplicaSet) SourceMongodbMongoDbInstanceType {
	typ := SourceMongodbMongoDbInstanceTypeTypeReplicaSet

	return SourceMongodbMongoDbInstanceType{
		ReplicaSet: &replicaSet,
		Type:       typ,
	}
}

func CreateSourceMongodbMongoDbInstanceTypeMongoDBAtlas(mongoDBAtlas SourceMongodbMongoDBAtlas) SourceMongodbMongoDbInstanceType {
	typ := SourceMongodbMongoDbInstanceTypeTypeMongoDBAtlas

	return SourceMongodbMongoDbInstanceType{
		MongoDBAtlas: &mongoDBAtlas,
		Type:         typ,
	}
}

func (u *SourceMongodbMongoDbInstanceType) UnmarshalJSON(data []byte) error {

	standaloneMongoDbInstance := new(SourceMongodbStandaloneMongoDbInstance)
	if err := utils.UnmarshalJSON(data, &standaloneMongoDbInstance, "", true, true); err == nil {
		u.StandaloneMongoDbInstance = standaloneMongoDbInstance
		u.Type = SourceMongodbMongoDbInstanceTypeTypeStandaloneMongoDbInstance
		return nil
	}

	replicaSet := new(SourceMongodbReplicaSet)
	if err := utils.UnmarshalJSON(data, &replicaSet, "", true, true); err == nil {
		u.ReplicaSet = replicaSet
		u.Type = SourceMongodbMongoDbInstanceTypeTypeReplicaSet
		return nil
	}

	mongoDBAtlas := new(SourceMongodbMongoDBAtlas)
	if err := utils.UnmarshalJSON(data, &mongoDBAtlas, "", true, true); err == nil {
		u.MongoDBAtlas = mongoDBAtlas
		u.Type = SourceMongodbMongoDbInstanceTypeTypeMongoDBAtlas
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceMongodbMongoDbInstanceType) MarshalJSON() ([]byte, error) {
	if u.StandaloneMongoDbInstance != nil {
		return utils.MarshalJSON(u.StandaloneMongoDbInstance, "", true)
	}

	if u.ReplicaSet != nil {
		return utils.MarshalJSON(u.ReplicaSet, "", true)
	}

	if u.MongoDBAtlas != nil {
		return utils.MarshalJSON(u.MongoDBAtlas, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

type SourceMongodbMongodb string

const (
	SourceMongodbMongodbMongodb SourceMongodbMongodb = "mongodb"
)

func (e SourceMongodbMongodb) ToPointer() *SourceMongodbMongodb {
	return &e
}

func (e *SourceMongodbMongodb) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mongodb":
		*e = SourceMongodbMongodb(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceMongodbMongodb: %v", v)
	}
}

type SourceMongodb struct {
	// The authentication source where the user information is stored.
	AuthSource *string `default:"admin" json:"auth_source"`
	// The database you want to replicate.
	Database string `json:"database"`
	// The MongoDb instance to connect to. For MongoDB Atlas and Replica Set TLS connection is used by default.
	InstanceType *SourceMongodbMongoDbInstanceType `json:"instance_type,omitempty"`
	// The password associated with this username.
	Password   *string              `json:"password,omitempty"`
	sourceType SourceMongodbMongodb `const:"mongodb" json:"sourceType"`
	// The username which is used to access the database.
	User *string `json:"user,omitempty"`
}

func (s SourceMongodb) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(s, "", false)
}

func (s *SourceMongodb) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &s, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *SourceMongodb) GetAuthSource() *string {
	if o == nil {
		return nil
	}
	return o.AuthSource
}

func (o *SourceMongodb) GetDatabase() string {
	if o == nil {
		return ""
	}
	return o.Database
}

func (o *SourceMongodb) GetInstanceType() *SourceMongodbMongoDbInstanceType {
	if o == nil {
		return nil
	}
	return o.InstanceType
}

func (o *SourceMongodb) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *SourceMongodb) GetSourceType() SourceMongodbMongodb {
	return SourceMongodbMongodbMongodb
}

func (o *SourceMongodb) GetUser() *string {
	if o == nil {
		return nil
	}
	return o.User
}
