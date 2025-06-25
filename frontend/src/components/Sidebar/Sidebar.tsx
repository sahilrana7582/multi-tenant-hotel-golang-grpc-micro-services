import type React from "react";
import { useState } from "react";
import { BsFillHouseFill, BsPersonCircle, BsPersonVideo3 } from "react-icons/bs";
import { MdDashboard, MdFastfood } from "react-icons/md";
import { SiGooglemessages } from "react-icons/si";
import styles from './Sidebar.module.scss'
import { FaChevronLeft, FaChevronRight, FaHotel } from "react-icons/fa";
import { NavLink } from "react-router-dom";
import { FaListCheck } from "react-icons/fa6";


interface Navitem {
  to: string;
  label: string;
  Icon: React.ComponentType<{ className?: string }>;
}


const navItems: Navitem[] = [
  { to: '/', label: 'Dashboard', Icon: MdDashboard },
  { to: '/rooms', label: 'Rooms', Icon: BsFillHouseFill  },
  { to: '/messages', label: 'Messages', Icon: SiGooglemessages  },
  { to: '/housekeeping', label: 'Housekeeping', Icon: BsPersonCircle },
  { to: '/kitchen', label: 'IRD Orders', Icon: MdFastfood },
  { to: '/frontdesk', label: 'Frontdesk', Icon: BsPersonVideo3 },
  { to: '/login', label: 'Check In', Icon: FaListCheck },
];


const Sidebar: React.FC = () => {
  const [collapsed, setCollapsed] = useState<boolean>(false);
  const toggleSidebar = () => setCollapsed(prev => !prev);


  return (
    <aside className={`${styles.sidebar} ${collapsed ? styles.collapsed : ''}`}>

      {/* Logo / Branding */}
      <div className={styles.logoSection}>
        <div className={styles.logoIcon}>
          <FaHotel/>
        </div>

        {!collapsed && <span className={styles.logoText}>ButlerAI</span>}
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
  )
}


export default Sidebar;