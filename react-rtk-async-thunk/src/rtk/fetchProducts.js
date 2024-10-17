import { createAsyncThunk } from '@reduxjs/toolkit';
import axios from 'axios';

const API_URL = 'http://localhost:8083/api/v1/products';

// Fetch products from the API
export const fetchProductsFromAPI = createAsyncThunk(
  'product/fetchProductsFromAPI',
  async () => {
    const response = await axios.get(API_URL);
    return response.data; // Return product data from the API
  }
);