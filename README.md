# goYun
## 接口
### 获取所有项目
* url: /get
* 返回： 所有项目名称
### 获取项目信息
* url: /procject/get
* 参数: project 项目名称
* 返回： 该项目的所有函数
### 新增项目
* url: /procject/add
* 参数: project 项目名称
* 返回： true ? false
### 获取项目的单个函数
* url: /func/get
* 参数: project 项目名称, func_name 函数名称
* 返回: 返回该函数的内容
### 新增函数、更新函数
* url: /func/add
* 参数：project 项目名称, func_name 函数名称, func_data 函数内容
* 返回：true ? false
### 删除函数
* url: /func/del
* 参数：project 项目名称, func_name 函数名称
* 返回：true ? false
### 编译项目
* url: /procject/build
* 参数: project 项目名称
* 返回：编辑结果
