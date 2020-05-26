import React from 'react';
import cx from 'classnames';

import AddItem from './AddItem';
import styles from './styles.module.css';

const List = ({ title, items, handleAdd, handleDelete, handleSelect }) => (
  <div>
    <p className={styles.title}>{title}</p>
    <ul className={styles.list}>
      {items.map((item) => (
        <li key={item.id} className={styles.item}>
          <button
            className={cx(styles.button, styles.content)}
            onClick={() => handleSelect(item)}
          >
            {item.name}
          </button>
          <button
            className={cx(styles.button, styles.control, styles.delete)}
            onClick={() => handleDelete(item.id)}
          >
            &#10799;
          </button>
        </li>
      ))}
      <li className={styles.item}>
        <AddItem onSubmit={handleAdd} />
      </li>
    </ul>
  </div>
);

export default List;
