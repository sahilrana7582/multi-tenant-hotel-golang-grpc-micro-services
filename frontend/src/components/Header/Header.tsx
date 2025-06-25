import { useLocation } from 'react-router-dom';
import styles from './Header.module.scss';

const routeTitles: Record<string, string> = {
  '/': 'Dashboard',
  '/rooms': 'Rooms',
  '/housekeeping': 'Housekeeping',
  '/messages': 'Messages',
  '/kitchen': 'Kitchen',
  '/frontdesk': 'Front Desk',
};

export const Header = () => {
  const location = useLocation();
  const pageTitle = routeTitles[location.pathname] || 'Dashboard';

  return (
    <header className={styles.header}>
      <h2 className={styles.pageTitle}>{pageTitle}</h2>

      <div className={styles.rightSection}>
        <input
          type="text"
          className={styles.searchInput}
          placeholder="Search room, guest, book, etc"
        />

       
      </div>
    </header>
  );
};
