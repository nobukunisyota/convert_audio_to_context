import React, { useState } from 'react';
import './App.css'
import ApiGet from './component/ApiGet'
import ApiGetByID from './component/ApiGetByID';

function App() {
    // Get API All Tasks
    const [apiGetEvent, setApiGetEvent] = useState(false);
    const getApiHandler = () => {
        setApiGetEvent(!apiGetEvent);
    }
    // Get API Tasks By ID
    const [getTask, setGetTask] = useState(false);
    const getTaskHandler = () => {
        setGetTask(!getTask);
    }
    return (
        <div>
            <button onClick={() => getApiHandler()}>Get All Task</button>
            {apiGetEvent ? <ApiGet /> : null}
            <button onClick={() => getTaskHandler()}>Get All ID</button>
            {getTask ? <ApiGetByID id_num="1"/> : null}
        </div>
    )
}

export default App
