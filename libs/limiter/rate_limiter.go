package limiter

import (
	"sync"
	"time"

	"go.uber.org/ratelimit"
	// "golang.org/x/time/rate"
	// "golang.org/x/sync/errgroup"
	// "golang.org/x/time/rate"
)

var (
	DefaultRateLimitTimeout = 2 * time.Second
)

// RateLimiter do rate limiter
type RateLimiter struct {
	mu      sync.RWMutex
	limiter ratelimit.Limiter
}

// rate is number of request per seconds
func NewRateLimiter(rateLimit int, burst int) *RateLimiter {
	// limiter := rate.NewLimiter(rate.Limit(rateLimit), burst)

	return &RateLimiter{
		mu:      sync.RWMutex{},
		limiter: ratelimit.New(rateLimit),
	}
}

// WaitN waits until enough resources are available for a request with given weight.
func (r *RateLimiter) WaitN(timeout time.Duration, weight int) error {
	r.limiter.Take()
	return nil
	// r.mu.Lock()
	// defer r.mu.Unlock()

	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	// errGr := errgroup.Group{}
	// errGr.Go(func() error {
	// 	if err := r.limiter.WaitN(ctx, weight); err != nil {
	// 		return err
	// 	}
	// 	return nil
	// },
	// )

	// return errGr.Wait()
}
