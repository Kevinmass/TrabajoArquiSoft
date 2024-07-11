import React, { useEffect, useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos";

export const Detalle = () => {
  const [objeto, setObjeto] = useState({});
  const [comentario, setComentario] = useState("");
const [mostrarComent, setmostrarComent] = useState ([]);

  useEffect(() => {
   
    const curso = localStorage.getItem("cursoInfo");
    if (curso) {
      setObjeto(JSON.parse(curso));
    }
  }, []);

  useEffect(()=>{
    if (objeto.id) {
      mostrarComentarios();
    } 
   },[objeto]);

  const handleInputChange = (event) => {
    setComentario(event.target.value);
  };

  const Comentario = async () => {
    if (comentario.trim() === "") {
        alert("El comentario no puede estar vacío");
        return;
      }
    const comentario1 = {
      curso_id: objeto.id,
      user: localStorage.getItem("user"),
      comment: comentario,
    };

    try {
      const response = await fetch("http://localhost:8080/POSTcomentario", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(comentario1),
      });
      if (!response.ok) {
        throw new Error("error al enviar el comentario");
      }
      const coment = await response.json();
      alert("comentario cargado correctamente!");
      mostrarComentarios();
    } catch {
      alert("parece que hubo un error, trataremos de solucionarlo, no te preocupes.");
    }
  };


const mostrarComentarios = async () => {
try{
  const response = await fetch("http://localhost:8080/GETcomentarios?curso_id="+ objeto.id)
  if (!response.ok){
    throw new Error("error al obtener los comentarios");
  }
  const data = await response.json();
  setmostrarComent(data);
} catch (error){console.error("error al mostrar los comentarios", error)};

}








  return (
  
    <div>
          <HeaderMisCursos />
      <div className="cursoIndividual">
        <h1>{objeto.nombre}</h1>
        <p>Descripción: {objeto.descripcion}</p>
        <p>
          Profesor: {objeto.user}
        </p>
        <p>Duracion: {objeto.duracion}</p>
        <p>Requisitos: {objeto.requisitos}</p>
        <form>
          <input
            type="text"
            name="comment"
            placeholder="Aqui puedes ingresar un comentario sobre el curso."
            value={comentario}
            onChange={handleInputChange}
            required
          />
        </form>
        <button onClick={Comentario}>
          Enviar comentario
        </button>
        <div className="div-comentarios">
          <h1>Comentarios sobre el curso:</h1>
        {mostrarComent.map((coment, index) => (
          <div key={index}>
            <div className="" >
              <p>User: {coment.user}</p>
              <p>Comentario: {coment.comment}</p>
            </div>
           </div>
          ))}

        
        
        </div>
      </div>
    </div>
  );

}