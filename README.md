# Go, Dzilla!

Simple URL shortener.

Current state:
- Minimalistic web interface
- Uses handmade token generation, collisions could exist.
- Allows to create custom links
- Redis as a storage
- Works properly only when URLs are provided with protocol (e.g. `https://`)

TODO:
- [ ] Better token generation with collision prevention
- [ ] URL validation
- [ ] URL unification (if necessary)
- [ ] Accept URN (and just basically URI without protocol, maybe?)
- [ ] URL lifetime
- [x] Make it possible to create custom tokens (like godzilla:12345/Github -> github.com)
- [x] Basic configuration
