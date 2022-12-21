package iwf

type ClientOptions struct {
	ServerUrl     string
	WorkerUrl     string
	ObjectEncoder ObjectEncoder
}

const (
	defaultWorkerUrl = "http://localhost:8803"
	defaultServerUrl = "http://localhost:8801"
)

var localDefaultClientOptions = ClientOptions{
	ServerUrl:     defaultServerUrl,
	WorkerUrl:     defaultWorkerUrl,
	ObjectEncoder: GetDefaultObjectEncoder(),
}

func GetLocalDefault() ClientOptions {
	return localDefaultClientOptions
}
