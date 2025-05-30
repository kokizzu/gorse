// Copyright 2024 gorse Project Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.0
// source: cache_store.proto

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	CacheStore_Ping_FullMethodName                = "/protocol.CacheStore/Ping"
	CacheStore_Get_FullMethodName                 = "/protocol.CacheStore/Get"
	CacheStore_Set_FullMethodName                 = "/protocol.CacheStore/Set"
	CacheStore_Delete_FullMethodName              = "/protocol.CacheStore/Delete"
	CacheStore_GetSet_FullMethodName              = "/protocol.CacheStore/GetSet"
	CacheStore_SetSet_FullMethodName              = "/protocol.CacheStore/SetSet"
	CacheStore_AddSet_FullMethodName              = "/protocol.CacheStore/AddSet"
	CacheStore_RemSet_FullMethodName              = "/protocol.CacheStore/RemSet"
	CacheStore_Push_FullMethodName                = "/protocol.CacheStore/Push"
	CacheStore_Pop_FullMethodName                 = "/protocol.CacheStore/Pop"
	CacheStore_Remain_FullMethodName              = "/protocol.CacheStore/Remain"
	CacheStore_AddScores_FullMethodName           = "/protocol.CacheStore/AddScores"
	CacheStore_SearchScores_FullMethodName        = "/protocol.CacheStore/SearchScores"
	CacheStore_DeleteScores_FullMethodName        = "/protocol.CacheStore/DeleteScores"
	CacheStore_UpdateScores_FullMethodName        = "/protocol.CacheStore/UpdateScores"
	CacheStore_ScanScores_FullMethodName          = "/protocol.CacheStore/ScanScores"
	CacheStore_AddTimeSeriesPoints_FullMethodName = "/protocol.CacheStore/AddTimeSeriesPoints"
	CacheStore_GetTimeSeriesPoints_FullMethodName = "/protocol.CacheStore/GetTimeSeriesPoints"
)

// CacheStoreClient is the client API for CacheStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CacheStoreClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	GetSet(ctx context.Context, in *GetSetRequest, opts ...grpc.CallOption) (*GetSetResponse, error)
	SetSet(ctx context.Context, in *SetSetRequest, opts ...grpc.CallOption) (*SetSetResponse, error)
	AddSet(ctx context.Context, in *AddSetRequest, opts ...grpc.CallOption) (*AddSetResponse, error)
	RemSet(ctx context.Context, in *RemSetRequest, opts ...grpc.CallOption) (*RemSetResponse, error)
	Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error)
	Pop(ctx context.Context, in *PopRequest, opts ...grpc.CallOption) (*PopResponse, error)
	Remain(ctx context.Context, in *RemainRequest, opts ...grpc.CallOption) (*RemainResponse, error)
	AddScores(ctx context.Context, in *AddScoresRequest, opts ...grpc.CallOption) (*AddScoresResponse, error)
	SearchScores(ctx context.Context, in *SearchScoresRequest, opts ...grpc.CallOption) (*SearchScoresResponse, error)
	DeleteScores(ctx context.Context, in *DeleteScoresRequest, opts ...grpc.CallOption) (*DeleteScoresResponse, error)
	UpdateScores(ctx context.Context, in *UpdateScoresRequest, opts ...grpc.CallOption) (*UpdateScoresResponse, error)
	ScanScores(ctx context.Context, in *ScanScoresRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ScanScoresResponse], error)
	AddTimeSeriesPoints(ctx context.Context, in *AddTimeSeriesPointsRequest, opts ...grpc.CallOption) (*AddTimeSeriesPointsResponse, error)
	GetTimeSeriesPoints(ctx context.Context, in *GetTimeSeriesPointsRequest, opts ...grpc.CallOption) (*GetTimeSeriesPointsResponse, error)
}

type cacheStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewCacheStoreClient(cc grpc.ClientConnInterface) CacheStoreClient {
	return &cacheStoreClient{cc}
}

