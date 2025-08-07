package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type GasReport struct {
	Timestamp      time.Time `bson:"timestamp"`
	BlockHash      string    `bson:"blockHash"`
	TransactionHash string    `bson:"transactionHash"`
	FromAddress     string    `bson:"fromAddress"`
	ToAddress      string    `bson:"toAddress"`
	GasUsed       uint64    `bson:"gasUsed"`
	GasPrice      *big.Int  `bson:"gasPrice"`
	ReportType     string    `bson:"reportType"` // e.g., "on-chain", "local"
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %v", err)
	}

	ethNodeURL := os.Getenv("ETH_NODE_URL")
	if ethNodeURL == "" {
		log.Fatal("ETH_NODE_URL environment variable not set (e.g., wss://sepolia.infura.io/ws/v3/YOUR_INFURA_PROJECT_ID)")
	}

	contractAddressStr := os.Getenv("CONTRACT_ADDRESS")
	if contractAddressStr == "" {
		log.Fatal("CONTRACT_ADDRESS environment variable not set")
	}

	kafkaBrokerURL := os.Getenv("KAFKA_BROKER_URL")
	if kafkaBrokerURL == "" {
		kafkaBrokerURL = "localhost:9092"
		log.Printf("KAFKA_BROKER_URL not set, defaulting to %s", kafkaBrokerURL)
	}

	kafkaTopic := os.Getenv("KAFKA_TOPIC")
	if kafkaTopic == "" {
		kafkaTopic = "blockchain-events"
		log.Printf("KAFKA_TOPIC not set, defaulting to %s", kafkaTopic)
	}

	client, err := ethclient.Dial(ethNodeURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer client.Close()
	log.Printf("Connected to Ethereum node: %s", ethNodeURL)

	contractAddress := common.HexToAddress(contractAddressStr)
	log.Printf("Monitoring contract at address: %s", contractAddress.Hex())

	contractAbi, err := abi.JSON(strings.NewReader(contracts.MyContractABI))
	if err != nil {
		log.Fatalf("Failed to parse contract ABI: %v", err)
	}

	valueUpdatedEventSignature := []byte("ValueUpdated(uint256,address)")
	valueUpdatedEventHash := crypto.Keccak256Hash(valueUpdatedEventSignature)
	log.Printf("Listening for event: ValueUpdated (Hash: %s)", valueUpdatedEventHash.Hex())

	writer := &kafka.Writer{
		Addr:         kafka.TCP(kafkaBrokerURL),
		Topic:        kafkaTopic,
		Balancer:     &kafka.LeastBytes{},
		RequiredAcks: kafka.RequireOne,
		Async:        false,
		BatchTimeout: time.Millisecond * 10,
	}
	defer writer.Close()
	log.Printf("Kafka producer connected to %s, topic %s", kafkaBrokerURL, kafkaTopic)

	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics:    [][]common.Hash{{valueUpdatedEventHash}},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatalf("Failed to subscribe to contract events: %v", err)
	}
	defer sub.Unsubscribe()
	log.Println("Subscribed to contract events.")

	for {
		select {
		case err := <-sub.Err():
			log.Fatalf("Subscription error: %v", err)
		case vLog := <-logs:
			log.Printf("Received new log from block %d, Tx Hash: %s", vLog.BlockNumber, vLog.TxHash.Hex())

			var eventData interface{}
			var eventType string

			if vLog.Topics[0] == valueUpdatedEventHash {
				eventType = "ValueUpdated"
				var valueUpdated ValueUpdatedEvent

				err := contractAbi.UnpackIntoInterface(&valueUpdated, "ValueUpdated", vLog.Data)
				if err != nil {
					log.Printf("Failed to unpack ValueUpdated event data: %v", err)
					continue
				}
				if len(vLog.Topics) >= 2 {
					valueUpdated.Sender = common.BytesToAddress(vLog.Topics[2].Bytes())
				}
				valueUpdated.Raw = vLog
				eventData = valueUpdated

			} else {
				log.Printf("Unknown event received (Topic: %s)", vLog.Topics[0].Hex())
				continue
			}

			msg := KafkaEventMessage{
				EventType:       eventType,
				ContractName:    "MyContract",
				ContractAddress: contractAddress.Hex(),
				TransactionHash: vLog.TxHash.Hex(),
				BlockNumber:     vLog.BlockNumber,
				Timestamp:       time.Now().Format(time.RFC3339),
				EventData:       eventData,
			}
			jsonMsg, err := json.Marshal(msg)
			if err != nil {
				log.Printf("Failed to marshal Kafka message: %v", err)
				continue
			}
			err = writer.WriteMessages(context.Background(), kafka.Message{
				Key:   []byte(vLog.TxHash.Hex()),
				Value: jsonMsg,
			})
			if err != nil {
				log.Printf("Failed to write message to Kafka: %v", err)
			} else {
				log.Printf("Successfully sent %s event from block %d to Kafka.", eventType, vLog.BlockNumber)
			}

		}
	}

}
