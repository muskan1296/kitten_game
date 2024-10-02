import React, { useState } from 'react';

const StartButton = () => {
    const [loading, setLoading] = useState(false);

    const handleStart = async () => {
        setLoading(true);
        console.log("Start button clicked"); // Log statement for debugging
        try {
            const response = await fetch('http://localhost:8080/start', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
            });

            if (response.ok) {
                const data = await response.json();
                console.log(data.message); // Check response message
                // Handle successful game start
            } else {
                console.error('Failed to start game:', response.statusText);
            }
        } catch (error) {
            console.error('Error starting game:', error);
        } finally {
            setLoading(false);
        }
    };

    return (
        <button onClick={handleStart} disabled={loading}>
            {loading ? 'Starting...' : 'Start Game'}
        </button>
    );
};

export default StartButton;
