package resources

import (
	"os"
	"path"
	"testing"
)

func TestConf(t *testing.T) {
	wd, _ := os.Getwd()
	fileName := path.Join(wd, "test.log")
	Conf(fileName)
	Logger.Debugf("hello debug.")
	Logger.Infof("hello info.")
}
