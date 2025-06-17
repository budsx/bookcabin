# BookCabin Docker Deployment Guide

This guide explains how to deploy the BookCabin application using Docker and Docker Compose.

## Prerequisites

- Docker (v20.10.0 or higher)
- Docker Compose (v2.0.0 or higher)

## Quick Start

1. **Clone the repository** (if not already done):
   ```bash
   git clone <repository-url>
   cd bookcabin
   ```

2. **Start all services**:
   ```bash
   docker-compose up -d
   ```

3. **Access the application**:
   - Frontend: http://localhost
   - Backend API: http://localhost:8080
   - Database: localhost:3306

## Services

### Frontend (React)
- **Container**: `bookcabin-frontend`
- **Port**: 80
- **Technology**: React 19 with Nginx
- **Features**: 
  - Production-optimized build
  - Gzip compression
  - API proxy to backend

### Backend (Go)
- **Container**: `bookcabin-backend`
- **Port**: 8080
- **Technology**: Go 1.21 with Chi router
- **Features**:
  - Multi-stage build for smaller image
  - Health checks via database connection
  - Environment-based configuration

### Database (MySQL)
- **Container**: `bookcabin-db`
- **Port**: 3306
- **Technology**: MySQL 8.0
- **Features**:
  - Automatic schema initialization
  - Persistent data storage
  - Health checks

## Environment Configuration

The application uses environment variables for configuration. The default values in `docker-compose.yml` are:

```env
SERVICE_PORT=8080
DB_HOST=database
DB_PORT=3306
DB_USER=bookcabin_user
DB_PASSWORD=bookcabin_password
DB_NAME=bookcabin
```

For production, create a `.env` file and override these values:

```env
# Copy from .env.example and modify as needed
SERVICE_PORT=8080
DB_HOST=database
DB_PORT=3306
DB_USER=your_secure_user
DB_PASSWORD=your_secure_password
DB_NAME=bookcabin
```

## Commands

### Development

```bash
# Start all services
docker-compose up

# Start in background
docker-compose up -d

# View logs
docker-compose logs -f

# View logs for specific service
docker-compose logs -f backend
```

### Management

```bash
# Stop all services
docker-compose down

# Stop and remove volumes (WARNING: deletes database data)
docker-compose down -v

# Rebuild images
docker-compose build

# Rebuild and start
docker-compose up --build
```

### Database

```bash
# Access database shell
docker-compose exec database mysql -u bookcabin_user -p bookcabin

# Import additional SQL
docker-compose exec database mysql -u bookcabin_user -p bookcabin < your-file.sql

# Backup database
docker-compose exec database mysqldump -u bookcabin_user -p bookcabin > backup.sql
```

## Troubleshooting

### Common Issues

1. **Port conflicts**: If ports 80, 8080, or 3306 are in use, modify the ports in `docker-compose.yml`
2. **Database connection issues**: Wait for the database health check to pass before the backend starts
3. **Build failures**: Ensure Docker has enough memory allocated (4GB recommended)

### Logs

```bash
# Check all service logs
docker-compose logs

# Check specific service
docker-compose logs backend
docker-compose logs frontend
docker-compose logs database
```

### Health Checks

```bash
# Check service status
docker-compose ps

# Manual health check
curl http://localhost:8080/api/v1/seat-map
```

## Production Considerations

1. **Security**:
   - Change default database passwords
   - Use secrets management for sensitive data
   - Run behind reverse proxy (Nginx/Traefik)
   - Enable HTTPS

2. **Performance**:
   - Increase database resources for production workloads
   - Configure proper database indexes
   - Consider database connection pooling

3. **Monitoring**:
   - Add health check endpoints
   - Configure log aggregation
   - Set up monitoring and alerting

4. **Backup**:
   - Regular database backups
   - Test backup restoration procedures

## Network Architecture

```
Internet → Frontend (Port 80) → Backend (Port 8080) → Database (Port 3306)
                ↓
         Static React App
                ↓
         /api/* → Proxy to Backend
```

The frontend serves the React application and proxies API requests to the backend service. All services communicate through a private Docker network.

## File Structure

```
bookcabin/
├── Dockerfile              # Backend Dockerfile
├── docker-compose.yml      # Orchestration file
├── .dockerignore           # Backend build exclusions
├── client/
│   ├── Dockerfile          # Frontend Dockerfile
│   ├── .dockerignore       # Frontend build exclusions
│   └── nginx.conf          # Nginx configuration
├── database.sql            # Database schema
└── table.sql              # Initial data
``` 