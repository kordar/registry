package registry_test

import (
	"encoding/json"
	"github.com/kordar/registry"
	"log"
	"testing"
)

type Node struct {
	Id       string   `json:"id"`
	Children []string `json:"children"`
}

var str = []string{"{\"id\":\"8001\"}", "{\"id\":\"8002\"}", "{\"id\":\"8003\",\"children\":[\"AAA\",\"BBB\"]}"}
var str2 = []string{"{\"id\":\"8001\"}", "{\"id\":\"8001\"}"}

func TestHashringRegistry_Load(t *testing.T) {
	hashringRegistry := registry.NewHashringRegistry(1, "8001")
	hashringRegistry.Load(str, func(v interface{}) (string, int, []string) {
		node := Node{}
		json.Unmarshal([]byte(v.(string)), &node)
		return node.Id, 1, node.Children
	})

	node := hashringRegistry.GetNode("8908897708970")

	log.Println("----------------->", node)

}
