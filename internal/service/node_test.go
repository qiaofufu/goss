package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	v1 "master/api/node/v1"
	"testing"
)

func TestNodeService_Node(t *testing.T) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		t.Error(err)
	}

	client := v1.NewNodeClient(conn)

	res, err := client.ListNode(context.Background(), &v1.ListNodeRequest{
		Page: 0,
		Size: 10,
	})
	if err != nil {
		t.Error(err)
	}
	if res.Status != 0 {
		t.Error(res.Msg)
	}
	for _, node := range res.Nodes {
		resp, err := client.DeleteNode(context.Background(), &v1.DeleteNodeRequest{
			Id: node.Id,
		})
		if err != nil {
			t.Error(err)
		}
		if resp.Status != 0 {
			t.Error(resp.Msg)
		}
	}

	node, err := client.CreateNode(context.Background(), &v1.CreateNodeRequest{
		Addr: "192.168.1.101",
	})
	if err != nil {
		t.Error(err)
	}
	if node.Status != 0 {
		t.Error(node.Msg)
	}

	node2, err := client.CreateNode(context.Background(), &v1.CreateNodeRequest{
		Addr: "192.168..102",
	})
	if err != nil {
		t.Error(err)
	}
	if node2.Status != 0 {
		t.Error(node2.Msg)
	}

	t.Log(node, node2)
	listNode, err := client.ListNode(context.Background(), &v1.ListNodeRequest{
		Page: 0,
		Size: 10,
	})
	if err != nil {
		t.Error(err)
	}
	if len(listNode.Nodes) != 2 {
		t.Error("list node error")
	}
	t.Log(listNode)
}

func TestNodeService_GetNode(t *testing.T) {
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint("127.0.0.1:9000"))
	if err != nil {
		t.Error(err)
	}

	client := v1.NewNodeClient(conn)

	resp, err := client.GetNode(context.Background(), &v1.GetNodeRequest{
		Key:      10,
		Replicas: 3,
	})
	if err != nil {
		t.Error(err)
	}
	if resp.Status != 0 {
		t.Error(resp.Msg)
	}
	if len(resp.Nodes) != 3 {
		t.Error("get node error")
	}
	t.Log(resp)
}
