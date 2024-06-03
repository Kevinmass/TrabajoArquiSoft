import logo from './logo.svg';
import './App.css';
import { useEffect, useState } from 'react';
import { queryByTestId } from '@testing-library/react';
function App() {

  const [query, setQuery] = useState([])
  const [cursos, setCursos] = useState('')
  const [buscar, setBuscar] = useState('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    setLoading(true);
    fetch('http://localhost:8080/cursos')
      .then(response => response.json())
      .then(data => setCursos(data))
      .catch(error => console.error('error al buscar el curso solicitado', error))
      .finally(() => setLoading(false))
  }, [])




  const cargarDatos = (e) => {
    setQuery(e.target.value)
    console.log('estas buscando: ' + query)

  }


  const Buscar = (e) => {
    setBuscar(query)
    console.log('se cargó la busqueda: ' + buscar)
  }

  return (
    <div className="App">

      <div className='buscador'>
        <input
          type='text'
          placeholder='Introducir curso deseado...'
          onChange={cargarDatos}
        />
        <button onClick={Buscar}>Buscar</button>
      </div>

      <div className='muestraCursos'>
          {loading && <li>Cargando...</li>}
      </div>

    </div>
  );
}

export default App;
