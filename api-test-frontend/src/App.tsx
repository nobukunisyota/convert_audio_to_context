import React, { useState } from 'react';
import './App.css'
import ApiGet from './component/ApiGet'

function App() {
    // Get API All Tasks
    const [apiGetEvent, setApiGetEvent] = useState(false);
    const getApiHandler = () => {
        setApiGetEvent(!apiGetEvent);
    }
    return (
        <div>
            <button onClick={() => getApiHandler()}>Get All Task</button>
            {apiGetEvent ? <ApiGet /> : null}
        </div>
    )
}

export default App
