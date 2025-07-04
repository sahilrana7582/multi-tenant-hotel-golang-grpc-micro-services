import React from 'react';
import styles from './RoomCard.module.scss';

interface Room {
  id: number;
  name: string;
  category: string;
}

const RoomCard: React.FC<{ room: Room }> = ({ room }) => {
  return (
    <div className={styles.card}>
      <h3>{room.name}</h3>
      <p>{room.category}</p>
    </div>
  );
};

export default RoomCard;