import React, { useState } from "react"
import { Header } from "../componentes/encabezado"
import '../paginas/estilo_paginas.css'
export const Register = () => {
    const [usuario, setUsuario] = useState({
        nombre: '',
        apellido: '',
        email: '',
        user: '',
        password: ''
    })


    const handleChange = (e) => {
        setUsuario({
            [e.target.name]: e.target.value,
        });
    };


    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await fetch('http://localhost:8080/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(usuario),
            });
            if (!response.ok) {
                throw new Error('error al crear el usuario')
            }
            alert('Usuario creado correctamente')
        } catch (error) {
            console.error('Error: ', error.message);
            alert('ocurrio un error al crear el usuario');
        }
    };

    return (
        <div className="container">
            <Header />

            <div className="registro">

                <h1>Register</h1>
                <form className="formulario" onSubmit={handleSubmit}>

                    <div className="cajas_inputs">
                        <h2>Nombre</h2>
                        <input
                            type="text"
                            name="nombre"
                            placeholder="Juan"
                            value={usuario.nombre}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Apellido</h2>
                        <input
                            type="text"
                            name="apellido"
                            placeholder="Muruzabal"
                            value={usuario.apellido}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Correo Electrónico</h2>
                        <input
                            type="email"
                            name="email"
                            placeholder="email@example.com"
                            value={usuario.email}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Usuario</h2>
                        <input
                            type="text"
                            name="user"
                            placeholder="usuario1234"
                            value={usuario.user}
                            onChange={handleChange}
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Contraseña</h2>
                        <input
                            type="password"
                            name="password"
                            placeholder="12345abcd"
                            value={usuario.password}
                            onChange={handleChange}
                        />
                    </div>
                    <button className="boton-sesion" type="submit" >Registrarse</button>
                </form>
            </div >
        </div >)
}