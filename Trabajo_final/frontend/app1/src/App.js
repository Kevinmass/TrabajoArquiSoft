import logo from './logo.svg';
import './App.css';
import React, { useEffect, useState } from 'react';
import { queryByTestId } from '@testing-library/react';
import { Header } from './componentes/encabezado.js';
import { BrowserRouter as Router, Route, Switch, Routes } from 'react-router-dom';
import { Login } from './paginas/login.js'
import { Register } from './paginas/signup.js';
import { Home } from './paginas/home.js';
import { PrimeraPag } from './paginas/primerapag.js'
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

          </Routes>
        </Router>
      </div>
    </div>
  )
}
export default App;
