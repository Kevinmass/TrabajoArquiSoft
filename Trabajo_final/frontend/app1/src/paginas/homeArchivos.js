import React, { useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos";
import { Footer } from "../componentes/footer";

export const HomeArchivos = () => {
  const [file, setFile] = useState(null);
  const [cursoId, setCursoId] = useState("");
  const [files, setFiles] = useState([]);
  const [fetchCursoId, setFetchCursoId] = useState("");

  const handleArchivo = (event) => {
    const selectedFile = event.target.files[0];
    setFile(selectedFile);
  };

  const handleCursoIdChange = (event) => {
    setCursoId(event.target.value);
  };

  const handleFetchCursoIdChange = (event) => {
    setFetchCursoId(event.target.value);
  };

  const handleUploadArchivo = async () => {
    if (!file) {
      alert("Por favor seleccione un archivo!");
      return;
    }

    if (!cursoId) {
      alert("Por favor coloque un id de curso!");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);


    try {
      const response = await fetch("http://localhost:8080/POSTarchivo/"+ cursoId, {
        method: "POST",
        body: formData,
      });

      if (response.status != 500) {
        alert("Archivo subido exitosamente");
      } else {
        alert("Fallo al subir el archivo");
      }
    } catch (error) {
      console.error("Error subiendo archivo: ", error);
    }
  };

  const BuscarArchivos = async () => {
    if (!fetchCursoId) {
      alert("Please enter a curso ID!");
      return;
    }

    try {
      const response = await fetch("http://localhost:8080/GETarchivo/"+ fetchCursoId);

      if (response.ok) {
        const data = await response.json();
        setFiles(data);
      } else {
        alert("Fallo al conseguir los archivos");
      }
    } catch (error) {
      console.error("Error buscando archivos: ", error);
    }
  };

  return (
    <div>
      <HeaderMisCursos />
      <div>
        <h1>Estás en la página de Archivos.</h1>
        <div className="ingresarDatos">
          <div>
            <input type="file" onChange={handleArchivo} />
            {file && <p>Archivo seleccionado: {file.name}</p>}
            <input
              type="text"
              placeholder="Ingrese el ID del curso"
              value={cursoId}
              onChange={handleCursoIdChange}
            />
            <button onClick={handleUploadArchivo}>Subir</button>
          </div>
          <div>
            <h2>Buscador de archivos por curso</h2>
            <input
              type="text"
              placeholder="Ingrese el ID del curso"
              value={fetchCursoId}
              onChange={handleFetchCursoIdChange}
            />
            <button onClick={BuscarArchivos}>Buscar archivos</button>
            <ul>
              {files.map((file, index) => (
                <li key={index}>
                  <p>{file.name}</p>
                  {file.base64 && (
                    <img src={`data:image/png;base64,${file.base64}`} alt={file.name} width="200" />
                  )}
                </li>
              ))}
            </ul>
          </div>
          <Footer />
        </div>
      </div>
    </div>
  );
};
