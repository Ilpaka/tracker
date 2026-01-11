import React, { useState, useEffect } from 'react'
import HabitCard from '../components/HabitCard'
import CreateHabitModal from '../components/CreateHabitModal'
import { GetHabits, GetTracks } from '../wailsjs/go/main/App'
import '../styles/Dashboard.css'

function Dashboard({ user, onLogout }) {
  const [habits, setHabits] = useState([])
  const [habitsWithTracks, setHabitsWithTracks] = useState([])
  const [showCreateModal, setShowCreateModal] = useState(false)
  const [loading, setLoading] = useState(true)

  const loadHabits = async () => {
    try {
      setLoading(true)
      const habitsData = await GetHabits(user.userId)
      setHabits(habitsData || [])

      // Load tracks for each habit
      const habitsWithTracksData = await Promise.all(
        (habitsData || []).map(async (habit) => {
          try {
            const tracks = await GetTracks(habit.id, user.userId)
            return { ...habit, tracks: tracks || [] }
          } catch (err) {
            console.error('Error loading tracks for habit:', habit.id, err)
            return { ...habit, tracks: [] }
          }
        })
      )

      setHabitsWithTracks(habitsWithTracksData)
    } catch (err) {
      console.error('Error loading habits:', err)
    } finally {
      setLoading(false)
    }
  }

  useEffect(() => {
    loadHabits()
  }, [user.userId])

  const handleHabitCreated = () => {
    setShowCreateModal(false)
    loadHabits()
  }

  const handleHabitDeleted = () => {
    loadHabits()
  }

  const handleTrackUpdated = () => {
    loadHabits()
  }

  return (
    <div className="dashboard">
      <header className="dashboard-header">
        <h1 className="dashboard-title">Life Tracker</h1>
        <button onClick={onLogout} className="logout-button">
          Logout
        </button>
      </header>

      <div className="dashboard-content">
        {loading ? (
          <div className="loading">Loading...</div>
        ) : habitsWithTracks.length === 0 ? (
          <div className="empty-state">
            <p>No habits yet. Create your first one!</p>
          </div>
        ) : (
          <div className="habits-list">
            {habitsWithTracks.map((habit) => (
              <HabitCard
                key={habit.id}
                habit={habit}
                userId={user.userId}
                onDelete={handleHabitDeleted}
                onTrackUpdate={handleTrackUpdated}
              />
            ))}
          </div>
        )}

        <button
          onClick={() => setShowCreateModal(true)}
          className="create-habit-button"
        >
          + New Habit
        </button>
      </div>

      {showCreateModal && (
        <CreateHabitModal
          userId={user.userId}
          onClose={() => setShowCreateModal(false)}
          onCreated={handleHabitCreated}
        />
      )}
    </div>
  )
}

export default Dashboard
