import React from 'react';
import styles from './Layout.module.scss';
import { Outlet } from 'react-router-dom';
import Sidebar from '../Sidebar/Sidebar';

const Layout: React.FC = () => {
  return (
    <div className={styles.layout}>
        <Sidebar/>

      <div className={styles.mainArea}>
        <div className={styles.header}>
          <h2>Dashboard</h2>
        </div>
        <main className={styles.content}>
          <Outlet />
        </main>
      </div>
    </div>
  );
};

export default Layout;
