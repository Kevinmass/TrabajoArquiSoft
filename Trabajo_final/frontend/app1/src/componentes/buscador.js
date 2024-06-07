import { useState, useEffect, Component } from "react"
import "./estilos.css"
export const Buscador = () => {

  const [query, setQuery] = useState("")
  const [cursos, setCursos] = useState([])
const [resultados,setResultados] = useState(false)
  
  useEffect(()=> {
    const mostrarCursos= () =>{

  fetch('http://localhost:8080/cursos')
  .then(response => response.json())
  .then(data => {
    setCursos(data);
  })
  .catch(error => console.error('error al cargar los cursos', error))
}
mostrarCursos();
},[] )
  

  const cargarDatos = (e) => {
    setQuery(e.target.value)
  }


  const Buscar = () => {
      fetch('http://localhost:8080/cursos/'+query)
        .then(response => response.json())
        .then(data => {
          setCursos(Array.isArray(data)?data : [data]);  //con esta linea aseguramos que el curso esté dentro de un array y no tire el error de cursos.map is not a function
          setResultados(data.length===0);
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

        <div className="muestraCursos">
          {resultados ? (
            <p>no se obtuvieron resultados para su busqueda</p>
          ) : (
              cursos.map((curso) => (
                <div key={curso.id}>
                  <h2>{curso.nombre}</h2>
                  <p>Descripción: {curso.descripcion}</p>
                  <p>Profesor: {curso.profesor_nombre} {curso.profesor_apellido}</p>
                  <p>Correo del profesor: {curso.profesor_correo}</p>
                  <p>Fecha de Creación: {curso.fecha_creacion}</p>
                  <p>Fecha de Modificación: {curso.fecha_actualizacion}</p>
                </div>
              )
            )
          )}
        </div>
      </div>
    </div>
  );
}