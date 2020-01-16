package settings


type Settings struct {
	RpcPort   string `yaml:"rpc-port"`
	MachineId uint8  `yaml:"machine-id"`
}
