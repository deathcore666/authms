package service

import "github.com/go-kit/kit/log"
import "time"

type loggingMiddleware struct {
	logger log.Logger
	AuthService
}

func (mw loggingMiddleware) Login(username, password string) (output bool, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "login",
			"input", username+" "+password,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.AuthService.Login(username, password)
	return
}

func (mw loggingMiddleware) Register(username, password string) (output int, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "login",
			"input", username+" "+password,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	output, err = mw.AuthService.Register(username, password)
	return
}
