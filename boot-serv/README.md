# boot-serv
Quorum 网络引导服务

## 接口目录

[1. 联盟ID列表 ](#1-联盟ID列表)\
[2. 创建联盟信息 ](#2-创建联盟信息)\
[3. 删除联盟信息 ](#3-删除联盟信息)\
[4. 修改联盟信息 ](#4-修改联盟信息)\
[5. 查询联盟信息 ](#5-查询联盟信息)\
[6. 联盟是否存在 ](#6-联盟是否存在)\
[7. 节点ID列表 ](#7-节点ID列表)\
[8. 创建节点信息 ](#8-创建节点信息)\
[9. 删除节点信息 ](#9-删除节点信息)\
[10. 修改节点信息 ](#10-修改节点信息)\
[11. 查询节点信息 ](#11-查询节点信息)\
[12. 节点是否存在 ](#12-节点是否存在)

## 接口描述

### [1. 联盟ID列表](#接口目录)
GET /consortiums
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums",
    "Method": "GET", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> 无

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": ["consortium01", "consortium02"]
}
```
响应说明：
> 无

### [2. 创建联盟信息](#接口目录)
POST /consortiums
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums",
    "Method": "POST", 
    "Header": {
        "Content-Type": "application/json"
    },
    "Body": {
        "id":"consortium01",
        "detail":"consortium01 detail",
        "chainId":10000,
        "consensus":"raft",
        "difficulty":"0x0",
        "gasLimit":"0xE0000000",
        "alloc": {
            "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
              "balance": "1000000000000000000000000000"
            }
        }
    }
}
```
参数说明:
> id: 唯一标识，必选\
> detail: 联盟描述，必选\
> chainId: 联盟ID，必选\
> consensus: 共识方式，必选\
> difficulty: 难度，必选\
> gasLimit: 手续费用，必选\
> alloc: 预定账户金额，可选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "id": "consortium01",
        "detail": "consortium01 detail",
        "chainId": 10000,
        "consensus": "raft",
        "difficulty": "0x0",
        "gasLimit": "0xE0000000",
        "alloc": {
            "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
                "balance": "1000000000000000000000000000"
            }
        }
    }
}
```
响应说明：
> id: 唯一标识\
> detail: 联盟描述\
> chainId: 联盟ID\
> consensus: 共识方式\
> difficulty: 难度\
> gasLimit: 手续费用\
> alloc: 预定账户金额

### [3. 删除联盟信息](#接口目录)
DELETE /consortiums/{id}
```
Request: {
    "URL": "http://HOSTNAME/consortiums/consortium01",
    "Method": "DELETE", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "Id": "consortium01"
    }
}
```
响应说明：
> id: 唯一标识

### [4. 修改联盟信息](#接口目录) 
PUT /consortiums/{id}
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium01",
    "Method": "PUT", 
    "Header": {
        "Content-Type": "application/json"
    },
    "Body": {
        "detail":"consortium01 detail",
        "chainId":10000,
        "consensus":"raft",
        "difficulty":"0x0",
        "gasLimit":"0xE0000000",
        "alloc": {
            "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
                "balance": "1000000000000000000000000000"
            },
            "0xed9d02e382b34818e88b88a309c7fe71e65f4193": {
                "balance": "1000000000000000000000000000"
            }
        }
    }
}
```
参数说明:
> id: 唯一标识，必选\
> detail: 联盟描述，必选\
> chainId: 联盟ID，必选\
> consensus: 共识方式，必选\
> difficulty: 难度，必选\
> gasLimit: 手续费用，必选\
> alloc: 预定账户金额，可选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "consortium": "consortium02",
        "detail": "consortium02 detail",
        "chainId": 10000,
        "consensus": "raft",
        "difficulty": "0x0",
        "gasLimit": "0xE0000000",
        "alloc": {
            "0xed9d02e382b34818e88b88a309c7fe71e65f4193": {
                "balance": "1000000000000000000000000000"
            },
            "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
                "balance": "1000000000000000000000000000"
            }
        }
    }
}
```
响应说明：
> id: 唯一标识\
> detail: 联盟描述\
> chainId: 联盟ID\
> consensus: 共识方式\
> difficulty: 难度\
> gasLimit: 手续费用\
> alloc: 预定账户金额

### [5. 查询联盟信息](#接口目录)
GET /consortiums/{id}
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02",
    "Method": "GET", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "id": "consortium02",
        "detail": "consortium02 detail",
        "chainId": 10000,
        "consensus": "raft",
        "difficulty": "0x0",
        "gasLimit": "0xE0000000",
        "alloc": {
            "0xed9d02e382b34818e88b88a309c7fe71e65f4193": {
                "balance": "1000000000000000000000000000"
            },
            "0xed9d02e382b34818e88b88a309c7fe71e65f419d": {
                "balance": "1000000000000000000000000000"
            }
        }
    }
}
```
响应说明：
> id: 唯一标识\
> detail: 联盟描述\
> chainId: 联盟ID\
> consensus: 共识方式\
> difficulty: 难度\
> gasLimit: 手续费用\
> alloc: 预定账户金额

### [6. 联盟是否存在](#接口目录)
GET /consortiums/{id}/exist
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02/exist",
    "Method": "GET", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "id": "consortium02",
        "exist": true
    }
}
```
响应说明：
> id: 唯一标识\
> exist: 是否存在

### [7. 节点ID列表](#接口目录)
GET /consortiums/{id}/nodes
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02/nodes",
    "Method": "GET", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": ["publicKey01","publicKey02"]
}
```
响应说明：
> 无

### [8. 创建节点信息](#接口目录)
POST /consortiums/{id}/nodes/
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02/nodes",
    "Method": "POST", 
    "Header": {
        "Content-Type": "application/json"
    },
    "Body": {
        "publicKey":"0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
        "host":"127.0.0.1",
        "port":8999,
        "raftport":9888
    }
}
```
参数说明:
> id: 唯一标识，必选\
> publicKey: 节点公钥，必选\
> host: 节点服务地址，必选\
> port: 节点服务端口，必选\
> raftport: 节点raft共识端口，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "publicKey":"0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
        "host":"127.0.0.1",
        "port":8999,
        "raftport":9888
    }
}
```
响应说明：
> publicKey: 节点公钥\
> host: 节点服务地址\
> port: 节点服务端口\
> raftport: 节点raft共识端口

### [9. 删除节点信息](#接口目录)
DELETE /consortiums/{id}/nodes/{publicKey}
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium01/nodes/0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
    "Method": "DELETE", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选\
> publicKey: 节点公钥，必选

```
Response: {
    "code": 200,
    "message": "Success."
    "data": {
        "id": "consortium01",
        "publicKey":"0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b"
    }
}
```
响应说明：
> id: 唯一标识\
> publicKey: 节点公钥

### [10. 修改节点信息](#接口目录)
PUT /consortiums/{id}/nodes/{publicKey}
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02/nodes/0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
    "Method": "PUT", 
    "Header": {
        "Content-Type": "application/json"
    },
    "Body": {
        "host":"127.0.0.1",
        "port":8999,
        "raftport":9888
    }
}
```
参数说明:
> id: 唯一标识，必选\
> publicKey: 节点公钥，必选\
> host: 节点服务地址，必选\
> port: 节点服务端口，必选\
> raftport: 节点raft共识端口，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "publicKey":"0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
        "host":"127.0.0.1",
        "port":8999,
        "raftport":9888
    }
}
```
响应说明：
> publicKey: 节点公钥\
> host: 节点服务地址\
> port: 节点服务端口\
> raftport: 节点raft共识端口

