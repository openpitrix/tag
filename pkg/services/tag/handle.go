package tag

import (
	"context"
	"openpitrix.io/logger"
	"openpitrix.io/tag/pkg/db/mysql"
	"openpitrix.io/tag/pkg/models"
	"openpitrix.io/tag/pkg/pb"
	"openpitrix.io/tag/pkg/util/pbutil"
)

func (s *Server) CreateTag(ctx context.Context, req *pb.CreateTagRequest) (*pb.CreateTagResponse, error) {
	tag := models.NewTag(req.Key.Value, req.Value.Value, req.Creator.Value)
	_, err := models.CreateTag(tag)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTagResponse{
		TagId: pbutil.ToProtoString(tag.TagId),
	}, nil
}

func (s *Server) DescribeTags(ctx context.Context, req *pb.DescribeTagsRequest) (*pb.DescribeTagsResponse, error) {
	conds := map[string]interface{}{}
	if req.TagId != nil {
		conds["tag_id in (?)"] = req.TagId
	}
	if req.Key != nil {
		conds["key in (?)"] = req.Key
	}
	if req.Value != nil {
		conds["value in (?)"] = req.Value
	}
	if req.CreateTime != nil {
		conds["create_time > (?)"] = req.CreateTime
	}
	total, tags, err := models.DescribeTags(mysql.Conditions(conds))
	if err != nil {
		return nil, err
	}
	pbTags := models.TagSetToPbSet(tags)
	return &pb.DescribeTagsResponse{
		TotalCount: uint32(total),
		Tags:       pbTags,
	}, nil
}

func (s *Server) ModifyTag(ctx context.Context, req *pb.ModifyTagRequest) (*pb.ModifyTagResponse, error) {
	fields := map[string]interface{}{}
	if req.Key != nil {
		fields["key"] = req.Key.Value
	}
	if req.Value != nil {
		fields["value"] = req.Value.Value
	}
	err := models.UpdateTag(req.TagId.Value, fields)
	if err != nil {
		return nil, err
	}
	return &pb.ModifyTagResponse{
		TagId: req.TagId,
	}, nil
}

func (s *Server) DeleteTags(ctx context.Context, req *pb.DeleteTagsRequest) (*pb.DeleteTagsResponse, error) {
	if len(req.TagId) == 0 {
		logger.Errorf(ctx, "")
		return nil, nil
	}
	err := models.DeleteTags(req.TagId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteTagsResponse{
		TagId: req.TagId,
	}, nil
}

func (s *Server) AttachTags(ctx context.Context, req *pb.AttachTagsRequest) (*pb.AttachTagsResponse, error) {
	err := models.AttachTags(req.ResourceId, req.TagId)
	if err != nil {
		return nil, err
	}
	return &pb.AttachTagsResponse{
		ResourceId: req.ResourceId,
		TagId:      req.TagId,
	}, nil
}

func (s *Server) DetachTags(ctx context.Context, req *pb.DetachTagsRequest) (*pb.DetachTagsResponse, error) {
	err := models.DetachTags(req.ResourceId, req.TagId)
	if err != nil {
		return nil, err
	}
	return &pb.DetachTagsResponse{
		ResourceIds: req.ResourceId,
		TagIds:      req.TagId,
	}, nil
}
