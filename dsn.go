package godsn

import (
	"net/url"
	"strconv"
	"strings"
	"time"
)

type DSN struct {
	url *url.URL
	*DSNValues
}

// parses dsn string and returns DSN instance
func Parse(dsn string) (*DSN, error) {
	parsed, err := url.Parse(dsn)
	if err != nil {
		return nil, err
	}
	d := DSN{
		parsed,
		&DSNValues{parsed.Query()},
	}
	return &d, nil
}

// Parses query and returns dsn values
func ParseQuery(query string) (*DSNValues, error) {
	parsed, err := url.ParseQuery(query)
	if err != nil {
		return nil, err
	}
	return &DSNValues{parsed}, nil
}

// returns DSNValues from url.Values
func NewValues(query url.Values) (*DSNValues, error) {
	return &DSNValues{query}, nil
}

// return Host
func (d *DSN) Host() string {
	return d.url.Host
}

// return Scheme
func (d *DSN) Scheme() string {
	return d.url.Scheme
}

// returns path
func (d *DSN) Path() string {
	return d.url.Path
}

// returns user
func (d *DSN) User() *url.Userinfo {
	return d.url.User
}

// DSN Values
type DSNValues struct {
	url.Values
}

// returns int value
func (d *DSNValues) GetInt(param string, def int) int {
	value := d.Get(param)
	if i, err := strconv.Atoi(value); err == nil {
		return i
	} else {
		return def
	}
}

// returns string value
func (d *DSNValues) GetString(param string, def string) string {
	value := d.Get(param)
	if value == "" {
		return def
	} else {
		return value
	}
}

// returns string value
func (d *DSNValues) GetBool(param string, def bool) bool {
	value := strings.ToLower(d.Get(param))
	if value == "true" || value == "1" {
		return true
	} else if value == "0" || value == "false" {
		return false
	} else {
		return def
	}
}

// returns string value
func (d *DSNValues) GetSeconds(param string, def time.Duration) time.Duration {
	if i, err := strconv.Atoi(d.Get(param)); err == nil {
		return time.Duration(i) * time.Second
	} else {
		return def
	}
}
