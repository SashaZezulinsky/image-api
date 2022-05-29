# image-api
image-api is a REST API service to get, upload and update images

# Architecture
We use Uncle Bob The Clean Arcirecture to represent service logic

# Storage
The best way to store images is to store them on a filesystem
For this purpose we use MongoDB GridFS
GridFS allows us to store images > 16MB
GridFS bucket has two collections: `fs.files` to storage metadata and `fs.chunks` to store files


# Limitations
Know limitation is updating existing image
GridFS doesn't have a function to update existing image
That is why update for us is delete + insert

# Build
`make build`

# Test
`make test`

# Run locally
`make local`
