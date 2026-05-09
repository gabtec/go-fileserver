# Go Fileserver

A simple static file server built in Go and designed for containerization.

## Usage

### Environment Variables

- `SHARE_FOLDER`: Path to the folder to serve files from.
- `PORT`: Port to listen on (default: 8080).

### Local Execution

```bash
export SHARE_FOLDER=./public
go run .
```

### Docker

1. Build the image:

   ```bash
   docker build -t go-fileserver .
   ```

2. Run the container, mounting your local folder:
   ```bash
   docker run -p 8080:8080 -e SHARE_FOLDER=/data -v /path/to/your/files:/data go-fileserver
   ```
