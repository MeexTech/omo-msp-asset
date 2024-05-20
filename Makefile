.PHONY: proto
proto:
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset/common.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset/asset.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset/thumb.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset/folder.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset/score.proto
	protoc --proto_path=. --micro_out=. --go_out=. proto/asset/label.proto
