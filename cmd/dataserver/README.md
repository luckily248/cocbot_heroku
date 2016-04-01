# DataServer


##数据服务器配置

Use bee run -downdoc=true -gendoc=true to run your API application and rebuild document automatically.
2. MongoDB环境配置
	1. 下载对应系统版本 已下载windows版本32位 64位安装包 注意版本风险
	2. 添加环境变量 
	3. bin同目录中建立/data/db
	4. 新建命令行窗口 运行后台服务 e.g. E:\MongoDB\Server\3.0\bin>mongod --dbpath E:\MongoDB\Server\3.0\data
	5. 新建命令行窗口 运行mongo -version 显示版本 则测试通过
3. 项目配置
	1. 配置conf/app.conf 
		2. 修改httpport 默认为8888,runmode 为dev时 需要配置gopath（参考下面章节 go开发环境配置） 而需要使用api页面 需要用dev 模式.
		3. 修改AdminHttpPort 默认为8038
	2. 配置conf/cityid.json 按格式修改需要采集的城市 已有默认数据
	3. 打开浏览器进入管理页面 查询服务器状态 e.g. http://127.0.0.1:8038/healthcheck
	4. 打开浏览器打开文档地址 查询API信息 e.g. http://127.0.0.1:8888/swagger/swagger-1/ (需要dev模式)


##开发文档
1. 系统结构
	1. 采集器
		1. 采用百度-易源作为天气数据采集
	2. 数据库 
		1. 采用mongoDB数据库 内存/文档数据库
	3. API web服务
		1. swagger 标准restful api接口
2. 源代码结构
	1. dataserver
		1. collector
			1. showapiCollector.go 采集器
		2. conf
			1. app.conf 服务器基本配置
			2. cityid.json 需采集的城市id model
		3. controllers
			1. ShowWeatherController.go 控制器 处理http请求根据注解生成 自动化api文档
		2. docs 自动化文档生成目录 自动生成不需修改
		3. models
			1. BaseDBmodel.go 基础类 处理基础数据库操作
			2. ShowWeatherRespModel.go 实体类 处理数据对接和传输
		3. routers
			1. comment.....自动生成 处理注解路由
			2. router.go  路由入口
		3. swagger 自动下载 不需修改
		4. tests
			1. routers 自动生成
			2. default_test.go 测试类 遵循tdd原则 
		3. dataserver/dataserver.exe 执行文件
		4. main.go 主程序入口
5. 后续开发 
	1. 数据库使用 
		1. C#开发指南 https://docs.mongodb.org/getting-started/csharp/client/
		2. 驱动包 CSharpDriver-2.0.1.zip
	2. GO开发环境配置
		1. 解压go主程序 添加目录下/bin到环境变量gopath中
		2. 新建命令行窗口 进入项目中 tests目录 运行 go test default_test.go -v 如下图则测试通过
	3. 项目打包发布
		1. 打开命令行窗口，进入项目目录，输入 bee pack 拷贝到目标服务器 解压即可

>     RUN TestApi-4
>     PASS: TestApi-4 (0.26s)
>     default_test.go:255: areaidtest:http://apis.baidu.com/showapi_open_bus/weather_showapi/address?areaid=101010100&needIndex=1&needMoreDay=1
>     default_test.go:281: addressResp:北京
>     default_test.go:282: addressResp index:白天光线弱不需要佩戴太阳镜
>     RUN TestUpsert-4
>     PASS: TestUpsert-4 (4.07s)
>     default_test.go:290: content _id:101280101
>     default_test.go:299: content get _id:37%.

2. 风险
	1. mongoDB数据库中 32位系统最高可用2G内存
	2. mongoDB数据库中 windows系统最高8T linux系统最高128T
	3. 目前数据库默认为127.0.0.1:27017
	4. 如果启动服务器出现 Detected unclean shutdown - E:\MongoDB\Server\3.0\data\mongod.lock is not empty.  说明数据库没正常关闭，需要删除路径中 mongod.lock文件 重新启动
