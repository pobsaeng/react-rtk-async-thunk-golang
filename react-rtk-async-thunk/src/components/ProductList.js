import React from "react";

const ProductList = ({ products }) => {
  return (
    <table className="table table-striped">
      <thead>
        <tr>
          <th scope="col">ID</th>
          <th scope="col">Code</th>
          <th scope="col">Name</th>
          <th scope="col">Description</th>
          <th scope="col">Active</th>
          <th scope="col">Price</th>
          <th scope="col">Stock</th>
          <th scope="col">Created By</th>
          <th scope="col">Created At</th>
        </tr>
      </thead>
      <tbody>
        {products.map((product) => (
          <tr key={product.id}>
            <td>{product.id}</td>
            <td>{product.code}</td>
            <td>{product.name}</td>
            <td>{product.description}</td>
            <td>{product.active ? 'Yes' : 'No'}</td>
            <td>{product.price}</td>
            <td>{product.stock}</td>
            <td>{product.created_by}</td>
            <td>{new Date(product.created_at).toLocaleDateString()}</td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default ProductList;