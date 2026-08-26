package trusted

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"

	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httperr"
)

// WithOnlyTrustedSubnet пропускает дальше только если ip юзера находится в подсети subnet
//
// subnet указан в формате CIDR, если он пустой, то middleware никого не пускает дальше
//
// В header указан заголовок, в котором искать ip пользователя
func WithOnlyTrustedSubnet(header string, subnet string) (func(http.Handler) http.Handler, error) {
	var trustedSubnet *net.IPNet

	subnet = strings.TrimSpace(subnet)
	if subnet != "" {
		var err error
		_, trustedSubnet, err = net.ParseCIDR(subnet)

		if err != nil {
			return nil, err
		}
	}

	return func(h http.Handler) http.Handler {
		return httperr.Adapt(func(w http.ResponseWriter, r *http.Request) error {
			if trustedSubnet == nil {
				w.WriteHeader(http.StatusForbidden)
				return errors.New("TrustedSubnet Middleware: no trusted subnet")
			}

			ip := r.Header.Get(header)

			if !trustedSubnet.Contains(net.ParseIP(ip)) {
				w.WriteHeader(http.StatusForbidden)
				return fmt.Errorf("TrustedSubnet Middleware: ip %s is not in trusted subnet %s", ip, trustedSubnet.String())
			}

			h.ServeHTTP(w, r)
			return nil
		})
	}, nil
}
