package registry

import (
	"fmt"
	"github.com/kordar/hashring"
)

type HashringRegistry struct {
	count    map[string]int
	hashring *hashring.HashRing
	node     string
}

func NewHashringRegistry(spots int, node string) HashringRegistry {
	return HashringRegistry{
		count:    map[string]int{},
		hashring: hashring.NewHashRing(spots),
		node:     node,
	}
}

func (h *HashringRegistry) Load(values []string, f func(v interface{}) (string, int, []string)) {
	weights := map[string]int{}
	tmp := map[string]int{}
	for _, value := range values {
		id, w, nodes := f(value)
		if id == "" {
			continue
		}
		if nodes == nil || len(nodes) == 0 {
			weights[id] = w
			tmp[id]++
		} else {
			for _, node := range nodes {
				key := fmt.Sprintf("%s:%s", id, node)
				weights[key] = w
				tmp[key]++
			}
		}
	}

	for k := range h.count {
		if _, exists := tmp[k]; !exists {
			h.hashring.RemoveNode(k)
		}
	}

	for k, v := range tmp {
		w := 1
		if weights[k] != 0 {
			w = weights[k]
		}
		h.hashring.AddNode(k, w*v)
	}

	h.count = tmp
}

func (h *HashringRegistry) GetNode(key string) string {
	return h.hashring.GetNode(key)
}

func (h *HashringRegistry) Can(key string) bool {
	return h.hashring.GetNode(key) == h.node
}
