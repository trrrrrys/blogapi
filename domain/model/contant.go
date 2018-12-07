package model

import "time"

type Content struct {
	ID            int
	Title         string
	Author        *User
	Tags          []*Tag
	PublishedDate uint
	Body          string
}

// CalcPublished Calculate published date.
func (c *Content) CalcPublished() uint {
	now := uint(time.Now().Unix())
	return uint((now - c.PublishedDate) / 86400)
}
