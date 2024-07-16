import { useState, useEffect } from "react";
import "./estilos.css";
import { Link } from "react-router-dom";

export const Buscador = () => {
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

  const inscribirCurso = async (datos) => {
    try {
      const response = await fetch("http://localhost:8080/inscribirse", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(datos),
      });
      if (!response.ok) {
        throw new Error("error al inscribirse");
      }
      const data = await response.json();
      alert("inscripción realizada correctamente!");
    } catch {
      alert("ya estás inscripto a este curso!");
    }
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
          <Link to='/homeArchivos' className="link-button">Subir archivos</Link>
        </div>
      </div>
      <div className="muestraCursos">
        {resultados ? (
          <p>no se obtuvieron resultados para su busqueda</p>
        ) : (
          cursos.map((curso) => (
            <div className="cursoIndividual" key={curso.id}>
              <h2>{curso.nombre}</h2>
              <p>Descripción: {curso.descripcion}</p>
              <p>
                Profesor: {curso.user} </p>
              <p>Duracion: {curso.duracion}</p>
              <p>Requisitos: {curso.requisitos}</p>
              <button
                onClick={() => {
                  var datos = {
                    User: localStorage.getItem("user"),
                    Curso_id: curso.id,
                  };
                  console.log(datos);
                  inscribirCurso(datos);
                }}
              >
                Inscribirse al curso
              </button>
         
                <Link
               onClick= { () => {
                localStorage.setItem("cursoInfo",JSON.stringify(curso))
               
                  console.log(curso)
               }}
                  to="/cursodetalle"  className="link-button">+ info</Link>
                
            
                

              
            </div>
          ))
        )}
      </div>
    </div>
  );
};
