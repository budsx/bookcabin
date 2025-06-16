import React, { useState } from 'react';
import './SeatMap.css';

const SeatMap = ({ data }) => {
  const [selectedSeat, setSelectedSeat] = useState(null);

  // Extract seat map data from the nested structure
  const seatMapData = data?.seatsItineraryParts?.[0]?.segmentSeatMaps?.[0]?.passengerSeatMaps?.[0]?.seatMap;
  const cabin = seatMapData?.cabins?.[0];

  if (!cabin) {
    return <div className="error">No seat map data available</div>;
  }

  const handleSeatClick = (seat) => {
    if (seat.storefrontSlotCode === 'SEAT' && seat.available) {
      setSelectedSeat(seat.code === selectedSeat ? null : seat.code);
    }
  };

  const getSeatClass = (seat) => {
    const classes = ['seat'];
    
    switch (seat.storefrontSlotCode) {
      case 'SEAT':
        if (!seat.available) {
          classes.push('occupied');
        } else if (seat.code === selectedSeat) {
          classes.push('selected');
        } else {
          classes.push('available');
        }
        
        // Add characteristic classes
        if (seat.seatCharacteristics?.includes('W')) {
          classes.push('window');
        }
        if (seat.seatCharacteristics?.includes('A')) {
          classes.push('aisle');
        }
        if (seat.seatCharacteristics?.includes('EXIT')) {
          classes.push('exit');
        }
        break;
      case 'AISLE':
        classes.push('aisle-space');
        break;
      case 'BLANK':
        classes.push('blank');
        break;
      case 'WING':
        classes.push('wing');
        break;
      case 'BULKHEAD':
        classes.push('bulkhead');
        break;
      default:
        classes.push('other');
    }
    
    return classes.join(' ');
  };

  const renderSeat = (seat, seatIndex) => {
    const isClickable = seat.storefrontSlotCode === 'SEAT' && seat.available;
    
    return (
      <div
        key={seatIndex}
        className={getSeatClass(seat)}
        onClick={() => isClickable && handleSeatClick(seat)}
        title={seat.code || seat.storefrontSlotCode}
      >
        {seat.storefrontSlotCode === 'SEAT' ? (
          <span className="seat-label">{seat.code}</span>
        ) : seat.storefrontSlotCode === 'AISLE' ? (
          <span className="aisle-label">|</span>
        ) : null}
      </div>
    );
  };

  return (
    <div className="seat-map">
      <div className="aircraft-info">
        <h3>Aircraft: {seatMapData.aircraft}</h3>
        <p>Deck: {cabin.deck}</p>
      </div>

      <div className="seat-legend">
        <div className="legend-item">
          <div className="seat available"></div>
          <span>Available</span>
        </div>
        <div className="legend-item">
          <div className="seat occupied"></div>
          <span>Occupied</span>
        </div>
        <div className="legend-item">
          <div className="seat selected"></div>
          <span>Selected</span>
        </div>
        <div className="legend-item">
          <div className="seat window"></div>
          <span>Window</span>
        </div>
        <div className="legend-item">
          <div className="seat aisle"></div>
          <span>Aisle</span>
        </div>
      </div>

      {selectedSeat && (
        <div className="selected-seat-info">
          <h4>Selected Seat: {selectedSeat}</h4>
          <button 
            className="confirm-btn"
            onClick={() => alert(`Seat ${selectedSeat} confirmed!`)}
          >
            Confirm Selection
          </button>
          <button 
            className="cancel-btn"
            onClick={() => setSelectedSeat(null)}
          >
            Clear Selection
          </button>
        </div>
      )}

      <div className="cabin">
        <div className="cabin-header">
          <h4>Main Cabin</h4>
          <div className="column-headers">
            {cabin.seatColumns.map((column, index) => (
              <div key={index} className="column-header">
                {column !== 'LEFT_SIDE' && column !== 'RIGHT_SIDE' && column !== 'AISLE' ? column : ''}
              </div>
            ))}
          </div>
        </div>

        <div className="rows">
          {cabin.seatRows.map((row, rowIndex) => (
            <div key={rowIndex} className="row">
              <div className="row-number">{row.rowNumber}</div>
              <div className="seats">
                {row.seats.map((seat, seatIndex) => renderSeat(seat, seatIndex))}
              </div>
              <div className="row-number">{row.rowNumber}</div>
            </div>
          ))}
        </div>
      </div>
    </div>
  );
};

export default SeatMap; 