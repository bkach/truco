import React from 'react';

const List = ({ items, handleAdd, handleSelect }) => (
  <ul>
    {items.map((item) => (
      <li key={item.id}>
        <button onClick={() => handleSelect(item)}>{item.name}</button>
      </li>
    ))}
    <li>
      <button onClick={handleAdd}>Add</button>
    </li>
  </ul>
);

export default List;
