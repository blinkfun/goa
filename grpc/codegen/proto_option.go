package codegen

type valueType int

const (
	valueTypeBool   = 1
	valueTypeString = 2
)

type option struct {
	name      string
	valueType valueType
}

var validProtoOptions = []option{
	{
		name:      "cc_enable_arenas",
		valueType: valueTypeBool,
	},
	{
		name:      "csharp_namespace",
		valueType: valueTypeString,
	},
	{
		name:      "go_package",
		valueType: valueTypeString,
	},
	{
		name:      "java_multiple_files",
		valueType: valueTypeBool,
	},
	{
		name:      "java_outer_classname",
		valueType: valueTypeString,
	},
	{
		name:      "java_package",
		valueType: valueTypeString,
	},
	{
		name:      "objc_class_prefix",
		valueType: valueTypeString,
	},
	{
		name:      "php_namespace",
		valueType: valueTypeString,
	},
	{
		name:      "ruby_package",
		valueType: valueTypeString,
	},
}

func validOption(name, value string) bool {
	for _, v := range validProtoOptions {
		if name == v.name && validOptionValue(v, value) {
			return true
		}
	}
	return false
}

func validOptionValue(opt option, value string) bool {
	if opt.valueType == valueTypeBool {
		return value == "true" || value == "false"
	}
	return true
}
