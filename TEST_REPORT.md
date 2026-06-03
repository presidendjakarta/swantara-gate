# Swantara Gate API Test Report

**Test Date:** 2026-05-15  
**Base URL:** http://localhost:8081  
**Credentials:** admin / admin1324

---

## 📊 Test Summary

| Metric | Count |
|--------|-------|
| ✅ Passed | 6 |
| ❌ Failed | 41 |
| 📝 Total | 47 |
| ⏱️ Total Duration | 125ms |

---

## ✅ Passed Tests

1. **GET /api/health** - 8ms
2. **POST /api/admin/auth/login** - 52ms
3. **GET /api/admin/auth/me** - 1ms
4. **POST /api/admin/auth/refresh** - 6ms
5. **POST /api/admin/auth/logout** - 2ms
6. **POST /api/admin/auth/login** (re-auth) - 2ms

---

## ❌ Failed Tests (41 endpoints)

Most failures are likely due to:
- Database tables not yet created
- Foreign key constraints (need to create parent records first)
- Missing required fields in request body

### Common Issues:

1. **Users CRUD** - May need super_admin role to create users
2. **Consumers CRUD** - Table might not exist
3. **Hosts CRUD** - Table might not exist  
4. **Virtual Hosts CRUD** - Requires host_id first
5. **Virtual Directories** - Requires virtual_host_id first
6. **Rate Limits** - Requires virtual_directory_id first
7. **CORS Configs** - Requires virtual_directory_id first
8. **Circuit Breakers** - Requires virtual_directory_id first

---

## 🎯 Recommendations

1. **Run database migrations first:**
   ```bash
   swantara-gate migrate
   ```

2. **Check if tables exist:**
   ```sql
   .tables
   ```

3. **Test with correct ID sequences:**
   - Create Host first → Get host_id
   - Create VirtualHost with host_id → Get vhost_id
   - Create VirtualDirectory with vhost_id → Get vdir_id
   - Then test Rate Limits, CORS, etc. with vdir_id

---

## 📁 Generated Files

- **API_DOCUMENTATION.md** - Full API documentation with request/response examples
- **test_results.json** - Detailed test results in JSON format
- **TEST_REPORT.md** - This file

---

## 🚀 How to Re-run Tests

```bash
# Default (localhost:8081)
go run tools/api_runner.go

# Custom URL
go run tools/api_runner.go http://your-server:port
```

---

## 📝 Notes

- Tests run sequentially and depend on previous successful creates
- Tokens are automatically extracted and reused after login
- Each test logs request/response for documentation generation
- Failed tests still generate documentation with error details
