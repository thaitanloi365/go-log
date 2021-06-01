package logger

import (
	"log"
	"testing"
	"time"

	"github.com/kjk/dailyrotate"

	"gopkg.in/natefinch/lumberjack.v2"
)

func TestRotateLogger(t *testing.T) {
	rotateLog, err := dailyrotate.NewFile("2006-01-02.log", func(path string, didRotate bool) {})
	if err != nil {
		panic(err)
	}

	var logger = New(&Config{
		BufferedSize: 100,
		Writer:       log.New(rotateLog, "\r\n", 0),
		Notifier:     NewSlackNotifier("", ""),
	})

	var errData = map[string]interface{}{
		"error_code":    500,
		"error_message": "internal server error",
	}

	var payload = map[string]interface{}{
		"user_id":   1,
		"user_name": "Loi",
	}
	for i := 0; i < 1; i++ {
		// logger.Debug("Test Debug", i)

		// logger.Debugf("Test Debugf %d", i)

		// logger.DebugWithTag("This is a tag", "Test Debug", i)

		// logger.DebugfWithTag("This is a tag", "Test Debugf %d", i)

		// logger.Info("Test Info", i)

		// logger.Infof("Test Infof %s %d", "1231231", i)

		// logger.DebugWithTag("This is a tag", "Test Debug", i)

		// logger.DebugfWithTag("This is a tag", "Test Debugf %d", i)

		// logger.InfoWithTag("This is a tag", "Test Info", i)

		// logger.InfofWithTag("This is a tag", "Test Infof %s %d", "1231231", i)

		// logger.InfoWithRequestInfo(&RequestInfo{
		// 	ReqID:      "reqID",
		// 	Status:     200,
		// 	Method:     "GET",
		// 	URI:        "https://test.com",
		// 	UserID:     "UserID",
		// 	RefErrorID: "RefErrorID",
		// }, "Test InfoWithRequestInfo", i)

		// logger.ErrorWithRequestInfo(&RequestInfo{
		// 	ReqID:      "reqID",
		// 	Status:     200,
		// 	Method:     "GET",
		// 	URI:        "https://test.com",
		// 	UserID:     "UserID",
		// 	RefErrorID: "RefErrorID",
		// }, "Test ErrorWithRequestInfo", payload, errData)

		logger.ErrorJSONWithRequestInfo(&RequestInfo{
			ReqID:      "reqID",
			Status:     200,
			Method:     "GET",
			URI:        "https://test.com",
			UserID:     "UserID",
			RefErrorID: "RefErrorID",
		}, "Test ErrorWithRequestInfo", payload, errData)

		logger.InfoWithRequestInfo(&RequestInfo{
			ReqID:      "reqID",
			Status:     200,
			Method:     "GET",
			URI:        "https://test.com",
			UserID:     "UserID",
			RefErrorID: "RefErrorID",
			Tag:        "This is a tag",
		}, "Test InfoWithRequestInfo", i)

		// logger.ErrorWithRequestInfo(&RequestInfo{
		// 	ReqID:      "reqID",
		// 	Status:     200,
		// 	Method:     "GET",
		// 	URI:        "https://test.com",
		// 	UserID:     "UserID",
		// 	RefErrorID: "RefErrorID",
		// 	Tag:        "This is a tag",
		// }, "Test ErrorWithRequestInfo", payload, errData)

		// logger.ErrorJSONWithRequestInfo(&RequestInfo{
		// 	ReqID:      "reqID",
		// 	Status:     200,
		// 	Method:     "GET",
		// 	URI:        "https://test.com",
		// 	UserID:     "UserID",
		// 	RefErrorID: "RefErrorID",
		// 	Tag:        "This is a tag",
		// }, "Test ErrorWithRequestInfo", payload, errData)

		time.Sleep(time.Second)
	}

	return

}

func TestLumperjackLogger(t *testing.T) {
	var writer = &lumberjack.Logger{
		Filename:   "foo.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	var logger = New(&Config{
		BufferedSize: 100,
		Writer:       log.New(writer, "\r\n", 0),
	})
	var data = []interface{}{
		"asdf", "ss", "sss",
	}

	logger.Debugf("%s\n[info] "+"asdf", append([]interface{}{"aaaaa"}, data...)...)
	for i := 0; i < 10; i++ {
		logger.Debugf("count %d \n", i)
		logger.Debug("count sssss", i, "asdfasdf")
		time.Sleep(time.Second)
	}
}
