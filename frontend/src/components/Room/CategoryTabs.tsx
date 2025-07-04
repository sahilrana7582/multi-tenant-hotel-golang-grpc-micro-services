import React from 'react';
import styles from './CategoryTabs.module.scss';

interface CategoryTabsProps {
  categories: string[];
  selectedCategory: string;
  onSelect: (category: string) => void;
}

const CategoryTabs: React.FC<CategoryTabsProps> = ({ categories, selectedCategory, onSelect }) => {
  return (
    <div className={styles.tabContainer}>
      {categories.map((category) => (
        <button
          key={category}
          onClick={() => onSelect(category)}
          className={`${styles.tabButton} ${selectedCategory === category ? styles.active : ''}`}
        >
          {category}
        </button>
      ))}
    </div>
  );
};

export default CategoryTabs;