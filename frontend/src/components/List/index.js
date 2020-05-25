import React from 'react';

import AddItem from './AddItem';

const List = ({ items, handleAdd, handleSelect }) => (
  <ul>
    {items.map((item) => (
      <li key={item.id}>
        <button onClick={() => handleSelect(item)}>{item.name}</button>
      </li>
    ))}
    <li>
      <AddItem onSubmit={handleAdd} />
    </li>
  </ul>
);

export default List;
