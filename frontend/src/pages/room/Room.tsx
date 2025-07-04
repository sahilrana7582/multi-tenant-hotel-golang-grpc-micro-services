import React, { useState } from 'react';
import CategoryTabs from '../../components/Room/CategoryTabs';
import RoomList from '../../components/Room/RoomList';

const categories = ['Deluxe', 'Suite', 'Twin Bed', 'Luxury', 'Executive'];

const rooms = [
  { id: 1, name: 'Deluxe Twilight Room', category: 'Deluxe' },
  { id: 2, name: 'Suite Serenity View', category: 'Suite' },
  { id: 3, name: 'Twin Bed Galaxy', category: 'Twin Bed' },
  { id: 4, name: 'Luxury Moonlight Stay', category: 'Luxury' },
  { id: 5, name: 'Executive Skyline Room', category: 'Executive' },
  { id: 6, name: 'Deluxe Starfall Room', category: 'Deluxe' },
];

const Room: React.FC = () => {
  const [selectedCategory, setSelectedCategory] = useState(categories[0]);
  const filteredRooms = rooms.filter((room) => room.category === selectedCategory);

  return (
    <div style={{ backgroundColor: '#1a142d', minHeight: '100vh', padding: '2rem' }}>
      <h1 style={{ color: '#cdc6f6', fontFamily: 'Inter, sans-serif', marginBottom: '1.5rem' }}>
        Explore Our Rooms
      </h1>
      <CategoryTabs
        categories={categories}
        selectedCategory={selectedCategory}
        onSelect={setSelectedCategory}
      />
      <RoomList rooms={filteredRooms} />
    </div>
  );
};

export default Room;
