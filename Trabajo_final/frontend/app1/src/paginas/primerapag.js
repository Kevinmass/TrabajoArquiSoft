import React from "react";
import '../componentes/estilos.css'
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { Footer } from "../componentes/footer";
import { Header} from "../componentes/encabezado"

export const PrimeraPag = () => {
   return  (
    <div>
        <Header/>
        <div className="container-horizontal">
        <div className="containers">
        <h5>Tienes una cuenta creada? </h5>
        <div className="botones-buscador">
        <Link to="/login" className="link-button">Iniciar sesión</Link>
        </div>
    </div>
    <br>
    </br>
    <div className="containers">
        <h5>No tienes cuenta?</h5>
        <div className="botones-buscador">
        <Link to="/register" className="link-button" >Registrate aqui</Link>
        </div>
        </div>
        </div>

        <Footer/>
    </div>

)
}