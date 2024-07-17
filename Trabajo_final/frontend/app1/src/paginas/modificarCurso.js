import { HeaderMisCursos } from "../componentes/encabezadoMisCursos"
import React, { useState } from "react"

    



export const ModificarCurso =()=> {
    const [cursomodif,setCurso] = useState({
        user: localStorage.getItem('user'),
        id: Number(localStorage.getItem('cursoID')),
        nombre:"",
        descripcion: "",
        duracion:"",
        requisitos: "",
    });
    
       
            
    const handleChange =(e)=>{
        setCurso({
            ...cursomodif,
            [e.target.name]: e.target.value,
        });
    }
const cursoModificado = async (e) => {
e.preventDefault();
    try {
        const response = await fetch('http://localhost:8080/modificarCurso',{
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(cursomodif)
        })
        if(!response.ok){
            throw new Error("ocurrio un error al consumir el endpoint para modificarCurso")
        }
        alert("Modificaste el curso con éxito!");
    window.location.pathname='/miscursos';
    } catch (err) {
        alert(err.message)
        setCurso({
            nombre:"",
            descripcion: "",
            duracion:"",
            requisitos: "",
        })
    }

}

return (

    <div>
        <HeaderMisCursos/>
        <div>           
            <h1>
               Estas en el recurso para modificar el curso que seleccionaste, por favor completa los siguentes campos:
            </h1>
            <div className="ingresarDatos">
                <form onSubmit={cursoModificado} >
                    <div className="cajas_inputs">
                        <h2>Nombre del curso nuevo</h2>
                        <input 
                        type="text"
                        name="nombre"
                        placeholder="Programación II"
                        value={cursomodif.nombre}
                        onChange={handleChange}
                        required
                        />
                    </div>
                        <div className="cajas_inputs">
                        <h2>Descripcion nueva del curso</h2>
                    <input 
                    type="text"
                    name="descripcion"
                    placeholder="Este curso trata sobre como aprender a programar"
                    value={cursomodif.descripcion}
                    onChange={handleChange}
                    required
                    />    
                    </div>

                    <div className="cajas_inputs">
                        <h2>Duracion nueva del curso</h2>
                        <input 
                        type="text"
                        name="duracion"
                        placeholder="3 meses"
                        value={cursomodif.duracion}
                        onChange={handleChange}
                        required    
                        />  
                        </div>

                      <div className="cajas_inputs">
                      <h2>Requisitos nuevos del curso.</h2>
                     <input 
                        type="text"
                        name="requisitos"
                        placeholder="Conocimientos en Fundamentos de Programacion"
                        value={cursomodif.requisitos}
                        onChange={handleChange}
                        required/>    
                        </div>
                    <button type="submit">Modificar curso</button>
                </form>
            </div>
        </div>
    </div>
    
)}
