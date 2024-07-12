import React, { useState } from "react";
import { Header } from "../componentes/encabezado";
import { Footer } from "../componentes/footer";

export const HomeArchivos = () => {
  const [content, setContent] = useState("");
  const [error, setError] = useState(null);

  const handleCreateAndWriteFile = async () => {
    try {
      console.log(content);
      const response = await fetch("http://localhost:8080/POSTarchivo", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ content }),
      });

      if (!response.ok) {
        throw new Error("Failed to create and write file");
      }

      setError(null);
    } catch (error) {
      setError(error.message);
    }
  };
/*
  const handleReadFile = async () => {
    try {
      const response = await fetch("http://localhost:8080/ReadFile", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ path: filePath }),
      });

      if (!response.ok) {
        throw new Error("Failed to read file");
      }

      const data = await response.json();
      // Handle the received file content data
      console.log(data);
      setError(null);
    } catch (error) {
      setError(error.message);
    }
  };
*/
  return (
    <div>
      <Header />
      <form>
        
        <label htmlFor="content">Contenido:</label>
        <textarea
          id="content"
          value={content}
          onChange={(e) => setContent(e.target.value)}
        ></textarea>
        <br />
        <button type="button" onClick={handleCreateAndWriteFile}>
          Crear archivo
        </button>
      </form>
      {error && <p>{error}</p>}
      <Footer />
    </div>
  );
};

