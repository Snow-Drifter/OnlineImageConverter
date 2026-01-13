# Private Image Converter

A web-based image converter that runs entirely in the browser using WebAssembly. Images never leave your the user's device. Private and cheap to host.

## Usage

Select an image and provide an output filename. The file extension determines the target format.

## Installation

1. Configure your hostname in `image-converter.service`:

```ini
Environment=IMAGE_CONVERTER_HOSTS=yourdomain.com,www.yourdomain.com
```

2. Run the install script as root:

```bash
./install.sh
```

3. Start the service:

```bash
systemctl start image-converter
```

## Configuration

`IMAGE_CONVERTER_HOSTS` — Comma-separated list of allowed hostnames for TLS certificate generation via Let's Encrypt.
