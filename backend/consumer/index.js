require('dotenv').config({ path: '../../.env' });

const { Kafka } = require('kafkajs');
const axios = require('axios');

// --- Configuration ---
// Get Kafka brokers and API URL from environment variables set in docker-compose.yml
const kafkaBrokers = process.env.KAFKA_BROKERS.split(',');
const apiUrl = process.env.API_URL;
const topic = 'gas-reports';
const groupId = 'gas-report-group';

// --- Kafka Client Initialization ---
const kafka = new Kafka({
  clientId: 'gas-report-consumer',
  brokers: kafkaBrokers,
  // Add retry logic to handle initial connection issues
  retry: {
    initialRetryTime: 300,
    retries: 8
  }
});

const consumer = kafka.consumer({ groupId: groupId });

// --- Main Execution Logic ---
const run = async () => {
  try {
    // 1. Connect the consumer to the Kafka cluster
    await consumer.connect();
    console.log('Consumer connected to Kafka');

    // 2. Subscribe to the specified topic
    // fromBeginning: true ensures we process messages from the start if the consumer is new
    await consumer.subscribe({ topic: topic, fromBeginning: true });
    console.log(`Consumer subscribed to topic: ${topic}`);

    // 3. Run the consumer, processing each message
    await consumer.run({
      eachMessage: async ({ topic, partition, message }) => {
        try {
          // The message value from Kafka is a Buffer, so we convert it to a string
          const reportJson = message.value.toString();
          console.log(`Received message from partition ${partition}: ${reportJson}`);
          console.log("------------");
          console.log(`API URL: ${apiUrl}/api/reports`);
          console.log("++++++++++++");

          // Parse the JSON string into an object
          const reportData = JSON.parse(reportJson);
          

          // Make a POST request to the backend API to save the report
          // The API endpoint should be designed to receive this data structure

          
          await axios.post(`${apiUrl}/api/reports`, reportData);

          console.log(`Successfully sent report for tenant ${reportData.tenantId} to API.`);
        } catch (error) {
          console.error('Error processing message or sending to API:', error.message);
          // Decide on an error handling strategy: retry, send to a dead-letter queue, etc.
        }
      },
    });
  } catch (error) {
    console.error('Failed to start the Kafka consumer:', error);
    // Exit the process if the consumer cannot start, so Docker can restart it
    process.exit(1);
  }
};

// --- Graceful Shutdown ---
const shutdown = async () => {
  console.log('Shutting down consumer...');
  await consumer.disconnect();
  process.exit(0);
};

// Listen for signals to gracefully shut down
process.on('SIGINT', shutdown);
process.on('SIGTERM', shutdown);

// Start the consumer
run();