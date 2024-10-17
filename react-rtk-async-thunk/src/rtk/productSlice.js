import { createSlice } from '@reduxjs/toolkit';
import { fetchProductsFromAPI } from './fetchProducts';

const initialState = [];
const productSlice = createSlice({
     name: "product",
     initialState: { products: initialState },
     loading: false,
     error: null,
     reducers: {},
     extraReducers: (builder) => {
          builder.addCase(fetchProductsFromAPI.pending, (state) => {
               state.loading = true;
               state.error = null;
          }).addCase(fetchProductsFromAPI.fulfilled, (state, action) => {
               state.loading = false;
               state.products = action.payload;
          }).addCase(fetchProductsFromAPI.rejected, (state, action) => {
               state.loading = false;
               state.error = action.error.message;
          });
     }
});
export default productSlice;


