import React from 'react';
import RoomCard from './RoomCard';
import styles from './RoomList.module.scss';

interface Room {
  id: number;
  name: string;
  category: string;
}

const RoomList: React.FC<{ rooms: Room[] }> = ({ rooms }) => {
  if (!rooms.length) return <p className={styles.empty}>No rooms available in this category.</p>;

  return (
    <div className={styles.grid}>
      {rooms.map((room) => (
        <RoomCard key={room.id} room={room} />
      ))}
    </div>
  );
};

export default RoomList;