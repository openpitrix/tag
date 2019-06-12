package models

import (
	"openpitrix.io/tag/pkg/global"
	"openpitrix.io/tag/pkg/pb"
	"openpitrix.io/tag/pkg/util/idutil"
	"openpitrix.io/tag/pkg/util/pbutil"
	"time"
)

type ResourceTag struct {
	ResourceTagId string
	TagId         string
	ResourceId    string
	CreateTime    time.Time
}

//table name
const (
	TableResourceTag = "resource_tag"
)

//ID Prefix
const (
	ResourceTagIdPrefix = "rt-"
)

func NewResourceTagId() string {
	return idutil.GetUuid(ResourceTagIdPrefix)
}

func NewResourceTag(resourceId, tagId string) *ResourceTag {
	resourceTagId := NewResourceTagId()
	return &ResourceTag{
		ResourceTagId: resourceTagId,
		ResourceId:    resourceId,
		TagId:         tagId,
		CreateTime:    time.Now(),
	}
}

func ResourceTagToPb(resourceTag *ResourceTag) *pb.ResourceTag {
	pbResourceTag := pb.ResourceTag{}
	pbResourceTag.ResourceTagId = pbutil.ToProtoString(resourceTag.ResourceTagId)
	pbResourceTag.ResourceId = pbutil.ToProtoString(resourceTag.ResourceId)
	pbResourceTag.TagId = pbutil.ToProtoString(resourceTag.TagId)
	pbResourceTag.CreateTime = pbutil.ToProtoTimestamp(resourceTag.CreateTime)
	return &pbResourceTag
}

func ResourceTagSetToPbSet(resourceTags []*ResourceTag) []*pb.ResourceTag {
	var pbResourceTags []*pb.ResourceTag
	for _, rt := range resourceTags {
		pbRt := ResourceTagToPb(rt)
		pbResourceTags = append(pbResourceTags, pbRt)
	}
	return pbResourceTags
}

func AttachTags(resourceIds, tagIds []string) error {
	db := global.GetInstance().GetDB()
	for _, rid := range resourceIds {
		for _, tid := range tagIds {
			resourceTag := NewResourceTag(rid, tid)
			err := db.Model(&ResourceTag{}).Create(resourceTag).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func DetachTags(resourceIds, tagIds []string) error {
	db := global.GetInstance().GetDB()
	for rid := range resourceIds {
		err := db.Delete(&Tag{}).Where("resource_id = (?) AND tag_id in (?)", rid, tagIds).Error
		if err != nil {
			return err
		}
	}
	return nil
}
