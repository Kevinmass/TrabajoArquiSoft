import './App.css';
import React, { useEffect, useState } from 'react';
import { BrowserRouter as Router, Route, Switch, Routes } from 'react-router-dom';
import { Login } from './paginas/login.js'
import { Register } from './paginas/signup.js';
import { Home } from './paginas/home.js';
import { PrimeraPag } from './paginas/primerapag.js'
import { Detalle } from './paginas/cursodetalle.js'
import { MisCursos } from './paginas/miscursos.js';
import { HomeProfesor } from './paginas/homeProfesor.js';
function App() {
  return (
    <div>

      <div>
        <Router>
          <Routes>
            <Route path="/" Component={PrimeraPag} />
            <Route path="/home" Component={Home} />
            <Route path="/login" Component={Login} />
            <Route path="/register" Component={Register} />
            <Route path='/cursodetalle' Component={Detalle} />
            <Route path="/miscursos" Component={MisCursos} />
            <Route path='/homeProfesor' Component={HomeProfesor}/>


          </Routes>
        </Router>
      </div>
    </div>
  )
}
export default App;
