import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";
import { Footer } from "../componentes/footer";
import { Header} from "../componentes/encabezado"

export const PrimeraPag = () => {
   return  (
    <div>
        <Header/>
        <li><Link to="/login" >Login</Link></li>
        <li><Link to="/register" >Registrarse</Link></li>
        <Footer/>
    </div>

)
}