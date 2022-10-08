package errors

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	NOERROR  string = "NOERROR"
	NXDOMAIN string = "NXDOMAIN"
	SERVFAIL string = "SERVFAIL"
	REFUSED  string = "REFUSED"
)

func HandleErrors(err error, context string) {
	logrus.Error(errors.Wrapf(err, "%s", context))
}
