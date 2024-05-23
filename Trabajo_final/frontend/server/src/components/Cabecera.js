import React from 'react'
import logo from '../recursos/gojo.png'
export const Cabecera = () => {
  return (
    <div className="container">
            <nav id="navbar">
                <a href="#tope" target="_self">
                    <img src={logo} className='logo' alt="Imagen Logo"/>
                </a>
            </nav>
        </div>
  )
}
