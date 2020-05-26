import React, { useState } from 'react';
import cx from 'classnames';

import styles from './styles.module.css';

const AddItem = ({ onSubmit }) => {
  const [name, setName] = useState('');

  return (
    <>
      <input
        className={cx(styles.input, styles.content)}
        type="text"
        placeholder="Add"
        value={name}
        onChange={(event) => setName(event.target.value)}
      />
      <button
        className={cx(styles.button, styles.control, styles.submit)}
        onClick={(event) => onSubmit(name)}
      >
        &#10003;
      </button>
    </>
  );
};

export default AddItem;
