# BookCabin Assessment


## TODO
- [√] Database Design
- [√] Insert Data
- [√] Create Endpoint Get SeatMap
- [√] Create Endpoint to Select SeatMap
- [√] Create Frontend
- [√] Integrate


### Prerequisites

- Docker
- Docker Compose

### Running the Application

1. **Clone the repository** (if not already done):
   ```bash
   git clone <repository-url>
   cd bookcabin
   ```

2. **Start all services**:
   ```bash
   docker-compose up -d
   ```

3. **View logs** (optional):
   ```bash
   # View all services logs
   docker-compose logs -f
   
   # View specific service logs
   docker-compose logs -f backend
   docker-compose logs -f frontend
   docker-compose logs -f database
   ```

4. **Access the application**:
   - Frontend: http://localhost
   - Backend API: http://localhost:8080
   - Database: localhost:3306

### Stopping the Application

```bash
# Stop all services
docker-compose down

# Stop and remove volumes (this will delete database data)
docker-compose down -v
```

### Database Initialization

The MySQL database is automatically initialized with:
- `dump.sql` - Schema and initial data

### Service Details

- **Database**: MySQL 8.0 on port 3306
- **Backend**: Go application on port 8080
- **Frontend**: React application served by Nginx on port 80

### Development

To rebuild services after code changes:
```bash
# Rebuild and restart specific service
docker-compose up -d --build backend

# Rebuild all services
docker-compose up -d --build
```

