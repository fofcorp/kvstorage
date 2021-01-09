## kvstorage

Tier1 features:  
- [x] basic key-value storage
- [ ] http api server with handlers
- [ ] request logging via middleware
- [ ] expiring keys with garbage collection
- [ ] graceful shutdown
- [ ] numeric values with atomic increment/decrement (special operations - /incr, /decr)
- [ ] service options via config package (from os environment(prioritized) && yml)
- [ ] simple user-load generator for tests
- [ ] prometheus integration (add custom metrics: count ops by type, count hit/miss, measure ops timings, garbage collection stats)
- [ ] dockerized from scratch with healthcheck
- [ ] simple ci with tests (travis?) + githooks for test before commit
- [ ] docker-compose for app, prometheus, grafana (with dashboard)
- [ ] max size limiting (number of keys or size in bytes)
- [ ] generate global auth token on first start && save to fs and authentication with it
- [ ] endpoint with personalized auth-token generator (you will need global auth token anyway here)

Tier2 features:
- [ ] data encryption
- [ ] authorization on ops with RBAC (set/get/delete)
- [ ] client golang library (integration with API)
- [ ] rate limiting (global + personalized by header)
- [ ] data persistence
- [ ] distributed running with peers (leader election or gossip or ???)
