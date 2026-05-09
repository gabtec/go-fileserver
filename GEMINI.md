# Go Fileserver project

This project goal is to build a simple static file server using golang.

The application should be prepared to be containarized:

- it will receive a SHARE_FOLDER variable that will point to a folder where static folders will be stored.

Then using golang standard lib it should serve every file and folders (recursively), inside that folder.
The index page should be customized with:

- a title "Gabtec Shared Files" preceeded by a folder emoji.
- a info with the folder being served
- a footer with the app version, which should be in a separate version file

Finally add a github workflow to build and push the container image into ghcr.
The action should only be triggered when a new tag and release is committed.

Consider it will be a public github repo

## Status

- [x] Initialize Go module.
- [x] Implement `main.go` with `http.FileServer`.
- [x] Create multi-stage `Dockerfile`.
- [x] Add `README.md` with usage instructions.
- [x] Create GitHub workflow for GHCR deployment.
