package pulic_type

type (
	// 公用配置文件
	ConfType struct {
		Name     string                  `json:"name"`
		MicroSer map[string]MicroSerType `json:"micro_ser" yaml:"micro_ser"`
		// 配置文件的属性
		Attr    map[string]interface{} `json:"attr"`
		Version string                 `json:"version"`
	}

	// 微服务配置文件,一个对象对应一个网络服务
	MicroSerType struct {
		//	服务的地址
		Addr string `json:"addr"`
		// 服务的用户名或者id
		Id string `json:"id"`
		// 服务的授权密码
		Secret string `json:"secret"`
		// 服务的属性
		Attr map[string]interface{} `json:"attr"`
	}
)
