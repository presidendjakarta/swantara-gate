# Swantara Gate API Testing Tools

## 🚀 Quick Start

### Run All Tests

```bash
# Using default URL (http://localhost:8081)
go run tools/api_runner.go

# Using custom URL
go run tools/api_runner.go http://your-server:port
```

## 📁 Generated Files

After running tests, you'll get:

1. **API_DOCUMENTATION.md** - Complete API documentation with:
   - All endpoints tested
   - Request bodies (if any)
   - Response bodies with status codes
   - Response times
   - Error messages (if any)

2. **test_results.json** - Machine-readable test results:
   - Total tests run
   - Pass/fail counts
   - Detailed results for each endpoint
   - Duration per request

3. **TEST_REPORT.md** - Human-readable summary:
   - Test statistics
   - Passed tests list
   - Failed tests analysis
   - Recommendations

## 🔧 Configuration

### Default Settings

- **Base URL:** `http://localhost:8081`
- **Username:** `admin`
- **Password:** `admin1324`

### Custom Credentials

Edit the test code in `tools/api_runner.go`:

```go
// Line ~133
suite.makeRequest("POST", "/api/admin/auth/login", map[string]interface{}{
    "username": "your-username",
    "password": "your-password",
}, nil)
```

## 📊 Test Coverage

Currently testing **47 endpoints** across **11 categories**:

| Category | Endpoints | Status |
|----------|-----------|--------|
| Health Check | 1 | ✅ |
| Authentication | 5 | ✅ |
| Configuration | 1 | ✅ |
| Users | 5 | ⚠️ |
| API Consumers | 5 | ⚠️ |
| Hosts | 5 | ⚠️ |
| Virtual Hosts | 5 | ⚠️ |
| Routes (Virtual Directories) | 5 | ⚠️ |
| Rate Limits | 5 | ⚠️ |
| CORS Configs | 5 | ⚠️ |
| Circuit Breakers | 5 | ⚠️ |

✅ = Working  
⚠️ = Needs database setup

## 🐛 Troubleshooting

### Tests Failing?

1. **Check if server is running:**
   ```bash
   curl http://localhost:8081/api/health
   ```

2. **Check database migrations:**
   ```bash
   swantara-gate migrate
   ```

3. **Verify credentials:**
   - Username: `admin`
   - Password: `admin1324`

4. **Check database tables:**
   ```bash
   sqlite3 data/database.db ".tables"
   ```

### Token Issues

If authentication fails:
- Check if user exists in database
- Verify password hash is correct
- Check JWT secret in `.env`

## 📝 Adding New Tests

To add tests for new endpoints:

```go
// In RunAllTests() method
fmt.Println("🆕 X. Your Category")

// Create
suite.makeRequest("POST", "/api/admin/your-endpoint", map[string]interface{}{
    "field1": "value1",
    "field2": "value2",
}, nil)

// Read All
suite.makeRequest("GET", "/api/admin/your-endpoint", nil, nil)

// Read One
suite.makeRequest("GET", "/api/admin/your-endpoint/1", nil, nil)

// Update
suite.makeRequest("PUT", "/api/admin/your-endpoint/1", map[string]interface{}{
    "field1": "updated",
}, nil)

// Delete
suite.makeRequest("DELETE", "/api/admin/your-endpoint/1", nil, nil)
```

## 🎯 Best Practices

1. **Run tests after migrations** - Ensure database schema is up to date
2. **Check test_results.json** - For detailed error messages
3. **Review API_DOCUMENTATION.md** - For API reference
4. **Test in order** - Some endpoints depend on others (e.g., need host_id before creating vhost)

## 📈 Future Enhancements

- [ ] Add more endpoints (SSL, ACME, Maintenance, etc.)
- [ ] Parallel test execution
- [ ] Data cleanup after tests
- [ ] HTML report generation
- [ ] CI/CD integration
- [ ] Load testing mode
- [ ] WebSocket testing

---

**Need help?** Check `API_DOCUMENTATION.md` for detailed API reference.
