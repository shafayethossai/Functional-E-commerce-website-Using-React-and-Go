// import React, { useState, useEffect } from 'react';
// import '../styles/AddEditProduct.css';

// function AddEditProduct({ product, onSave, onCancel }) {
//   const [formData, setFormData] = useState({
//     title: '',
//     description: '',
//     price: '',
//     imageUrl: '',
//   });

//   useEffect(() => {
//     if (product) {
//       setFormData({
//         title: product.title,
//         description: product.description,
//         price: product.price,
//         imageUrl: product.imageUrl,
//       });
//     }
//   }, [product]);

//   const handleChange = (e) => {
//     const { name, value } = e.target;
//     setFormData((prevData) => ({
//       ...prevData,
//       [name]: value,
//     }));
//   };

//   const handleSubmit = (e) => {
//     e.preventDefault();
//     onSave(formData);
//   };

//   return (
//     <div className="add-edit-product">
//       <h2>{product ? 'Edit Product' : 'Add Product'}</h2>
//       <form onSubmit={handleSubmit}>
//         <label>
//           Title:
//           <input
//             type="text"
//             name="title"
//             value={formData.title}
//             onChange={handleChange}
//           />
//         </label>
//         <label>
//           Description:
//           <textarea
//             name="description"
//             value={formData.description}
//             onChange={handleChange}
//           />
//         </label>
//         <label>
//           Price:
//           <input
//             type="number"
//             name="price"
//             value={formData.price}
//             onChange={handleChange}
//           />
//         </label>
//         <label>
//           Image URL:
//           <input
//             type="text"
//             name="imageUrl"
//             value={formData.imageUrl}
//             onChange={handleChange}
//           />
//         </label>
//         <div className="buttons">
//           <button type="submit">{product ? 'Save Changes' : 'Add Product'}</button>
//           <button type="button" onClick={onCancel}>
//             Cancel
//           </button>
//         </div>
//       </form>
//     </div>
//   );
// }

// export default AddEditProduct;

import React, { useState, useEffect } from 'react';
import '../styles/AddEditProduct.css';

function AddEditProduct({ product, onSave, onCancel }) {
  const [formData, setFormData] = useState({
    title: '',
    description: '',
    price: '',
    imageUrl: '',
  });

  useEffect(() => {
    if (product) {
      setFormData({
        title: product.title,
        description: product.description,
        price: product.price,
        imageUrl: product.imageUrl,
      });
    } else {
      // Reset form when adding new product
      setFormData({
        title: '',
        description: '',
        price: '',
        imageUrl: '',
      });
    }
  }, [product]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    
    // Convert price to number
    const productData = {
      ...formData,
      price: Number(formData.price)
    };
    
    onSave(productData);
    
    // Only reset form if not editing
    if (!product) {
      setFormData({
        title: '',
        description: '',
        price: '',
        imageUrl: '',
      });
    }
  };

  return (
    <div className="add-edit-product">
      <h2>{product ? 'Edit Product' : 'Add Product'}</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Title:
          <input
            type="text"
            name="title"
            value={formData.title}
            onChange={handleChange}
            required
          />
        </label>
        <label>
          Description:
          <textarea
            name="description"
            value={formData.description}
            onChange={handleChange}
            required
          />
        </label>
        <label>
          Price:
          <input
            type="number"
            name="price"
            value={formData.price}
            onChange={handleChange}
            step="0.01"
            min="0"
            required
          />
        </label>
        <label>
          Image URL:
          <input
            type="text"
            name="imageUrl"
            value={formData.imageUrl}
            onChange={handleChange}
          />
        </label>
        <div className="buttons">
          <button type="submit">{product ? 'Save Changes' : 'Add Product'}</button>
          <button type="button" onClick={onCancel}>
            Cancel
          </button>
        </div>
      </form>
    </div>
  );
}

export default AddEditProduct;