import React, { useState } from 'react';
import { NavLink } from 'react-router-dom';
import {
  FaHome,
  FaBed,
  FaBroom,
  FaChartPie,
  FaUsers,
  FaCog,
  FaChevronLeft,
  FaChevronRight,
} from 'react-icons/fa';
import styles from './Sidebar.module.scss';

interface NavItem {
  to: string;
  label: string;
  Icon: React.ComponentType<{ className?: string }>;
}

const navItems: NavItem[] = [
  { to: '/', label: 'Dashboard', Icon: FaHome },
  { to: '/rooms', label: 'Rooms', Icon: FaBed },
  { to: '/housekeeping', label: 'Housekeeping', Icon: FaBroom },
  // Example additional items:
  // { to: '/analytics', label: 'Analytics', Icon: FaChartPie },
  // { to: '/users', label: 'Users', Icon: FaUsers },
  // { to: '/settings', label: 'Settings', Icon: FaCog },
];

const Sidebar: React.FC = () => {
  const [collapsed, setCollapsed] = useState(false);

  const toggleSidebar = () => setCollapsed(prev => !prev);

  return (
    <aside
      className={`${styles.sidebar} ${collapsed ? styles.collapsed : ''}`}
    >
      {/* Logo / Branding */}
      <div className={styles.logoSection}>
        {/* You can swap in your own logo/icon here */}
        <div className={styles.logoIcon}> {/* optional icon */}
          {/* e.g. <YourLogoIcon /> */}
          🏨
        </div>
        {!collapsed && <span className={styles.logoText}>Lodgify</span>}
      </div>

      {/* Navigation Links */}
      <nav className={styles.nav}>
        {navItems.map(({ to, label, Icon }) => (
          <NavLink
            to={to}
            key={to}
            className={({ isActive }) =>
              [
                styles.link,
                isActive ? styles.active : '',
                collapsed ? styles.collapsedLink : '',
              ]
                .filter(Boolean)
                .join(' ')
            }
            title={collapsed ? label : undefined} // show tooltip text when collapsed
          >
            <Icon className={styles.linkIcon} />
            {!collapsed && <span className={styles.linkText}>{label}</span>}
          </NavLink>
        ))}
      </nav>

      {/* Spacer grows to push toggle to bottom */}
      <div className={styles.spacer} />

      {/* Collapse / Expand Toggle */}
      <button
        className={styles.toggleBtn}
        onClick={toggleSidebar}
        aria-label={collapsed ? 'Expand sidebar' : 'Collapse sidebar'}
      >
        {collapsed ? <FaChevronRight /> : <FaChevronLeft />}
      </button>
    </aside>
  );
};

export default Sidebar;
