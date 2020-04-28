// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/health_checker/redis/v2/redis.proto

package envoy_config_health_checker_redis_v2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = types.DynamicAny{}
)

// define the regex for a UUID once up-front
var _redis_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Redis with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Redis) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Key

	return nil
}

// RedisValidationError is the validation error returned by Redis.Validate if
// the designated constraints aren't met.
type RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RedisValidationError) ErrorName() string { return "RedisValidationError" }

// Error satisfies the builtin error interface
func (e RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRedis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RedisValidationError{}
