import React from "react";
import './estilos.css'
import { Link } from 'react-router-dom'


export const Header = () => {

    return (
        <header className="encabezadoPrincipal">

            <div className="titulo-encabezado">
                <h1> CURSOS ONLINE  DE LA UCC</h1>
            </div>


            <div className="links-encabezado">
                <ul>

                    <li><Link to="/login" className="custom-link">Login</Link></li>
                    <li><Link to="/register"className="custom-link" >Registrarse</Link></li>
                </ul>
            </div>

        </header>

    );
}