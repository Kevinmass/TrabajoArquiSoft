import React, { useState } from "react";
import { HeaderMisCursos } from "../componentes/encabezadoMisCursos"
import { Footer } from "../componentes/footer";
export const HomeArchivos = () => {
  
  const [file, setFile] = useState(null);

  const handleFileChange = (event) => {
    const selectedFile = event.target.files[0];
    setFile(selectedFile);
  };

  const handleFileUpload = async () => {
    if (!file) {
      alert("Please select a file first!");
      return;
    }
  
    const formData = new FormData();
    formData.append("file", file);
  
      console.log("Uploading file:", file);
    try {
      const response = await fetch('http://localhost:8080/POSTarchivo', {
        method: 'POST',
        body: formData
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
  
      return (
  
          <div>
              <HeaderMisCursos/>
              <div>           
                  <h1>
                      Estas en la página de crear cursos, a continuacion se te pediran los datos que se necesita para crear un curso.
                  </h1>
                  <div className="ingresarDatos">
                          <div>
                          <input type="file" onChange={handleFileChange}  />
                          {file && <p>Archivo seleccionado: {file.name}</p>}
                          <button onClick={handleFileUpload}>Subir</button>
                          </div>
                          <Footer/>
                  </div>
              </div>
          </div>
          
      );
    }