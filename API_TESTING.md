# API Testing Guide - Matrimonial Service

## Quick Start

1. **Start the services:**

   ```bash
   make docker-up
   # OR manually:
   docker-compose -f deployments/docker/docker-compose.yml up -d
   ```

2. **Run database migrations (if not auto-run):**

   ```bash
   bash scripts/setup_db.sh
   ```

3. **Access the API:**
   - REST API: http://localhost:8080
   - gRPC: localhost:50051

---

## REST API Endpoints (Use in Postman)

### 1. Location Service

#### Get All Divisions

```
GET http://localhost:8080/api/v1/locations/divisions
```

**Response Example:**

```json
{
  "divisions": [
    {
      "id": "1",
      "name_en": "Dhaka",
      "name_bn": ""
    }
  ]
}
```

#### Get Districts by Division

```
GET http://localhost:8080/api/v1/locations/divisions/3/districts
```

**Path Parameter:**

- `division_id`: The ID of the division (e.g., 3 for Dhaka)

#### Get Upazilas by District

```
GET http://localhost:8080/api/v1/locations/districts/47/upazilas
```

**Path Parameter:**

- `district_id`: The ID of the district

---

### 2. Education Service

#### Get All Qualifications

```
GET http://localhost:8080/api/v1/education/qualifications
```

**Response Example:**

```json
{
  "qualifications": [
    {
      "id": "1",
      "name": "SSC",
      "common_name": "Secondary School Certificate",
      "level_id": "1",
      "stream_id": "1"
    }
  ]
}
```

---

### 3. Static Data Service

#### Get Religions

```
GET http://localhost:8080/api/v1/static/religions
```

**Response Example:**

```json
{
  "religions": [
    {
      "id": "1",
      "name": "Islam"
    },
    {
      "id": "2",
      "name": "Hinduism"
    }
  ]
}
```

#### Get Professions

```
GET http://localhost:8080/api/v1/static/professions
```

#### Get Marital Statuses

```
GET http://localhost:8080/api/v1/static/marital-statuses
```

---

## Postman Collection

### Import these as a collection:

1. **Create New Collection:** "Matrimonial Service API"

2. **Add Environment Variables:**

   - `base_url`: http://localhost:8080
   - `grpc_url`: localhost:50051

3. **Sample Requests:**

#### Request 1: Get Divisions

```
Method: GET
URL: {{base_url}}/api/v1/locations/divisions
```

#### Request 2: Get Districts for Dhaka Division (ID=3)

```
Method: GET
URL: {{base_url}}/api/v1/locations/divisions/3/districts
```

#### Request 3: Get All Qualifications

```
Method: GET
URL: {{base_url}}/api/v1/education/qualifications
```

#### Request 4: Get Religions

```
Method: GET
URL: {{base_url}}/api/v1/static/religions
```

#### Request 5: Get Professions

```
Method: GET
URL: {{base_url}}/api/v1/static/professions
```

#### Request 6: Get Marital Statuses

```
Method: GET
URL: {{base_url}}/api/v1/static/marital-statuses
```

---

## Testing with cURL

```bash
# Get all divisions
curl http://localhost:8080/api/v1/locations/divisions

# Get districts for division 3
curl http://localhost:8080/api/v1/locations/divisions/3/districts

# Get qualifications
curl http://localhost:8080/api/v1/education/qualifications

# Get religions
curl http://localhost:8080/api/v1/static/religions

# Get professions
curl http://localhost:8080/api/v1/static/professions

# Get marital statuses
curl http://localhost:8080/api/v1/static/marital-statuses
```

---

## Common Issues

### Database Connection Error

- Ensure PostgreSQL is running: `docker-compose ps`
- Check connection string in `config.yaml`

### Port Already in Use

- Change ports in `config.yaml` or `docker-compose.yml`
- Kill existing processes: `lsof -ti:8080 | xargs kill -9`

### Migration Errors

- Run: `bash scripts/setup_db.sh`
- Or manually: `psql -U postgres -d matrimonial_db -f internal/infrastructure/database/migrations/001_init.sql`

---

## Health Check

```bash
# Check if services are running
docker-compose ps

# Check API is responding
curl http://localhost:8080/api/v1/static/religions

# Check gRPC is running
grpcurl -plaintext localhost:50051 list
```
