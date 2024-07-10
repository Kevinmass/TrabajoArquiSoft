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
        <h5>Ya tiene una cuenta creada ? </h5>
        <div className="botones-buscador">
        <button><Link to="/login" className="custom-link">Iniciar sesion con tu cuenta</Link></button>
        </div>
    </div>
    <br>
    </br>
    <div className="containers">
        <h5>No tienes cuenta?</h5>
        <div className="botones-buscador">
        <button><Link to="/register" className="custom-link" >Registrate aqui</Link></button>
        </div>
        </div>
        </div>

        <Footer/>
    </div>

)
}