package vegetaas

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ViBiOh/httputils/v4/pkg/httperror"
	"github.com/ViBiOh/httputils/v4/pkg/logger"
	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// Handler for request. Should be use with net/http
func Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		url, rps, duration, err := parseQuery(r)
		if err != nil {
			httperror.BadRequest(w, err)
			return
		}

		go attack(url, rps, duration)
	})
}

func attack(url string, rps int, duration time.Duration) {
	logger.Info("Attacking `%s` at %d requets per second during %s", url, rps, duration)
	defer logger.Info("Attack of `%s` is done!", url)

	for range vegeta.NewAttacker().Attack(
		vegeta.NewStaticTargeter(vegeta.Target{
			Method: "GET",
			URL:    url,
		}), vegeta.Rate{
			Freq: rps,
			Per:  time.Second,
		}, duration, "Boom!") {
	}
}

func parseQuery(r *http.Request) (url string, rps int, duration time.Duration, err error) {
	url = strings.TrimSpace(r.URL.Query().Get("url"))
	if len(url) == 0 {
		err = errors.New("url is required")
		return
	}

	rps, err = strconv.Atoi(r.URL.Query().Get("rps"))
	if err != nil {
		err = fmt.Errorf("unable to parse request per second: %s", err)
		return
	}

	duration, err = time.ParseDuration(r.URL.Query().Get("duration"))
	if err != nil {
		err = fmt.Errorf("unable to parse duration: %s", err)
		return
	}

	return
}
