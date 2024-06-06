import { useState, useEffect } from "react"
import "./estilos.css"
export const Buscador = () => {

  const [query, setQuery] = useState("")
  const [cursos, setCursos] = useState([])
  const [buscar, setBuscar] = useState("")
  const [loading, setLoading] = useState(true)

  




  const cargarDatos = (e) => {
    setQuery(e.target.value)
    console.log('estas buscando: ' + query)

  }


  const Buscar = () => {
      setLoading(true);
      fetch('http://localhost:8080/cursos?' +buscar)
        .then(response => response.json())
        .then(data => {
          setCursos(data);
          setLoading(false);
        })
        .catch(error => console.error('error al buscar el curso solicitado', error))
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
        {loading ? (
            <p>Cargando...</p>
          ) : (
            cursos.map((curso) => (    // aca me tira un error rari
              <div key={curso.id}>
                <h2>{curso.nombre}</h2>
                <p>Profesor: {curso.nombre_profesor} {curso.apellido_profesor}</p>
                <p>Correo del profesor: {curso.correo}</p>
                <p>Descripción: {curso.descripcion}</p>
                <p>Fecha de Creación: {curso.fecha_creacion}</p>
                <p>Fecha de Modificación: {curso.fecha_modificacion}</p>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  );
}