import React, { useState } from "react"
import { Header } from "../componentes/encabezado"
import '../paginas/estilo_paginas.css'
export const Register = () => {
    const [usuario, setUsuario] = useState({
        nombre: '',
        apellido: '',
        email: '',
        user: '',
        password: '',
    })


    const handleChange = (e) => {
        setUsuario({
            ...usuario,
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
            alert('Usuario creado correctamente');
            window.location.pathname = '/'
        } catch (error) {
            console.error('Error: ', error.message);
            alert('ocurrio un error al crear el usuario');
            setUsuario({
                nombre: '',
                apellido: '',
                email: '',
                user: '',
                password: '',
            });
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
                            onChange={handleChange}
                            value={usuario.nombre}
                            required
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Apellido</h2>
                        <input
                            type="text"
                            name="apellido"
                            placeholder="Muruzabal"
                            onChange={handleChange}
                            value={usuario.apellido}
                            required

                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Correo Electrónico</h2>
                        <input
                            type="email"
                            name="email"
                            placeholder="email@example.com"
                            onChange={handleChange}
                            value={usuario.email}
                            required

                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Usuario</h2>
                        <input
                            type="text"
                            name="user"
                            placeholder="usuario1234"
                            onChange={handleChange}
                            value={usuario.user}
                            required

                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Contraseña</h2>
                        <input
                            type="password"
                            name="password"
                            placeholder="12345abcd"
                            onChange={handleChange}
                            value={usuario.password}
                            required

                        />
                    </div>

                    <button className="boton-sesion" type="submit" >Registrarse</button>
                </form>
            </div >
        </div >)
}