import React, { useState } from "react"
import { Header } from "../componentes/encabezado"
import { Footer } from "../componentes/footer";


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

    const esProfesor = async () => {
        try{
            const response = await fetch("http://localhost:8080/isAdmin/"+ usuario.user)
        if(!response.ok){
            throw new Error('error al verificar si es administrador o estudiante.')
        }
        const data = await response.json();
        return data;
        
        } catch (error ){
        
            alert(error.message,": no se pudo verificar si es o no administrador de algun curso.")
            return false;
        }}
        
    const handleSubmit = async (e) => {
        e.preventDefault()

        try {
            const isProfesor = await esProfesor();
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
            localStorage.setItem('sesionCookie', data.token)
            localStorage.setItem('user', usuario.user)
            localStorage.setItem('profesor',isProfesor)
            alert('bienvenido ' + usuario.user)
            if(isProfesor){
                window.location.pathname = '/homeProfesor'
            } else{
            window.location.pathname = '/home'
            }
                
         
        
        } catch (error) {
            alert(error.message + ': los datos ingresados son incorrectos');
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
            <Footer/>
        </div>)

}