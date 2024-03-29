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
		w.WriteHeader(http.StatusOK)
	})
}

func attack(url string, rps int, duration time.Duration) {
	logger.Info("Attacking `%s` at %d requets per second during %s", url, rps, duration)
	defer logger.Info("Attack of `%s` is done!", url)

	attacker := vegeta.NewAttacker(vegeta.HTTP2(true), vegeta.KeepAlive(false), vegeta.Connections(5), vegeta.MaxConnections(100), vegeta.Workers(3), vegeta.MaxWorkers(10))
	defer attacker.Stop()

	for range attacker.Attack(
		vegeta.NewStaticTargeter(vegeta.Target{
			Method: "GET",
			URL:    url,
		}), vegeta.Rate{
			Freq: rps,
			Per:  time.Second,
		}, duration, "Boom!") {
		// no need to process metrics
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
		err = fmt.Errorf("parse request per second: %w", err)
		return
	}

	duration, err = time.ParseDuration(r.URL.Query().Get("duration"))
	if err != nil {
		err = fmt.Errorf("parse duration: %w", err)
		return
	}

	return
}
