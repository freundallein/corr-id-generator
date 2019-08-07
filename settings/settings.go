package settings

type Settings struct {
	Name      string `yaml:"name"`
	RpcPort   string `yaml:"rpc-port"`
	MachineId uint8  `yaml:"machine-id"`
}
