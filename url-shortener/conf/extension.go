package conf

type ConfigExtension int

const (
	ConfigExtensionYaml ConfigExtension = iota
	ConfigExtensionJson
	ConfigExtensionDb
)

var ConfigExtensionHashmap = map[string]ConfigExtension{
	"yaml": ConfigExtensionYaml,
	"yml":  ConfigExtensionYaml,
	"json": ConfigExtensionJson,
	"db":   ConfigExtensionDb,
}
