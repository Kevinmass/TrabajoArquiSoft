import './App.css';
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import { Login } from './paginas/login.js'
import { Register } from './paginas/signup.js';
import { Home } from './paginas/home.js';
import { PrimeraPag } from './paginas/primerapag.js'
import { Detalle } from './paginas/cursodetalle.js'
import { MisCursos } from './paginas/miscursos.js';
import { HomeProfesor } from './paginas/homeProfesor.js';
import {CrearCurso} from './paginas/crearCurso.js';
import { HomeArchivos } from './paginas/homeArchivos.js';
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
            <Route path='/crearCurso' Component={CrearCurso}/>
            <Route path='/homeArchivos' Component={HomeArchivos}/>


          </Routes>
        </Router>
      </div>
    </div>
  )
}
export default App;
