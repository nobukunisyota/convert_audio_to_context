import React, { useState, useEffect }   from 'react'
import axios from 'axios';

function ApiGet() {
    const [data, setData] = useState([]);
    useEffect(() => {
        const backendApi = "http://127.0.0.1:8001/api/v1/tasks";
        axios.get(backendApi, {
            headers: {
            'Content-Type': 'application/json'
            }
        })
        .then(response => setData(response.data));
    }, []);
    return (
        <div>
            {data.map((task: any) => (
                    <div key={task.id}>
                        <p>ID : {task.id}</p>
                        <p>Name : {task.name}</p>
                        <p>Context : {task.context}</p>
                        <p>Tag : {task.tag}</p>
                        <p>CreatedAt : {task.created_at}</p>
                        <p>UpdatedAt : {task.updated_at}</p>
                    </div>
            ))}
        </div>
    )
}

export default ApiGet
