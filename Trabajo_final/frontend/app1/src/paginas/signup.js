import React, { useState } from "react"
import { Header } from "../componentes/encabezado"
import '../paginas/estilo_paginas.css'
import { Footer } from "../componentes/footer"
export const Register = () => {
    const [usuario, setUsuario] = useState({
        nombre: '',
        apellido: '',
        email: '',
        user: '',
        password: '',
        profesor: false
    })


    const handleChange = (e) => {
        const value = e.target.type === 'checkbox' ? e.target.checked : e.target.value;
        setUsuario({
            ...usuario,
            [e.target.name]: value,
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
            localStorage.setItem('user', usuario.user)
            alert('Usuario creado correctamente');
            window.location.pathname = '/home'
        } catch (error) {
            console.error('Error: ', error.message);
            alert('error: ya existe una persona con ese usuario/email');
            setUsuario({
                nombre: '',
                apellido: '',
                email: '',
                user: '',
                password: '',
                profesor: false
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

                    <div className="cajas_inputs">
                        <h2>Es profesor?</h2>
                        <label><input
                            type="checkbox"
                            name="profesor"
                            checked={usuario.profesor}
                            onChange={handleChange}                        
                        />
                            Si
                        </label>
                    </div>

                    <button className="boton-sesion" type="submit" >Registrarse</button>
                </form>
            </div >
            <Footer/>

        </div >)
}