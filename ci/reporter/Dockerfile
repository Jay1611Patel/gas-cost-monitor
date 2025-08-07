FROM node:18

WORKDIR /usr/app/reporter

COPY package*.json ./
RUN npm install

COPY . .

CMD sh -c "while [ ! -f /usr/app/contracts/gasReporterOutput.json ]; do echo '⏳ Waiting for gas report...'; sleep 2; done && node reporter.js"
