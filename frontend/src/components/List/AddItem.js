import React, { useState } from 'react';

const AddItem = ({ onSubmit }) => {
  const [name, setName] = useState('');

  return (
    <>
      <input
        type="text"
        placeholder="Add"
        value={name}
        onChange={(event) => setName(event.target.value)}
      />
      <button onClick={(event) => onSubmit(name)}>&#10003;</button>
    </>
  );
};

export default AddItem;
