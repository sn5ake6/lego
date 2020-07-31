package mysqlconnector

import (
	"github.com/go-sql-driver/mysql"
	"github.com/vseinstrumentiru/lego/tools/lego"
	"logur.dev/logur"
)

// SetLogger configures the global database logger.
func SetLogger(logger lego.LogErr) {
	logger = logger.WithFields(map[string]interface{}{"component": "mysql"})

	_ = mysql.SetLogger(logur.NewErrorPrintLogger(logger))
}
