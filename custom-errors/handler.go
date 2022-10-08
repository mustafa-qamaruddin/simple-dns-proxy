package custom_errors

import (
	"github.com/mustafa-qamaruddin/simple-dns-proxy/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

const (
	NOERROR  string = "NOERROR"
	NXDOMAIN string = "NXDOMAIN"
	SERVFAIL string = "SERVFAIL"
	REFUSED  string = "REFUSED"
)

// DNSResponseCode provides response codes for question answers.
type DNSResponseCode uint8

// DNSResponseCode known values.
const (
	DNSResponseCodeNoErr     DNSResponseCode = 0  // No error
	DNSResponseCodeFormErr   DNSResponseCode = 1  // Format Error                       [RFC1035]
	DNSResponseCodeServFail  DNSResponseCode = 2  // Server Failure                     [RFC1035]
	DNSResponseCodeNXDomain  DNSResponseCode = 3  // Non-Existent Domain                [RFC1035]
	DNSResponseCodeNotImp    DNSResponseCode = 4  // Not Implemented                    [RFC1035]
	DNSResponseCodeRefused   DNSResponseCode = 5  // Query Refused                      [RFC1035]
	DNSResponseCodeYXDomain  DNSResponseCode = 6  // Name Exists when it should not     [RFC2136]
	DNSResponseCodeYXRRSet   DNSResponseCode = 7  // RR Set Exists when it should not   [RFC2136]
	DNSResponseCodeNXRRSet   DNSResponseCode = 8  // RR Set that should exist does not  [RFC2136]
	DNSResponseCodeNotAuth   DNSResponseCode = 9  // Server Not Authoritative for zone  [RFC2136]
	DNSResponseCodeNotZone   DNSResponseCode = 10 // Name not contained in zone         [RFC2136]
	DNSResponseCodeBadVers   DNSResponseCode = 16 // Bad OPT Version                    [RFC2671]
	DNSResponseCodeBadSig    DNSResponseCode = 16 // TSIG Signature Failure             [RFC2845]
	DNSResponseCodeBadKey    DNSResponseCode = 17 // Key not recognized                 [RFC2845]
	DNSResponseCodeBadTime   DNSResponseCode = 18 // Signature out of time window       [RFC2845]
	DNSResponseCodeBadMode   DNSResponseCode = 19 // Bad TKEY Mode                      [RFC2930]
	DNSResponseCodeBadName   DNSResponseCode = 20 // Duplicate key name                 [RFC2930]
	DNSResponseCodeBadAlg    DNSResponseCode = 21 // Algorithm not supported            [RFC2930]
	DNSResponseCodeBadTruc   DNSResponseCode = 22 // Bad Truncation                     [RFC4635]
	DNSResponseCodeBadCookie DNSResponseCode = 23 // Bad/missing Server Cookie          [RFC7873]
)

func HandleErrors(err error, context common.Error) {
	logrus.Error(errors.Wrapf(err, "%v", context))
}
