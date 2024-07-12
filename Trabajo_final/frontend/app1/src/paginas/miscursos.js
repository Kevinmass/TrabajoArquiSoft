import React, { useEffect, useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos";
import { Link } from "react-router-dom";

export const MisCursos = () => {

    const [resultados, setResultados] = useState(false)
    const [cursos, setCursos] = useState([])

    const usuario = localStorage.getItem('user');

    const esProfesor = localStorage.getItem('profesor') === 'true';

useEffect (()=>{
    fetch('http://localhost:8080/GetCursosdelCliente/' + usuario)
    .then(response => response.json())
    .then(data => {
        setCursos(Array.isArray(data) ? data : [data]);
        setResultados(data.length === 0);

    })
    .catch(error => console.error('ocurrió un error al obtener tus cursos, lo siento.', error))
}, [usuario]);
   

const eliminarCurso = async (cursoid) =>{

try {
    const response = await fetch('http://localhost:8080/eliminarCurso',{
method: 'POST',
headers: {
    "Content-Type": "application/json",
  },
  body: JSON.stringify({
    user: usuario,
    id: cursoid
  }),
    })
    if (!response.ok) {
        throw new Error("error al eliminar el curso");
      }
      setCursos(cursos.filter(curso => curso.id !== cursoid));

} catch(error){
    alert (error.message)
}

}
    return <div>
        <HeaderMisCursos />
        <div className="muestraCursos">
            {resultados ? (
                <p>no estás inscripto a ningún curso todavía...</p>
            ) : (
                cursos.map((curso) => (
                    <div className="cursoIndividual" key={curso.id}>
                        <h2>{curso.nombre}</h2>
                        <p>Descripción: {curso.descripcion}</p>
                        <p>Profesor: {curso.user}</p>
                        <p>Duracion: {curso.duracion}</p>
                        <p>Requisitos: {curso.requisitos}</p>
                        {esProfesor ? (
                            <div>
                            <button onClick={() => eliminarCurso(curso.id)}>Eliminar curso</button>  
                        <button><Link to='/modificarCurso' onClick= { () => {
                localStorage.setItem('cursoID',JSON.stringify(curso.id))
               }}>Modificar curso</Link></button></div> ):(<div>no es profesor ok </div>
                       ) }
                    </div>

                )))}
        </div>
    </div>
}