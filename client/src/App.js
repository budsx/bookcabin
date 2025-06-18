import React, { useState, useEffect } from 'react';
import './App.css';
import SeatMap from './components/SeatMap';
import { seatMapService } from './services/seatMapService';

function App() {
  const [seatMapData, setSeatMapData] = useState(null);
  const [loading, setLoading] = useState(true);
  const [flightId, setFlightId] = useState(312); // HARDCODED FLIGHT ID FOR NOW

  useEffect(() => {
    const fetchSeatMapData = async () => {
      try {
        const seatMapData = await seatMapService.fetchSeatMap(flightId);
        console.log(seatMapData);

        setSeatMapData(seatMapData);
        setLoading(false);
      } catch (error) {
        console.error('Error fetching seat map data:', error);
        setLoading(false);
      }
    };

    fetchSeatMapData();
  }, [flightId]);

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
        {seatMapData && <SeatMap data={seatMapData} flightId={flightId} />}
      </main>
    </div>
  );
}

export default App;
