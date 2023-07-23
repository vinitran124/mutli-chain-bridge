// api.js
import axios from 'axios';

// Create a new instance of axios with the base URL from .env
export const api = axios.create({
  baseURL: 'https://54.255.20.201',
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000,
});
