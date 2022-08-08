import './App.css';

function App() {
    return (
        <div className='App'>
            <h1>Servido desde go?</h1>
            <button
                onClick={async () => {
                    const response = await fetch('http://localhost:3000/users');
                    const users = await response.json();
                    console.log(users);
                }}
            >
                Obtener datos
            </button>
        </div>
    );
}

export default App;
