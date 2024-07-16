import React, { useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos";
import { Footer } from "../componentes/footer";

export const HomeArchivos = () => {
  const [file, setFile] = useState(null);
  const [cursoId, setCursoId] = useState("");
  const [files, setFiles] = useState([]);
  const [fetchCursoId, setFetchCursoId] = useState("");

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    setFile(selectedFile);
  };

  const handleCursoIdChange = (event) => {
    setCursoId(event.target.value);
  };

  const handleFetchCursoIdChange = (event) => {
    setFetchCursoId(event.target.value);
  };

  const handleFileUpload = async () => {
    if (!file) {
      alert("Please select a file first!");
      return;
    }

    if (!cursoId) {
      alert("Please enter a curso ID!");
      return;
    }

    const formData = new FormData();
    formData.append("file", file);


    try {
      const response = await fetch("http://localhost:8080/POSTarchivo/"+ cursoId, {
        method: "POST",
        body: formData,
      });

      if (response.status === 201) {
        alert("File uploaded successfully!");
      } else {
        alert("Failed to upload file.");
      }
    } catch (error) {
      console.error("Error uploading file:", error);
    }
  };

  const fetchFiles = async () => {
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
        alert("Failed to fetch files.");
      }
    } catch (error) {
      console.error("Error fetching files:", error);
    }
  };

  return (
    <div>
      <HeaderMisCursos />
      <div>
        <h1>Estás en la página de Archivos.</h1>
        <div className="ingresarDatos">
          <div>
            <input type="file" onChange={handleFileChange} />
            {file && <p>Archivo seleccionado: {file.name}</p>}
            <input
              type="text"
              placeholder="Ingrese el ID del curso"
              value={cursoId}
              onChange={handleCursoIdChange}
            />
            <button onClick={handleFileUpload}>Subir</button>
          </div>
          <div>
            <h2>Fetch Files</h2>
            <input
              type="text"
              placeholder="Ingrese el ID del curso"
              value={fetchCursoId}
              onChange={handleFetchCursoIdChange}
            />
            <button onClick={fetchFiles}>Fetch Files</button>
            <ul>
              {files.map((file, index) => (
                <li key={index}>{file.name}</li>
              ))}
            </ul>
          </div>
          <Footer />
        </div>
      </div>
    </div>
  );
};
