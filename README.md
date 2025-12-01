# Minimalist Life Tracker

A super minimalist habit tracker built with Wails (Go + React). Track your daily habits with a clean, simple interface inspired by GitHub's contribution graph.

## Features

- **Authentication**: Simple login/register system
- **Habit Management**: Create habits with custom names, descriptions, and icons
- **Icon Picker**: Choose from 90+ emojis across 9 categories (Fitness, Health, Food, Sleep, Study, Work, Habits, Nature, Other)
- **Visual Progress Tracking**: Year-long grid visualization (365 days) showing your habit completion
- **Minimalist Design**: Clean, distraction-free interface focused on simplicity

## Technology Stack

- **Backend**: Go (Golang)
- **Frontend**: React + Vite
- **Desktop Framework**: Wails v2
- **Database**: SQLite with GORM (using modernc.org/sqlite - pure Go, no CGO required)
- **Authentication**: bcrypt password hashing

## Prerequisites

Before running this application, make sure you have:

- [Go](https://golang.org/dl/) (1.21 or later)
- [Node.js](https://nodejs.org/) (18 or later)
- [Wails CLI](https://wails.io/docs/gettingstarted/installation)

**Note**: This application uses a pure Go SQLite driver (modernc.org/sqlite), so **no CGO or GCC is required**. It works out of the box on Windows, macOS, and Linux without any C compiler.

### Install Wails CLI

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

## Installation

1. Clone the repository:
```bash
git clone <your-repo-url>
cd tracker
```

2. Install Go dependencies:
```bash
go mod download
```

3. Install frontend dependencies:
```bash
cd frontend
npm install
cd ..
```

## Development

Run the application in development mode with hot reload:

```bash
wails dev
```

This will:
- Start the Go backend
- Start the Vite dev server
- Open the application window
- Enable hot reload for both frontend and backend

## Building

Build the application for production:

```bash
wails build
```

The compiled application will be in the `build/bin` directory.

### Build for specific platforms:

```bash
# Windows
wails build -platform windows/amd64

# macOS
wails build -platform darwin/universal

# Linux
wails build -platform linux/amd64
```

## Usage

### First Time Setup

1. Launch the application
2. Click "Don't have an account? Register"
3. Create a username and password
4. Login with your credentials

### Creating Habits

1. Click the "+ New Habit" button
2. Select an icon from the icon picker
3. Enter a habit name (e.g., "200 toe lifts")
4. Optionally add a description (e.g., "90 push ups daily")
5. Click "Create Habit"

### Tracking Habits

- Click on any cell in the grid to mark that day as complete
- Click again to unmark
- The grid shows the last 365 days
- Darker cells indicate completed days

### Managing Habits

- Click the × button on a habit card to delete it
- Hover over grid cells to see the date

## Project Structure

```
tracker/
├── main.go                 # Application entry point
├── app.go                  # Application logic and API methods
├── database.go             # Database models and initialization
├── wails.json              # Wails configuration
├── go.mod                  # Go dependencies
├── frontend/
│   ├── src/
│   │   ├── components/     # React components
│   │   │   ├── HabitCard.jsx
│   │   │   ├── HabitGrid.jsx
│   │   │   ├── CreateHabitModal.jsx
│   │   │   └── IconPicker.jsx
│   │   ├── pages/          # Page components
│   │   │   ├── Login.jsx
│   │   │   └── Dashboard.jsx
│   │   ├── styles/         # CSS files
│   │   ├── wailsjs/        # Generated Wails bindings
│   │   ├── App.jsx
│   │   └── main.jsx
│   ├── index.html
│   ├── package.json
│   └── vite.config.js
└── README.md
```

## Database Schema

### Users
- `id`: Primary key
- `username`: Unique username
- `password`: Bcrypt hashed password
- `created_at`: Registration timestamp

### Habits
- `id`: Primary key
- `user_id`: Foreign key to users
- `name`: Habit name
- `description`: Optional description
- `icon`: Emoji icon
- `created_at`: Creation timestamp

### Tracks
- `id`: Primary key
- `habit_id`: Foreign key to habits
- `date`: Date of completion
- `completed`: Boolean flag
- `created_at`: Timestamp

## API Methods (Go ↔ React)

### Authentication
- `Login(username, password)` → LoginResponse
- `Register(username, password)` → LoginResponse

### Habit Management
- `CreateHabit(userId, name, description, icon)` → HabitResponse
- `GetHabits(userId)` → []HabitResponse
- `DeleteHabit(habitId, userId)` → error

### Tracking
- `ToggleTrack(habitId, userId, date)` → TrackResponse
- `GetTracks(habitId, userId)` → []TrackResponse

## Design Philosophy

This application follows a **minimalist** design philosophy:

- **No unnecessary features**: Only what you need to track habits
- **Clean interface**: No clutter, no distractions
- **Simple interactions**: Click to complete, that's it
- **Visual clarity**: Grid visualization makes patterns obvious
- **Fast and lightweight**: Desktop app, no web server needed

## Security Notes

- Passwords are hashed using bcrypt
- Local SQLite database stores all data
- No telemetry or external connections
- All data stays on your machine

## Contributing

Feel free to open issues or submit pull requests!

## License

MIT License - feel free to use this project however you like.

## Screenshots

The app features:
- Clean login screen with simple username/password
- Dashboard with habit cards
- Grid visualization (52 weeks × 7 days)
- Icon picker with 90+ emojis
- Smooth, minimalist animations

## Troubleshooting

### "wails: command not found"
Make sure `$GOPATH/bin` is in your PATH:
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

### Build errors
Try cleaning and rebuilding:
```bash
wails clean
wails build
```

### Frontend not updating
Clear the build cache:
```bash
cd frontend
rm -rf node_modules dist
npm install
cd ..
wails dev
```

## Future Ideas

Potential enhancements (keeping minimalism in mind):
- Export data to JSON/CSV
- Dark mode
- Habit streaks counter
- Weekly/monthly summary views
- Custom grid color themes
- Habit categories/tags

---

Built with ❤️ using Wails, Go, and React