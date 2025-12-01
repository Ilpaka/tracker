import React, { useMemo } from 'react'
import { ToggleTrack } from '../wailsjs/go/main/App'
import '../styles/HabitGrid.css'

function HabitGrid({ habitId, userId, tracks, onUpdate }) {
  // Generate grid for last 365 days (52 weeks * 7 days)
  const gridData = useMemo(() => {
    const today = new Date()
    const days = []
    const trackMap = new Map()

    // Create a map of completed dates
    tracks.forEach(track => {
      if (track.completed) {
        const date = new Date(track.date)
        const dateStr = date.toISOString().split('T')[0]
        trackMap.set(dateStr, track)
      }
    })

    // Generate last 365 days
    for (let i = 364; i >= 0; i--) {
      const date = new Date(today)
      date.setDate(date.getDate() - i)
      const dateStr = date.toISOString().split('T')[0]

      days.push({
        date: dateStr,
        completed: trackMap.has(dateStr),
        dayOfWeek: date.getDay()
      })
    }

    return days
  }, [tracks])

  const handleCellClick = async (dateStr) => {
    try {
      await ToggleTrack({
        habitId: habitId,
        userId: userId,
        date: dateStr
      })
      onUpdate()
    } catch (err) {
      console.error('Error toggling track:', err)
    }
  }

  return (
    <div className="habit-grid-container">
      <div className="habit-grid">
        {gridData.map((day, index) => (
          <div
            key={index}
            className={`grid-cell ${day.completed ? 'completed' : ''}`}
            onClick={() => handleCellClick(day.date)}
            title={day.date}
          />
        ))}
      </div>
    </div>
  )
}

export default HabitGrid
