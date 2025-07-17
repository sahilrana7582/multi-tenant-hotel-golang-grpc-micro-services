import { useState } from 'react';
import styles from './Accordion.module.scss';
import type { AccordionItemBase, AccordionProps } from './types';
import { FaHandsHoldingCircle } from 'react-icons/fa6';

const Accordion = <T extends AccordionItemBase>({
    icon,
    title,
    data,
  }: AccordionProps<T>) => {
  const [isOpen, setIsOpen] = useState(false);

  const toggleAccordion = () => setIsOpen(prev => !prev);

  return (
    <div
      className={styles.accordion}
    >
      <div className={styles.header} onClick={toggleAccordion}>
        <div className={styles.icon}>{icon}</div>
        {title}
      </div>

      {isOpen && (
        <ul className={styles.content}>
          {data.map((dept) => (
            <li key={dept.name} className={styles.item}>
                <FaHandsHoldingCircle />
              {dept.name}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Accordion;
