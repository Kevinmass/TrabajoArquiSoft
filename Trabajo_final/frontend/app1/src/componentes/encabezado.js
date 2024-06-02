import React from "react";
import './estilos.css'
export const Header = () => {

    return (
        <header className="encabezadoPrincipal">

            <div className="titulo-encabezado">
                <h1> CURSOS ONLINE  DE LA UCC</h1>
            </div>


            <div className="botones-encabezado">
                <button className="boton-login ">Ingresar</button>
                <button className="boton-signup">Registrarse</button>

            </div>

        </header>
    )
}