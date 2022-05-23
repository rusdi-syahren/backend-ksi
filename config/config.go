package config

import (
	"errors"
	"os"
	"strconv"
	"time"

	dotenv "github.com/joho/godotenv"
)

var (

	// Development env checking, this env for debug purpose
	Development string
	// Hello config
	Hello string

	// HTTPPort config
	HTTPPort uint16
	// GRPCPort Config
	GRPCPort uint16

	// BasicAuthUsername config
	BasicAuthUsername string
	// BasicAuthPassword config
	BasicAuthPassword string

	// RedisHost config
	RedisHost string
	// RedisPort config
	RedisPort string
	// RedisPassword string
	RedisPassword string
	// RedisTLS string
	RedisTLS string

	// CacheExpired config
	CacheExpired time.Duration

	// WriteDBHost config
	WriteDBHost string
	// WriteDBName config
	WriteDBName string
	// WriteDBUser config
	WriteDBUser string
	// WriteDBPassword config
	WriteDBPassword string

	// ReadDBHost config
	ReadDBHost string
	// ReadDBName config
	ReadDBName string
	// ReadDBUser config
	ReadDBUser string
	// ReadDBPassword config
	ReadDBPassword string
)

// Load function will load all config from environment variable
func Load() error {
	// load .env
	err := dotenv.Load(".env")
	if err != nil {
		return errors.New(".env is not loaded properly")
	}

	development, ok := os.LookupEnv("DEVELOPMENT")
	if !ok {
		return errors.New("DEVELOPMENT env is not loaded")
	}

	// set Development
	Development = development

	hello, ok := os.LookupEnv("HELLO")
	if !ok {
		return errors.New("HELLO env is not loaded")
	}

	// set Hello
	Hello = hello

	// ------------------------------------
	redisHost, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		return errors.New("REDIS_HOST env is not loaded")
	}

	// set REDIS_TLS
	RedisHost = redisHost

	redisPort, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		return errors.New("REDIS_PORT env is not loaded")
	}

	// set RedisPort
	RedisPort = redisPort

	redisPassword, ok := os.LookupEnv("REDIS_AUTH")
	if !ok {
		return errors.New("REDIS_PORT env is not loaded")
	}

	// set RedisPassword
	RedisPassword = redisPassword

	redisTLS, ok := os.LookupEnv("REDIS_TLS")
	if !ok {
		return errors.New("REDIS_PORT env is not loaded")
	}

	// set RedisTLS
	RedisTLS = redisTLS

	cacheExpiredStr, ok := os.LookupEnv("CACHE_EXPIRED")

	if !ok {
		return errors.New("CACHE_EXPIRED is not loaded")
	}

	cacheExpired, err := time.ParseDuration(cacheExpiredStr)

	if err != nil {
		return errors.New("CACHE_EXPIRED is not valid")
	}

	CacheExpired = cacheExpired

	// ------------------------------------

	httpPortStr, ok := os.LookupEnv("HTTP_PORT")
	if !ok {
		return errors.New("HTTP_PORT env is not loaded")
	}

	httpPort, err := strconv.Atoi(httpPortStr)
	if err != nil {
		return errors.New("HTTP_PORT env is not valid")
	}

	// set http port
	HTTPPort = uint16(httpPort)

	grpcPortStr, ok := os.LookupEnv("GRPC_PORT")
	if !ok {
		return errors.New("GRPC_PORT env is not loaded")
	}

	grpcPort, err := strconv.Atoi(grpcPortStr)
	if err != nil {
		return errors.New("GRPC_PORT env is not valid")
	}

	// set grpc port
	GRPCPort = uint16(grpcPort)
	// ------------------------------------

	basicAuthUsername, ok := os.LookupEnv("BASIC_AUTH_USER")
	if !ok {
		return errors.New("BASIC_AUTH_USER env is not loaded")
	}

	// set BasicAuthUsername
	BasicAuthUsername = basicAuthUsername

	basicAuthPassword, ok := os.LookupEnv("BASIC_AUTH_PASS")
	if !ok {
		return errors.New("BASIC_AUTH_PASS env is not loaded")
	}

	// set BasicAuthPassword
	BasicAuthPassword = basicAuthPassword
	// ------------------------------------

	writeDBHost, ok := os.LookupEnv("WRITE_DB_HOST")
	if !ok {
		return errors.New("WRITE_DB_HOST env is not loaded")
	}

	// set WriteDBHost
	WriteDBHost = writeDBHost

	writeDBName, ok := os.LookupEnv("WRITE_DB_NAME")
	if !ok {
		return errors.New("WRITE_DB_NAME env is not loaded")
	}

	// set WriteDBName
	WriteDBName = writeDBName

	writeDBUser, ok := os.LookupEnv("WRITE_DB_USER")
	if !ok {
		return errors.New("WRITE_DB_USER env is not loaded")
	}

	// set WriteDBUser
	WriteDBUser = writeDBUser

	writeDBPassword, ok := os.LookupEnv("WRITE_DB_PASSWORD")
	if !ok {
		return errors.New("WRITE_DB_PASSWORD env is not loaded")
	}

	// set WriteDBPassword
	WriteDBPassword = writeDBPassword
	// ------------------------------------

	readDBHost, ok := os.LookupEnv("READ_DB_HOST")
	if !ok {
		return errors.New("READ_DB_HOST env is not loaded")
	}

	// set ReadDBHost
	ReadDBHost = readDBHost

	readDBName, ok := os.LookupEnv("READ_DB_NAME")
	if !ok {
		return errors.New("READ_DB_NAME env is not loaded")
	}

	// set ReadDBName
	ReadDBName = readDBName

	readDBUser, ok := os.LookupEnv("READ_DB_USER")
	if !ok {
		return errors.New("READ_DB_USER env is not loaded")
	}

	// set ReadDBUser
	ReadDBUser = readDBUser

	readDBPassword, ok := os.LookupEnv("READ_DB_PASSWORD")
	if !ok {
		return errors.New("READ_DB_PASSWORD env is not loaded")
	}

	// set ReadDBPassword
	ReadDBPassword = readDBPassword
	// ------------------------------------

	return nil
}
