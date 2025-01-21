package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	blockHeight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoin_block_height",
		Help: "Current block height",
	})
	numPeers = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "bitcoin_num_peers",
		Help: "Number of peers connected",
	})
)

func init() {
	prometheus.MustRegister(blockHeight, numPeers)
}

func main() {
	// set  environment variables 
	bitcoinRPCURL := os.Getenv("BITCOIN_RPC_URL")
	bitcoinRPCUSER := os.Getenv("BITCOIN_RPC_USER")
	bitcoinRPCPASSWORD := os.Getenv("BITCOIN_RPC_PASSWORD")

	if bitcoinRPCURL == "" || bitcoinRPCUSER == "" || bitcoinRPCPASSWORD == "" {
		log.Fatal("BITCOIN_RPC_URL, BITCOIN_RPC_USER, and BITCOIN_RPC_PASSWORD must be set")
	}

	// Set up RPC client
	connCfg := &rpcclient.ConnConfig{
		Host:         bitcoinRPCURL,
		User:         bitcoinRPCUSER,
		Pass:         bitcoinRPCPASSWORD,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatalf("Error connecting to Bitcoin RPC server: %v", err)
	}
	defer client.Shutdown()

	// Set default app port
	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8335"
	}

	// expose metrics endpoint for prometheus
	http.Handle("/metrics", promhttp.Handler())

	// Update metrics periodically 
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			updateMetrics(client)
		}
	}()

	log.Printf("Bitcoin Node Exporter listening on port :%s", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}

func updateMetrics(client *rpcclient.Client) {
	blockCount, err := client.GetBlockCount()
	if err != nil {
		log.Printf("Error getting block count: %v", err)
	} else {
		blockHeight.Set(float64(blockCount))
	}

	peerInfo, err := client.GetPeerInfo()
	if err != nil {
		log.Printf("Error getting peer info: %v", err)
	} else {
		numPeers.Set(float64(len(peerInfo)))
	}

}
