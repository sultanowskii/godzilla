# Go, Dzilla!

Simple URL shortener.

Current state:
- Uses handmade token generation, therefore there might be collisions.
- Redis as a storage
- No configuration
- Minimalistic web interface

TODO:
- [ ] Use PostgreSQL as a storage
- [ ] Better token generation with collision prevention
- [ ] URL validation
- [ ] URL unification (if necessary)
- [ ] Make it possible to create custom tokens (like godzilla:12345/Github -> github.com)
- [ ] Configuration
