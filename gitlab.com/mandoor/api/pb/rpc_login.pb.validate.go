// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: rpc_login.proto

package pb

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on LoginUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginUserRequestMultiError, or nil if none found.
func (m *LoginUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Email

	// no validation rules for Password

	if len(errors) > 0 {
		return LoginUserRequestMultiError(errors)
	}

	return nil
}

// LoginUserRequestMultiError is an error wrapping multiple validation errors
// returned by LoginUserRequest.ValidateAll() if the designated constraints
// aren't met.
type LoginUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUserRequestMultiError) AllErrors() []error { return m }

// LoginUserRequestValidationError is the validation error returned by
// LoginUserRequest.Validate if the designated constraints aren't met.
type LoginUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserRequestValidationError) ErrorName() string { return "LoginUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserRequestValidationError{}

// Validate checks the field values on LoginUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginUserResponseMultiError, or nil if none found.
func (m *LoginUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	// no validation rules for Msg

	// no validation rules for Id

	if len(errors) > 0 {
		return LoginUserResponseMultiError(errors)
	}

	return nil
}

// LoginUserResponseMultiError is an error wrapping multiple validation errors
// returned by LoginUserResponse.ValidateAll() if the designated constraints
// aren't met.
type LoginUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginUserResponseMultiError) AllErrors() []error { return m }

// LoginUserResponseValidationError is the validation error returned by
// LoginUserResponse.Validate if the designated constraints aren't met.
type LoginUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginUserResponseValidationError) ErrorName() string {
	return "LoginUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e LoginUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginUserResponseValidationError{}
