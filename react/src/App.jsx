import React, { useState } from 'react';
import { useSpring, animated } from 'react-spring';
import BoggleBoard from './BoggleBoard';
import Sidebar from './SideBar';
import './index.css'; // Import your custom CSS file

export const HoveredWordContext = React.createContext();

const App = () => {
  const [showInterface, setShowInterface] = useState(false);
  const [hoveredWord, setHoveredWord] = useState(null);
  const [hoveredPath, setHoveredPath] = useState([]);
  const [wordsData, setWordsData] = useState([]);
  const [board, setBoard] = useState(Array(16).fill(""));

  const fadeIn = useSpring({
    opacity: showInterface ? 1 : 0,
  });

  const floatUp = useSpring({
    transform: showInterface ? 'translateY(0%)' : 'translateY(100%)',
  });

  const handleBoardChange = (index, value) => {
    const newBoard = [...board];
    newBoard[index] = value.toUpperCase();
    setBoard(newBoard);
  };


  const handleSolveClick = () => {
    var words
    const boardString = board.join('').toLowerCase();

    fetch('http://localhost:8080/solve', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ board: boardString }),
    })
    .then(response => response.json())
    .then(data => {
      console.log('Response JSON:', data); // Log the response JSON directly here
      setWordsData(data.words); // Update the state with received words
    })
    .catch(error => console.error('Error solving board:', error));
  };

  return (
    <div className="bg-gradient min-h-screen flex flex-col justify-between">
      <header className="p-4">
        <h1 className="text-white text-3xl">A work in progress...</h1>
      </header>
      <main className="p-4">
        <button
          onClick={() => setShowInterface(!showInterface)}
          className="bg-blue-500 text-white py-2 px-4 rounded"
        >
          {showInterface ? 'Hide' : 'Show'} Word Solver
        </button>

        <animated.div style={fadeIn} className="interface-container mt-4">
          <animated.div style={floatUp} className="interface-content">
            <HoveredWordContext.Provider value={{ hoveredWord, setHoveredWord, hoveredPath, setHoveredPath }}>
              <BoggleBoard board={board} onBoardChange={handleBoardChange} />
              <button
                onClick={handleSolveClick}
                className="bg-green-500 text-white py-2 px-4 rounded mt-4"
              >
                Solve
              </button>
              <Sidebar wordsData={wordsData} />
            </HoveredWordContext.Provider>
          </animated.div>
        </animated.div>
      </main>
      <footer className="p-4">
        <p className="text-white">&copy; wordSolver</p>
      </footer>
    </div>
  );
};

export default App;