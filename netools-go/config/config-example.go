package config

import "time"

const ServerAddr string = "localhost:8081"

const RedisAddr string = "x.x.x.x:6379"
const RedisPwd string = "xxxxxxxxxx"

var TextKeyPrefix string = "text_"
var FileKeyPrefix string = "file_"
var KeyExpireTime time.Duration = 60 * time.Minute //60min

var TempFilePath string = "./upload/"
