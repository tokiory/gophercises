package database

import "errors"

type DatabaseError error

var ErrPathNotFound DatabaseError = errors.New("url not found in database paths")
