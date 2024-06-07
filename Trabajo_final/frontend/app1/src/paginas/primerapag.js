import React from "react";
import '../paginas/estilo_paginas.css'
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { Footer } from "../componentes/footer";
import { Header} from "../componentes/encabezado"

export const PrimeraPag = () => {
   return  (
    <div>
        <Header/>
<<<<<<< HEAD
        <div className="container-horizontal">
        <div class="izquierda">
        <h5>Ya tiene una cuenta creada ?</h5>
        <button><Link to="/login" >Iniciar sesion con tu cuenta</Link></button>
    </div>
    <br>
    </br>
    <div class="derecha">
        <h5>Acaso no tienes cuenta?</h5>
        
        <button><Link to="/register" >Registrate aqui</Link></button>
    </div>
    </div>
=======
        <li><Link to="/login" >Iniciar sesion</Link></li>
        <li><Link to="/register" >Quiero registrarme</Link></li>
>>>>>>> df67c94e23762101cdac8fb33eb4225dc5b52ab7
        <Footer/>
    </div>

)
}