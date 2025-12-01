import React, { useState } from 'react'
import '../styles/IconPicker.css'

const ICON_CATEGORIES = {
  Fitness: ['🏋️', '💪', '🤸', '🧘', '🏃', '🚴', '🏊', '⚽', '🏀', '🎾'],
  Health: ['🩺', '💊', '🧠', '❤️', '🫀', '🫁', '🦷', '👁️', '🌡️', '🩹'],
  Food: ['🍎', '🥗', '🥑', '🥦', '🥕', '🍊', '🍇', '🥤', '☕', '🍵'],
  Sleep: ['😴', '🛏️', '🌙', '⭐', '💤', '🌃', '🕐', '⏰', '⏱️', '⌛'],
  Study: ['📚', '📖', '✏️', '📝', '🎓', '🧑‍💻', '💻', '🖊️', '📄', '📋'],
  Work: ['💼', '🏢', '📊', '📈', '📉', '💰', '🎯', '✅', '📌', '🔔'],
  Habits: ['🔥', '⚡', '✨', '🌟', '💫', '🎨', '🎭', '🎪', '🎬', '🎮'],
  Nature: ['🌱', '🌿', '🌳', '🌲', '🌴', '🌵', '🌾', '🍀', '🌺', '🌻'],
  Other: ['🎵', '🎶', '📷', '🎥', '🎤', '🎧', '📱', '⌚', '🔑', '🎁']
}

function IconPicker({ selectedIcon, onSelect }) {
  const [activeCategory, setActiveCategory] = useState('Fitness')

  return (
    <div className="icon-picker">
      <div className="icon-selected">
        <span className="icon-large">{selectedIcon}</span>
      </div>

      <div className="icon-categories">
        {Object.keys(ICON_CATEGORIES).map((category) => (
          <button
            key={category}
            type="button"
            className={`category-button ${activeCategory === category ? 'active' : ''}`}
            onClick={() => setActiveCategory(category)}
          >
            {category}
          </button>
        ))}
      </div>

      <div className="icon-grid">
        {ICON_CATEGORIES[activeCategory].map((icon) => (
          <button
            key={icon}
            type="button"
            className={`icon-button ${selectedIcon === icon ? 'selected' : ''}`}
            onClick={() => onSelect(icon)}
          >
            {icon}
          </button>
        ))}
      </div>
    </div>
  )
}

export default IconPicker
