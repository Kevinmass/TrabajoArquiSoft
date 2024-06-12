import React, { useEffect, useState } from "react";
import { HeaderMisCursos } from "./encabezadoMisCursos";

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
                    <div classname="cursoIndividual" key={curso.id}>
                        <h2>{curso.nombre}</h2>
                        <p>Descripción: {curso.descripcion}</p>
                        <p>Profesor: {curso.profesor_nombre} {curso.profesor_apellido}</p>
                        <p>Correo del profesor: {curso.profesor_correo}</p>
                        <p>Fecha de Creación: {curso.fecha_creacion}</p>
                        <p>Fecha de Modificación: {curso.fecha_actualizacion}</p>
                    </div>

                )))}
        </div>
    </div>
}