import {  type ReactNode } from "react";
import styles from "./StatCard.module.scss"



interface StatCardProps {
    title: string;
    value: number | string;
    change?: string;
    icon?: ReactNode;
    positive?: boolean;
    active?: boolean;
    onClick?: () => void;
}


const StatCard: React.FC<StatCardProps> = ({title, value, change, icon, positive, active, onClick}) => {


    return (
        <div
            className={`${styles.card} ${active ? styles.active : ""}`}
            onClick={onClick}
        >
        <div className={styles.topRow}>
          <span className={styles.title}>{title}</span>
          {icon && <span className={styles.icon}>{icon}</span>}
        </div>
        <div className={styles.value}>{value}</div>
        {change && (
            <div className={styles.changeWrapper}>
                <div className={`${styles.change} ${positive ? styles.positive : styles.negative}`}>
                {change}
                </div>
                <span className={styles.subText}>from last week</span>
            </div>
        )}

      </div>
    );
}
export default StatCard;