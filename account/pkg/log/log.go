package log

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"regexp"

	"go.uber.org/zap"
)

// Print hold the log key and values to be logged
type Print map[string]interface{}

var (
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger

	loggerHandler      *zap.Logger
	sugarLoggerHandler *zap.SugaredLogger
)

var CLIENT_ERROR = "CLIENT_ERROR"
var SERVER_ERROR = "SERVER_ERROR"

// SetLogLevel set level for log
//
// info | debug # default is info level
func SetLogLevel(level string) {
	switch level {
	case "info":
		InitProd()
	case "debug":
		InitDev()
	default:
		InitProd()
	}
}

func InitProd() {
	initial(zap.NewAtomicLevelAt(zap.InfoLevel))
	initialForHandler(zap.NewAtomicLevelAt(zap.InfoLevel))
}

func InitDev() {
	initial(zap.NewAtomicLevelAt(zap.DebugLevel))
	initialForHandler(zap.NewAtomicLevelAt(zap.DebugLevel))
}

func initial(level zap.AtomicLevel) {
	fmt.Println("Init Zap logger")

	var err error
	zapConfig := CreateZapConfig(false)
	zapConfig.Level = level

	logger, err = zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal("Creating Zap Logger error.")
	}
	sugarLogger = logger.Sugar()

	defer func() {
		_ = logger.Sync()
		_ = sugarLogger.Sync()
	}()
}

func initialForHandler(level zap.AtomicLevel) {
	fmt.Println("Init Zap logger")

	var err error
	zapConfig := CreateZapConfig(true)
	zapConfig.Level = level

	loggerHandler, err = zapConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatal("Creating Zap Logger error.")
	}
	sugarLoggerHandler = loggerHandler.Sugar()

	defer func() {
		_ = loggerHandler.Sync()
		_ = sugarLoggerHandler.Sync()
	}()
}

func CreateZapConfig(isHandler bool) *zap.Config {
	zapConfig := zap.NewDevelopmentConfig()

	zapConfig.EncoderConfig = GetEncoderConfig(isHandler)
	zapConfig.Encoding = "console"

	return &zapConfig
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Infof(msg string, args ...interface{}) {
	msg = initLogMessage(msg)
	sugarLogger.Infof(msg, args...)
}

func InfoCtx(ctx context.Context, msg string) {
	msg = initLogCtxMessage(ctx, msg)

	logger.Info(msg)
}

func InfoCtxf(ctx context.Context, msg string, args ...interface{}) {
	msg = initLogCtxMessage(ctx, msg)

	sugarLogger.Infof(msg, args...)
}

func InfoCtxHandlerf(ctx context.Context, msg string, args ...interface{}) {
	msg = initCtxHandlerMessage(ctx, msg)
	sugarLoggerHandler.Infof(msg, args...)
}

func InfoCtxNoFuncf(ctx context.Context, msg string, args ...interface{}) {
	msg = initLogCtxWithoutFuncMessage(ctx, msg)

	sugarLoggerHandler.Infof(msg, args...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Debugf(msg string, args ...interface{}) {
	msg = initLogMessage(msg)
	sugarLogger.Debugf(msg, args...)
}

func DebugCtxf(ctx context.Context, msg string, args ...interface{}) {
	msg = initLogCtxMessage(ctx, msg)

	sugarLogger.Debugf(msg, args...)
}

func Warnf(msg string, args ...interface{}) {
	msg = initLogMessage(msg)
	sugarLogger.Warnf(msg, args...)
}

func WarnCtxf(ctx context.Context, msg string, args ...interface{}) {
	msg = initLogCtxMessage(ctx, msg)
	sugarLogger.Warnf(msg, args...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Errorf(msg string, args ...interface{}) {
	msg = initLogMessage(msg)
	sugarLogger.Errorf(msg, args...)
}

func ErrorCtxf(ctx context.Context, msg string, args ...interface{}) {
	msg = initLogCtxMessage(ctx, msg)
	sugarLogger.Errorf(msg, args...)
}

func ErrorCtxHandlerf(ctx context.Context, msg string, args ...interface{}) {
	msg = initCtxHandlerMessage(ctx, msg)
	sugarLoggerHandler.Errorf(msg, args...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func Fatalf(msg string, args ...interface{}) {
	msg = initLogMessage(msg)
	sugarLogger.Fatalf(msg, args...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Panicf(msg string, args ...interface{}) {
	msg = initLogMessage(msg)
	sugarLogger.Panicf(msg, args...)
}

func initLogMessage(msg string) string {
	caller := GetFunctionNameAtRuntime(3)
	msg = msg + getfunctionName(caller.FunctionName)
	return msg
}

func initLogCtxMessage(ctx context.Context, msg string) string {
	caller := GetFunctionNameAtRuntime(3)
	final := msg + getfunctionName(caller.FunctionName) + getUserID(ctx) + getRequestID(ctx) + getTraceID(ctx)

	return final
}

func initLogCtxWithoutFuncMessage(ctx context.Context, msg string) string {
	final := msg + getUserID(ctx) + getRequestID(ctx) + getTraceID(ctx)

	return final
}

func initCtxHandlerMessage(ctx context.Context, msg string) string {
	caller := GetFunctionNameAtRuntime(4)
	final := msg + getfunctionName(caller.FunctionName) + getUserID(ctx) + getRequestID(ctx) + getTraceID(ctx)

	return final
}

// getfunctionName
func getfunctionName(functionName string) string {
	return fmt.Sprintf("	func:%s() ", functionName)
}

func getUserID(ctx context.Context) string {
	prefix := GetLogPrefix(ctx)
	if prefix.UserID == "" {
		return ""
	}

	return fmt.Sprintf("user_id=%s ", prefix.UserID)
}

// getTraceID
func getTraceID(ctx context.Context) string {
	prefix := GetLogPrefix(ctx)
	if prefix.TraceID == "" {
		return ""
	}

	return fmt.Sprintf("	trace_id=%s ", prefix.TraceID)
}

// getRequestID
func getRequestID(ctx context.Context) string {
	prefix := GetLogPrefix(ctx)
	if prefix.RequestID == "" {
		return ""
	}

	return fmt.Sprintf("	request_id=%s ", prefix.RequestID)
}

// ToJsonString converts an object to a JSON string. Returns an empty string if the object is nil.
// Sensitive data is replaced with ****.
func ToJsonString(obj interface{}) string {
	if obj == nil {
		return ""
	}

	// Create a copy of the object to avoid modifying the original
	copyObj := reflect.ValueOf(obj).Interface()

	// Convert the object to JSON
	jsonBytes, err := json.Marshal(copyObj)
	if err != nil {
		return ""
	}

	// Replace sensitive data with ****
	jsonString := string(jsonBytes)
	jsonString = replaceSensitiveData(jsonString)

	return jsonString
}

// Helper function to replace data of sensitive field with ****
func replaceSensitiveData(jsonString string) string {
	// Define your sensitive fields here
	sensitiveFields := []string{"creditCardNumber", "ssn",
		"Password", "password", "HashPassword", "hash_password",
		"otp", "OTP", "Otp", "OtpCode", "otpCode", "Otpcode", "otpcode",
		"AccessToken", "RefreshToken", "access_token", "refresh_token", "Authorization", "authorization"}

	for _, field := range sensitiveFields {
		re := regexp.MustCompile(`"` + field + `"\s*:\s*"[^"]*"`)
		jsonString = re.ReplaceAllString(jsonString, `"`+field+`":"****"`)
	}

	return jsonString
}
