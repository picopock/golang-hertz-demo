install:
	go install -v github.com/go-delve/delve/cmd/dlv@latest
	go install github.com/cloudwego/hertz/cmd/hz@latest

init_swagger:
	swag init --parseDependency --parseInternal --parseDepth 5 --instanceName "swagger"

init_api:
	hz new --model_dir gen/model --idl idl/hello.thrift --customize_layout=template/layout.yaml --customize_package=template/package.yaml

update_api:
	hz update --model_dir gen/model --idl idl/hello.thrift  --customize_package=template/package.yaml
