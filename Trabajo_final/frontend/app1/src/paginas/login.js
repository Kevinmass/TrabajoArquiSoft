import React, { useState } from "react"
import { Header } from "../componentes/encabezado"

export const Login = () => {
    const [usuario, setUsuario] = useState({
        user: '',
        email: '',
        password: '',
    })


    const handleChange = (e) => {
        setUsuario({
            ...usuario,
            [e.target.name]: e.target.value,
        });
    };

    const handleSubmit = async (e) => {
        e.preventDefault()

        try {
            const response = await fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify( usuario )
            });
            if (!response.ok) {
                throw new Error('error al iniciar sesión')
            }
            const data = await response.json();
            const token = data.token;
            localStorage.setItem('sesionCookie', data.token)
            localStorage.setItem('user', usuario.user)

            alert('bienvenido ' + usuario.user)
            window.location.pathname = '/'
        } catch (error) {
            alert('error: ' + error.message);
            setUsuario({
                user: '',
                email: '',
                password: '',
            });
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
                            name="user"
                            placeholder="usuario"
                            onChange={handleChange}
                            value={usuario.user}
                            required
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Correo electrónico</h2>
                        <input
                            type="email"
                            name="email"
                            placeholder="email@gmail.com"
                            onChange={handleChange
                            }
                            value={usuario.email}
                            required
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Contraseña</h2>
                        <input
                            type="password"
                            name="password"
                            placeholder="contraseña"
                            onChange={handleChange}
                            value={usuario.password}
                            required
                        />
                    </div>

                    <button className="boton-sesion" type="submit">Iniciar sesión</button>
                </form>
            </div>
        </div>)

}