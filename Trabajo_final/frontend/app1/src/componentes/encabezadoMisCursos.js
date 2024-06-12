import React from "react";
import './estilos.css'
import { Link } from 'react-router-dom'


export const HeaderMisCursos = () => {

    return (
        <header className="encabezadoPrincipal">

            <div className="titulo-encabezado">
                <h1> CURSOS ONLINE  DE LA UCC</h1>
            </div>


            <div className="links-encabezado">
                <ul>
                    <li><Link to="/home" >Home</Link></li>

                </ul>
            </div>

        </header>

    );
}