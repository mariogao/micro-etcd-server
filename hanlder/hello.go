package hanlder

import (
	"context"
	"fmt"

	proto "github.com/mariogao/micro-etcd-server/proto"
)

type MyhelloStruct struct {
}

func (m *MyhelloStruct) MyHello(ctx context.Context, req *proto.HelloReq, res *proto.HelloRes) error {
	fmt.Print("11111")
	res.ResName = "hello" + req.Name
	return nil

}
