import React, { useState } from 'react';
import './SeatMap.css';
import { seatMapService } from '../services/seatMapService';

const SeatMap = ({ data, flightId }) => {
  const [selectedSeat, setSelectedSeat] = useState(null);
  const [showSeatInfo, setShowSeatInfo] = useState(false);
  const [seatInfo, setSeatInfo] = useState(null);
  const [isSelecting, setIsSelecting] = useState(false);
  const [selectionError, setSelectionError] = useState(null);

  const seatMapData = data?.seatsItineraryParts?.[0]?.segmentSeatMaps?.[0]?.passengerSeatMaps?.[0]?.seatMap;
  const cabin = seatMapData?.cabins?.[0];

  if (!cabin) {
    return <div className="error">No seat map data available</div>;
  }

  const handleSeatClick = (seat) => {
    if (seat.storefrontSlotCode === 'SEAT') {
      setSeatInfo(seat);
      setShowSeatInfo(true);
      
      if (seat.available) {
        setSelectedSeat(seat.code === selectedSeat ? null : seat.code);
      }
    }
  };

  const formatPrice = (priceInfo) => {
    if (!priceInfo?.alternatives?.[0]?.[0]) return 'Free';
    const price = priceInfo.alternatives[0][0];
    return `${price.currency} ${price.amount.toFixed(2)}`;
  };

  const closeSeatInfo = () => {
    setShowSeatInfo(false);
    setSeatInfo(null);
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
    const isClickable = seat.storefrontSlotCode === 'SEAT';
    
    return (
      <div
        key={seatIndex}
        className={getSeatClass(seat)}
        onClick={() => isClickable && handleSeatClick(seat)}
        title={seat.code || seat.storefrontSlotCode}
      >
        {seat.storefrontSlotCode === 'SEAT' ? (
          <div className="seat-container">
            <div className="seat-back"></div>
            <div className="seat-cushion"></div>
            <span className="seat-label">{seat.code}</span>
          </div>
        ) : seat.storefrontSlotCode === 'AISLE' ? (
          <span className="aisle-label"></span>
        ) : null}
      </div>
    );
  };

  const handleConfirmSelection = async () => {
    if (!selectedSeat || !flightId) {
      alert('Unable to confirm seat selection');
      return;
    }

    setIsSelecting(true);
    setSelectionError(null);

    try {
      const payload = {
        flightId: flightId,
        seatCode: selectedSeat,
        passengerInfo: {
          firstName: 'Rutwik',
          lastName: 'Budhekar',
        },
      };

      const result = await seatMapService.selectSeat(payload);
      
      alert(`Seat ${selectedSeat} successfully selected!`);
      console.log('Seat selection result:', result);
    } catch (error) {
      setSelectionError(`Failed to select seat: ${error.message}`);
      console.error('Seat selection error:', error);
    } finally {
      setIsSelecting(false);
    }
  };

  return (
    <div className="seat-map">
      <div className="aircraft-info">
        <h3>Aircraft: {seatMapData.aircraft}</h3>
        <p>Deck: {cabin.deck}</p>
      </div>

      <div className="seat-legend">
        <div className="legend-item">
          <div className="seat available">
            <div className="seat-container">
              <div className="seat-back"></div>
              <div className="seat-cushion"></div>
            </div>
          </div>
          <span>Available</span>
        </div>
        <div className="legend-item">
          <div className="seat occupied">
            <div className="seat-container">
              <div className="seat-back"></div>
              <div className="seat-cushion"></div>
            </div>
          </div>
          <span>Occupied</span>
        </div>
        <div className="legend-item">
          <div className="seat selected">
            <div className="seat-container">
              <div className="seat-back"></div>
              <div className="seat-cushion"></div>
            </div>
          </div>
          <span>Selected</span>
        </div>
        <div className="legend-item">
          <div className="seat window">
            <div className="seat-container">
              <div className="seat-back"></div>
              <div className="seat-cushion"></div>
            </div>
          </div>
          <span>Window</span>
        </div>
      </div>

      {selectedSeat && (
        <div className="selected-seat-info">
          <h4>Selected Seat: {selectedSeat}</h4>
          {selectionError && (
            <div className="error-message">
              {selectionError}
            </div>
          )}
          <button 
            className="confirm-btn"
            onClick={handleConfirmSelection}
            disabled={isSelecting}
          >
            {isSelecting ? 'Confirming...' : 'Confirm Selection'}
          </button>
          <button 
            className="cancel-btn"
            onClick={() => {
              setSelectedSeat(null);
              setSelectionError(null);
            }}
            disabled={isSelecting}
          >
            Clear Selection
          </button>
        </div>
      )}

      {/* Seat Info Modal */}
      {showSeatInfo && seatInfo && (
        <div className="seat-info-modal">
          <div className="modal-overlay" onClick={closeSeatInfo}></div>
          <div className="modal-content">
            <div className="modal-header">
              <h3>Seat Information</h3>
              <button className="close-btn" onClick={closeSeatInfo}>×</button>
            </div>
            
            <div className="modal-body">
              <div className="seat-detail-section">
                <h4>📍 Seat Details</h4>
                <div className="detail-grid">
                  <div className="detail-item">
                    <span className="label">Seat Code:</span>
                    <span className="value">{seatInfo.code || 'N/A'}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Status:</span>
                    <span className={`value status ${seatInfo.available ? 'available' : 'occupied'}`}>
                      {seatInfo.available ? '✅ Available' : '❌ Occupied'}
                    </span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Free of Charge:</span>
                    <span className="value">{seatInfo.freeOfCharge ? '✅ Yes' : '❌ No'}</span>
                  </div>
                  <div className="detail-item">
                    <span className="label">Entitled:</span>
                    <span className="value">{seatInfo.entitled ? '✅ Yes' : '❌ No'}</span>
                  </div>
                </div>
              </div>

              {seatInfo.prices && (
                <div className="seat-detail-section">
                  <h4>💰 Pricing</h4>
                  <div className="detail-grid">
                    <div className="detail-item">
                      <span className="label">Base Price:</span>
                      <span className="value price">{formatPrice(seatInfo.prices)}</span>
                    </div>
                    {seatInfo.taxes && (
                      <div className="detail-item">
                        <span className="label">Taxes:</span>
                        <span className="value price">{formatPrice(seatInfo.taxes)}</span>
                      </div>
                    )}
                    {seatInfo.total && (
                      <div className="detail-item">
                        <span className="label">Total Price:</span>
                        <span className="value price total">{formatPrice(seatInfo.total)}</span>
                      </div>
                    )}
                  </div>
                </div>
              )}

              {seatInfo.seatCharacteristics && seatInfo.seatCharacteristics.length > 0 && (
                <div className="seat-detail-section">
                  <h4>🎯 Seat Features</h4>
                  <div className="characteristics">
                    {seatInfo.seatCharacteristics.map((char, index) => (
                      <span key={index} className="characteristic-tag">
                        {char === 'W' && '🪟 Window'}
                        {char === 'A' && '🚶 Aisle'}
                        {char === 'EXIT' && '🚪 Exit Row'}
                        {char === 'EXTRA_LEGROOM' && '📏 Extra Legroom'}
                        {char === 'PREMIUM' && '⭐ Premium'}
                        {char !== 'W' && char !== 'A' && char !== 'EXIT' && char !== 'EXTRA_LEGROOM' && char !== 'PREMIUM' && char}
                      </span>
                    ))}
                  </div>
                </div>
              )}

              {seatInfo.designations && seatInfo.designations.length > 0 && (
                <div className="seat-detail-section">
                  <h4>🏷️ Designations</h4>
                  <div className="designations">
                    {seatInfo.designations.map((designation, index) => (
                      <span key={index} className="designation-tag">{designation}</span>
                    ))}
                  </div>
                </div>
              )}

              {seatInfo.limitations && seatInfo.limitations.length > 0 && (
                <div className="seat-detail-section">
                  <h4>⚠️ Limitations</h4>
                  <div className="limitations">
                    {seatInfo.limitations.map((limitation, index) => (
                      <div key={index} className="limitation-item">
                        <span className="limitation-text">{limitation}</span>
                      </div>
                    ))}
                  </div>
                </div>
              )}
            </div>

            <div className="modal-footer">
              {seatInfo.available && (
                <button 
                  className="select-seat-btn"
                  onClick={() => {
                    setSelectedSeat(seatInfo.code);
                    closeSeatInfo();
                  }}
                >
                  Select This Seat
                </button>
              )}
              <button className="close-modal-btn" onClick={closeSeatInfo}>
                Close
              </button>
            </div>
          </div>
        </div>
      )}

      <div className="cabin">
        <div className="cabin-header">
          <h4>Main Cabin</h4>
          <div className="column-headers">
            {cabin.seatColumns.map((column, index) => (
              <div key={index} className="column-header">
                {column !== 'LEFT_SIDE' && column !== 'RIGHT_SIDE' && column !== 'AISLE' ? column :   ''}
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