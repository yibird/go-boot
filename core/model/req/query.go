package req

// BaseQueryModel 通用查询模型
type BaseQueryModel struct {
	Current  int64 `form:"current" json:"current"`   // 当前页数
	PageSize int64 `form:"pageSize" json:"pageSize"` // 每页显示条数
}

type Sort struct {
	// 需要排序的列集合
	Columns []string `mapstructure:"columns" json:"columns" yaml:"columns"`
	// 排序方式,支持desc和aes两种
	Mode string `mapstructure:"mode" json:"mode" yaml:"mode"`
}

type Conditional struct {
	// 查询列名
	Column string `mapstructure:"column" json:"column" yaml:"column"`
	// 操作符,例如 like、not、eq等等
	Operator string `mapstructure:"operator" json:"operator" yaml:"operator"`
	// 查询值
	Value any `mapstructure:"value" json:"value" yaml:"value"`
}

type ConditionalQuery struct {
	// 条件集合
	Conditionals []Conditional `mapstructure:"conditionals" json:"conditionals" yaml:"conditionals"`
	// 条件连接符,例如 and、or等等
	Connector string `mapstructure:"connector" json:"connector" yaml:"connector"`
}

// AdvancedQuery 高级查询模型
type AdvancedQuery struct {
	// 列排序
	Sort Sort `mapstructure:"sort" json:"sort" yaml:"sort"`
	// 列分组
	Group []string `mapstructure:"group" json:"group" yaml:"group"`
	// 列查询
	Querys []ConditionalQuery `mapstructure:"querys" json:"querys" yaml:"querys"`
}
