import React from 'react'
import HabitGrid from './HabitGrid'
import { DeleteHabit } from '../wailsjs/go/main/App'
import '../styles/HabitCard.css'

function HabitCard({ habit, userId, onDelete, onTrackUpdate }) {
  const handleDelete = async () => {
    if (window.confirm('Are you sure you want to delete this habit?')) {
      try {
        await DeleteHabit(habit.id, userId)
        onDelete()
      } catch (err) {
        console.error('Error deleting habit:', err)
        alert('Failed to delete habit')
      }
    }
  }

  return (
    <div className="habit-card">
      <div className="habit-header">
        <div className="habit-icon">{habit.icon}</div>
        <div className="habit-info">
          <h3 className="habit-name">{habit.name}</h3>
          {habit.description && (
            <p className="habit-description">{habit.description}</p>
          )}
        </div>
        <button onClick={handleDelete} className="habit-delete">
          ×
        </button>
      </div>

      <HabitGrid
        habitId={habit.id}
        userId={userId}
        tracks={habit.tracks}
        onUpdate={onTrackUpdate}
      />
    </div>
  )
}

export default HabitCard
