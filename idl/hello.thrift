namespace go hello

struct HelloReq {
    1: string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct HelloResp {
    1: string RespBody;
}

struct WorldReq {
    1: string Name (api.query="name"); // 添加 api 注解为方便进行参数绑定
}

struct WorldResp {
    1: string RespBody;
}


service HelloService {
    HelloResp Hello(1: HelloReq request) (api.get="/hello");
    WorldResp World(1: WorldReq request) (api.get="/world");
}