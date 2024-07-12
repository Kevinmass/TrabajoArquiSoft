import { useState,useEffect } from "react";
import React from "react";
import { Link } from "react-router-dom";


export const BuscadorProfesor = () =>{


const [query, setQuery] = useState("");
  const [cursos, setCursos] = useState([]);
  const [resultados, setResultados] = useState(false);

  const MostrarCursos = () => {
    fetch("http://localhost:8080/cursos")
      .then((response) => response.json())
      .then((data) => {
        setCursos(data);
      })
      .catch((error) => console.error("error al cargar los cursos", error));
  };

  useEffect(() => {
    MostrarCursos();
  }, []);

  const LimpiarBusqueda = () => {
    setResultados(false);
    MostrarCursos();
    setQuery("");
  };

  
  const cargarDatos = (e) => {
    setQuery(e.target.value);
  };

  const Buscar = () => {
    fetch("http://localhost:8080/cursosPorNombre/" + query)
      .then((response) => response.json())
      .then((data) => {
        setCursos(Array.isArray(data) ? data : [data]); // Asegura que el curso esté dentro de un array
        setResultados(data.length === 0);
      })
      .catch((error) =>
        alert("por favor introduce el nombre del curso deseado", error)
      );
  };

  return (
    <div className="App">
      <div className="caja-buscador">
        <div className="buscador">
          <input
            type="text"
            value={query}
            placeholder="Qué curso estás buscando?"
            onChange={cargarDatos}
          />
        </div>
        <div className="botones-buscador">
          <button onClick={Buscar}>Buscar</button>
          <button onClick={LimpiarBusqueda}>Limpiar búsqueda</button>
          <Link to="/miscursos" className="link-button">
            Mis cursos
          </Link>
          
        <Link to='/crearCurso' className="link-button">
        Crear curso
        </Link>
        </div>
      </div>
      <div className="muestraCursos">
        {resultados ? (
          <p>no se obtuvieron resultados para su busqueda</p>
        ) : (
          cursos && cursos.length > 0 ? (
            cursos.map((curso) => (
              <div className="cursoIndividual" key={curso.id}>
                <h2>{curso.nombre}</h2>
                <p>Descripción: {curso.descripcion}</p>
                <p>Profesor: {curso.user}</p>
                <p>Duracion: {curso.duracion}</p>
                <p>Requisitos: {curso.requisitos}</p>             
              </div>
            ))
          ) : (
            <p>No hay cursos disponibles</p>
          )
        )}
      </div>
    </div>
  );
};
