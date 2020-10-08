package log

import "os"
import "github.com/sirupsen/logrus"

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
}
