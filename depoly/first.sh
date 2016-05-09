#步骤：
# 检查mocha是否已经安装，没有安装则安装mocha
# git clone项目所有源文件
# 切换到项目目录，安装pkg
# 进入到静态文件目录里，fis3构建静态文件
# 切换到项目根目录，启动端口监听
# 端口监听20s,然后模拟输出Ctrl+c,主要是确保bable编译完成
# 运行mocha单元测试
git clone https://github.com/zuoyanart/pizzaCmsApi.git pizzatest && cd pizzatest && godep restore && neo run -c config.toml;
