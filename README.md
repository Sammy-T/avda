# avda

Aegis Vault Desktop App is a desktop application for viewing one-time passwords generated from an [Aegis Authenticator](https://github.com/beemdevelopment/Aegis) vault backup or export file. 

The app is built with [Wails](https://github.com/wailsapp/wails) and [Svelte](https://github.com/sveltejs/svelte) and uses the [avdu](https://github.com/Sammy-T/avdu) module for OTP handling.

> [!NOTE]
> I built this app as a helper utility that can be used for convenient OTP access while on your personal desktop. While it does feature some of the basic functionality of Aegis Authenticator, it isn't intended to be a standalone 2FA app.

> [!NOTE]
> HOTP is not implemented due to syncing concerns.

![avda_screen_04](https://github.com/user-attachments/assets/7b997ab6-3cfc-4b28-90a2-5e0e4d262ca5)
![avda_screen_05](https://github.com/user-attachments/assets/b41efff8-f6ca-4aba-be3d-17de060a46bb)


## Getting started

Just run the binary matching your OS and open your local vault file.

> [!NOTE]
> The application is not code signed so a warning may pop up when running it.

## Development

Requirements:

- Node.js
- Go
- Wails

Check <https://wails.io/docs/gettingstarted/installation> for OS specific requirements.

### Run the dev environment

```bash
wails dev
```

### Build the app

```bash
wails build
```

This will build the app binary to the `build/bin` directory.

