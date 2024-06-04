import React, { useState } from "react"
import { Header } from "../componentes/encabezado"

export const Login = () => {
    const [usuario, setUsuario] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState(false)

    const handleSubmit = (e) => {
        e.preventDefault()
        if (usuario === "" || password === "") {
            setError(true)
            return
        }
        console.log(usuario + password)
    }

    const iniciarSesion = () => {

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
                            onChange={e => setUsuario(e.target.value)}
                        />
                    </div>
                    <div className="cajas_inputs">
                        <h2>Contraseña</h2>
                        <input
                            type="password"
                            placeholder="contraseña"
                            value={password}
                            onChange={e => setPassword(e.target.value)}
                        />
                    </div>

                    <button className="boton-sesion" type="submit">Iniciar sesión</button>
                </form>
            </div>
        </div>)

}