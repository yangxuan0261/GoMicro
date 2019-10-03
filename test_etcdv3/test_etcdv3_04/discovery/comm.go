package discovery

const ServiceKey = "services/"
const ServiceAPI = "/api/test"

// ServiceInfo 需要上报的节点信息
type ServiceInfo struct {
	Name string
	Host string
}
