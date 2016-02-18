package pulic_type

type (
	ConfType struct {
		Name     string                  `json:"name"`
		MicroSer map[string]MicroSerType `json:"micro_ser"`
		Version  string                  `json:"version"`
	}

	MicroSerType struct {
		Addr   string                 `json:"addr"`
		Id     string                 `json:"id"`
		Secret string                 `json:"secret"`
		Attr   map[string]interface{} `json:"attr"`
	}
)
