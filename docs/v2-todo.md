## feature for V2

- TTL support for preview instance
- merge preview instance with main instance
- transfer prod instance to preview instance
- Restart Current action: force-update the swarm service with the Current Deployment image and latest env/config, without creating a new Deployment record.
- Retry Candidate action: retry the latest non-current Deployment only when it already has a built image; apply that image as a new attempt without rebuilding from source.
- Initial deploy safety cleanup: align first deploy with the Current Deployment definition by keeping `is_current=false` until the first image is successfully created and applied.
