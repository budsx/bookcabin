import React, { useState, useEffect } from 'react';
import './App.css';
import SeatMap from './components/SeatMap';
import { seatMapService } from './services/seatMapService';

function App() {
  const [seatMapData, setSeatMapData] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchSeatMapData = async () => {
      try {
        // HARDCODED FLIGHT ID FOR NOW
        const seatMapData = await seatMapService.fetchSeatMap(312);
        console.log(seatMapData);

        setSeatMapData(seatMapData);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching seat map data:', error);
        setLoading(false);
      }
    };

    fetchSeatMapData();
  }, []);

  if (loading) {
    return (
      <div className="App">
        <div className="loading">Loading seat map...</div>
      </div>
    );
  }

  return (
    <div className="App">
      <header className="App-header">
        <h1>Flight Seat Map</h1>
        <p>Select your preferred seat</p>
      </header>
      
      <main className="seat-map-container">
        {seatMapData && <SeatMap data={seatMapData} />}
      </main>
    </div>
  );
}

export default App;