func (c *cacheStoreClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, CacheStore_Ping_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, CacheStore_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*SetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetResponse)
	err := c.cc.Invoke(ctx, CacheStore_Set_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, CacheStore_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) GetSet(ctx context.Context, in *GetSetRequest, opts ...grpc.CallOption) (*GetSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSetResponse)
	err := c.cc.Invoke(ctx, CacheStore_GetSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) SetSet(ctx context.Context, in *SetSetRequest, opts ...grpc.CallOption) (*SetSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SetSetResponse)
	err := c.cc.Invoke(ctx, CacheStore_SetSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) AddSet(ctx context.Context, in *AddSetRequest, opts ...grpc.CallOption) (*AddSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddSetResponse)
	err := c.cc.Invoke(ctx, CacheStore_AddSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) RemSet(ctx context.Context, in *RemSetRequest, opts ...grpc.CallOption) (*RemSetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemSetResponse)
	err := c.cc.Invoke(ctx, CacheStore_RemSet_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) Push(ctx context.Context, in *PushRequest, opts ...grpc.CallOption) (*PushResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PushResponse)
	err := c.cc.Invoke(ctx, CacheStore_Push_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) Pop(ctx context.Context, in *PopRequest, opts ...grpc.CallOption) (*PopResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PopResponse)
	err := c.cc.Invoke(ctx, CacheStore_Pop_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) Remain(ctx context.Context, in *RemainRequest, opts ...grpc.CallOption) (*RemainResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemainResponse)
	err := c.cc.Invoke(ctx, CacheStore_Remain_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) AddScores(ctx context.Context, in *AddScoresRequest, opts ...grpc.CallOption) (*AddScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddScoresResponse)
	err := c.cc.Invoke(ctx, CacheStore_AddScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) SearchScores(ctx context.Context, in *SearchScoresRequest, opts ...grpc.CallOption) (*SearchScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchScoresResponse)
	err := c.cc.Invoke(ctx, CacheStore_SearchScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) DeleteScores(ctx context.Context, in *DeleteScoresRequest, opts ...grpc.CallOption) (*DeleteScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteScoresResponse)
	err := c.cc.Invoke(ctx, CacheStore_DeleteScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) UpdateScores(ctx context.Context, in *UpdateScoresRequest, opts ...grpc.CallOption) (*UpdateScoresResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateScoresResponse)
	err := c.cc.Invoke(ctx, CacheStore_UpdateScores_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) ScanScores(ctx context.Context, in *ScanScoresRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[ScanScoresResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &CacheStore_ServiceDesc.Streams[0], CacheStore_ScanScores_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ScanScoresRequest, ScanScoresResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CacheStore_ScanScoresClient = grpc.ServerStreamingClient[ScanScoresResponse]

func (c *cacheStoreClient) AddTimeSeriesPoints(ctx context.Context, in *AddTimeSeriesPointsRequest, opts ...grpc.CallOption) (*AddTimeSeriesPointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddTimeSeriesPointsResponse)
	err := c.cc.Invoke(ctx, CacheStore_AddTimeSeriesPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheStoreClient) GetTimeSeriesPoints(ctx context.Context, in *GetTimeSeriesPointsRequest, opts ...grpc.CallOption) (*GetTimeSeriesPointsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTimeSeriesPointsResponse)
	err := c.cc.Invoke(ctx, CacheStore_GetTimeSeriesPoints_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheStoreServer is the server API for CacheStore service.
// All implementations must embed UnimplementedCacheStoreServer
// for forward compatibility.
type CacheStoreServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Set(context.Context, *SetRequest) (*SetResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	GetSet(context.Context, *GetSetRequest) (*GetSetResponse, error)
	SetSet(context.Context, *SetSetRequest) (*SetSetResponse, error)
	AddSet(context.Context, *AddSetRequest) (*AddSetResponse, error)
	RemSet(context.Context, *RemSetRequest) (*RemSetResponse, error)
	Push(context.Context, *PushRequest) (*PushResponse, error)
	Pop(context.Context, *PopRequest) (*PopResponse, error)
	Remain(context.Context, *RemainRequest) (*RemainResponse, error)
	AddScores(context.Context, *AddScoresRequest) (*AddScoresResponse, error)
	SearchScores(context.Context, *SearchScoresRequest) (*SearchScoresResponse, error)
	DeleteScores(context.Context, *DeleteScoresRequest) (*DeleteScoresResponse, error)
	UpdateScores(context.Context, *UpdateScoresRequest) (*UpdateScoresResponse, error)
	ScanScores(*ScanScoresRequest, grpc.ServerStreamingServer[ScanScoresResponse]) error
	AddTimeSeriesPoints(context.Context, *AddTimeSeriesPointsRequest) (*AddTimeSeriesPointsResponse, error)
	GetTimeSeriesPoints(context.Context, *GetTimeSeriesPointsRequest) (*GetTimeSeriesPointsResponse, error)
	mustEmbedUnimplementedCacheStoreServer()
}

// UnimplementedCacheStoreServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedCacheStoreServer struct{}

func (UnimplementedCacheStoreServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCacheStoreServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCacheStoreServer) Set(context.Context, *SetRequest) (*SetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedCacheStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCacheStoreServer) GetSet(context.Context, *GetSetRequest) (*GetSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSet not implemented")
}
func (UnimplementedCacheStoreServer) SetSet(context.Context, *SetSetRequest) (*SetSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetSet not implemented")
}
func (UnimplementedCacheStoreServer) AddSet(context.Context, *AddSetRequest) (*AddSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSet not implemented")
}
func (UnimplementedCacheStoreServer) RemSet(context.Context, *RemSetRequest) (*RemSetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemSet not implemented")
}
func (UnimplementedCacheStoreServer) Push(context.Context, *PushRequest) (*PushResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Push not implemented")
}
func (UnimplementedCacheStoreServer) Pop(context.Context, *PopRequest) (*PopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pop not implemented")
}
func (UnimplementedCacheStoreServer) Remain(context.Context, *RemainRequest) (*RemainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remain not implemented")
}
func (UnimplementedCacheStoreServer) AddScores(context.Context, *AddScoresRequest) (*AddScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddScores not implemented")
}
func (UnimplementedCacheStoreServer) SearchScores(context.Context, *SearchScoresRequest) (*SearchScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchScores not implemented")
}
func (UnimplementedCacheStoreServer) DeleteScores(context.Context, *DeleteScoresRequest) (*DeleteScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteScores not implemented")
}
func (UnimplementedCacheStoreServer) UpdateScores(context.Context, *UpdateScoresRequest) (*UpdateScoresResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateScores not implemented")
}
func (UnimplementedCacheStoreServer) ScanScores(*ScanScoresRequest, grpc.ServerStreamingServer[ScanScoresResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ScanScores not implemented")
}
func (UnimplementedCacheStoreServer) AddTimeSeriesPoints(context.Context, *AddTimeSeriesPointsRequest) (*AddTimeSeriesPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTimeSeriesPoints not implemented")
}
func (UnimplementedCacheStoreServer) GetTimeSeriesPoints(context.Context, *GetTimeSeriesPointsRequest) (*GetTimeSeriesPointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTimeSeriesPoints not implemented")
}
func (UnimplementedCacheStoreServer) mustEmbedUnimplementedCacheStoreServer() {}
func (UnimplementedCacheStoreServer) testEmbeddedByValue()                    {}

// UnsafeCacheStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CacheStoreServer will
// result in compilation errors.
type UnsafeCacheStoreServer interface {
	mustEmbedUnimplementedCacheStoreServer()
}

func RegisterCacheStoreServer(s grpc.ServiceRegistrar, srv CacheStoreServer) {
	// If the following call pancis, it indicates UnimplementedCacheStoreServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&CacheStore_ServiceDesc, srv)
}

func _CacheStore_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Set_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_GetSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).GetSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_GetSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).GetSet(ctx, req.(*GetSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_SetSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).SetSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_SetSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).SetSet(ctx, req.(*SetSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_AddSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).AddSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_AddSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).AddSet(ctx, req.(*AddSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_RemSet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemSetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).RemSet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_RemSet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).RemSet(ctx, req.(*RemSetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Push_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Push(ctx, req.(*PushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_Pop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Pop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Pop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Pop(ctx, req.(*PopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_Remain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).Remain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_Remain_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).Remain(ctx, req.(*RemainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_AddScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).AddScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_AddScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).AddScores(ctx, req.(*AddScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_SearchScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).SearchScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_SearchScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).SearchScores(ctx, req.(*SearchScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_DeleteScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).DeleteScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_DeleteScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).DeleteScores(ctx, req.(*DeleteScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_UpdateScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateScoresRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).UpdateScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_UpdateScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).UpdateScores(ctx, req.(*UpdateScoresRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_ScanScores_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScanScoresRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CacheStoreServer).ScanScores(m, &grpc.GenericServerStream[ScanScoresRequest, ScanScoresResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type CacheStore_ScanScoresServer = grpc.ServerStreamingServer[ScanScoresResponse]

func _CacheStore_AddTimeSeriesPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTimeSeriesPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).AddTimeSeriesPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_AddTimeSeriesPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).AddTimeSeriesPoints(ctx, req.(*AddTimeSeriesPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CacheStore_GetTimeSeriesPoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTimeSeriesPointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheStoreServer).GetTimeSeriesPoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CacheStore_GetTimeSeriesPoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheStoreServer).GetTimeSeriesPoints(ctx, req.(*GetTimeSeriesPointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CacheStore_ServiceDesc is the grpc.ServiceDesc for CacheStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CacheStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol.CacheStore",
	HandlerType: (*CacheStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CacheStore_Ping_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CacheStore_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _CacheStore_Set_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CacheStore_Delete_Handler,
		},
		{
			MethodName: "GetSet",
			Handler:    _CacheStore_GetSet_Handler,
		},
		{
			MethodName: "SetSet",
			Handler:    _CacheStore_SetSet_Handler,
		},
		{
			MethodName: "AddSet",
			Handler:    _CacheStore_AddSet_Handler,
		},
		{
			MethodName: "RemSet",
			Handler:    _CacheStore_RemSet_Handler,
		},
		{
			MethodName: "Push",
			Handler:    _CacheStore_Push_Handler,
		},
		{
			MethodName: "Pop",
			Handler:    _CacheStore_Pop_Handler,
		},
		{
			MethodName: "Remain",
			Handler:    _CacheStore_Remain_Handler,
		},
		{
			MethodName: "AddScores",
			Handler:    _CacheStore_AddScores_Handler,
		},
		{
			MethodName: "SearchScores",
			Handler:    _CacheStore_SearchScores_Handler,
		},
		{
			MethodName: "DeleteScores",
			Handler:    _CacheStore_DeleteScores_Handler,
		},
		{
			MethodName: "UpdateScores",
			Handler:    _CacheStore_UpdateScores_Handler,
		},
		{
			MethodName: "AddTimeSeriesPoints",
			Handler:    _CacheStore_AddTimeSeriesPoints_Handler,
		},
		{
			MethodName: "GetTimeSeriesPoints",
			Handler:    _CacheStore_GetTimeSeriesPoints_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ScanScores",
			Handler:       _CacheStore_ScanScores_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "cache_store.proto",
}
