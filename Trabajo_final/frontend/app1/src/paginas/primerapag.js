import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { Footer } from "../componentes/footer";
import { Header} from "../componentes/encabezado"

export const PrimeraPag = () => {
   return  (
    <div>
        <Header/>
        <li><Link to="/login" >Iniciar sesion</Link></li>
        <li><Link to="/register" >Quiero registrarme</Link></li>
        <Footer/>
    </div>

)
}