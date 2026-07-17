import dotenv from "dotenv";

dotenv.config({
  path: "./.env",
});

import { generateEvent } from "./eventGenerator.js";

const interval = process.env.INTERVAL || 1000;

console.log("Order Event Generator Started...");
console.log(`Sending events every ${interval} ms`);

setInterval(async () => {
  await generateEvent();
}, interval);