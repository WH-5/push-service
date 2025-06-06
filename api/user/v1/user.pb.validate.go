// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/user/v1/user.proto

package v1

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

// Validate checks the field values on GetIdByUniqueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetIdByUniqueRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetIdByUniqueRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetIdByUniqueRequestMultiError, or nil if none found.
func (m *GetIdByUniqueRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetIdByUniqueRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetUniqueId()); l < 1 || l > 20 {
		err := GetIdByUniqueRequestValidationError{
			field:  "UniqueId",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if !_GetIdByUniqueRequest_UniqueId_Pattern.MatchString(m.GetUniqueId()) {
		err := GetIdByUniqueRequestValidationError{
			field:  "UniqueId",
			reason: "value does not match regex pattern \"^[a-zA-Z0-9_-]{1,20}$\"",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetIdByUniqueRequestMultiError(errors)
	}

	return nil
}

// GetIdByUniqueRequestMultiError is an error wrapping multiple validation
// errors returned by GetIdByUniqueRequest.ValidateAll() if the designated
// constraints aren't met.
type GetIdByUniqueRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetIdByUniqueRequestMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetIdByUniqueRequestMultiError) AllErrors() []error { return m }

// GetIdByUniqueRequestValidationError is the validation error returned by
// GetIdByUniqueRequest.Validate if the designated constraints aren't met.
type GetIdByUniqueRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIdByUniqueRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIdByUniqueRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIdByUniqueRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIdByUniqueRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIdByUniqueRequestValidationError) ErrorName() string {
	return "GetIdByUniqueRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetIdByUniqueRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIdByUniqueRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIdByUniqueRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIdByUniqueRequestValidationError{}

var _GetIdByUniqueRequest_UniqueId_Pattern = regexp.MustCompile("^[a-zA-Z0-9_-]{1,20}$")

// Validate checks the field values on GetIdByUniqueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetIdByUniqueReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetIdByUniqueReply with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetIdByUniqueReplyMultiError, or nil if none found.
func (m *GetIdByUniqueReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetIdByUniqueReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for UserId

	if len(errors) > 0 {
		return GetIdByUniqueReplyMultiError(errors)
	}

	return nil
}

// GetIdByUniqueReplyMultiError is an error wrapping multiple validation errors
// returned by GetIdByUniqueReply.ValidateAll() if the designated constraints
// aren't met.
type GetIdByUniqueReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetIdByUniqueReplyMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetIdByUniqueReplyMultiError) AllErrors() []error { return m }

// GetIdByUniqueReplyValidationError is the validation error returned by
// GetIdByUniqueReply.Validate if the designated constraints aren't met.
type GetIdByUniqueReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetIdByUniqueReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetIdByUniqueReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetIdByUniqueReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetIdByUniqueReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetIdByUniqueReplyValidationError) ErrorName() string {
	return "GetIdByUniqueReplyValidationError"
}

// Error satisfies the builtin error interface
func (e GetIdByUniqueReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetIdByUniqueReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetIdByUniqueReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetIdByUniqueReplyValidationError{}
