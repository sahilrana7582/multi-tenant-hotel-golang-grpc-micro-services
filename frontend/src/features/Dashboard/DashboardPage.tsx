// src/features/Dashboard/Dashboard.tsx
import React, { useState } from "react";
import StatCard from "../../components/StatsCard/StatCard";
import { MdOutlineDateRange, MdLogin, MdLogout, MdAttachMoney } from "react-icons/md";
import styles from "./Dashboard.module.scss";

const Dashboard: React.FC = () => {

    const [active, setActive] = useState("bookings");

  return (
    <div className={styles.dashboard}>
      <div className={styles.statGrid}>
            <StatCard
                title="New Bookings"
                value={840}
                change="+8.70%"
                positive
                icon={<MdOutlineDateRange />}
                active={active === "bookings"}
                onClick={() => setActive("bookings")}
            />
            
            <StatCard
                title="Check-In"
                value={231}
                change="+3.56%"
                positive
                icon={<MdLogin />}
                active={active === "checkin"}
                onClick={() => setActive("checkin")}
            />            
            
            <StatCard
                title="Check-Out"
                value={231}
                change="-1.56%"
                positive={false}
                icon={<MdLogout />}
                active={active === "checkout"}
                onClick={() => setActive("checkout")}
            />   

            <StatCard
                title="Check-In"
                value={231}
                change="-0.56%"
                positive={false}
                icon={< MdAttachMoney />}
                active={active === "revenue"}
                onClick={() => setActive("revenue")}
            />
      </div>
    </div>
  );
};

export default Dashboard;
