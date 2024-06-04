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
        <div>
            <Header/>
            <h1>Login</h1>

            <form className="formulario_login" onSubmit={handleSubmit}>
                <h2>ingrese el usuario</h2>
                <input
                    type="text"
                    placeholder="usuario"
                    value={usuario}
                    onChange={e => setUsuario(e.target.value)}
                />
                <h2>ingrese la contraseña</h2>
                <input
                    type="password"
                    placeholder="contraseña"
                    value={password}
                    onChange={e => setPassword(e.target.value)}
                />

            <button className="boton-sesion" type="submit">Iniciar sesión</button>
            </form>
        </div>)

}