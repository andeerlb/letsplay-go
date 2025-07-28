package client

import (
	"time"

	"github.com/go-resty/resty/v2"
)

func NewRestyClientWithRetry(timeout time.Duration, retryCount int) *resty.Client {
	client := resty.New().
		SetTimeout(timeout).
		SetRetryCount(retryCount).
		SetRetryWaitTime(2 * time.Second).
		SetRetryMaxWaitTime(5 * time.Second).
		AddRetryCondition(func(r *resty.Response, err error) bool {
			return err != nil || (r.StatusCode() >= 500 && r.StatusCode() != 501)
		})
	return client
}
