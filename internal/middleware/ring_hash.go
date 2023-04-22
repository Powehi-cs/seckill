package middleware

import (
	"github.com/Powehi-cs/seckill/pkg/errors"
	"github.com/Powehi-cs/seckill/pkg/logger"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hash/crc32"
	"log"
	"sort"
	"strconv"
	"strings"
)

// Node 哈希环上的每一个节点
type Node struct {
	IP        string
	HashValue uint32
}

// Ring 哈希环
type Ring struct {
	Nodes    []Node // 真实节点+虚拟节点
	Replicas int    // 虚拟节点数量
}

func (ring *Ring) Len() int {
	return len(ring.Nodes)
}

func (ring *Ring) Swap(i, j int) {
	ring.Nodes[i], ring.Nodes[j] = ring.Nodes[j], ring.Nodes[i]
}

func (ring *Ring) Less(i, j int) bool {
	if ring.Nodes[i].HashValue < ring.Nodes[j].HashValue {
		return true
	} else if ring.Nodes[i].HashValue == ring.Nodes[j].HashValue {
		return ring.Nodes[i].IP < ring.Nodes[j].IP
	}

	return false
}

func NewRing(replicas int) *Ring {
	return &Ring{
		Replicas: replicas,
	}
}

func (ring *Ring) Add(ips []string) {
	for _, ip := range ips {
		for i := 0; i < ring.Replicas; i++ {
			ring.Nodes = append(ring.Nodes, Node{
				IP:        ip,
				HashValue: hashFunc(ip + strconv.Itoa(i)),
			})
		}
	}

	sort.Sort(ring)
}

func (ring *Ring) Delete(ips []string) {
	for _, ip := range ips {
		var newNodes []Node
		for _, node := range ring.Nodes {
			if !strings.Contains(node.IP, ip) {
				newNodes = append(newNodes, node)
			}
		}
		ring.Nodes = newNodes
	}
}

// Get 对用户名称做负载均衡
func (ring *Ring) Get(name string) string {
	if ring.Len() == 0 {
		return ""
	}

	hash := hashFunc(name)
	idx := sort.Search(len(ring.Nodes), func(i int) bool {
		node := ring.Nodes[i]
		return (node.HashValue > hash) || (node.HashValue == hash && node.IP >= name)
	})

	if idx == len(ring.Nodes) {
		idx = 0
	}

	return ring.Nodes[idx].IP
}

// 哈希函数
func hashFunc(data string) uint32 {
	hash := crc32.NewIEEE()
	_, err := hash.Write([]byte(data))
	errors.PrintInStdout(err)
	return hash.Sum32()
}

func ConsistentHash() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ring := NewRing(viper.GetInt("ringHash.replicas"))
		ring.Add(viper.GetStringSlice("ringHash.ips"))
		var ip string
		if name, ok := ctx.Get("name"); ok {
			ip = ring.Get(name.(string))
		}
		if ip == "" {
			logger.Fail(ctx, 500, "一致性哈希出错")
			ctx.Abort()
			return
		}

		log.Println(ctx.RemoteIP(), ip, ip == ctx.RemoteIP(), ctx.ClientIP(), ctx.Request.URL.Host)
		if ip == ctx.RemoteIP() {
			ctx.Next()
			return
		}

		newURL := "http://" + ip + ctx.Request.URL.Path
		logger.Redirect(ctx, newURL)
		ctx.Abort()
	}
}
