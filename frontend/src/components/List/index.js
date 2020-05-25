import React from 'react';

import AddItem from './AddItem';

const List = ({ title, items, handleAdd, handleDelete, handleSelect }) => (
  <>
    <p>{title}</p>
    <ul>
      {items.map((item) => (
        <li key={item.id}>
          <button onClick={() => handleSelect(item)}>{item.name}</button>
          <button onClick={() => handleDelete(item.id)}>&#10799;</button>
        </li>
      ))}
      <li>
        <AddItem onSubmit={handleAdd} />
      </li>
    </ul>
  </>
);

export default List;
