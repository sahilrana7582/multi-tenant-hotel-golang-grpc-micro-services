import React from 'react';
import styles from './Layout.module.scss';
import { Outlet } from 'react-router-dom';
import Sidebar from '../Sidebar/Sidebar';
import { Header } from '../Header/Header';

const Layout: React.FC = () => {
  return (
    <div className={styles.layout}>
        <Sidebar/>

      <div className={styles.mainArea}>
        <Header/>
        <main className={styles.content}>
          <Outlet />
        </main>
      </div>
    </div>
  );
};

export default Layout;
