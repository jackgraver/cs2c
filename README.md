# CS2 Demo Viewer

A comprehensive web-based viewer for 2D visualization and analysis of Counter-Strike 2 demo files. This project provides an intuitive interface for watching, analyzing, and managing CS2 demos with real-time playback, team statistics, and interactive map visualization.

## ✨ Features

### 🎮 Demo Playback

-   **Real-time Visualization**: Smooth 2D map rendering with PixiJS
-   **Interactive Playback Controls**: Play/pause, speed control (0.5x to 4x), and timeline scrubbing
-   **Round Navigation**: Easy switching between rounds with visual indicators
-   **Multi-map Support**: Support for all major CS2 maps including Nuke (upper/lower levels)

### 📊 Team & Player Analytics

-   **Live Team Stats**: Real-time health, armor, and equipment tracking
-   **Player Information**: Detailed player cards showing health, armor, weapons, and status
-   **Kill Feed**: Real-time kill notifications with weapon and headshot indicators
-   **Score Tracking**: Live round and match score updates

### 🗺️ Map Visualization

-   **2D Top-down View**: Clear visualization of player positions and movements
-   **Grenade Tracking**: Smoke, molotov, and in-air grenade visualization
-   **Bomb Plant Indicators**: Visual markers for bomb plants and defuses
-   **Shot Trails**: Weapon shot visualization and tracking

### 📁 Demo Management

-   **Upload System**: Support for both individual `.dem` files and `.zip` archives
-   **Series Organization**: Group related demos by series/tournament
-   **Filtering & Search**: Filter demos by map, team, or date
-   **Admin Panel**: Demo management, storage monitoring, and system administration

### 🔧 Technical Features

-   **Tick-based Data**: Precise game state tracking at 64-tick intervals
-   **Efficient Caching**: Smart round data caching for smooth navigation
-   **Responsive Design**: Works on desktop and mobile devices
-   **Real-time Updates**: Live data streaming during demo playback

## 🏗️ Architecture

### Frontend

-   **React 19** with TypeScript for type safety
-   **PixiJS** for high-performance 2D rendering
-   **Tailwind CSS** for modern, responsive styling
-   **React Router** for navigation
-   **Vite** for fast development and building

### Backend

-   **Go** with Gin framework for high-performance API
-   **demoinfocs-golang** for CS2 demo parsing
-   **SQLite** for data persistence
-   **CORS-enabled** for cross-origin requests

### Data Processing

-   **Tick-by-tick Analysis**: Captures player positions, health, weapons, and actions
-   **Event Tracking**: Kills, bomb plants, grenade throws, and weapon shots
-   **Round Management**: Automatic round detection and scoring
-   **JSON Storage**: Efficient round data storage and retrieval

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

-   Inspired by [cslens.com](https://cslens.com) - an excellent CS demo viewer
-   Built with [demoinfocs-golang](https://github.com/markus-wa/demoinfocs-golang) for demo parsing
-   Uses [PixiJS](https://pixijs.com/) for high-performance 2D rendering
