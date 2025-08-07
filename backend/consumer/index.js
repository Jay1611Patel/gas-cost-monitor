require('dotenv').config({ path: '../../.env' });

const { Kafka } = require('kafkajs');
const axios = require('axios');

const API_URL = process.env.API_URL || 'http://localhost:3000';
const API_GAS_REPORT_ENDPOINT = `${API_URL}/api/gas-report`;

const kafka = new Kafka({
    clientId: 'gas-reporter-consumer',
    brokers: process.env.KAFKA_BROKERS ? process.env.KAFKA_BROKERS.split(',') : ['localhost:9092'],
});

const consumer = kafka.consumer({ groupId: 'gas-report-group' });
const topic = 'gas-reports';

const run = async () => {
    await consumer.connect();
    await consumer.subscribe({ topic: topic, fromBeginning: true });

    await consumer.run({
        eachMessage: async ({ topic, partition, message }) => {
            try {
                const gasReport = JSON.parse(message.value.toString());
                console.log(`Received message from Kafka:`, gasReport);

                await axios.post(API_GAS_REPORT_ENDPOINT, gasReport);
                console.log(`Successfully sent gas report to API for tenantId: ${gasReport.tenantId}`);
            } catch (error) {
                console.error(`Error processing Kafka message:`, error.message);
            }
        }
    });
}

run().catch(console.error);

const signalHandler = async (signal) => {
  console.log(`Received signal: ${signal}. Disconnecting consumer.`);
  await consumer.disconnect();
  process.exit(0);
};

process.on('SIGINT', signalHandler);
process.on('SIGTERM', signalHandler);
