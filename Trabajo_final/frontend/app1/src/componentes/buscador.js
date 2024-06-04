import { useState, useEffect } from "react"
import "./estilos.css"
export const Buscador = () => {

  const [query, setQuery] = useState([])
  const [cursos, setCursos] = useState('')
  const [buscar, setBuscar] = useState('')
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    setLoading(true);
    fetch('http://localhost:8080/cursos/' + query)
      .then(response => response.json())
      .then(data => setCursos(data))
      .catch(error => console.error('error al buscar el curso solicitado', error))
      .finally(() => setLoading(false))
  }, [buscar])




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
      <div className="caja-buscador">
        <div className='buscador'>
          <input
            type='text'
            placeholder='Introducir el id del curso deseado...'
            onChange={cargarDatos}
          />
          <button onClick={Buscar}>Buscar</button>
        </div>

        <div className='muestraCursos'>
          {console.log(cursos)}
        </div>
      </div>
    </div>
  )
}