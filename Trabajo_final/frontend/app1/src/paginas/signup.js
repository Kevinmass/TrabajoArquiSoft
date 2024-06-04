import React, { useState } from "react"
import { Header } from "../componentes/encabezado"

export const Register = () => {
    const [usuario, setUsuario] = useState({
        nombre: '',
        apellido: '',
        email: 'a',
        user: '',
        password:'' 
    })


    const handleChange = (e) => {
        setUsuario({
          [e.target.name]: e.target.value,
        });
      };


    const handleSubmit = async (e) => {
        e.preventDefault();
       try{ 
        const response =  await fetch('http://localhost:8080/register', {
            method :'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(usuario),
        });
        if(!response.ok){
            throw new Error('error al crear el usuario')
        }
alert('Usuario creado correctamente')
    } catch(error) {
        console.error('Error: ', error.message);
        alert('ocurrio un error al crear el usuario');
    }
    };

    return (
        <div>
            <Header />
            <h1>Register</h1>

            <form className="formulario_login" onSubmit={handleSubmit}>
                <h2>Nombre</h2>
                <input
                    type="text"
                    name="nombre"
                    placeholder="nombre"
                    value={usuario.nombre}
                    onChange={handleChange}
                />

                <h2>Apellido</h2>
                <input
                    type="text"
                    name="apellido"
                    placeholder="apellido"
                    value={usuario.apellido}
                    onChange={handleChange}
                />
                <h2>Correo Electronico</h2>
                <input
                    type="email"
                    name="email"
                    placeholder="email@example.com"
                    value={usuario.email}
                    onChange={handleChange}
                />
                <h2>Usuario</h2>
                <input
                    type="text"
                    name="user"
                    placeholder="usuario"
                    value={usuario.user}
                    onChange={handleChange}
                />
                <h2>Contraseña</h2>
                <input
                    type="password"
                    name="password"
                    placeholder="contraseña"
                    value={usuario.password}
                    onChange={handleChange}
                />

                <button className="boton-sesion" type="submit" >Registrarse</button>
            </form>
        </div>)

}