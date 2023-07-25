// api.js
import axios from 'axios';

// Create a new instance of axios with the base URL from .env
export const api = axios.create({
  baseURL: 'http://13.215.62.47:3030',
  headers: {
    'Content-Type': 'application/json',
  },
  timeout: 10000,
});
