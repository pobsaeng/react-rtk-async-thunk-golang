import React from 'react';                                              
import ReactDOM from 'react-dom/client';                                
import reportWebVitals from './reportWebVitals';
import { Provider } from 'react-redux';                             
import "bootstrap/dist/css/bootstrap.min.css";
import store from "./rtk/store";
import ProductComponent from './components/ProductComponent';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <Provider store={store}>
    <ProductComponent />
    </Provider>
  </React.StrictMode>
);
reportWebVitals();







