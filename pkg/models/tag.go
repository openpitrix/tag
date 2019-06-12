// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package models

import (
	"openpitrix.io/tag/pkg/db/mysql"
	"openpitrix.io/tag/pkg/global"
	"openpitrix.io/tag/pkg/pb"
	"openpitrix.io/tag/pkg/util/idutil"
	"openpitrix.io/tag/pkg/util/pbutil"
	"time"
)

type Tag struct {
	TagId      string    `gorm:"column:tag_id"`
	Key        string    `gorm:"column:key"`
	Value      string    `gorm:"column:value"`
	Creator    string    `gorm:"column:creator"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}

//table name
const (
	TableTag = "tag"
)

//ID Prefix
const (
	TagIdPrefix = "t-"
)

func NewTagId() string {
	return idutil.GetUuid(TagIdPrefix)
}

func NewTag(key, value, creator string) *Tag {
	tagId := NewTagId()
	return &Tag{
		TagId:      tagId,
		Key:        key,
		Value:      value,
		Creator:    creator,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
}

func TagToPb(tag *Tag) *pb.Tag {
	pbTag := pb.Tag{}
	pbTag.TagId = pbutil.ToProtoString(tag.TagId)
	pbTag.Key = pbutil.ToProtoString(tag.Key)
	pbTag.Value = pbutil.ToProtoString(tag.Value)
	pbTag.CreateTime = pbutil.ToProtoTimestamp(tag.CreateTime)
	pbTag.UpdateTime = pbutil.ToProtoTimestamp(tag.UpdateTime)
	return &pbTag
}

func TagSetToPbSet(tags []*Tag) []*pb.Tag {
	var pbTags []*pb.Tag
	for _, t := range tags {
		pbT := TagToPb(t)
		pbTags = append(pbTags, pbT)
	}
	return pbTags
}

func CreateTag(tag *Tag) (*Tag, error) {
	db := global.GetInstance().GetDB()
	err := db.Model(&Tag{}).Create(tag).Error
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func DescribeTags(args ...mysql.Query) (int, []*Tag, error) {
	var total int
	tags := []*Tag{}
	option := &mysql.QueryOption{Offset: 0, Limit: -1}
	db := global.GetInstance().GetDB()

	if len(args) > 0 {
		for _, opt := range args {
			opt.Fn(option)
		}
		if option.Columns != "" {
			db = db.Select(option.Columns)
		}
		if len(option.Conditions) > 0 {
			keys, values := option.GetWhere()
			db = db.Where(keys, values...)
		}
		if len(option.Preload) > 0 {
			for _, preload := range option.Preload {
				db = db.Preload(preload)
			}
		}
		db.Group(option.GroupBy).Count(&total)
		db = db.Group(option.GroupBy).Order(option.OrderBy).Offset(option.Offset).Limit(option.Limit)
	}
	err := db.Find(&tags).Error
	return total, tags, err
}

func UpdateTag(tagId string, fields map[string]interface{}) error {
	db := global.GetInstance().GetDB()
	err := db.Model(&Tag{}).Where("tag_id = ?", tagId).Update(fields).Error
	return err
}

func DeleteTags(tagIds []string) error {
	db := global.GetInstance().GetDB()
	err := db.Delete(&Tag{}).Where("tag_id in (?)", tagIds).Error
	return err
}
