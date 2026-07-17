// This file will: Generate random events, Keep track of existing orders, Ensure updates only happen on existing orders, Send events to the Go backend
import { customers, restaurants, menu, statuses } from "./data.js";

import { sendEvent } from "./api.js";

//store order IDs
const ordersIds = [];

//Helper Function
function randomFromArray(array) {
  return array[Math.floor(Math.random() * array.length)];
}

function generateItems() {
  const items = [];
  const numberOfItems = Math.floor(Math.random() * 3) + 1;

  for (let i = 0; i < numberOfItems; i++) {
    items.push({
      itemId: randomFromArray(menu),
      qty: Math.floor(Math.random() * 3) + 1,
    });
  }

  return items;
}

//Event Functions

async function createOrderEvent() {
  const event = {
    type: "order.create",
    customerId: randomFromArray(customers),
    restaurantId: randomFromArray(restaurants),
    items: generateItems(),
  };

  const response = await sendEvent(event);

  if (response && response.orderId) {
    ordersIds.push(response.orderId);
    console.log("Order Created:", response.orderId);
  }
}

async function updateStatusEvent() {
  if (ordersIds.length === 0) return;

  const event = {
    type: "order.update.status",
    orderId: randomFromArray(ordersIds),
    status: randomFromArray(statuses),
  };

  await sendEvent(event);

  console.log("Status Updated");
}

async function updateItemsEvent() {
  if (ordersIds.length === 0) return;

  const event = {
    type: "order.update.items",
    orderId: randomFromArray(ordersIds),
    items: generateItems(),
  };

  await sendEvent(event);

  console.log("Items Updated");
}

//Main Generator

async function generateEvent() {
  const random = Math.random();

  if (ordersIds.length < 5) {
    await createOrderEvent();
    return;
  }

  if (random < 0.4) {
    await createOrderEvent();
  } else if (random < 0.7) {
    await updateStatusEvent();
  } else {
    await updateItemsEvent();
  }
}

export { generateEvent };
