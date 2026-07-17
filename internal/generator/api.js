//This file is responsible for only one thing Sending HTTP requests to Go Backend
import axios from "axios";

async function sendEvent(event) {
  try {
    const response = await axios.post(process.env.API_URL, event);

    return response.data;
  } catch (error) {
    console.error("API Error:", error.message);

    return null;
  }
}

export { sendEvent };
