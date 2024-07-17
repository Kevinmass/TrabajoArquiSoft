import React, { useState } from "react"
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos"

export const CrearCurso = () => {


    
const [curso,setCurso] = useState({
    user: "",
    nombre:"",
    descripcion: "",
    duracion:"",
    requisitos: "",
});

   
        
    const handleChange =(e)=>{
        setCurso({
            ...curso,
            [e.target.name]: e.target.value,
        });
    }

    const cursoNuevo = async(e) =>{
        e.preventDefault();
        try{
            
            const usuario = localStorage.getItem('user');

            curso.user=usuario;
            const response = await fetch ('http://localhost:8080/crearCurso',{
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(curso)
        });
        if (!response.ok) {
            throw new Error('error al crear el curso')
        }
        alert('curso creado correctamente')
        setCurso({
            user: "",
            nombre: "",
            descripcion: "",
            duracion: "",
            requisitos: "",
        });
    } catch(error){
        alert(error.message + ", no se pudo crear el curso")
        setCurso({
            user:"",
            nombre: "",
            descripcion: "",
            duracion: "",
            requisitos: "",
        });
    }
        
    
    }
    return (

        <div>
            <HeaderMisCursos/>
            <div>           
                <h1>
                    Complete los siguientes campos para poder crear su curso
                </h1>
                <div className="ingresarDatos">
                    <form onSubmit={cursoNuevo} >
                        <div className="cajas_inputs">
                        <h2>Nombre del curso</h2>
                        <input 
                        type="text"
                        name="nombre"
                        placeholder="Programación II"
                        value={curso.nombre}
                        onChange={handleChange}
                        required
                        />
                        </div>
                        <div className="cajas_inputs">
                        <h2>Descripcion</h2>
                        <input 
                        type="text"
                        name="descripcion"
                        placeholder="Este curso trata sobre como aprender a programar"
                        value={curso.descripcion}
                        onChange={handleChange}
                        required
                        />    
                        </div>

                        <div className="cajas_inputs">
                        <h2>Duracion</h2>
                        <input 
                        type="text"
                        name="duracion"
                        placeholder="3 meses"
                        value={curso.duracion}
                        onChange={handleChange}
                        required    
                        />  
                        </div>

                          <div className="cajas_inputs">
                          <h2>Requisitos</h2>
                        <input 
                        type="text"
                        name="requisitos"
                        placeholder="Conocimientos en Fundamentos de Programacion"
                        value={curso.requisitos}
                        onChange={handleChange}
                        required
                        />    
                        </div>
                        <button type="submit">Crear curso</button>
                    </form>
                </div>
            </div>
        </div>
        
    )}
