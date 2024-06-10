import { useState, useEffect } from "react";
import "./estilos.css";
import { Link } from "react-router-dom";

export const Buscador = () => {
    const [query, setQuery] = useState("");
    const [cursos, setCursos] = useState([]);
    const [resultados, setResultados] = useState(false);

    useEffect(() => {
        const mostrarCursos = () => {
            fetch('http://localhost:8080/cursos')
                .then(response => response.json())
                .then(data => {
                    setCursos(data);
                })
                .catch(error => console.error('error al cargar los cursos', error));
        };
        mostrarCursos();
    }, []);

    const inscribirCurso = async (datos) => {
        try {
            const response = await fetch('http://localhost:8080/inscribirse', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(datos)
            });
            if (!response.ok) {
                throw new Error('error al inscribirse');
            }
            const data = await response.json();
            alert('inscripción realizada correctamente!');
        } catch {
            alert('error al inscribirse');
        }
    };

    const cargarDatos = (e) => {
        setQuery(e.target.value);
    };

    const Buscar = () => {
        fetch('http://localhost:8080/cursosPorNombre/' + query)
            .then(response => response.json())
            .then(data => {
                setCursos(Array.isArray(data) ? data : [data]);
                setResultados(data.length === 0);
            })
            .catch(error => console.error('error al buscar el curso solicitado', error));
    };

    return (
        <div className="App">
            <div className="caja-buscador">
                <div className='buscador'>
                    <input
                        type='text'
                        placeholder='Qué curso estás buscando?'
                        onChange={cargarDatos}
                    />
                    <button onClick={Buscar}>Buscar</button>
                </div>
            </div>
            <div className="muestraCursos">
                {resultados ? (
                    <p>No se obtuvieron resultados para su búsqueda</p>
                ) : (
                    cursos.map((curso) => (
                        <div className="cursoIndividual" key={curso.id}>
                            <h2>{curso.nombre}</h2>
                            <p>Descripción: {curso.descripcion}</p>
                            <p>Profesor: {curso.profesor_nombre} {curso.profesor_apellido}</p>
                            <p>Correo del profesor: {curso.profesor_correo}</p>
                            <p>Fecha de Creación: {curso.fecha_creacion}</p>
                            <p>Fecha de Modificación: {curso.fecha_actualizacion}</p>
                            <button onClick={() => {
                                const datos = {
                                    User: localStorage.getItem('user'),
                                    Curso_id: curso.id
                                };
                                console.log(datos);
                                inscribirCurso(datos);
                            }}>Inscribirse al curso</button>
                        </div>
                    ))
                )}
            </div>
        </div>
    );
};
