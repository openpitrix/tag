// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package constants

const (
	ServiceName = "tag"
)

const (
	TagName = "json"
)

const (
	DESC = "desc"
	ASC  = "asc"
)

const (
	DefaultOffset = uint32(0)
	DefaultLimit  = uint32(20)
)

const (
	DefaultSelectLimit = 200
)

const (
	StatusActive   = "active"
	StatusDisabled = "disabled"
	StatusDeleted  = "deleted"
)

var RecordStatuses = []string{
	StatusActive,
	StatusDisabled,
	StatusDeleted,
}
