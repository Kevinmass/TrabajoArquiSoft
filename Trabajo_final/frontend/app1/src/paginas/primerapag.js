import React from "react";
import '../paginas/estilo_paginas.css'
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
        <button><Link to="/login" >Iniciar sesion con tu cuenta</Link></button>
    </div>
    <br>
    </br>
    <div class="containers">
        <h5>Acaso no tienes cuenta?</h5>
        
        <button><Link to="/register" >Registrate aqui</Link></button>
        </div>
        </div>

        <Footer/>
    </div>

)
}