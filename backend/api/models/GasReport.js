const mongoose = require('mongoose');

const gasReportSchema = new mongoose.Schema({
  tenantId: {
    type: String,
    required: true,
  },
  contractName: {
    type: String,
    required: true,
  },
  method: {
    type: String,
    required: true,
  },
  gasUsed: {
    type: Number,
    required: true,
  },
  timestamp: {
    type: Date,
    default: Date.now,
  },
});

module.exports = mongoose.model('GasReport', gasReportSchema);