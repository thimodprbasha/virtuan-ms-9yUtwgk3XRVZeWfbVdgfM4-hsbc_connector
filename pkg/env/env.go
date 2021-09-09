package env

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
	gorm "gorm.io/gorm"
	model "hsbc_connector/pkg/model"
)

var RDB *gorm.DB
var MDB *mongo.Database
var CMDCtrl *model.StorageCtrl
var QueryCtrl *model.StorageCtrl
var ESCtrl *model.MessageCtrl

const GRPC_PORT = "GRPC_PORT"
const REST_PORT = "REST_PORT"

var GRPCPORT string
var RESTPORT string
