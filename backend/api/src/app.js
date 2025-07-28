require("dotenv").config();

const express = require("express");
const app = express();
const PORT = process.env.API_PORT || 3000;

app.use(express.json());

app.get("/", (req, res) => {
  res.status(200).json({ message: "Gas Cost Monitor API is running!" });
});

app.post("/api/gas-report", (req, res) => {
  const gasReport = req.body;
  console.log("Received gas report:");
  console.log(JSON.stringify(gasReport, null, 2)); // Pretty print JSON

  // In a real application, you would save this to a database
  // For now, just acknowledge receipt
  res.status(200).json({
    message: "Gas report received successfully!",
    status: "ok",
    receivedAt: new Date().toISOString(),
  });
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
