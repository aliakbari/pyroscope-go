package pyroscope

import (
	"github.com/grafana/pyroscope-go"
	"github.com/zeromicro/go-zero/zrpc"
)

func PyroscopeInit(config2 zrpc.RpcServerConf) {
	pyroscope.Start(pyroscope.Config{
		ApplicationName: config2.Name + ".messenger.app",

		Logger: pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,
		},
	})
}
