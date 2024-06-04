import React, { useState } from "react"
import { Header } from "../componentes/encabezado"

export const Login = () => {
    const [usuario, setUsuario] = useState('')
    const [password, setPassword] = useState('')


    const handleSubmit = async (e) => {
        e.preventDefault()

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ usuario, password })
            });
            if (!response.ok) {
                throw new Error('error al iniciar sesión')
            }
            const data =  response.json();
            const token = data.token;
            localStorage.setItem('sesionCookie', data.token)
        } catch (error) {
            alert('error: ' + error.message)
        }
    }




        return (
            <div className="container">
                <Header />
                <div className="registro">
                    <h1>Login</h1>

                    <form className="formulario" onSubmit={handleSubmit}>
                        <div className="cajas_inputs">
                            <h2>Usuario</h2>
                            <input
                                type="text"
                                placeholder="usuario"
                                value={usuario}
                                onChange={e => setUsuario(e.target.value)
                                }
                                required
                            />
                        </div>
                        <div className="cajas_inputs">
                            <h2>Contraseña</h2>
                            <input
                                type="password"
                                placeholder="contraseña"
                                value={password}
                                onChange={e => setPassword(e.target.value)}
                                required
                            />
                        </div>

                        <button className="boton-sesion" type="submit">Iniciar sesión</button>
                    </form>
                </div>
            </div>)

    }