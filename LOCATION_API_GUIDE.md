# Bangladesh Location Hierarchy - Postman Guide

## 📍 Complete Hierarchy File

The complete hierarchy with all divisions → districts → upazilas is saved in:

```
location_hierarchy.json
```

## 🌐 API Endpoints

### Method 1: Fetch Complete Hierarchy (Using Script)

```bash
# Run the Python script to fetch all data:
python3 scripts/get_full_hierarchy.py

# Output: location_hierarchy.json with nested structure:
# divisions → districts → upazilas
```

### Method 2: Query Individual Levels in Postman

#### 1️⃣ Get All Divisions

```
GET http://localhost:8080/api/v1/locations/divisions
```

**Response Example:**

```json
{
  "divisions": [
    {"id": "1", "nameEn": "Barishal", "nameBn": ""},
    {"id": "5", "nameEn": "Mymensingh", "nameBn": ""},
    ...
  ]
}
```

---

#### 2️⃣ Get Districts Under a Division

```
GET http://localhost:8080/api/v1/locations/divisions/{division_id}/districts
```

**Example - Mymensingh Districts:**

```
GET http://localhost:8080/api/v1/locations/divisions/5/districts
```

**Response:**

```json
{
  "districts": [
    { "id": "41", "divisionId": "5", "nameEn": "Jamalpur", "nameBn": "" },
    { "id": "42", "divisionId": "5", "nameEn": "Mymensingh", "nameBn": "" },
    { "id": "43", "divisionId": "5", "nameEn": "Netrokona", "nameBn": "" },
    { "id": "44", "divisionId": "5", "nameEn": "Sherpur", "nameBn": "" }
  ]
}
```

---

#### 3️⃣ Get Upazilas Under a District

```
GET http://localhost:8080/api/v1/locations/districts/{district_id}/upazilas
```

**Example - Jamalpur Upazilas:**

```
GET http://localhost:8080/api/v1/locations/districts/41/upazilas
```

**Response:**

```json
{
  "upazilas": [
    { "id": "354", "districtId": "41", "nameEn": "Baksiganj", "nameBn": "" },
    { "id": "355", "districtId": "41", "nameEn": "Dewanganj", "nameBn": "" },
    { "id": "356", "districtId": "41", "nameEn": "Islampur", "nameBn": "" },
    {
      "id": "357",
      "districtId": "41",
      "nameEn": "Jamalpur Sadar",
      "nameBn": ""
    },
    { "id": "360", "districtId": "41", "nameEn": "Madarganj", "nameBn": "" },
    { "id": "358", "districtId": "41", "nameEn": "Melandah", "nameBn": "" },
    { "id": "359", "districtId": "41", "nameEn": "Sharishabari", "nameBn": "" }
  ]
}
```

---

## 📊 Complete Hierarchy Summary

**Total Locations:**

- ✅ **8 Divisions**
- ✅ **64 Districts** (exactly as per Bangladesh official count)
- ✅ **557 Upazilas**

**Hierarchy Breakdown by Division:**

| Division   | Districts | Upazilas |
| ---------- | --------- | -------- |
| Barishal   | 6         | 42       |
| Chattogram | 11        | 113      |
| Dhaka      | 13        | 136      |
| Khulna     | 10        | 64       |
| Mymensingh | 4         | 35       |
| Rajshahi   | 8         | 71       |
| Rangpur    | 8         | 58       |
| Sylhet     | 4         | 38       |

---

## 🔍 Finding Specific Locations

### Example: Madarganj in Jamalpur, Mymensingh

**Step 1:** Get Mymensingh Division

```
GET http://localhost:8080/api/v1/locations/divisions
→ Find: {"id": "5", "nameEn": "Mymensingh"}
```

**Step 2:** Get Districts in Mymensingh

```
GET http://localhost:8080/api/v1/locations/divisions/5/districts
→ Find: {"id": "41", "nameEn": "Jamalpur"}
```

**Step 3:** Get Upazilas in Jamalpur

```
GET http://localhost:8080/api/v1/locations/districts/41/upazilas
→ Find: {"id": "360", "nameEn": "Madarganj"}
```

**Complete Path:**

```
Mymensingh (Division ID: 5)
  └─ Jamalpur (District ID: 41)
      └─ Madarganj (Upazila ID: 360)
```

---

## 📁 View Complete Hierarchy File

```bash
# View entire hierarchy
cat location_hierarchy.json | jq '.'

# Find specific division
cat location_hierarchy.json | jq '.divisions[] | select(.nameEn=="Mymensingh")'

# Get all districts in Dhaka
cat location_hierarchy.json | jq '.divisions[] | select(.nameEn=="Dhaka") | .districts[].nameEn'

# Find Madarganj
cat location_hierarchy.json | jq '.divisions[].districts[].upazilas[] | select(.nameEn=="Madarganj")'
```

---

## 🚀 Server Information

- **REST API**: http://localhost:8080
- **gRPC**: localhost:50051
- **PostgreSQL**: localhost:5433
- **Redis**: localhost:6380

---

## ⚡ Quick Commands

```bash
# Regenerate complete hierarchy
python3 scripts/get_full_hierarchy.py

# Check server status
ps aux | grep matrimonial-service

# View server logs
tail -f server.log

# Test API
curl http://localhost:8080/api/v1/locations/divisions | jq '.'
```
