import React, { useContext } from 'react';
import { HoveredWordContext } from './App';
import './index.css'; // Import your custom CSS file

function Sidebar({ wordsData }) {
  const { setHoveredWord, setHoveredPath } = useContext(HoveredWordContext);

  return (
    <div className="sidebar">
      <div className="word-list">
        {wordsData.map((data, index) => (
          <div
            key={index}
            className="word"
            onMouseEnter={() => { setHoveredWord(data.word); setHoveredPath(data.path); }}
            onMouseLeave={() => { setHoveredWord(null); setHoveredPath([]); }}
          >
            {data.word}
          </div>
        ))}
      </div>
    </div>
  );
}

export default Sidebar;