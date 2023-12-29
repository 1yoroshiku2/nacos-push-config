package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

type NacosClient struct {
	clientConfig *constant.ClientConfig
}

func (c *NacosClient) GetClientConfig() *constant.ClientConfig {
	return c.clientConfig
}

func main() {
	// 从环境变量获取 Nacos 配置信息
	nacosEndpoint := os.Getenv("NACOS_ENDPOINT")
	namespaceID := os.Getenv("NAMESPACE_ID")
	group := os.Getenv("GROUP")
	dataID := os.Getenv("DATA_ID")
	Type := os.Getenv("TYPE")
	// nacosEndpoint := "101.91.152.37"
	// namespaceID := "675313de-7d3f-488f-8494-ce4b34afd4ed"
	// group := "DEFAULT_GROUP"
	// dataID := "hs-test999"

	// 检查必要参数是否存在
	if nacosEndpoint == "" || namespaceID == "" || group == "" || dataID == "" || Type == "" {
		log.Fatal("Missing required environment variables: NACOS_ENDPOINT, NAMESPACE_ID, GROUP, DATA_ID,TYPE")
	}

	//create ServerConfig
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(nacosEndpoint, 30210, constant.WithContextPath("/nacos")),
	}

	//create ClientConfig
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(namespaceID),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		//constant.WithLogDir("C:/Users/hs/Desktop/test-go/nacos/log"),
		//constant.WithCacheDir("/tmp/nacos/cache"),
		//constant.WithLogLevel("debug"),
	)

	// create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	// 从环境变量获取本地 YAML 文件路径
	yamlFilePath := os.Getenv("YAML_FILE")
	//yamlFilePath := "C:/Users/hs/Desktop/test-go/nacos/public-center.yaml"

	// 检查必要参数是否存在
	if yamlFilePath == "" {
		log.Fatal("Missing required environment variable: YAML_FILE")
	}

	// 读取本地 YAML 文件
	yamlFile, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		log.Fatalf("Failed to read YAML file: %v", err)
	}

	// 推送配置到 Nacos
	// success, err := client.PublishConfig(
	// 	&.ConfigParam{
	// 		DataId:  dataID,
	// 		Group:   group,
	// 		Content: string(yamlFile),
	// 		Type:    "yaml",
	// 	})

	_, err = client.PublishConfig(vo.ConfigParam{
		//tenant:  namespaceID,
		DataId:  dataID,
		Group:   group,
		Content: string(yamlFile),
		Type:    Type,
	})

	if err != nil {
		log.Fatalf("Failed to publish config to Nacos: %v", err)
	}

}
