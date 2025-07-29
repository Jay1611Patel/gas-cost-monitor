require('dotenv').config({ path: '../../../.env' });

const fs = require('fs');
const path = require('path');
const axios = require('axios');
const { getPackedSettings } = require('http2');

const GAS_REPORT_FILE_PATH = path.resolve(__dirname, '../../../contracts/gasReporterOutput.json');
const API_URL = process.env.API_URL || 'http://localhost:3000';
const API_GAS_REPORT_ENDPOINT = `${API_URL}/api/gas-report`;

async function sendGasReport () {
    console.log(`[CI Reporter] Starting to send gas report...`);

    if (!fs.existsSync(GAS_REPORT_FILE_PATH)) {
        console.error(`[CI Reporter] Gas report file does not exist at path: ${GAS_REPORT_FILE_PATH}`);
        console.error(`[CI Reporter] Please ensure Hardhat tests run successfully first.`);
        process.exit(1);
    }

    let rawGasReport;
    try {
        rawGasReport = fs.readFileSync(GAS_REPORT_FILE_PATH, 'utf8');
    } catch (error) {
        console.error(`[CI Reporter] Error reading gas report file: ${error.message}`);
        process.exit(1);
    }

    let gasReport;
    try {
        gasReport = JSON.parse(rawGasReport);
    } catch (error) {
        console.error(`[CI Reporter] Error parsing gas report: ${error.message}`);
        process.exit(1);
    }

    const recordsToSend = [];

    for (contractName in gasReport) {
        // checks if 
        if (Object.prototype.hasOwnProperty.call(gasReport, contractName)) {
            const contractMethods = gasReport[contractName];
            for (const methodName in contractMethods) {
                if (Object.prototype.hasOwnProperty.call(contractMethods, methodName)) {
                    const methodData = contractMethods[methodName];
                    recordsToSend.push({
                        contractName: contractName,
                        method: methodName,
                        gasUsed: methodData.gasUsed,
                        costUsd: methodData.usdCost,
                        timestamp: new Date().toISOString(),
                    });
                }
            }
        }
    }

    if (recordsToSend.length === 0) {
        console.warn("[CI Reporter] No gas data found to send in the report.");
        return;
    }

    console.log(`[CI Reporter] Found ${recordsToSend.length} gas records to send.`);

    for (const record of recordsToSend) {
        try {
            console.log(`[CI Reporter] Sending gas report for ${record.contractName}.${record.method}...`);
            const response = await axios.post(API_GAS_REPORT_ENDPOINT, record);
            console.log(`[CI Reporter] Successfully sent ${record.contractName}.${record.method}:`, response.data.message);
        } catch (error) {
            console.error(`[CI Reporter] Error sending record for ${record.contractName}.${record.method}:`);
            if (error.response) {
                console.error(`  Status: ${error.response.status}`);
                console.error(`  Data: ${JSON.stringify(error.response.data, null, 2)}`);
            } else if (error.request) {
                console.error(`  No response received. Request made but no response:`, error.request);
            } else {
                console.error(`  Error in request setup:`, error.message);
            }
        }
    }
    console.log(`[CI Reporter] Gas report sending process complete.`);


}

sendGasReport();
