### 配置目录说明

- 动态配置项基于yaml文件，编译后启动时可传入config目录启动;
- 根据debug、release、test三种环境载入配置，配置文件于对应环境的目录中；
- 可使用ginger-gen工具，根据yaml文件生成配置项解析代码；
