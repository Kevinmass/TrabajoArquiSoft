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
      fetch('http://localhost:8080/cursos/'+query)
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
               // aca me tira un error rari
              <div key={cursos.id}>
                <h2>{cursos.nombre}</h2>
                <p>Descripción: {cursos.descripcion}</p>
                <p>Profesor: {cursos.rofesor_nombre} {cursos.profesor_apellido}</p>
                <p>Correo del profesor: {cursos.profesor_correo}</p>
              
                <p>Fecha de Creación: {cursos.fecha_creacion}</p>
                <p>Fecha de Modificación: {cursos.fecha_actualizacion}</p>
              </div>
            )
          }
        </div>
      </div>
    </div>
  );
}