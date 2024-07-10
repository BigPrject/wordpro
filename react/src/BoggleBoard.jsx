import React from 'react';
import './index.css';

function BoggleBoard({ board, onBoardChange }) {
  const renderBoard = () => {
    return (
      <div className="boggle-grid">
        {board.map((letter, index) => (
          <input
            key={index}
            type="text"
            maxLength="1"
            value={letter}
            onChange={(e) => onBoardChange(index, e.target.value)}
            className="grid-cell"
            style={{
              textTransform: 'uppercase',
            }}
          />
        ))}
      </div>
    );
  };

  return (
    <div className="boggle-board">
      {renderBoard()}
    </div>
  );
}

export default BoggleBoard;