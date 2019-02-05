define({ "api": [  {    "type": "post",    "url": "/comment/add",    "title": "add comment to a post by id",    "name": "AddComment",    "group": "Comment",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "postId",            "description": "<p>target post id</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "content",            "description": "<p>comment contents</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\tpostId : \"bhc5sih5vl33qmk8p5t0\",\n\tcontent: \"nice writeup!\"\n}",          "type": "json"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>success message</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\" : \"success\",\n\t\"data\" : {\n\t\t\"message\": \"Add comment success!\"\n\t}\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/comment.go",    "groupTitle": "Comment"  },  {    "type": "post",    "url": "/post/create",    "title": "create a post",    "name": "CreatePost",    "group": "Post",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "content",            "description": "<p>post content</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\tcontent : \"my AWESOME post!\"\n}",          "type": "json"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>post data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"id\": \"bhc5sih5vl33qmk8p5t0\",\n\t\t\"content\": \"my AWESOME post!\",\n\t\t\"creatorId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\"createdAt\": \"2019-02-04T23:46:18.156289+08:00\",\n\t\t\"createdTime\": \"2019-02-04 23:46:18\",\n\t\t\"modifiedAt\": \"2019-02-04T23:46:18.156289+08:00\",\n\t\t\"modifiedTime\": \"2019-02-04 23:46:18\",\n\t\t\"comments\": []\n\t}\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/post.go",    "groupTitle": "Post"  },  {    "type": "delete",    "url": "/post/delete/{id}",    "title": "delete post by id",    "name": "DeletePost",    "group": "Post",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "id",            "description": "<p>post id</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n    id : \"bhc5vup5vl33qmk8p5u0\"\n}",          "type": "type"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>success data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"message\": \"Delete success!\"\n\t}\n}",          "type": "type"        }      ]    },    "filename": "./src/handlers/post.go",    "groupTitle": "Post"  },  {    "type": "get",    "url": "/post/all",    "title": "get post list in descending order",    "name": "GetAllPosts",    "group": "Post",    "version": "0.1.0",    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>post list</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": [\n\t\t{\n\t\t\t\"id\": \"bhc5vup5vl33qmk8p5u0\",\n\t\t\t\"content\": \"my GREAT GREAT post!\",\n\t\t\t\"creator\": \"john doe\",\n\t\t\t\"creatorId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\t\"createdAt\": \"2019-02-04T15:53:31.474Z\",\n\t\t\t\"createdTime\": \"2019-02-04 23:53:31\",\n\t\t\t\"modifiedAt\": \"2019-02-04T15:53:31.474Z\",\n\t\t\t\"modifiedTime\": \"2019-02-04 23:53:31\",\n\t\t\t\"comments\": []\n\t\t},\n\t\t{\n\t\t\t\"id\": \"bhc5sih5vl33qmk8p5t0\",\n\t\t\t\"content\": \"my AWESOME post!\",\n\t\t\t\"creator\": \"john doe\",\n\t\t\t\"creatorId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\t\"createdAt\": \"2019-02-04T15:46:18.156Z\",\n\t\t\t\"createdTime\": \"2019-02-04 23:46:18\",\n\t\t\t\"modifiedAt\": \"2019-02-04T15:46:18.156Z\",\n\t\t\t\"modifiedTime\": \"2019-02-04 23:46:18\",\n\t\t\t\"comments\": []\n\t\t}\n\t]\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/post.go",    "groupTitle": "Post"  },  {    "type": "get",    "url": "/post/{id}",    "title": "get a post by id",    "name": "GetPost",    "group": "Post",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "id",            "description": "<p>post id</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\tid : \"bhc5sih5vl33qmk8p5t0\"\n}",          "type": "json"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>post data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"id\": \"bhc5sih5vl33qmk8p5t0\",\n\t\t\"content\": \"my AWESOME post!\",\n\t\t\"creator\": \"john doe\",\n\t\t\"creatorId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\"createdAt\": \"2019-02-04T15:46:18.156Z\",\n\t\t\"createdTime\": \"2019-02-04 23:46:18\",\n\t\t\"modifiedAt\": \"2019-02-04T15:46:18.156Z\",\n\t\t\"modifiedTime\": \"2019-02-04 23:46:18\",\n\t\t\"comments\": []\n\t}\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/post.go",    "groupTitle": "Post"  },  {    "type": "patch",    "url": "/post/update",    "title": "update post content by id",    "name": "UpdatePist",    "group": "Post",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "id",            "description": "<p>post id</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\tid : \"bhc5vup5vl33qmk8p5u0\",\n\tcontent: \"my GREAT GREAT GREAT post\"\n}",          "type": "json"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>success data</p>"          }        ]      },      "examples": [        {          "title": "jSuccess-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"message\": \"Post update success!\"\n\t}\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/post.go",    "groupTitle": "Post"  },  {    "type": "get",    "url": "/user/all",    "title": "get user list",    "name": "GetAllUsers",    "group": "User",    "version": "0.1.0",    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "user",            "description": "<p>list</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": [\n\t\t{\n\t\t\t\"userId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\t\"username\": \"john doe\",\n\t\t\t\"createdAt\": \"2019-02-04T14:23:26.062Z\",\n\t\t\t\"createdTime\": \"2019-02-04 22:23:26\"\n\t\t}\n\t]\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/user.go",    "groupTitle": "User"  },  {    "type": "get",    "url": "/user/{id}",    "title": "get user by id",    "name": "GetUser",    "group": "User",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "id",            "description": "<p>user id</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\tid : \"bhc4lnh5vl33qmk8p5r0\"\n}",          "type": "json"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>user data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"userId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\"username\": \"john doe\",\n\t\t\"createdAt\": \"2019-02-04T14:23:26.062Z\",\n\t\t\"createdTime\": \"2019-02-04 22:23:26\"\n\t}\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/user.go",    "groupTitle": "User"  },  {    "type": "post",    "url": "/user/login",    "title": "user login",    "name": "Login",    "group": "User",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "username",            "description": "<p>username</p>"          },          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "password",            "description": "<p>password</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\tusername : \"john doe\",\n\tpassword: \"123123123\"\n}",          "type": "json"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "success",            "description": "<p>data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"userId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\"username\": \"john doe\",\n\t\t\"createdAt\": \"2019-02-04T14:23:26.062Z\",\n\t\t\"createdTime\": \"2019-02-04 22:23:26\"\n\t}\n}",          "type": "json"        }      ]    },    "filename": "./src/handlers/user.go",    "groupTitle": "User"  },  {    "type": "post",    "url": "/user/logout",    "title": "user logout",    "name": "Logout",    "group": "User",    "version": "0.1.0",    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "String",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "Object",            "optional": false,            "field": "data",            "description": "<p>success data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"message\": \"Logout success!\"\n\t}\n}",          "type": "Object"        }      ]    },    "filename": "./src/handlers/user.go",    "groupTitle": "User"  },  {    "type": "post",    "url": "/user/signup",    "title": "user sign up",    "name": "Signup",    "group": "User",    "version": "0.1.0",    "parameter": {      "fields": {        "Parameter": [          {            "group": "Parameter",            "type": "String",            "optional": false,            "field": "username",            "description": "<p>username</p>"          }        ]      },      "examples": [        {          "title": "Request-Example:",          "content": "{\n\t\"username\": \"john doe\",\n\t\"password\": \"123123123\"\n}",          "type": "type"        }      ]    },    "success": {      "fields": {        "200": [          {            "group": "200",            "type": "string",            "optional": false,            "field": "status",            "description": "<p>success status</p>"          },          {            "group": "200",            "type": "object",            "optional": false,            "field": "data",            "description": "<p>user data</p>"          }        ]      },      "examples": [        {          "title": "Success-Response:",          "content": "HTTP/1.1 200 OK\n{\n\t\"status\": \"success\",\n\t\"data\": {\n\t\t\"userId\": \"bhc4lnh5vl33qmk8p5r0\",\n\t\t\"username\": \"john doe\",\n\t\t\"createdAt\": \"2019-02-04T22:23:26.06252+08:00\",\n\t\t\"createdTime\": \"2019-02-04 22:23:26\"\n\t}\n}",          "type": "type"        }      ]    },    "filename": "./src/handlers/user.go",    "groupTitle": "User"  },  {    "success": {      "fields": {        "Success 200": [          {            "group": "Success 200",            "optional": false,            "field": "varname1",            "description": "<p>No type.</p>"          },          {            "group": "Success 200",            "type": "String",            "optional": false,            "field": "varname2",            "description": "<p>With type.</p>"          }        ]      }    },    "type": "",    "url": "",    "version": "0.0.0",    "filename": "./doc/main.js",    "group": "_Users_levblanc_projects_golang_restful_api_doc_main_js",    "groupTitle": "_Users_levblanc_projects_golang_restful_api_doc_main_js",    "name": ""  }] });
