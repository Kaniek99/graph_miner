import { useState } from 'react';
import './App.css';

function App() {
  interface ResponseData {
    message: string;
  }

  const [data, setData] = useState<ResponseData | null>(null);
  const [loading, setLoading] = useState<boolean>(false);

  const sendRequest = async () => {
    setLoading(true);
    try {
      console.log('Sending request to backend');
      const response = await fetch('http://host.docker.internal:8080/', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        body: null,
    });
      const result = await response.json();
      // const result = await response.text();
      console.log('Response:', result);
      setData(result);
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="container">
      <div className="content">
        <h1>Hello World</h1>
        <p>Welcome to graph_miner app built with Vite</p>
        <button
          className="fetch-button"
          onClick={sendRequest}
          disabled={loading}
        >
          {loading ? 'Loading...' : 'Fetch Data'}
        </button>
        {data && (
          <div className="data-container">
            <h2>Response Data:</h2>
            <pre>{JSON.stringify(data, null, 2)}</pre>
          </div>
        )}
      </div>
    </div>
  );
}

export default App;