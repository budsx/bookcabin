const API_BASE_URL = process.env.REACT_APP_API_URL || '';

export const seatMapService = {
  async fetchSeatMap(flightId) {
    try {
      const response = await fetch(`${API_BASE_URL}/api/v1/seat-map?flightId=${flightId}`, {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      return data;
    } catch (error) {
      console.error('Error fetching seat map:', error);
      throw error;
    }
  },

  async selectSeat(payload) {
    try {
      const response = await fetch(`${API_BASE_URL}/api/v1/seat-map/select`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(payload),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      return await response.json();
    } catch (error) {
      console.error('Error selecting seat:', error);
      throw error;
    }
  },
};

export default seatMapService; 