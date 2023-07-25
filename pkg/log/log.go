package pkg

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
)

func PrintErr(err error) {
	if err != nil {
		logrus.Error(err)
	}
}

func CheckAndExit(err error) {
	if err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}

func PrintErrWithPrefix(prefix string, err error) {
	if err != nil {
		logrus.Error(fmt.Errorf("%s: %s", prefix, err.Error()))
	}
}

func PrintErrWithPrefixAndExit(prefix string, err error) {
	if err != nil {
		logrus.Error(fmt.Errorf("%s: %s", prefix, err.Error()))
		os.Exit(-1)
	}
}

func ExitWithCode(message string, code int) {
	if strings.TrimSpace(message) == "" {
		message = "No message"
	}
	logrus.Println(message)
	os.Exit(code)
}