### [11. 查询节点信息](#接口目录)
GET /consortiums/{id}/nodes/{publicKey}
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02/nodes/0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
    "Method": "GET", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选\
> publicKey: 节点公钥，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "publicKey": "0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
        "host": "127.0.0.2",
        "port": 8998,
        "raftport": 9889
    }
}
```
响应说明：
> publicKey: 节点公钥\
> host: 节点服务地址\
> port: 节点服务端口\
> raftport: 节点raft共识端口

### [12. 节点是否存在](#接口目录)
GET /consortiums/{id}/nodes/{publicKey}/exist
```
Request: {
    "URL": "http://{{HOSTNAME}}/consortiums/consortium02/nodes/0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b/exist",
    "Method": "GET", 
    "Header": {
        "Content-Type": "application/json"
    }
}
```
参数说明:
> id: 唯一标识，必选\
> publicKey: 节点公钥，必选

```
Response: {
    "code": 200,
    "message": "Success.",
    "data": {
        "id": "consortium02",
        "publicKey": "0x0427dc06acd0873ee3fc01bea8ad918a7059f8b7c7403ba964c02ec23560878abde3688f72cbb8fe7dfbc6312367d2578608e4c55402a52cb6cd053434fcc7b22b",
        "exist": true
    }
}
```
响应说明：
> id: 唯一标识\
> publicKey: 节点公钥\
> exist: 是否存在