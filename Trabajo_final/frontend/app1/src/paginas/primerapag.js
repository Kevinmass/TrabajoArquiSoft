import React from "react";
import { BrowserRouter as Router, Routes, Route, Link } from "react-router-dom";

export const PrimeraPag = () => {
   return  (
    <div>
        <li><Link to="/login" >Login</Link></li>
        <li><Link to="/register" >Registrarse</Link></li>
    </div>
)
}