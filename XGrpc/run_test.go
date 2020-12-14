package XGrpc

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	testGrpcPb "goinfras/XGrpc/test_protof"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"net"
	"testing"
)

const TestGrpcPort = 8899

// 先定义一个测试的实现特定接口的服务端类型
type MyRpcTestServiceServer struct{}

func (*MyRpcTestServiceServer) GetUser(ctx context.Context, userId *testGrpcPb.TestGetUserRequestMessage) (*testGrpcPb.TestGetUserResponseMessage, error) {
	fmt.Println("Receive Request:", userId.Id)
	return &testGrpcPb.TestGetUserResponseMessage{
		Id:   userId.Id,
		Name: "test",
	}, nil
}

func (*MyRpcTestServiceServer) SetUser(ctx context.Context, user *testGrpcPb.TestSetUserRequestMessage) (*testGrpcPb.TestSetUserResponseMessage, error) {
	return &testGrpcPb.TestSetUserResponseMessage{Result: true}, nil
}

func (*MyRpcTestServiceServer) SendSome(empty *testGrpcPb.TestEmptyReqMessage, stream testGrpcPb.TestRpcService_SendSomeServer) error {
	for {
		err := stream.Send(&testGrpcPb.TestStreamMessage{Bit: 1})
		if err != nil {
			return err
		}
	}
}

// 测试运行GRPC服务端
func TestGrpcServer(t *testing.T) {
	Convey("TestGrpcServer", t, func() {
		// 先启动一个监听端口
		listen, err := net.Listen("tcp", fmt.Sprintf(":%d", TestGrpcPort))
		So(err, ShouldBeNil)

		// 创建grpc服务并注册
		grpcServer := grpc.NewServer()
		testGrpcPb.RegisterTestRpcServiceServer(grpcServer, &MyRpcTestServiceServer{})

		// 服务运行
		err = grpcServer.Serve(listen)
		So(err, ShouldBeNil)

	})
}

// 测试运行一个客户端
func TestGrpcClient(t *testing.T) {
	Convey("TestGrpcClient", t, func() {
		// 获得一个grpc服务端的连接
		conn, err := grpc.Dial(fmt.Sprintf(":%d", TestGrpcPort), grpc.WithInsecure())
		So(err, ShouldBeNil)
		defer conn.Close()

		// 创建客户端
		client := testGrpcPb.NewTestRpcServiceClient(conn)

		message, err := client.GetUser(context.TODO(), &testGrpcPb.TestGetUserRequestMessage{Id: 1})
		So(err, ShouldBeNil)
		Println("Response Message:", message)
	})
}
