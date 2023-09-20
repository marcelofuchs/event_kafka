package main

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
)

func main() {

	fmt.Printf("CHEGUEI")

	// Configurações do produtor
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // Versão do Kafka

	// Crie o cliente Kafka
	client, err := sarama.NewClient([]string{"localhost:9091"}, config)
	if err != nil {
		fmt.Printf("Erro ao criar o cliente Kafka: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	// Nome do tópico a ser criado
	topicName := "NUM_OC_123"

	// Crie um administrador Kafka
	admin, err := sarama.NewClusterAdminFromClient(client)
	if err != nil {
		fmt.Printf("Erro ao criar o administrador Kafka: %v\n", err)
		os.Exit(1)
	}
	defer admin.Close()

	// Crie o tópico
	topicDetail := &sarama.TopicDetail{
		NumPartitions:     1,   // Número de partições
		ReplicationFactor: 1,   // Fator de replicação
		ConfigEntries:     nil, // Configurações adicionais (opcional)
	}

	err = admin.CreateTopic(topicName, topicDetail, false)
	if err != nil {
		fmt.Printf("Erro ao criar o tópico: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Tópico '%s' criado com sucesso!\n", topicName)
}
