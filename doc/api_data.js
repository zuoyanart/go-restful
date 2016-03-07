define({ "api": [
  {
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "optional": false,
            "field": "varname1",
            "description": "<p>No type.</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "varname2",
            "description": "<p>With type.</p>"
          }
        ]
      }
    },
    "type": "",
    "url": "",
    "version": "0.0.0",
    "filename": "./doc/main.js",
    "group": "F__golang_src_pizzaCmsApi_doc_main_js",
    "groupTitle": "F__golang_src_pizzaCmsApi_doc_main_js",
    "name": ""
  },
  {
    "type": "get",
    "url": "/article/:id",
    "title": "获取文章信息",
    "name": "______by_path",
    "group": "article",
    "version": "1.0.0",
    "description": "<p>用于后台管理员获取文章信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/article/:id"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>文章id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "filename": "./controller/article.go",
    "groupTitle": "article"
  },
  {
    "type": "PUT",
    "url": "/article/:id",
    "title": "更新article信息",
    "name": "__article__",
    "group": "article",
    "version": "1.0.0",
    "description": "<p>后台管理员更新文章信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/article/:id"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>文章名</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>文章的id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/article.go",
    "groupTitle": "article"
  },
  {
    "type": "post",
    "url": "/article/",
    "title": "创建article信息1",
    "name": "__article__1",
    "group": "article",
    "version": "1.0.0",
    "description": "<p>创建文章信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/article/"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>title</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "brief",
            "description": "<p>brief</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "content",
            "description": "<p>content</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/article.go",
    "groupTitle": "article"
  },
  {
    "type": "post",
    "url": "/article/page",
    "title": "page article",
    "name": "page_article",
    "group": "article",
    "version": "1.0.0",
    "description": "<p>page article</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/article/page"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "kw",
            "description": "<p>关键字</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "cp",
            "description": "<p>cp</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "mp",
            "description": "<p>mp</p>"
          },
          {
            "group": "Parameter",
            "type": "nodeid",
            "optional": false,
            "field": "nodeid",
            "description": "<p>节点id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/article.go",
    "groupTitle": "article"
  },
  {
    "type": "get",
    "url": "/node/:id",
    "title": "获取节点信息",
    "name": "______by_path",
    "group": "node",
    "version": "1.0.0",
    "description": "<p>用于后台管理员获取节点信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/node/:id"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>节点id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "filename": "./controller/node.go",
    "groupTitle": "node"
  },
  {
    "type": "PUT",
    "url": "/node/:id",
    "title": "更新node信息",
    "name": "__node__",
    "group": "node",
    "version": "1.0.0",
    "description": "<p>后台管理员更新节点信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/node/:id"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "title",
            "description": "<p>节点名</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>节点的id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/node.go",
    "groupTitle": "node"
  },
  {
    "type": "post",
    "url": "/node/",
    "title": "创建node信息1",
    "name": "__node__1",
    "group": "node",
    "version": "1.0.0",
    "description": "<p>创建节点信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/node/"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "name",
            "description": "<p>name</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "brief",
            "description": "<p>brief</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "pid",
            "description": "<p>pid</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/node.go",
    "groupTitle": "node"
  },
  {
    "type": "post",
    "url": "/node/page",
    "title": "page node",
    "name": "page_node",
    "group": "node",
    "version": "1.0.0",
    "description": "<p>page node</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/node/page"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "pid",
            "optional": false,
            "field": "pid",
            "description": "<p>节点id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/node.go",
    "groupTitle": "node"
  },
  {
    "type": "get",
    "url": "/user?id=:id",
    "title": "获取用户信息",
    "name": "______",
    "group": "user",
    "version": "1.0.0",
    "description": "<p>用于后台管理员获取用户信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/user"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>用户id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息内容</p>"
          }
        ]
      }
    },
    "filename": "./controller/user.go",
    "groupTitle": "user"
  },
  {
    "type": "get",
    "url": "/user/:id",
    "title": "获取用户信息",
    "name": "______by_path",
    "group": "user",
    "version": "1.0.0",
    "description": "<p>用于后台管理员获取用户信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/user/:id"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>用户id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "filename": "./controller/user.go",
    "groupTitle": "user"
  },
  {
    "type": "PUT",
    "url": "/user/:id",
    "title": "更新user信息",
    "name": "__user__",
    "group": "user",
    "version": "1.0.0",
    "description": "<p>后台管理员更新用户信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/user/:id"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "id",
            "description": "<p>用户的id</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/user.go",
    "groupTitle": "user"
  },
  {
    "type": "post",
    "url": "/user/",
    "title": "创建user信息1",
    "name": "__user__1",
    "group": "user",
    "version": "1.0.0",
    "description": "<p>创建用户信息</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/user/"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "username",
            "description": "<p>用户名</p>"
          },
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "password",
            "description": "<p>密码</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/user.go",
    "groupTitle": "user"
  },
  {
    "type": "post",
    "url": "/user/page",
    "title": "page user",
    "name": "page_user",
    "group": "user",
    "version": "1.0.0",
    "description": "<p>page user</p>",
    "sampleRequest": [
      {
        "url": "http://192.168.1.134:3000/v1/user/page"
      }
    ],
    "parameter": {
      "fields": {
        "Parameter": [
          {
            "group": "Parameter",
            "type": "string",
            "optional": false,
            "field": "kw",
            "description": "<p>关键字</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "cp",
            "description": "<p>cp</p>"
          },
          {
            "group": "Parameter",
            "type": "int",
            "optional": false,
            "field": "mp",
            "description": "<p>mp</p>"
          }
        ]
      }
    },
    "success": {
      "fields": {
        "Success 200": [
          {
            "group": "Success 200",
            "type": "bool",
            "optional": false,
            "field": "state",
            "description": "<p>状态</p>"
          },
          {
            "group": "Success 200",
            "type": "String",
            "optional": false,
            "field": "msg",
            "description": "<p>消息</p>"
          }
        ]
      }
    },
    "permission": [
      {
        "name": "admin"
      }
    ],
    "filename": "./controller/user.go",
    "groupTitle": "user"
  }
] });
