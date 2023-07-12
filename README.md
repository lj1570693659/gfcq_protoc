# gfcq_protoc
重庆赣锋项目管理部后端代码仓库

## 生成PB文件
*  protoc -I .\protobuf\ --go_out=.\pb\ .\protobuf\user\v1\user.proto 
*  protoc --go_out=.\pb\ .\protobuf\pbentity\user.proto 

*  protoc --proto_path=.\ --go_out=plugins=grpc:user user\v1\user.proto
*  protoc --proto_path=.\ --go_out=plugins=grpc:pbentity pbentity\user.proto
*  protoc --proto_path=.\ --go_out=plugins=grpc:common common\v1\department.proto
*  protoc --proto_path=.\ --go_out=plugins=grpc:config config\product\v1\level_assess.proto
  