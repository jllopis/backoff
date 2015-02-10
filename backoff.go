package backoff

import "time"

// TODO. Implementar Exponential Backoff : http://javadoc.google-http-java-client.googlecode.com/hg/1.18.0-rc/com/google/api/client/util/ExponentialBackOff.html
// https://code.google.com/p/google-http-java-client/wiki/ExponentialBackoff
// http://en.wikipedia.org/wiki/Exponential_backoff
// http://docs.aws.amazon.com/general/latest/gr/api-retries.html

type Backoff interface {
	Duration(n int) time.Duration
}
