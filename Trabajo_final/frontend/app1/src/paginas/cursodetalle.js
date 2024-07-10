import React, { useEffect, useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos";
export const Detalle = () => {
  const [objeto, setObjeto] = useState({});
  const [comentario, setComentario] = useState("");

  useEffect(() => {
    const curso = localStorage.getItem("cursoInfo");
    if (curso) {
      setObjeto(JSON.parse(curso));
    }
  }, []);

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
    } catch {
      alert("parece que hubo un error, trataremos de solucionarlo, no te preocupes.");
    }
  };

  return (
  
    <div>
          <HeaderMisCursos />
      <div className="cursoIndividual">
        <h1>{objeto.nombre}</h1>
        <p>Descripción: {objeto.descripcion}</p>
        <p>
          Profesor: {objeto.profesor_nombre} {objeto.profesor_apellido}
        </p>
        <p>Correo del profesor: {objeto.profesor_correo}</p>
        <p>Fecha de Creación: {objeto.fecha_creacion}</p>
        <p>Fecha de Modificación: {objeto.fecha_actualizacion}</p>
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
      </div>
    </div>
  );
};
