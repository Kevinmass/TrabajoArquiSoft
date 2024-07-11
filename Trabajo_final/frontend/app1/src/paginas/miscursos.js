import React, { useEffect, useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos";

export const MisCursos = () => {

    const [resultados, setResultados] = useState(false)
    const [cursos, setCursos] = useState([])

    const usuario = localStorage.getItem('user')
useEffect (()=>{
    
    fetch('http://localhost:8080/GetCursosdelCliente/' + usuario)
    .then(response => response.json())
    .then(data => {
        setCursos(Array.isArray(data) ? data : [data]);
        setResultados(data.length === 0);

    })
    .catch(error => console.error('ocurrió un error al obtener tus cursos, lo siento.', error))
}, [usuario]);
   


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
                    </div>

                )))}
        </div>
    </div>
}