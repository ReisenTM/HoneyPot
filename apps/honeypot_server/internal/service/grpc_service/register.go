package grpc_service

import (
	"context"
	"errors"
	"github.com/sirupsen/logrus"
	"honeypot_server/internal/global"
	"honeypot_server/internal/models"
	"honeypot_server/internal/rpc/node_rpc"
)

func (NodeService) Register(ctx context.Context, request *node_rpc.RegisterRequest) (pd *node_rpc.BaseResponse, err error) {
	logrus.Info("Node Register")
	pd = new(node_rpc.BaseResponse)
	// 节点不存在，需要创建
	uid := request.NodeUid
	var model models.NodeModel
	err1 := global.DB.Take(&model, "uid = ?", uid).Error
	if err1 != nil {
		// 创建节点
		model = models.NodeModel{
			Title:  request.SystemInfo.HostName,
			Uid:    uid,
			IP:     request.Ip,
			Mac:    request.Mac,
			Status: 1,
			SystemInfo: models.NodeSystemInfo{
				NodeVersion:         request.Version,
				NodeCommit:          request.Commit,
				HostName:            request.SystemInfo.HostName,
				DistributionVersion: request.SystemInfo.DistributionVersion,
				CoreVersion:         request.SystemInfo.CoreVersion,
				SystemType:          request.SystemInfo.SystemType,
				StartTime:           request.SystemInfo.StartTime,
			},
		}
		err1 = global.DB.Create(&model).Error
		if err1 != nil {
			logrus.Errorf("节点创建失败 %s", err)
			return nil, errors.New("节点创建失败")
		}
		// 创建网卡记录
		//var networkList []models.NodeNetworkModel
		//for _, message := range request.NetworkList {
		//	networkList = append(networkList, models.NodeNetworkModel{
		//		NodeID:  model.ID,
		//		Network: message.Network,
		//		IP:      message.Ip,
		//		Mask:    int8(message.Mask),
		//		Status:  2,
		//	})
		//}
		//if len(networkList) > 0 {
		//	err = global.DB.Create(&networkList).Error
		//	if err != nil {
		//		logrus.Errorf("节点网卡保存失败 %s", err)
		//		return nil, errors.New("节点网卡保存失败")
		//	}
		//}
	}
	if model.Status != 1 {
		// 改状态
		global.DB.Model(&model).Update("status", 1)
	}
	return &node_rpc.BaseResponse{
		Code: 0,
		Msg:  "Register success",
	}, nil
}
