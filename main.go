package main

import (
	"log"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type TestObject struct {
	UserName   string
	FirstName  string
	LastName   string
	UserID     int
	HandleName string
}

func (t TestObject) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("userName", t.UserName)
	enc.AddInt("userID", t.UserID)
	return nil
}

func main() {
	// Example 1
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		// Structured context as strongly typed Field values.
		zap.Object("object", TestObject),
		zap.String("url", "http://google.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	logger2, err := zap.Config{
		Level: zap.NewAtomicLevelAt(zapcore.InfoLevel), Development: true, Encoding: "json", OutputPaths: []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger2.Sync()
	logger2.Error("Error")
	logger2.Error("Error")
	logger2.Info("Info", zap.Int("attempt", 3))
	logger2.Debug("Debug", zap.Duration("backoff", time.Second))
}
